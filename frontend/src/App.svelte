<script lang="ts">
    import {Generate, ReadClipboard, WriteClipboard} from '../wailsjs/go/main/App.js'
    import Editor from "./Editor.svelte";

    let input: { set(s: string): void; get(): string }
    let output: { set(s: string): void; get(): string }

    function generate(): void {
        Generate(input.get()).then(result => {
            output.set(result)
        })
    }

    function pasteJSON() {
        ReadClipboard().then(result => {
            input.set(result)
            generate()
        })
    }

    function copyCode(): void {
        WriteClipboard(output.get())
    }
</script>

<main>
    <div class="flex flex-col h-screen w-screen max-h-screen overflow-clip">
        <div class="self-center justify-self-center">
            <button class="button" on:click={generate}>Generate</button>
        </div>

        <div class="columns-2 gap-6 mx-8 mb-8 mt-4 h-full max-h-full">
            <div class="code h-full max-h-full flex flex-col overflow-clip">
                <div class="w-full bg-white/50 border-b border-t flex flex-row">
                    <span class="py-1 px-4 select-none text-yellow-600 font-mono">JSON</span>
                    <button on:click={pasteJSON} class="font-mono">Paste</button>
                </div>
                <Editor class="w-full h-full border-b" lang="json" bind:value={input}/>
            </div>
            <div class="code h-full max-h-full flex flex-col overflow-clip">
                <div class="w-full bg-white/50 border-b border-t flex flex-row">
                    <span class="py-1 px-4 select-none text-purple-600 font-mono">Go</span>
                    <button on:click={copyCode} class="font-mono">Copy</button>
                </div>
                <Editor class="w-full h-full border-b" lang="go" bind:value={output}/>
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

    .button {
        @apply w-fit px-4 py-1 bg-white/75 border rounded-md transition select-none;
    }

    .button:hover {
        @apply bg-gray-100;
    }
</style>
