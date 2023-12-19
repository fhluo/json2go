import {MenubarContent, MenubarItem, MenubarMenu, MenubarTrigger} from "@/components/ui/menubar.tsx"
import {useTranslation} from "react-i18next"

interface Props {
    increaseFontSize: () => void
    decreaseFontSize: () => void
    resetFontSize: () => void
}

export default function ({increaseFontSize, decreaseFontSize, resetFontSize}: Props) {
    const {t} = useTranslation()

    return (
        <MenubarMenu>
            <MenubarTrigger>
                {t("font.title", "Font")}
            </MenubarTrigger>
            <MenubarContent>
                <MenubarItem onClick={increaseFontSize}>
                    {t("font.increase", "Increase size")}
                </MenubarItem>
                <MenubarItem onClick={decreaseFontSize}>
                    {t("font.decrease", "Decrease size")}
                </MenubarItem>
                <MenubarItem onClick={resetFontSize}>
                    {t("font.reset", "Reset size")}
                </MenubarItem>
            </MenubarContent>
        </MenubarMenu>
    )
}
