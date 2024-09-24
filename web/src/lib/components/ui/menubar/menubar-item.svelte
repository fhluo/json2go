<script lang="ts">
import { cn } from "$lib/utils.js";
import type { Menubar as MenubarPrimitive } from "bits-ui";

type $$Props = MenubarPrimitive.ItemProps & {
	inset?: boolean;
};
type $$Events = MenubarPrimitive.ItemEvents;

const className: $$Props["class"] = undefined;
export const inset: $$Props["inset"] = undefined;
export { className as class };
</script>

<MenubarPrimitive.Item
	class={cn(
		"relative flex cursor-default select-none items-center rounded-sm px-2 py-1.5 text-sm outline-none data-[disabled]:pointer-events-none data-[highlighted]:bg-accent data-[highlighted]:text-accent-foreground data-[disabled]:opacity-50",
		inset && "pl-8",
		className
	)}
	{...$$restProps}
	on:click
	on:keydown
	on:focusin
	on:focusout
	on:pointerleave
	on:pointermove
	on:pointerdown
>
	<slot />
</MenubarPrimitive.Item>
