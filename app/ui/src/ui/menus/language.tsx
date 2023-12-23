import {
    MenubarContent,
    MenubarMenu,
    MenubarRadioGroup,
    MenubarRadioItem,
    MenubarTrigger
} from "@/components/ui/menubar.tsx"
import {useTranslation} from "react-i18next"
import {languages} from "@/lib/i18n.ts"
import {useLanguageStore} from "@/lib/store.ts"

export default function () {
    const {t} = useTranslation()
    let {language, setLanguage} = useLanguageStore()

    return (
        <MenubarMenu>
            <MenubarTrigger>
                {t("language.title", "Language")}
            </MenubarTrigger>
            <MenubarContent>
                <MenubarRadioGroup value={language} onValueChange={value => setLanguage(value)}>
                    {languages.map(language =>
                        <MenubarRadioItem key={language} value={language}>
                            {t(`language.${language}`)}
                        </MenubarRadioItem>
                    )}
                </MenubarRadioGroup>
            </MenubarContent>
        </MenubarMenu>
    )
}
