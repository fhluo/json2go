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
import {useEffect, useState} from "react"
import {Menubar} from "@/components/ui/menubar.tsx"
import {editor} from "monaco-editor"
import {useTranslation} from "react-i18next"
import {cn, copy, editorOptions, getMonacoLocale, paste, replaceContent} from "@/lib/utils.ts"
import FileMenu from "@/ui/menus/file.tsx"
import OptionsMenu from "@/ui/menus/options.tsx"
import ViewMenu from "@/ui/menus/view.tsx"
import FontMenu from "@/ui/menus/font.tsx"
import LanguageMenu from "@/ui/menus/language.tsx"
import ExamplesMenu from "@/ui/menus/examples.tsx"
import HelpMenu from "@/ui/menus/help.tsx"
import {Dialog} from "@/components/ui/dialog.tsx"
import JSON from "@/ui/json.tsx"
import Go from "@/ui/go.tsx"
import Message from "@/ui/message.tsx"
import {Button} from "@/components/ui/button.tsx"
import loader from "@monaco-editor/loader"
import {EventsEmit, EventsOn} from "../wailsjs/runtime"
import IStandaloneCodeEditor = editor.IStandaloneCodeEditor

const defaultFontSize = 16

function App() {
    const {i18n} = useTranslation()
    const [fontSize, setFontSize] = useState(defaultFontSize)
    const [validJSON, setValidJSON] = useState(false)
    const [realTime, setRealTime] = useState(false)
    const [language, setLanguage] = useState("en")

    useEffect(() => {
        Promise.all([
            GetLocale(), GetFontSize(),
            GetOptionsValidJSONBeforeGeneration(),
            GetOptionsGenerateInRealTime()
        ]).then(
            ([locale, fontSize, validJSON, realTime]) => {
                if (locale) {
                    setLanguage(locale)
                }
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

    let [jsonEditor, setJSONEditor] = useState<IStandaloneCodeEditor | null>(null)
    let [goEditor, setGoEditor] = useState<IStandaloneCodeEditor | null>(null)

    useEffect(() => {
        loader.config({
            paths: {vs: "monaco-editor/min/vs"},
            "vs/nls": {
                availableLanguages: {
                    "*": getMonacoLocale(language)
                }
            }
        })
        loader.init().then(monaco => {
            document.getElementById("json-editor")!.innerHTML = ""
            document.getElementById("go-editor")!.innerHTML = ""
            monaco.editor.getModels().forEach(model => model.dispose())

            setJSONEditor(monaco.editor.create(
                document.getElementById("json-editor")!,
                editorOptions("json", fontSize, "")
            ))
            setGoEditor(monaco.editor.create(
                document.getElementById("go-editor")!,
                editorOptions("go", fontSize, "")
            ))

            // remeasure fonts after creating editors and fonts are loaded to avoid rendering issues
            document.fonts.ready.then(() => {
                monaco.editor.remeasureFonts()
            })
        })
    }, [])

    let [message, setMessage] = useState("")

    function generate() {
        setMessage("")
        if (!jsonEditor) {
            console.log("no json editor")
        }
        Generate(jsonEditor?.getValue() || "").then(result => {
            goEditor?.setValue(result)
        })
    }

    useEffect(() => {
        if (jsonEditor) {
            jsonEditor.getModel()?.onDidChangeContent(() => {
                if (realTime) {
                    generate()
                }
            })
        }
    }, [jsonEditor])

    useEffect(() => {
        jsonEditor?.updateOptions({fontSize})
        goEditor?.updateOptions({fontSize})
    }, [fontSize, jsonEditor, goEditor])

    useEffect(() => {
        document.defaultView?.addEventListener("resize", () => {
            jsonEditor?.layout()
            goEditor?.layout()
            EventsEmit("resize")
        })
    }, [])

    useEffect(() => {
        EventsOn("error", (message: string) => {
            setMessage(message)
        })
    }, [])

    let replaceJSONEditorContent = (content: string) => {
        replaceContent(jsonEditor!, content)
    }
    let openJSON = () => void OpenJSONFile().then(value => {
        if (value !== "") {
            replaceJSONEditorContent(value)
        }
    })
    let saveGo = () => void SaveGoSourceFile(goEditor!.getValue())

    let [layout, setLayout] = useState(Layout.TwoColumns)
    let [view, setView] = useState(View.JSONAndGo)
    let [dialog, setDialog] = useState(<></>)

    return (
        <main className="w-screen h-screen flex flex-col">
            <Dialog>
                <Menubar className="rounded-none border-none bg-transparent">
                    <FileMenu setDialog={dialog => setDialog(dialog)} openJSON={openJSON} saveGo={saveGo}/>
                    <OptionsMenu validJSON={validJSON}
                                 setValidJSON={setValidJSON}
                                 realTime={realTime}
                                 setRealTime={setRealTime}/>
                    <ViewMenu view={view} setView={setView} layout={layout} setLayout={setLayout}/>
                    <FontMenu increaseFontSize={() => setFontSize(fontSize + 1)}
                              decreaseFontSize={() => setFontSize(fontSize - 1)}
                              resetFontSize={() => setFontSize(defaultFontSize)}/>
                    <LanguageMenu/>
                    <ExamplesMenu replaceContent={replaceJSONEditorContent}/>
                    <HelpMenu setDialog={setDialog}/>
                </Menubar>
                {dialog}
            </Dialog>
            <div className={cn("grid h-64 grow border-t border-b", {
                "grid-cols-2": layout === Layout.TwoColumns,
                "grid-rows-2": layout === Layout.TwoRows,
            })}>
                <JSON view={view} layout={layout} pasteJSON={() => paste(jsonEditor!)}/>
                <Go view={view} layout={layout} copyGoCode={() => copy(goEditor!)}/>
            </div>
            <div className="flex flex-row px-4 py-2 justify-end items-center h-12">
                {message && <Message message={message} clearMessage={() => setMessage("")}/>}
                <Button size="sm" onClick={generate} className="mr-2">{"Generate"}</Button>
            </div>
        </main>
    )
}

export default App
