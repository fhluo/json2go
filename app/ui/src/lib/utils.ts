import {type ClassValue, clsx} from "clsx"
import {twMerge} from "tailwind-merge"
import {BrowserOpenURL} from "../../wailsjs/runtime"
import {editor} from "monaco-editor"
import IStandaloneCodeEditor = editor.IStandaloneCodeEditor

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

export function openHomePage() {
    BrowserOpenURL("https://github.com/fhluo/json2go")
}

export function openRelease(version: string) {
    BrowserOpenURL(`https://github.com/fhluo/json2go/releases/tag/v${version}`)
}

export function replaceContent(editor: IStandaloneCodeEditor, content: string) {
    editor!.executeEdits("", [{
        range: editor!.getModel()!.getFullModelRange(),
        text: content,
    }])
}
