import {
	GetAllCapsWords,
	GetFontSize,
	GetLocale,
	GetOptionsGenerateInRealTime,
	GetOptionsValidJSONBeforeGeneration,
	GetVersion,
	SetAllCapsWords,
	SetFontSize,
	SetLocale,
	SetOptionsGenerateInRealTime,
	SetOptionsValidJSONBeforeGeneration,
} from "../../wailsjs/go/main/App";
import { BrowserOpenURL } from "../../wailsjs/runtime";

function splitWords(words: string): string[] {
	return words
		.split(",")
		.map((word) => word.trim())
		.filter((word) => word !== "");
}

export function createConfigState() {
	const defaultFontSize = 16;

	let locale = $state("en");
	let fontSize = $state(defaultFontSize);
	let allCapsWords = $state([] as string[]);
	let validJSON = $state(false);
	let realTime = $state(false);
	let version = $state("");

	return {
		async init() {
			locale = await GetLocale();
			fontSize = await GetFontSize();
			allCapsWords = await GetAllCapsWords();
			validJSON = await GetOptionsValidJSONBeforeGeneration();
			realTime = await GetOptionsGenerateInRealTime();
			version = await GetVersion();
		},
		get locale() {
			return locale;
		},
		set locale(value) {
			locale = value;
			void SetLocale(locale);
		},
		get fontSize() {
			return fontSize;
		},
		set fontSize(value) {
			fontSize = value;
			void SetFontSize(fontSize);
		},
		increaseFontSize() {
			this.fontSize += 1;
		},
		decreaseFontSize() {
			this.fontSize -= 1;
		},
		resetFontSize() {
			this.fontSize = defaultFontSize;
		},
		get allCapsWords() {
			return allCapsWords;
		},
		set allCapsWords(value) {
			allCapsWords = value;
			void SetAllCapsWords(allCapsWords);
		},
		addAllCapsWords(words: string) {
			this.allCapsWords = Array.from(
				new Set(allCapsWords.concat(splitWords(words))),
			);
		},
		removeAllCapsWords(words: string) {
			const wordList = splitWords(words);
			this.allCapsWords = Array.from(
				new Set(allCapsWords.filter((word) => !wordList.includes(word))),
			);
		},
		get validJSON() {
			return validJSON;
		},
		set validJSON(value) {
			validJSON = value;
			void SetOptionsValidJSONBeforeGeneration(validJSON);
		},
		get realTime() {
			return realTime;
		},
		set realTime(value) {
			realTime = value;
			void SetOptionsGenerateInRealTime(realTime);
		},
		get version() {
			return version;
		},
	};
}

export type ConfigState = ReturnType<typeof createConfigState>;

export function openHomePage() {
	BrowserOpenURL("https://github.com/fhluo/json2go");
}

export function openRelease(version: string) {
	BrowserOpenURL(`https://github.com/fhluo/json2go/releases/tag/v${version}`);
}
