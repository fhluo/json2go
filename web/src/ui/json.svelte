<script lang="ts">
    import {getContext} from "svelte";
    import {_} from "svelte-i18n";
    import {type UIState} from "../state/ui.svelte";
    import type {EditorsState} from "../state/editors.svelte";

    const ui = getContext<UIState>("ui");
    const editors = getContext<EditorsState>("editors");

    export interface Props {
        jsonElement: HTMLElement;
    }

    let {jsonElement = $bindable<HTMLElement>()}: Props = $props();
</script>

<div id="container-json" class="group" class:hidden={ui.hideJSONEditor} class:col-span-2={ui.showTwoColumns}
     class:row-span-2={ui.showTwoRows}>
    <div class="w-full bg-white/50 flex flex-row">
        <span class="py-1 px-4 select-none text-yellow-700 font-mono group-focus-within:text-yellow-500">
            JSON
        </span>
        <button type="button" on:click={editors.pasteToJSONEditor}>
            {$_("Paste")}
        </button>
    </div>
    <div class="w-full h-32 grow" id="json-editor" bind:this={jsonElement}></div>
</div>
