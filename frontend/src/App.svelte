<script lang="ts">
    import {Generate, ReadClipboard, WriteClipboard} from '../wailsjs/go/main/App.js'

    let code: string
    let json: string

    function generate(): void {
        Generate(json).then(result => code = result)
    }

    function readClipboard() {
        ReadClipboard().then(result => json = result)
    }

    function writeClipboard(): void {
        WriteClipboard(code)
    }
</script>

<main>
    <div class="flex flex-col h-screen w-screen">
        <div>
            <button class="w-fit px-5 py-1.5 bg-white/75 border rounded-md hover:bg-gray-100 transition self-center"
                    on:click={generate}>
                Generate
            </button>

            <button class="w-fit px-5 py-1.5 bg-white/75 border rounded-md hover:bg-gray-100 transition self-center"
                    on:click={writeClipboard}>
                Copy Code
            </button>

            <button class="w-fit px-5 py-1.5 bg-white/75 border rounded-md hover:bg-gray-100 transition self-center"
                    on:click={readClipboard}>
                Paste JSON
            </button>
        </div>

        <div class="columns-2 gap-6 mx-8 grow mb-8 mt-4">
            <div class="h-full">
                <textarea bind:value={json} class="w-full h-full" autocomplete="off" on:change={generate}
                          spellcheck="false"></textarea>
            </div>
            <div class="h-full">
                <textarea bind:value={code} class="w-full h-full" autocomplete="off" spellcheck="false"></textarea>
            </div>
        </div>
    </div>
</main>

<style>
    textarea {
        @apply rounded-md border-gray-300 shadow-inner resize-none;
        @apply font-mono;
    }

    textarea:focus {
        @apply ring-1 ring-blue-400;
    }
</style>
