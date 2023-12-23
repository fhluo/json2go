import {MenubarCheckboxItem, MenubarContent, MenubarMenu, MenubarTrigger} from "@/components/ui/menubar.tsx"
import {useTranslation} from "react-i18next"
import {useRealTimeStore, useValidJSONStore} from "@/lib/store.ts"

export default function () {
    const {t} = useTranslation()
    const {validJSON, setValidJSON} = useValidJSONStore()
    const {realTime, setRealTime} = useRealTimeStore()

    return (
        <MenubarMenu>
            <MenubarTrigger>
                {t("options.title", "Options")}
            </MenubarTrigger>
            <MenubarContent>
                <MenubarCheckboxItem checked={validJSON} onCheckedChange={checked => setValidJSON(checked)}>
                    {t("options.valid", "Validate JSON before generation")}
                </MenubarCheckboxItem>
                <MenubarCheckboxItem checked={realTime} onCheckedChange={checked => setRealTime(checked)}>
                    {t("options.realTime", "Generate in real time")}
                </MenubarCheckboxItem>
            </MenubarContent>
        </MenubarMenu>
    )
}
