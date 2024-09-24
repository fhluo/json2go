<script lang="ts">
import { cn, flyAndScale } from "$lib/utils.js";
import type { Menubar as MenubarPrimitive } from "bits-ui";

type $$Props = MenubarPrimitive.ContentProps;
const className: $$Props["class"] = undefined;
export const align: $$Props["align"] = "start";
export const alignOffset: $$Props["alignOffset"] = -4;
export const sideOffset: $$Props["sideOffset"] = 8;
export const transition: $$Props["transition"] = flyAndScale;
export const transitionConfig: $$Props["transitionConfig"] = undefined;
export { className as class };
</script>

<MenubarPrimitive.Content
	{transition}
	{transitionConfig}
	{sideOffset}
	{align}
	{alignOffset}
	class={cn(
		"z-50 min-w-[12rem] rounded-md border bg-popover p-1 text-popover-foreground shadow-md focus:outline-none",
		className
	)}
	{...$$restProps}
>
	<slot />
</MenubarPrimitive.Content>
