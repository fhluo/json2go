<script lang="ts">
    import { getContext } from "svelte";
    import { Layout } from "../state/ui.svelte";
    import type { UIState } from "../state/ui.svelte";

    const ui = getContext<UIState>("ui");
</script>

<div class={`grid h-64 grow border-t border-b ${ui.layout === Layout.TwoColumns ? 'grid-cols-2' : ''} ${ui.layout === Layout.TwoRows ? 'grid-rows-2' : ''}`}>
    <slot></slot>
</div>
