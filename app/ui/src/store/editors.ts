import loader from "@monaco-editor/loader";
import { editor } from "monaco-editor";
import { create } from "zustand";
import {
	Generate,
	OpenJSONFile,
	ReadClipboard,
	SaveGoSourceFile,
	WriteClipboard,
} from "../../wailsjs/go/main/App";
import IStandaloneCodeEditor = editor.IStandaloneCodeEditor;
import IStandaloneEditorConstructionOptions = editor.IStandaloneEditorConstructionOptions;

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

	if (locale == "zh") {
		return "zh-cn";
	}

	return "en";
}

function editorOptions(
	language: string,
	fontSize: number,
	value: string,
): IStandaloneEditorConstructionOptions {
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
	editor!.executeEdits("", [
		{
			range: editor!.getModel()!.getFullModelRange(),
			text: content,
		},
	]);
}

interface EditorsState {
	jsonEditor: IStandaloneCodeEditor | null;
	goEditor: IStandaloneCodeEditor | null;
	init: (
		language: string,
		fontSize: number,
		jsonID: string,
		goID: string,
	) => void;
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
		loader.init().then((monaco) => {
			const jsonElement = document.getElementById(jsonID);
			const goElement = document.getElementById(goID);

			jsonElement!.innerHTML = "";
			goElement!.innerHTML = "";
			monaco.editor.getModels().forEach((model) => model.dispose());

			const jsonEditor = monaco.editor.create(
				jsonElement!,
				editorOptions("json", fontSize, ""),
			);
			const goEditor = monaco.editor.create(
				goElement!,
				editorOptions("go", fontSize, ""),
			);

			set({
				jsonEditor: jsonEditor,
				goEditor: goEditor,
			});

			// remeasure fonts after creating editors and fonts are loaded to avoid rendering issues
			document.fonts.ready.then(() => {
				monaco.editor.remeasureFonts();
			});
		});
	},
	generate: async () => {
		get().goEditor?.setValue(
			await Generate(get().jsonEditor?.getValue() || ""),
		);
	},
	setJSON: (value) => {
		replaceContent(get().jsonEditor!, value);
	},
	pasteJSON: async () => {
		replaceContent(get().jsonEditor!, await ReadClipboard());
	},
	copyGo: async () => {
		void WriteClipboard(get().goEditor!.getValue());
	},
	openJSON: async () => {
		const content = await OpenJSONFile();
		if (content) {
			replaceContent(get().jsonEditor!, content);
		}
	},
	saveGo: async () => {
		void SaveGoSourceFile(get().goEditor!.getValue());
	},
}));
