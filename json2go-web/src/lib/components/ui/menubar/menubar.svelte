<script lang="ts">
import { cn } from "$lib/utils.js";
import type { Menubar as MenubarPrimitive } from "bits-ui";

type $$Props = MenubarPrimitive.Props;

const className: $$Props["class"] = undefined;
export { className as class };
</script>

<MenubarPrimitive.Root
	class={cn(
		"flex h-9 items-center space-x-1 rounded-md border bg-background p-1 shadow-sm",
		className
	)}
>
	<slot />
</MenubarPrimitive.Root>
