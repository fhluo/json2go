import {MenubarContent, MenubarItem, MenubarMenu, MenubarSeparator, MenubarTrigger} from "@/components/ui/menubar.tsx"
import {DialogTrigger} from "@/components/ui/dialog.tsx"
import Settings from "@/ui/dialogs/settings.tsx"
import {EventsEmit} from "../../../wailsjs/runtime"
import {useTranslation} from "react-i18next"
import {useEditorsStore} from "@/store/editors.ts"
import {useDialogStore} from "@/store/dialog.ts"

export default function () {
    const {t} = useTranslation()
    const setDialog = useDialogStore(state => state.setDialog)
    const {openJSON, saveGo} = useEditorsStore()

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
