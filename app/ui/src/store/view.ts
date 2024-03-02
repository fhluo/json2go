import { create } from "zustand";

export enum View {
	JSONAndGo = "JSON and Go",
	JSONOnly = "JSON Only",
	GoOnly = "Go Only",
}

interface ViewState {
	view: View;
	setView: (view: View) => void;
}

export const useViewStore = create<ViewState>((set) => ({
	view: View.JSONAndGo,
	setView: (view) =>
		set({
			view: view,
		}),
}));
