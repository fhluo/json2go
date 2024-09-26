<script lang="ts">
    import { getContext } from "svelte";
    import { _ } from "svelte-i18n";
    import type { UIState } from "../state/ui.svelte";
    import type { EditorsState } from "../state/editors.svelte";

    const ui = getContext<UIState>("ui");
    const editors = getContext<EditorsState>("editors");

    function copyGo() {
        editors.copyFromGoEditor();
    }
</script>

<div id="container-go" class="group" class:hidden={ui.hideGoEditor} class:col-span-2={ui.showTwoColumns}
     class:row-span-2={ui.showTwoRows}>
    <div class="w-full bg-white/50 flex flex-row">
        <span class="py-1 px-4 select-none text-purple-700 font-mono group-focus-within:text-purple-500">
            Go
        </span>
        <button type="button" on:click={copyGo}>
            {$_("Copy")}
        </button>
    </div>
    <div class="w-full h-32 grow" id="go-editor"></div>
</div>
