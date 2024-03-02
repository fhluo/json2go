import { Dialog } from "@/components/ui/dialog.tsx";
import { Menubar } from "@/components/ui/menubar.tsx";
import { useDialogStore } from "@/store/dialog.ts";
import { PropsWithChildren } from "react";

type Props = {};

export default function ({ children }: PropsWithChildren<Props>) {
	const dialog = useDialogStore((state) => state.dialog);

	return (
		<Dialog>
			<Menubar className="rounded-none border-none bg-transparent">
				{children}
			</Menubar>
			{dialog}
		</Dialog>
	);
}
