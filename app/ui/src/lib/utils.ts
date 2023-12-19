import {type ClassValue, clsx} from "clsx"
import {twMerge} from "tailwind-merge"
import {BrowserOpenURL} from "../../wailsjs/runtime"

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

export function OpenRelease(version: string) {
    BrowserOpenURL(`https://github.com/fhluo/json2go/releases/tag/v${version}`)
}
