<script lang="ts">
    import { getContext } from "svelte";
    import { _ } from "svelte-i18n";
    import type { UIState } from "../state/ui.svelte";
    import type { EditorsState } from "../state/editors.svelte";

    const ui = getContext<UIState>("ui");
    const editors = getContext<EditorsState>("editors");

    function pasteJSON() {
        editors.pasteToJSONEditor();
    }
</script>

<div
    id="container-json"
    style="display: {ui.view === 'Go Only' ? 'none' : ''}"
    class="group {ui.layout === 'Two Columns' && ui.view === 'JSON Only' ? 'col-span-2' : ''} {ui.layout === 'Two Rows' && ui.view === 'JSON Only' ? 'row-span-2' : ''}"
>
    <div class="w-full bg-white/50 flex flex-row">
        <span class="py-1 px-4 select-none text-yellow-700 font-mono group-focus-within:text-yellow-500">
            JSON
        </span>
        <button type="button" on:click={pasteJSON}>
            {$_("Paste")}
        </button>
    </div>
    <div class="w-full h-32 grow" id="json-editor"></div>
</div>
