import {type ClassValue, clsx} from "clsx"
import {twMerge} from "tailwind-merge"
import {BrowserOpenURL} from "../../wailsjs/runtime"

export function cn(...inputs: ClassValue[]) {
    return twMerge(clsx(inputs))
}

export function openHomePage() {
    BrowserOpenURL("https://github.com/fhluo/json2go")
}

export function openRelease(version: string) {
    BrowserOpenURL(`https://github.com/fhluo/json2go/releases/tag/v${version}`)
}
