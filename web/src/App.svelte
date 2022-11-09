<script lang="ts">
    import {Generate, GetAcronyms, ReadClipboard, SetAcronyms, WriteClipboard} from '../wailsjs/go/main/App.js'
    import Editor from "./Editor.svelte";
    import Button from "./Button.svelte";
    import {onMount} from "svelte";

    let input: { text: string }
    let output: { text: string }

    let acronyms: string

    function setAcronyms() {
        SetAcronyms(acronyms.split(",").map(acronym => acronym.trim()))
    }

    function getAcronyms() {
        GetAcronyms().then(result => {
            acronyms = result.join(", ")
        })
    }

    function generate(): void {
        Generate(input.text).then(result => {
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

    onMount(()=>{
        getAcronyms()
    })
</script>

<main>
    <div class="flex flex-col h-screen w-screen max-h-screen overflow-clip">
        <div class="flex flex-row mx-8 mt-4 gap-3 items-center">
            <label for="acronyms">Acronyms</label>
            <input type="text" id="acronyms"
                   class="rounded-md flex-grow border-gray-300 focus:ring-1 focus:ring-cyan-300 py-1 transition mr-3"
                   bind:value={acronyms} on:change={setAcronyms}>
            <Button on:click={generate}>Generate</Button>
        </div>

        <div class="columns-2 gap-6 mx-8 mb-8 mt-4 h-full max-h-full">
            <div class="code h-full max-h-full flex flex-col overflow-clip">
                <div class="w-full bg-white/50 border-b border-t flex flex-row">
                    <span class="py-1 px-4 select-none text-yellow-600 font-mono">JSON</span>
                    <button on:click={pasteJSON} class="font-mono">Paste</button>
                </div>
                <Editor class="w-full h-full border-b" lang="json" bind:editor={input}/>
            </div>
            <div class="code h-full max-h-full flex flex-col overflow-clip">
                <div class="w-full bg-white/50 border-b border-t flex flex-row">
                    <span class="py-1 px-4 select-none text-purple-600 font-mono">Go</span>
                    <button on:click={copyCode} class="font-mono">Copy</button>
                </div>
                <Editor class="w-full h-full border-b" lang="go" bind:editor={output}/>
            </div>
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
