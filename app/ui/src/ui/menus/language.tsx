import {
    MenubarContent,
    MenubarMenu,
    MenubarRadioGroup,
    MenubarRadioItem,
    MenubarTrigger
} from "@/components/ui/menubar.tsx"
import {useTranslation} from "react-i18next"
import {useEffect, useState} from "react"
import {SetLocale} from "../../../wailsjs/go/main/App"
import {languages} from "@/lib/i18n.ts"

export default function () {
    const {t, i18n} = useTranslation()

    let [language, setLanguage] = useState(i18n.language)
    useEffect(() => {
        void i18n.changeLanguage(language)
        void SetLocale(language)
    }, [language])

    return (
        <MenubarMenu>
            <MenubarTrigger>
                {t("language.title", "Language")}
            </MenubarTrigger>
            <MenubarContent>
                <MenubarRadioGroup value={language} onValueChange={value => setLanguage(value)}>
                    {languages.map(language =>
                        <MenubarRadioItem value={language}>
                            {t(`language.${language}`)}
                        </MenubarRadioItem>
                    )}
                </MenubarRadioGroup>
            </MenubarContent>
        </MenubarMenu>
    )
}
