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
} from "../../wailsjs/go/main/App";

interface LanguageState {
	language: string;
	init: () => void;
	setLanguage: (language: string) => void;
}

export const useLanguageStore = create<LanguageState>((set) => ({
	language: "",
	init: async () => {
		const language = (await GetLocale()) || "en";
		void i18n.changeLanguage(language);
		set({
			language: language,
		});
	},
	setLanguage: (language) => {
		void i18n.changeLanguage(language);
		void SetLocale(language);
		set({
			language: language,
		});
	},
}));

interface FontSizeState {
	fontSize: number;
	init: () => void;
	increase: () => void;
	decrease: () => void;
	reset: () => void;
}

const defaultFontSize = 16;

export const useFontSizeStore = create<FontSizeState>((set) => ({
	fontSize: defaultFontSize,
	init: async () =>
		set({
			fontSize: await GetFontSize(),
		}),
	increase: () =>
		set((state) => {
			void SetFontSize(state.fontSize + 1);
			return {
				fontSize: state.fontSize + 1,
			};
		}),
	decrease: () =>
		set((state) => {
			void SetFontSize(state.fontSize - 1);
			return {
				fontSize: state.fontSize - 1,
			};
		}),
	reset: () => {
		void SetFontSize(defaultFontSize);
		set({
			fontSize: defaultFontSize,
		});
	},
}));

interface AllCapsWordsState {
	words: string[];
	init: () => void;
	add: (words: string) => void;
	remove: (words: string) => void;
}

function splitWords(words: string): string[] {
	return words
		.split(",")
		.map((word) => word.trim())
		.filter((word) => word !== "");
}

function unique(items: string[]): string[] {
	return Array.from(new Set(items));
}

export const useAllCapsWordsStore = create<AllCapsWordsState>((set) => ({
	words: [],
	init: async () =>
		set({
			words: (await GetAllCapsWords()) || [],
		}),
	add: (words) =>
		set((state) => {
			const result = unique(state.words.concat(splitWords(words)));
			void SetAllCapsWords(result);
			return {
				words: result,
			};
		}),
	remove: (words) =>
		set((state) => {
			const items = splitWords(words);
			const result = unique(
				state.words.filter((word) => !items.includes(word)),
			);
			void SetAllCapsWords(result);
			return {
				words: result,
			};
		}),
}));

interface ValidJSONState {
	validJSON: boolean;
	init: () => void;
	setValidJSON: (validJSON: boolean) => void;
}

export const useValidJSONStore = create<ValidJSONState>((set) => ({
	validJSON: false,
	init: async () =>
		set({
			validJSON: await GetOptionsValidJSONBeforeGeneration(),
		}),
	setValidJSON: (validJSON) => {
		void SetOptionsValidJSONBeforeGeneration(validJSON);
		set({
			validJSON: validJSON,
		});
	},
}));

interface RealTimeState {
	realTime: boolean;
	init: () => void;
	setRealTime: (realTime: boolean) => void;
}

export const useRealTimeStore = create<RealTimeState>((set) => ({
	realTime: false,
	init: async () =>
		set({
			realTime: await GetOptionsGenerateInRealTime(),
		}),
	setRealTime: (realTime) => {
		void SetOptionsGenerateInRealTime(realTime);
		set({
			realTime: realTime,
		});
	},
}));
