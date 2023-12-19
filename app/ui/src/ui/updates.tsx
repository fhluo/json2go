import {DialogClose, DialogContent, DialogFooter, DialogHeader, DialogTitle,} from "@/components/ui/dialog.tsx"
import {useEffect, useState} from "react"
import {GetLatestVersion, GetVersion} from "../../wailsjs/go/main/App"
import {Button} from "@/components/ui/button.tsx"
import {useTranslation} from "react-i18next"
import {OpenRelease} from "@/lib/utils.ts"

export default function () {
    const {t} = useTranslation()

    const [version, setVersion] = useState("")
    const [latestVersion, setLatestVersion] = useState("")
    const [message, setMessage] = useState(t("updates.checking", "Checking..."))

    useEffect(() => {
        Promise.all([GetVersion(), GetLatestVersion()]).then(
            ([version, latestVersion]) => {
                setVersion(version)
                setLatestVersion(latestVersion)
            }
        )

        if (!version || !latestVersion) {
            setMessage(t("updates.failed", "Unable to check for updates."))
            return
        }

        if (version === latestVersion) {
            setMessage(t("updates.none", "You are using the latest version."))
        } else {
            setMessage(`${t("updates.available", "A new version is available: ")}v${latestVersion}.`)
        }
    }, [])

    return (
        <DialogContent className="select-none">
            <DialogHeader>
                <DialogTitle>{t("updates.title", "Check for updates")}</DialogTitle>
            </DialogHeader>
            <p className="text-gray-900 leading-relaxed">{message}</p>
            <DialogFooter>
                {version && latestVersion && version !== latestVersion &&
                  <Button variant="secondary" onClick={() => OpenRelease(latestVersion)}>
                      {t("Download")}
                  </Button>
                }
                <DialogClose asChild>
                    <Button variant="secondary">{t("OK")}</Button>
                </DialogClose>
            </DialogFooter>
        </DialogContent>
    )
}
