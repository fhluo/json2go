import {Dialog} from "@/components/ui/dialog.tsx"
import {Menubar} from "@/components/ui/menubar.tsx"
import {PropsWithChildren} from "react"
import {useDialogStore} from "@/store/dialog.ts"

interface Props {
}

export default function ({children}: PropsWithChildren<Props>) {
    const dialog = useDialogStore(state => state.dialog)

    return (
        <Dialog>
            <Menubar className="rounded-none border-none bg-transparent">
                {children}
            </Menubar>
            {dialog}
        </Dialog>
    )
}
