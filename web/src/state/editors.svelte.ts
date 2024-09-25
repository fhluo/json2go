import {editor} from "monaco-editor";
import loader from "@monaco-editor/loader"
import {Generate, OpenJSONFile, ReadClipboard, SaveGoSourceFile, WriteClipboard} from "../wailsjs/go/main/App";

function loaderConfig(language: string) {
    return {
        paths: {vs: "monaco-editor/min/vs"},
        "vs/nls": {
            availableLanguages: {
                "*": getMonacoLocale(language),
            },
        },
    };
}

const monacoLocales = [
    "en",
    "de",
    "es",
    "fr",
    "it",
    "ja",
    "ko",
    "ru",
    "zh-cn",
    "zh-tw",
];

function getMonacoLocale(locale: string): string {
    locale = locale.toLowerCase();
    if (monacoLocales.includes(locale)) {
        return locale;
    }

    if (locale === "zh") {
        return "zh-cn";
    }

    return "en";
}

function editorOptions(
    language: string,
    fontSize: number,
    value: string,
): editor.IStandaloneEditorConstructionOptions {
    return {
        value: value,
        language: language,
        fontFamily: "Jetbrains Mono, monospace",
        fontSize: fontSize,
        minimap: {
            enabled: false,
        },
        lineHeight: 25,
        automaticLayout: true,
    };
}

function replaceContent(editor: editor.IStandaloneCodeEditor | null, content: string) {
    const range = editor?.getModel()?.getFullModelRange();
    if (range) {
        editor?.executeEdits("", [
            {
                range: range,
                text: content,
            },
        ]);
    }
}


export function createEditorsState(
    initial: {
        language: string,
        fontSize: number,
        jsonElement: HTMLElement,
        goElement: HTMLElement
    }
) {
    let jsonEditor: editor.IStandaloneCodeEditor | null = null
    let goEditor: editor.IStandaloneCodeEditor | null = null

    const jsonElement = initial.jsonElement
    const goElement = initial.goElement
    const fontSize = initial.fontSize

    return {
        init() {
            loader.config(loaderConfig(initial.language))
            loader.init().then(monaco => {
                jsonElement.innerHTML = "";
                goElement.innerHTML = "";
                for (const model of monaco.editor.getModels()) {
                    model.dispose();
                }

                jsonEditor = monaco.editor.create(
                    jsonElement,
                    editorOptions("json", fontSize, ""),
                );
                goEditor = monaco.editor.create(
                    goElement,
                    editorOptions("go", fontSize, ""),
                );

                document.fonts.ready.then(() => monaco.editor.remeasureFonts());
            })
        },
        get json() {
            return jsonEditor?.getValue() || ""
        },
        set json(value) {
            if (jsonEditor) {
                replaceContent(jsonEditor, value)
            }
        },
        get go() {
            return goEditor?.getValue() || ""
        },
        set go(value) {
            if (goEditor) {
                replaceContent(goEditor, value)
            }
        },
        async json2go() {
            this.go = await Generate(this.json)
        },
        async pasteToJSONEditor() {
            this.json = await ReadClipboard()
        },
        async copyFromGoEditor() {
            await WriteClipboard(this.go)
        },
        async openJSONFile() {
            this.json = await OpenJSONFile()
        },
        async saveGoFile() {
            await SaveGoSourceFile(this.go)
        }
    }
}

export type EditorsState = ReturnType<typeof createEditorsState>
