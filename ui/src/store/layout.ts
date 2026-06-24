import { create } from "zustand";

export const Layout = {
	TwoColumns: "Two Columns",
	TwoRows: "Two Rows",
} as const;
export type Layout = (typeof Layout)[keyof typeof Layout];

interface LayoutState {
	layout: Layout;
	setLayout: (layout: Layout) => void;
}

export const useLayoutStore = create<LayoutState>((set) => ({
	layout: Layout.TwoColumns,
	setLayout: (layout) => set({ layout }),
}));
