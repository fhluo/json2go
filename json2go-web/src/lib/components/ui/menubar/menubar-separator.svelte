<script lang="ts">
import { cn } from "$lib/utils.js";
import type { Menubar as MenubarPrimitive } from "bits-ui";

type $$Props = MenubarPrimitive.SeparatorProps;
const className: $$Props["class"] = undefined;
export { className as class };
</script>

<MenubarPrimitive.Separator class={cn("-mx-1 my-1 h-px bg-muted", className)} {...$$restProps} />
