import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"
import {Browser} from "@wailsio/runtime";

export function cn(...inputs: ClassValue[]) {
    return twMerge(clsx(inputs))
}

export function openHomePage() {
    void Browser.OpenURL("https://github.com/fhluo/json2go");
}

export function openRelease(version: string) {
    void Browser.OpenURL(`https://github.com/fhluo/json2go/releases/tag/v${version}`);
}
