import { create } from "zustand";

export enum Layout {
	TwoColumns = "Two Columns",
	TwoRows = "Two Rows",
}

interface LayoutState {
	layout: Layout;
	setLayout: (layout: Layout) => void;
}

export const useLayoutStore = create<LayoutState>((set) => ({
	layout: Layout.TwoColumns,
	setLayout: (layout) =>
		set({
			layout: layout,
		}),
}));
