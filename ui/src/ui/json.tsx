import { cn } from "@/lib/utils.ts";
import { useEditorsStore } from "@/store/editors.ts";
import { useEditorLayout } from "@/store/ui.ts";
import { useTranslation } from "react-i18next";

export default function JSON() {
    const { t } = useTranslation();
    const { hideJSONEditor, editorSpan, isColumns, isRows } = useEditorLayout();
    const pasteJSON = useEditorsStore((state) => state.pasteJSON);

    return (
        <div
            id="container-json"
            style={{ display: hideJSONEditor ? "none" : "" }}
            className={cn("group", {
                "col-span-2": isColumns && editorSpan,
                "row-span-2": isRows && editorSpan,
            })}
        >
            <div className="w-full bg-white/50 flex flex-row">
                <span className="py-1 px-4 select-none text-yellow-700 font-mono group-focus-within:text-yellow-500">
                    JSON
                </span>
                <button type={"button"} onClick={pasteJSON}>
                    {t("Paste")}
                </button>
            </div>
            <div className="w-full h-32 grow" id="json-editor" />
        </div>
    );
}
