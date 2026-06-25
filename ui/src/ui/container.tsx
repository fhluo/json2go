import { cn } from "@/lib/utils.ts";
import { useEditorLayout } from "@/store/ui.ts";
import type { PropsWithChildren } from "react";

export default function Container({ children }: PropsWithChildren) {
	const { isColumns, isRows } = useEditorLayout();

	return (
		<div
			className={cn("grid h-64 grow border-t border-b", {
				"grid-cols-2": isColumns,
				"grid-rows-2": isRows,
			})}
		>
			{children}
		</div>
	);
}
