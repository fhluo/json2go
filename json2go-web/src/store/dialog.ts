import { ReactNode } from "react";
import { create } from "zustand";

interface DialogState {
	dialog: ReactNode;
	setDialog: (dialog: ReactNode) => void;
}

export const useDialogStore = create<DialogState>((set) => ({
	dialog: null,
	setDialog: (dialog) =>
		set({
			dialog: dialog,
		}),
}));
