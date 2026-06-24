import { create } from "zustand";

export const View = {
	JSONAndGo: "JSON and Go",
	JSONOnly: "JSON Only",
	GoOnly: "Go Only",
} as const;
export type View = (typeof View)[keyof typeof View];

interface ViewState {
	view: View;
	setView: (view: View) => void;
}

export const useViewStore = create<ViewState>((set) => ({
	view: View.JSONAndGo,
	setView: (view) => set({ view }),
}));
