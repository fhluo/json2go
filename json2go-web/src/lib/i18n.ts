import i18n from "i18next";
import { initReactI18next } from "react-i18next";
import en from "../../locales/en.json";
import zh from "../../locales/zh.json";

const resources = {
	en: {
		translation: en,
	},
	zh: {
		translation: zh,
	},
};

const languages = Object.keys(resources);

void i18n.use(initReactI18next).init({
	resources,
	lng: "en",
	interpolation: {
		escapeValue: false,
	},
});

export { i18n, languages };
