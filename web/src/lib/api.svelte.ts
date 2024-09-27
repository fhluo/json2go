import {Config, Version} from "@api/app/services";
import {Browser} from "@wailsio/runtime";

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
			locale = await Config.GetLocale();
			fontSize = await Config.GetFontSize();
			allCapsWords = await Config.GetAllCapsWords();
			validJSON = await Config.GetOptionsValidJSONBeforeGeneration();
			realTime = await Config.GetOptionsGenerateInRealTime();
			version = await Version.GetVersion();
		},
		get locale() {
			return locale;
		},
		set locale(value) {
			locale = value;
			void Config.SetLocale(locale);
		},
		get fontSize() {
			return fontSize;
		},
		set fontSize(value) {
			fontSize = value;
			void Config.SetFontSize(fontSize);
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
			void Config.SetAllCapsWords(allCapsWords);
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
			void Config.SetOptionsValidJSONBeforeGeneration(validJSON);
		},
		get realTime() {
			return realTime;
		},
		set realTime(value) {
			realTime = value;
			void Config.SetOptionsGenerateInRealTime(realTime);
		},
		get version() {
			return version;
		},
	};
}

export type ConfigState = ReturnType<typeof createConfigState>;

export function openHomePage() {
	void Browser.OpenURL("https://github.com/fhluo/json2go");
}

export function openRelease(version: string) {
	void Browser.OpenURL(`https://github.com/fhluo/json2go/releases/tag/v${version}`);
}
