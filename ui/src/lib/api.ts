import { Browser } from "@wailsio/runtime";
import i18n from "i18next";
import { create } from "zustand";
import {
	GetAllCapsWords,
	GetFontSize,
	GetLocale,
	GetOptionsGenerateInRealTime,
	GetOptionsValidJSONBeforeGeneration,
	SetAllCapsWords,
	SetFontSize,
	SetLocale,
	SetOptionsGenerateInRealTime,
	SetOptionsValidJSONBeforeGeneration,
} from "@api/app/services/config";

const defaultFontSize = 16;

function splitWords(words: string): string[] {
	return words
		.split(",")
		.map((word) => word.trim())
		.filter((word) => word !== "");
}

function unique(items: string[]): string[] {
	return Array.from(new Set(items));
}

interface ConfigState {
	language: string;
	fontSize: number;
	allCapsWords: string[];
	validJSON: boolean;
	realTime: boolean;
	init: () => Promise<void>;
	setLanguage: (language: string) => void;
	increaseFontSize: () => void;
	decreaseFontSize: () => void;
	resetFontSize: () => void;
	addAllCapsWords: (words: string) => void;
	removeAllCapsWords: (words: string) => void;
	setValidJSON: (validJSON: boolean) => void;
	setRealTime: (realTime: boolean) => void;
}

export const useConfigStore = create<ConfigState>((set) => ({
	language: "",
	fontSize: defaultFontSize,
	allCapsWords: [],
	validJSON: false,
	realTime: false,

	init: async () => {
		const [language, fontSize, allCapsWords, validJSON, realTime] =
			await Promise.all([
				GetLocale(),
				GetFontSize(),
				GetAllCapsWords(),
				GetOptionsValidJSONBeforeGeneration(),
				GetOptionsGenerateInRealTime(),
			]);
		void i18n.changeLanguage(language || "en");
		set({
			language: language || "en",
			fontSize,
			allCapsWords: allCapsWords || [],
			validJSON,
			realTime,
		});
	},

	setLanguage: (language) => {
		void i18n.changeLanguage(language);
		void SetLocale(language);
		set({ language });
	},

	increaseFontSize: () =>
		set((state) => {
			void SetFontSize(state.fontSize + 1);
			return { fontSize: state.fontSize + 1 };
		}),
	decreaseFontSize: () =>
		set((state) => {
			void SetFontSize(state.fontSize - 1);
			return { fontSize: state.fontSize - 1 };
		}),
	resetFontSize: () => {
		void SetFontSize(defaultFontSize);
		set({ fontSize: defaultFontSize });
	},

	addAllCapsWords: (words) =>
		set((state) => {
			const result = unique(
				state.allCapsWords.concat(splitWords(words)),
			);
			void SetAllCapsWords(result);
			return { allCapsWords: result };
		}),
	removeAllCapsWords: (words) =>
		set((state) => {
			const items = splitWords(words);
			const result = unique(
				state.allCapsWords.filter((word) => !items.includes(word)),
			);
			void SetAllCapsWords(result);
			return { allCapsWords: result };
		}),

	setValidJSON: (validJSON) => {
		void SetOptionsValidJSONBeforeGeneration(validJSON);
		set({ validJSON });
	},
	setRealTime: (realTime) => {
		void SetOptionsGenerateInRealTime(realTime);
		set({ realTime });
	},
}));

export function openHomePage() {
	void Browser.OpenURL("https://github.com/fhluo/json2go");
}

export function openRelease(version: string) {
	void Browser.OpenURL(
		`https://github.com/fhluo/json2go/releases/tag/v${version}`,
	);
}
