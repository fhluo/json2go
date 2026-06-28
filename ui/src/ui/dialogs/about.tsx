import { Button } from "@/components/ui/button.tsx";
import { DialogClose, DialogContent, DialogFooter, DialogHeader, DialogTitle } from "@/components/ui/dialog.tsx";
import { useEffect } from "react";
import { useVersionStore } from "@/lib/api.ts";
import { useTranslation } from "react-i18next";

export default function About() {
    const { t } = useTranslation();

    const version = useVersionStore((s) => s.version);
    const fetchVersion = useVersionStore((s) => s.fetchVersion);
    useEffect(() => {
        void fetchVersion();
    }, []);

    return (
        <DialogContent className="select-none">
            <DialogHeader>
                <DialogTitle>{t("about.title", "About")}</DialogTitle>
            </DialogHeader>
            <div className="flex flex-col items-center justify-center space-y-2 text-gray-900">
                <p className="text-lg font-semibold">json2go</p>
                <p className="text-sm leading-relaxed text-muted-foreground">
                    {t("about.description", "Generate Go type definitions from JSON")}
                </p>
                <div className="leading-relaxed text-sm">
                    <p>
                        <span className="select-none">{t("about.license", "License: ")}</span>
                        MIT
                    </p>
                    <p>
                        <span className="select-none">{t("about.version", "Version: ")}</span>
                        {version}
                    </p>
                </div>
                <p className="text-sm">Copyright © 2022-2026 fhluo</p>
            </div>
            <DialogFooter>
                <DialogClose render={<Button variant="secondary">{t("OK")}</Button>} />
            </DialogFooter>
        </DialogContent>
    );
}
