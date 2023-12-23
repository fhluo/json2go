import {cn} from "@/lib/utils.ts"
import {useTranslation} from "react-i18next"
import {Layout, useLayoutStore} from "@/store/layout.ts"
import {useViewStore, View} from "@/store/view.ts"
import {useEditorsStore} from "@/store/editors.ts"

export default function () {
    const {t} = useTranslation()
    const layout = useLayoutStore(state => state.layout)
    const view = useViewStore(state => state.view)
    const copyGo = useEditorsStore(state => state.copyGo)

    return (
        <div id="container-go"
             style={{display: view === View.JSONOnly ? "none" : ""}}
             className={cn("group", {
                 "col-span-2": layout === Layout.TwoColumns && view === View.GoOnly,
                 "row-span-2": layout === Layout.TwoRows && view === View.GoOnly,
                 "border-l": layout === Layout.TwoColumns && view === View.JSONAndGo,
                 "border-t": layout === Layout.TwoRows && view === View.JSONAndGo,
             })}>
            <div className="w-full bg-white/50 flex flex-row">
                <span className="py-1 px-4 select-none text-purple-700 font-mono group-focus-within:text-purple-500">
                    Go
                </span>
                <button onClick={copyGo}>{t("Copy")}</button>
            </div>
            <div className="w-full h-32 grow" id="go-editor"></div>
        </div>
    )
}
