import { create } from "zustand";

interface MessageState {
	message: string;
	setMessage: (message: string) => void;
	clear: () => void;
}

export const useMessageStore = create<MessageState>((set) => ({
	message: "",
	setMessage: (message) =>
		set({
			message: message,
		}),
	clear: () =>
		set({
			message: "",
		}),
}));
