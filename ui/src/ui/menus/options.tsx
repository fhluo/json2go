import { MenubarCheckboxItem, MenubarContent, MenubarMenu, MenubarTrigger } from "@/components/ui/menubar.tsx";
import { useConfigStore } from "@/lib/api.ts";
import { useTranslation } from "react-i18next";

export default function Options() {
    const { t } = useTranslation();
    const validJSON = useConfigStore((s) => s.validJSON);
    const setValidJSON = useConfigStore((s) => s.setValidJSON);
    const realTime = useConfigStore((s) => s.realTime);
    const setRealTime = useConfigStore((s) => s.setRealTime);

    return (
        <MenubarMenu>
            <MenubarTrigger>{t("options.title", "Options")}</MenubarTrigger>
            <MenubarContent>
                <MenubarCheckboxItem checked={validJSON} onCheckedChange={(checked) => setValidJSON(checked)}>
                    {t("options.valid", "Validate JSON before generation")}
                </MenubarCheckboxItem>
                <MenubarCheckboxItem checked={realTime} onCheckedChange={(checked) => setRealTime(checked)}>
                    {t("options.realTime", "Generate in real time")}
                </MenubarCheckboxItem>
            </MenubarContent>
        </MenubarMenu>
    );
}
