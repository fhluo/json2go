import "./app.css"
import {
    Generate,
    GetFontSize,
    GetLocale,
    GetOptionsGenerateInRealTime,
    GetOptionsValidJSONBeforeGeneration,
    OpenJSONFile,
    SaveGoSourceFile,
    SetFontSize,
    SetOptionsGenerateInRealTime,
    SetOptionsValidJSONBeforeGeneration
} from "../wailsjs/go/main/App.js"
import {Layout, View} from "./lib/types"
import {ReactNode, useEffect, useRef, useState} from "react"
import Menubar from "@/ui/menubar.tsx"
import {editor} from "monaco-editor"
import {useTranslation} from "react-i18next"
import {copy, editorOptions, loaderConfig, paste, replaceContent} from "@/lib/utils.ts"
import FileMenu from "@/ui/menus/file.tsx"
import OptionsMenu from "@/ui/menus/options.tsx"
import ViewMenu from "@/ui/menus/view.tsx"
import FontMenu from "@/ui/menus/font.tsx"
import LanguageMenu from "@/ui/menus/language.tsx"
import ExamplesMenu from "@/ui/menus/examples.tsx"
import HelpMenu from "@/ui/menus/help.tsx"
import JSON from "@/ui/json.tsx"
import Go from "@/ui/go.tsx"
import loader from "@monaco-editor/loader"
import {EventsEmit, EventsOn} from "../wailsjs/runtime"
import Container from "@/ui/container.tsx"
import Footer from "@/ui/footer.tsx"
import IStandaloneCodeEditor = editor.IStandaloneCodeEditor

const defaultFontSize = 16

function App() {
    const {i18n} = useTranslation()
    const [fontSize, setFontSize] = useState(defaultFontSize)
    const [language, setLanguage] = useState("en")
    const [message, setMessage] = useState("")

    const [validJSON, setValidJSON] = useState(false)
    const [realTime, setRealTime] = useState(false)

    const [layout, setLayout] = useState(Layout.TwoColumns)
    const [view, setView] = useState(View.JSONAndGo)
    const [dialog, setDialog] = useState<ReactNode>(<></>)

    useEffect(() => {
        Promise.all([
            GetLocale(), GetFontSize(),
            GetOptionsValidJSONBeforeGeneration(),
            GetOptionsGenerateInRealTime()
        ]).then(
            ([locale, fontSize, validJSON, realTime]) => {
                locale && setLanguage(locale)
                setFontSize(fontSize)
                setValidJSON(validJSON)
                setRealTime(realTime)
            }
        )
    }, [])
    useEffect(() => void i18n.changeLanguage(language), [language])
    useEffect(() => void SetFontSize(fontSize), [fontSize])
    useEffect(() => void SetOptionsValidJSONBeforeGeneration(validJSON), [validJSON])
    useEffect(() => void SetOptionsGenerateInRealTime(realTime), [realTime])

    let jsonEditor = useRef<IStandaloneCodeEditor | null>(null)
    let goEditor = useRef<IStandaloneCodeEditor | null>(null)

    const generate = () => {
        setMessage("")
        Generate(jsonEditor.current?.getValue() || "").then(result => {
            goEditor.current?.setValue(result)
        })
    }

    useEffect(() => {
        loader.config(loaderConfig(language))
        loader.init().then(monaco => {
            document.getElementById("json-editor")!.innerHTML = ""
            document.getElementById("go-editor")!.innerHTML = ""
            monaco.editor.getModels().forEach(model => model.dispose())

            jsonEditor.current = monaco.editor.create(
                document.getElementById("json-editor")!,
                editorOptions("json", fontSize, "")
            )
            goEditor.current = monaco.editor.create(
                document.getElementById("go-editor")!,
                editorOptions("go", fontSize, "")
            )

            // remeasure fonts after creating editors and fonts are loaded to avoid rendering issues
            document.fonts.ready.then(() => {
                monaco.editor.remeasureFonts()
            })

            if (jsonEditor.current) {
                jsonEditor.current.getModel()?.onDidChangeContent(() => {
                    if (realTime) {
                        generate()
                    }
                })
            }
        })

        document.defaultView?.addEventListener("resize", () => {
            jsonEditor.current?.layout()
            goEditor.current?.layout()
            EventsEmit("resize")
        })

        EventsOn("error", (message: string) => {
            setMessage(message)
        })
    }, [])


    useEffect(() => {
        jsonEditor.current?.updateOptions({fontSize})
        goEditor.current?.updateOptions({fontSize})
    }, [fontSize])

    return (
        <main className="w-screen h-screen flex flex-col">
            <Menubar dialog={dialog}>
                <FileMenu setDialog={dialog => setDialog(dialog)} openJSON={() =>
                    OpenJSONFile().then(value =>
                        value !== "" && replaceContent(jsonEditor.current!, value)
                    )
                } saveGo={() => SaveGoSourceFile(goEditor.current!.getValue())}/>
                <OptionsMenu validJSON={validJSON} setValidJSON={setValidJSON}
                             realTime={realTime} setRealTime={setRealTime}/>
                <ViewMenu view={view} setView={setView} layout={layout} setLayout={setLayout}/>
                <FontMenu increaseFontSize={() => setFontSize(fontSize + 1)}
                          decreaseFontSize={() => setFontSize(fontSize - 1)}
                          resetFontSize={() => setFontSize(defaultFontSize)}/>
                <LanguageMenu/>
                <ExamplesMenu replaceContent={content => replaceContent(jsonEditor.current!, content)}/>
                <HelpMenu setDialog={setDialog}/>
            </Menubar>

            <Container layout={layout}>
                <JSON view={view} layout={layout} pasteJSON={() => paste(jsonEditor.current!)}/>
                <Go view={view} layout={layout} copyGoCode={() => copy(goEditor.current!)}/>
            </Container>

            <Footer message={message} clearMessage={() => setMessage("")} generate={generate}/>
        </main>
    )
}

export default App
