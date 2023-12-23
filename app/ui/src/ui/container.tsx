import {cn} from "@/lib/utils.ts"
import {PropsWithChildren} from "react"
import {Layout, useLayoutStore} from "@/store/layout.ts"

interface Props {
}

export default function ({children}: PropsWithChildren<Props>) {
    const layout = useLayoutStore(state => state.layout)

    return (
        <div className={cn("grid h-64 grow border-t border-b", {
            "grid-cols-2": layout === Layout.TwoColumns,
            "grid-rows-2": layout === Layout.TwoRows,
        })}>
            {children}
        </div>
    )
}
