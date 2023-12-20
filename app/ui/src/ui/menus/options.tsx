import {MenubarCheckboxItem, MenubarContent, MenubarMenu, MenubarTrigger} from "@/components/ui/menubar.tsx"
import {useTranslation} from "react-i18next"

interface Props {
    validJSON: boolean
    setValidJSON: (v: boolean) => void

    realTime: boolean
    setRealTime: (v: boolean) => void
}

export default function ({validJSON, setValidJSON, realTime, setRealTime}: Props) {
    const {t} = useTranslation()

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
