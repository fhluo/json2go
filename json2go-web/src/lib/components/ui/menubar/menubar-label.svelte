<script lang="ts">
import { cn } from "$lib/utils.js";
import type { Menubar as MenubarPrimitive } from "bits-ui";

type $$Props = MenubarPrimitive.LabelProps & {
	inset?: boolean;
};
const className: $$Props["class"] = undefined;
export const inset: $$Props["inset"] = undefined;
export { className as class };
</script>

<MenubarPrimitive.Label
	class={cn("px-2 py-1.5 text-sm font-semibold", inset && "pl-8", className)}
	{...$$restProps}
>
	<slot />
</MenubarPrimitive.Label>
