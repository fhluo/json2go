<script lang="ts">
    import {Generate, ReadClipboard, WriteClipboard} from '../wailsjs/go/main/App.js'
    import {basicSetup, EditorView} from "codemirror";
    import {json} from "@codemirror/lang-json";
    import {drawSelection, keymap} from "@codemirror/view";
    import {indentWithTab} from "@codemirror/commands";
    import {onMount} from "svelte";
    import {go} from "@codemirror/legacy-modes/mode/go";
    import {StreamLanguage} from "@codemirror/language";
    import {EditorState} from "@codemirror/state";

    let jsonView: EditorView
    let goView: EditorView

    let jsonState = EditorState.create({
        extensions: [basicSetup, drawSelection(), keymap.of([indentWithTab]), json()],
    })
    let goState = EditorState.create({
        extensions: [
            EditorState.tabSize.of(4), basicSetup, drawSelection(),
            keymap.of([indentWithTab]), StreamLanguage.define(go)
        ],
    })

    onMount(async () => {
        jsonView = new EditorView({
            state: jsonState,
            parent: document.querySelector("#json")
        })

        goView = new EditorView({
            state: goState,
            parent: document.querySelector("#go")
        })
    })


    function generate(): void {
        Generate(jsonView.state.doc.toString()).then(result => {
            goView.dispatch({
                changes: {from: 0, to: goView.state.doc.length, insert: result}
            })
        })
    }

    function pasteJSON() {
        ReadClipboard().then(result => {
            jsonView.dispatch({
                changes: {from: 0, to: jsonView.state.doc.length, insert: result}
            })
            generate()
        })
    }

    function copyCode(): void {
        WriteClipboard(goView.state.doc.toString())
    }
</script>

<main>
    <div class="flex flex-col h-screen w-screen">
        <div class="self-center justify-self-center">
            <button class="button" on:click={generate}>Generate</button>
        </div>

        <div class="columns-2 gap-6 mx-8 grow mb-8 mt-4">
            <div class="code h-full max-h-full flex flex-col" id="json">
                <div class="w-full bg-white/50 border-b border-t flex flex-row">
                    <span class="py-1 px-4 select-none text-yellow-600 font-mono">JSON</span>
                    <button on:click={pasteJSON} class="font-mono">Paste</button>
                </div>
            </div>
            <div class="code h-full flex flex-col" id="go">
                <div class="w-full bg-white/50 border-b flex flex-row">
                    <span class="py-1 px-4 select-none text-purple-600 font-mono">Go</span>
                    <button on:click={copyCode} class="font-mono">Copy</button>
                </div>
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
