import {MenubarContent, MenubarItem, MenubarMenu, MenubarSeparator, MenubarTrigger} from "@/components/ui/menubar.tsx"
import {DialogTrigger} from "@/components/ui/dialog.tsx"
import Updates from "@/ui/dialogs/updates.tsx"
import About from "@/ui/dialogs/about.tsx"
import {useTranslation} from "react-i18next"
import {ReactNode} from "react"

interface Props {
    setDialog: (element: ReactNode) => void
}

export default function ({setDialog}: Props) {
    const {t} = useTranslation()

    return (
        <MenubarMenu>
            <MenubarTrigger>
                {t("help.title", "Help")}
            </MenubarTrigger>
            <MenubarContent>
                <DialogTrigger asChild>
                    <MenubarItem onClick={() => setDialog(<Updates/>)}>
                        {t("help.updates", "Check for updates")}
                    </MenubarItem>
                </DialogTrigger>
                <MenubarSeparator/>
                <DialogTrigger asChild>
                    <MenubarItem onClick={() => setDialog(<About/>)}>
                        {t("help.about", "About")}
                    </MenubarItem>
                </DialogTrigger>
            </MenubarContent>
        </MenubarMenu>
    )
}
