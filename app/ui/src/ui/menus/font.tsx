import {MenubarContent, MenubarItem, MenubarMenu, MenubarTrigger} from "@/components/ui/menubar.tsx"
import {useTranslation} from "react-i18next"
import {useFontSizeStore} from "@/lib/store.ts"

export default function () {
    const {t} = useTranslation()
    const {increase, decrease, reset} = useFontSizeStore()

    return (
        <MenubarMenu>
            <MenubarTrigger>
                {t("font.title", "Font")}
            </MenubarTrigger>
            <MenubarContent>
                <MenubarItem onClick={increase}>
                    {t("font.increase", "Increase size")}
                </MenubarItem>
                <MenubarItem onClick={decrease}>
                    {t("font.decrease", "Decrease size")}
                </MenubarItem>
                <MenubarItem onClick={reset}>
                    {t("font.reset", "Reset size")}
                </MenubarItem>
            </MenubarContent>
        </MenubarMenu>
    )
}
