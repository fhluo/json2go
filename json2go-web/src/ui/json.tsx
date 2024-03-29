import { cn } from "@/lib/utils.ts";
import { useEditorsStore } from "@/store/editors.ts";
import { Layout, useLayoutStore } from "@/store/layout.ts";
import { View, useViewStore } from "@/store/view.ts";
import { useTranslation } from "react-i18next";

export default function () {
	const { t } = useTranslation();
	const layout = useLayoutStore((state) => state.layout);
	const view = useViewStore((state) => state.view);
	const pasteJSON = useEditorsStore((state) => state.pasteJSON);

	return (
		<div
			id="container-json"
			style={{ display: view === View.GoOnly ? "none" : "" }}
			className={cn("group", {
				"col-span-2": layout === Layout.TwoColumns && view === View.JSONOnly,
				"row-span-2": layout === Layout.TwoRows && view === View.JSONOnly,
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
