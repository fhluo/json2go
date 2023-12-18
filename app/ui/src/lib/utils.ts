import {type ClassValue, clsx} from "clsx"
import {twMerge} from "tailwind-merge"

export function cn(...inputs: ClassValue[]) {
    return twMerge(clsx(inputs))
}

const monacoLocales = ["en", "de", "es", "fr", "it", "ja", "ko", "ru", "zh-cn", "zh-tw"]

export function getMonacoLocale(locale: string): string {
    locale = locale.toLowerCase()
    if (monacoLocales.includes(locale)) {
        return locale
    }

    if (locale == "zh") {
        return "zh-cn"
    }

    return "en"
}
