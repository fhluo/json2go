<script lang="ts">
    import {Generate, ReadClipboard, WriteClipboard} from '../wailsjs/go/main/App.js'
    import Editor from "./Editor.svelte";
    import Button from "./Button.svelte";
    import "./i18n"
    import {_, locale, locales} from "svelte-i18n";


    let input: { text: string }
    let output: { text: string }

    let allCaps: string = "ID"

    function generate(): void {
        Generate(input.text, allCaps.split(",").map((v) => v.trim())).then(result => {
            output.text = result
        })
    }

    function pasteJSON() {
        ReadClipboard().then(result => {
            input.text = result
            generate()
        })
    }

    function copyCode(): void {
        WriteClipboard(output.text)
    }

</script>

<main class="flex flex-col w-screen h-screen px-6 py-4 space-y-4">
    <div class="columns-2 gap-6 mt-4 h-full max-h-full mx-2">
        <div class="code h-full max-h-full flex flex-col overflow-clip">
            <div class="w-full bg-white/50 border-b border-t flex flex-row">
                <span class="py-1 px-4 select-none text-yellow-600 font-mono">JSON</span>
                <button on:click={pasteJSON}>{$_('paste', {default: 'Paste'})}</button>
            </div>
            <Editor class="w-full h-full border-b" lang="json" bind:editor={input}/>
        </div>
        <div class="code h-full max-h-full flex flex-col overflow-clip">
            <div class="w-full bg-white/50 border-b border-t flex flex-row">
                <span class="py-1 px-4 select-none text-purple-600 font-mono">Go</span>
                <button on:click={copyCode}>{$_('copy', {default: 'Copy'})}</button>
            </div>
            <Editor class="w-full h-full border-b" lang="go" bind:editor={output}/>
        </div>
    </div>

    <div class="flex flex-col space-y-2 mx-2 border rounded-lg px-4 py-3 bg-white/50">
        <div class="flex flex-row items-center gap-x-5">
            <div class="space-x-1.5">
                <label>{$_('language', {default: 'Language'})}</label>
                <select bind:value={$locale} class="w-fit py-1 border-gray-300 focus:ring-1 focus:ring-cyan-300 rounded">
                    {#each $locales as locale}
                        <option value={locale}>{locale}</option>
                    {/each}
                </select>
            </div>
            <div class="space-x-1.5 flex-grow flex flex-row items-center">
                <label for="allCaps" class="select-none">{$_('allCaps', {default: 'All caps'})}</label>
                <input type="text" id="allCaps"
                       class="rounded-md flex-grow border-gray-300 focus:ring-1 focus:ring-cyan-300 py-1 transition"
                       bind:value={allCaps}>
            </div>
            <Button on:click={generate}>{$_('generate', {default: 'Generate'})}</Button>
        </div>
    </div>
</main>

<style>
    .code {
        @apply border rounded-lg py-2.5 shadow-sm transition duration-200;
    }

    .code:focus-within {
        @apply ring-1 shadow-md;
    }

    .code button {
        @apply px-5 py-1 ml-auto text-gray-800 select-none transition;
    }

    .code button:hover {
        @apply bg-gray-300/25;
    }
</style>
