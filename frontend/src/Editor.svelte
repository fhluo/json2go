<script lang="ts">
    import {basicSetup, EditorView} from "codemirror";
    import {EditorState, type Extension} from "@codemirror/state";
    import {drawSelection, keymap} from "@codemirror/view";
    import {indentWithTab} from "@codemirror/commands";
    import {json} from "@codemirror/lang-json";
    import {StreamLanguage} from "@codemirror/language";
    import {go} from "@codemirror/legacy-modes/mode/go";
    import {onMount} from "svelte";

    let className: string
    export {className as class}

    let element: HTMLDivElement

    export let lang: string = ""

    let langSupport: Extension
    switch (lang) {
        case "json":
            langSupport = json()
            break
        case "go":
            langSupport = StreamLanguage.define(go)
    }

    let view: EditorView
    let state: EditorState = EditorState.create({
        extensions: [basicSetup, drawSelection(), keymap.of([indentWithTab]), langSupport]
    })

    onMount(() => {
        view = new EditorView({
            state,
            parent: element
        })
    })

    export let editor: { text: string } = {
        get text() {
            return view.state.doc.toString()
        },
        set text(s: string) {
            view.dispatch({
                changes: {from: 0, to: view.state.doc.length, insert: s}
            })
        }
    }

</script>

<div class="editor {className}" bind:this={element}></div>

<style global>
    @font-face {
        font-family: "JetBrains Mono";
        font-style: normal;
        font-weight: 400;
        src: local(""), url("assets/fonts/JetBrainsMono-Regular.woff2") format("woff2");;
    }

    .cm-editor {
        @apply h-full max-h-full;
    }

    .cm-editor .cm-scroller {
        @apply overflow-auto;
    }

    .cm-editor.cm-focused {
        outline: 0 !important;
    }

    .cm-editor.cm-focused .cm-selectionBackground, ::selection {
        @apply bg-blue-500;
    }

    .cm-editor .cm-content {
        font-family: "JetBrains Mono", ui-monospace, monospace;
        @apply bg-white/75;
    }

    .cm-editor .cm-gutters {
        @apply bg-white/75 font-mono pl-2 leading-7;
    }

    .cm-editor .cm-line {
        @apply leading-7;
    }
</style>