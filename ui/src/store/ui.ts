import { Layout, useLayoutStore } from "@/store/layout";
import { View, useViewStore } from "@/store/view";

export function useEditorLayout() {
	const layout = useLayoutStore((s) => s.layout);
	const view = useViewStore((s) => s.view);

	return {
		layout,
		view,
		isColumns: layout === Layout.TwoColumns,
		isRows: layout === Layout.TwoRows,
		editorSpan: view !== View.JSONAndGo,
		hideJSONEditor: view === View.GoOnly,
		hideGoEditor: view === View.JSONOnly,
		goBorderLeft: view === View.JSONAndGo && layout === Layout.TwoColumns,
		goBorderTop: view === View.JSONAndGo && layout === Layout.TwoRows,
	};
}
