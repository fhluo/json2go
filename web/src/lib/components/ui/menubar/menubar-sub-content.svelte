<script lang="ts">
import { cn, flyAndScale } from "$lib/utils.js";
import type { Menubar as MenubarPrimitive } from "bits-ui";

type $$Props = MenubarPrimitive.SubContentProps;
const className: $$Props["class"] = undefined;
export const transition: $$Props["transition"] = flyAndScale;
export const transitionConfig: $$Props["transitionConfig"] = {
	x: -10,
	y: 0,
};
export { className as class };
</script>

<MenubarPrimitive.SubContent
	{transition}
	{transitionConfig}
	class={cn(
		"z-50 min-w-max rounded-md border bg-popover p-1 text-popover-foreground shadow-lg focus:outline-none",
		className
	)}
	{...$$restProps}
>
	<slot />
</MenubarPrimitive.SubContent>
