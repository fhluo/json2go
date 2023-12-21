import {Dialog} from "@/components/ui/dialog.tsx"
import {Menubar} from "@/components/ui/menubar.tsx"
import {PropsWithChildren, ReactNode} from "react"

interface Props {
    dialog: ReactNode
}

export default function ({children, dialog}: PropsWithChildren<Props>) {
    return (
        <Dialog>
            <Menubar className="rounded-none border-none bg-transparent">
                {children}
            </Menubar>
            {dialog}
        </Dialog>
    )
}
