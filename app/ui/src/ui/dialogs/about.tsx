import {DialogClose, DialogContent, DialogFooter, DialogHeader, DialogTitle,} from "@/components/ui/dialog.tsx"
import {Button} from "@/components/ui/button.tsx"
import {useEffect, useState} from "react"
import {GetVersion} from "../../../wailsjs/go/main/App"
import {useTranslation} from "react-i18next"

export default function () {
    const {t} = useTranslation()

    const [version, setVersion] = useState("")
    useEffect(() => {
        GetVersion().then(value => setVersion(value))
    }, [])

    return (
        <DialogContent className="select-none">
            <DialogHeader>
                <DialogTitle>{t("about.title", "About")}</DialogTitle>
            </DialogHeader>
            <div className="flex flex-col space-y-2 items-center justify-center text-gray-900">
                <p className="text-lg font-semibold">json2go</p>
                <p className="leading-relaxed">{t("about.description", "Generate Go type definitions from JSON")}</p>
                <div className="leading-relaxed">
                    <p><span className="font-semibold select-none">{t("about.license", "License: ")}</span>MIT</p>
                    <p><span className="font-semibold select-none">{t("about.version", "Version: ")}</span>{version}</p>
                </div>
                <p>Copyright © 2022 fhluo</p>
            </div>
            <DialogFooter>
                <DialogClose asChild>
                    <Button variant="secondary">{t("OK")}</Button>
                </DialogClose>
            </DialogFooter>
        </DialogContent>
    )
}
