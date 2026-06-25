import { Button } from "@/components/ui/button.tsx";
import { DialogClose, DialogContent, DialogFooter, DialogHeader, DialogTitle } from "@/components/ui/dialog.tsx";
import { openRelease, useVersionStore } from "@/lib/api.ts";
import { useEffect, useState } from "react";
import { useTranslation } from "react-i18next";

export default function Updates() {
    const { t } = useTranslation();

    const version = useVersionStore((s) => s.version);
    const latestVersion = useVersionStore((s) => s.latestVersion);
    const hasUpdate = useVersionStore((s) => s.hasUpdate);
    const checking = useVersionStore((s) => s.checking);
    const checkForUpdate = useVersionStore((s) => s.checkForUpdate);
    const [message, setMessage] = useState(t("updates.checking", "Checking for updates..."));

    useEffect(() => {
        void checkForUpdate();
    }, []);

    useEffect(() => {
        if (checking) {
            return;
        }

        if (!version || !latestVersion) {
            setMessage(t("updates.failed", "Unable to check for updates."));
            return;
        }

        if (hasUpdate) {
            setMessage(
                t("updates.available", "A new version is available: v{{version}}.", {
                    version: latestVersion,
                }),
            );
        } else {
            setMessage(t("updates.upToDate", "You're up to date."));
        }
    }, [checking, version, latestVersion, hasUpdate]);

    return (
        <DialogContent className="select-none">
            <DialogHeader>
                <DialogTitle>{t("updates.title", "Check for updates")}</DialogTitle>
            </DialogHeader>
            <p className="text-gray-900 leading-relaxed">{message}</p>
            <DialogFooter>
                {version && latestVersion && hasUpdate && (
                    <Button variant="secondary" onClick={() => openRelease(latestVersion)}>
                        {t("Download")}
                    </Button>
                )}
                <DialogClose render={<Button variant="secondary">{t("OK")}</Button>} />
            </DialogFooter>
        </DialogContent>
    );
}
