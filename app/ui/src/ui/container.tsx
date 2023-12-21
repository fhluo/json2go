import {cn} from "@/lib/utils.ts"
import {Layout} from "@/lib/types.ts"
import {PropsWithChildren} from "react"

interface Props {
    layout: Layout,
}

export default function ({layout, children}: PropsWithChildren<Props>) {
    return (
        <div className={cn("grid h-64 grow border-t border-b", {
            "grid-cols-2": layout === Layout.TwoColumns,
            "grid-rows-2": layout === Layout.TwoRows,
        })}>
            {children}
        </div>
    )
}
