import loader from "@monaco-editor/loader";
import { editor } from "monaco-editor";
import { create } from "zustand";
import { toast } from "sonner";
import { Clipboard, Dialogs, JSON2Go } from "@api/app/services";
type IStandaloneCodeEditor = editor.IStandaloneCodeEditor;
type IStandaloneEditorConstructionOptions = editor.IStandaloneEditorConstructionOptions;

function loaderConfig(language: string) {
    return {
        paths: { vs: "monaco-editor/min/vs" },
        "vs/nls": {
            availableLanguages: {
                "*": getMonacoLocale(language),
            },
        },
    };
}

const monacoLocales = ["en", "de", "es", "fr", "it", "ja", "ko", "ru", "zh-cn", "zh-tw"];

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

function editorOptions(language: string, fontSize: number, value: string): IStandaloneEditorConstructionOptions {
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

function replaceContent(editor: IStandaloneCodeEditor, content: string) {
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

interface EditorsState {
    jsonEditor: IStandaloneCodeEditor | null;
    goEditor: IStandaloneCodeEditor | null;
    init: (language: string, fontSize: number, jsonID: string, goID: string) => void;
    generate: () => void;
    setJSON: (value: string) => void;
    pasteJSON: () => void;
    copyGo: () => void;
    openJSON: () => void;
    saveGo: () => void;
}

export const useEditorsStore = create<EditorsState>((set, get) => ({
    jsonEditor: null,
    goEditor: null,
    init: (language, fontSize, jsonID, goID) => {
        loader.config(loaderConfig(language));
        void loader.init().then((monaco) => {
            const jsonElement = document.getElementById(jsonID);
            const goElement = document.getElementById(goID);

            if (!(jsonElement && goElement)) {
                return;
            }

            const jsonValue = get().jsonEditor?.getValue() || "";
            const goValue = get().goEditor?.getValue() || "";

            get().jsonEditor?.dispose();
            get().goEditor?.dispose();
            jsonElement.innerHTML = "";
            goElement.innerHTML = "";
            for (const model of monaco.editor.getModels()) {
                model.dispose();
            }

            const jsonEditor = monaco.editor.create(jsonElement, editorOptions("json", fontSize, jsonValue));
            const goEditor = monaco.editor.create(goElement, editorOptions("go", fontSize, goValue));

            set({
                jsonEditor: jsonEditor,
                goEditor: goEditor,
            });

            // remeasure fonts after creating editors and fonts are loaded to avoid rendering issues
            void document.fonts.ready.then(() => {
                monaco.editor.remeasureFonts();
            });
        });
    },
    generate: async () => {
        toast.dismiss();
        get().goEditor?.setValue(await JSON2Go.Generate(get().jsonEditor?.getValue() || ""));
    },
    setJSON: (value) => {
        const editor = get().jsonEditor;
        if (editor) {
            replaceContent(editor, value);
        }
    },
    pasteJSON: async () => {
        const editor = get().jsonEditor;
        if (editor) {
            replaceContent(editor, await Clipboard.Read());
        }
    },
    copyGo: async () => {
        const editor = get().goEditor;
        if (editor) {
            void Clipboard.Write(editor.getValue());
        }
    },
    openJSON: async () => {
        const content = await Dialogs.OpenJSONFile();
        const editor = get().jsonEditor;
        if (content && editor) {
            replaceContent(editor, content);
        }
    },
    saveGo: async () => {
        const editor = get().goEditor;
        if (editor) {
            void Dialogs.SaveGoSourceFile(editor.getValue());
        }
    },
}));
