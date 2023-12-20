import {Layout, View} from "@/lib/types.ts"
import {cn} from "@/lib/utils.ts"
import {useTranslation} from "react-i18next"

interface Props {
    view: View
    layout: Layout
    copyGoCode: () => void
}

export default function ({view, layout, copyGoCode}: Props) {
    const {t} = useTranslation()

    return (
        <div id="container-go"
             style={{display: view === View.JSONOnly ? "none" : ""}}
             className={cn({
                 "col-span-2": layout === Layout.TwoColumns && view === View.GoOnly,
                 "row-span-2": layout === Layout.TwoRows && view === View.GoOnly,
                 "border-l": layout === Layout.TwoColumns && view === View.JSONAndGo,
                 "border-t": layout === Layout.TwoRows && view === View.JSONAndGo,
             })}>
            <div className="w-full bg-white/50 flex flex-row">
                <span className="py-1 px-4 select-none text-purple-700 font-mono">Go</span>
                <button onClick={copyGoCode}>{t("Copy")}</button>
            </div>
            <div className="w-full h-32 grow" id="go-editor"></div>
        </div>
    )
}