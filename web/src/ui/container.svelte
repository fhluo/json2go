<script lang="ts">
    import {getContext} from "svelte";
    import type {UIState} from "../state/ui.svelte";

    let {children} = $props();

    const ui = getContext<UIState>("ui");
</script>

<div class="grid h-64 grow border-t border-b" class:grid-cols-2={ui.showTwoColumns} class:grid-rows-2={ui.showTwoRows}>
    {@render children()}
</div>
