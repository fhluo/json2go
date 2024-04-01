import { cn } from "@/lib/utils.ts";
import { Layout, useLayoutStore } from "@/store/layout.ts";
import type { PropsWithChildren } from "react";

export default function ({ children }: PropsWithChildren) {
	const layout = useLayoutStore((state) => state.layout);

	return (
		<div
			className={cn("grid h-64 grow border-t border-b", {
				"grid-cols-2": layout === Layout.TwoColumns,
				"grid-rows-2": layout === Layout.TwoRows,
			})}
		>
			{children}
		</div>
	);
}
