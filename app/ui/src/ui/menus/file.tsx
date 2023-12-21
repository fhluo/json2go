import {MenubarContent, MenubarItem, MenubarMenu, MenubarSeparator, MenubarTrigger} from "@/components/ui/menubar.tsx"
import {DialogTrigger} from "@/components/ui/dialog.tsx"
import Settings from "@/ui/dialogs/settings.tsx"
import {EventsEmit} from "../../../wailsjs/runtime"
import {useTranslation} from "react-i18next"
import {ReactNode} from "react"

interface Props {
    setDialog: (element: ReactNode) => void
    openJSON: () => void
    saveGo: () => void
}

export default function ({setDialog, openJSON, saveGo}: Props) {
    const {t} = useTranslation()

    return (
        <MenubarMenu>
            <MenubarTrigger>
                {t("file.title", "File")}
            </MenubarTrigger>
            <MenubarContent>
                <MenubarItem onClick={openJSON}>
                    {t("file.openJSON", "Open JSON file")}
                </MenubarItem>
                <MenubarItem onClick={saveGo}>
                    {t("file.saveGo", "Save Go source file")}
                </MenubarItem>
                <MenubarSeparator/>
                <DialogTrigger asChild>
                    <MenubarItem onClick={() => setDialog(<Settings/>)}>
                        {t("file.settings", "Settings")}
                    </MenubarItem>
                </DialogTrigger>
                <MenubarSeparator/>
                <MenubarItem onClick={() => EventsEmit("exit")}>
                    {t("file.exit", "Exit")}
                </MenubarItem>
            </MenubarContent>
        </MenubarMenu>
    )
}
