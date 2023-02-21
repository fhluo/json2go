<script lang="ts">
    import {
        GetFontSize,
        GetLocale,
        ReadClipboard,
        SetFontSize,
        SetLocale,
        WriteClipboard
    } from "../../wailsjs/go/main/App";
    import {_, locale} from "svelte-i18n";
    import {onMount} from "svelte";
    import loader from "@monaco-editor/loader";
    import {editor} from "monaco-editor/esm/vs/editor/editor.api";
    import IStandaloneCodeEditor = editor.IStandaloneCodeEditor;
    import {Layout, View} from "./types";

    export let layout: Layout
    export let view: View

    export let jsonEditor: IStandaloneCodeEditor
    export let goEditor: IStandaloneCodeEditor

    export const defaultFontSize = 16
    export let fontSize = defaultFontSize

    function pasteJSON() {
        ReadClipboard().then(result => {
            jsonEditor.executeEdits("", [{
                range: jsonEditor.getModel()!.getFullModelRange(),
                text: result,
            }])
        })
    }

    function copyCode(): void {
        WriteClipboard(goEditor.getValue())
    }

    const monacoLocales = ["en", "de", "es", "fr", "it", "ja", "ko", "ru", "zh-cn", "zh-tw"]

    function getMonacoLocale(locale: string): string {
        locale = locale.toLowerCase()
        if (monacoLocales.includes(locale)) {
            return locale
        }

        if (locale == "zh") {
            return "zh-cn"
        }

        return "en"
    }

    GetLocale().then(result => {
        if (result !== "") {
            $locale = result
        }

        loader.config({
            'vs/nls': {
                availableLanguages: {
                    '*': getMonacoLocale($locale)
                }
            }
        })
    })

    $: SetLocale($locale)
    $: {
        jsonEditor?.updateOptions({fontSize})
        goEditor?.updateOptions({fontSize})
        SetFontSize(fontSize)
    }

    loader.config({paths: {vs: 'monaco-editor/min/vs'}})

    onMount(() => {
        loader.init().then(monaco => {
            function createEditor(domElement: HTMLElement, language: string, value: string): IStandaloneCodeEditor {
                return monaco.editor.create(domElement, {
                    value: value,
                    language: language,
                    fontFamily: 'Jetbrains Mono, monospace',
                    fontSize: fontSize,
                    minimap: {
                        enabled: false
                    },
                    lineHeight: 25,
                    automaticLayout: true,
                })
            }

            jsonEditor = createEditor(document.getElementById('json-editor')!, 'json', '')
            goEditor = createEditor(document.getElementById('go-editor')!, 'go', '')

            // remeasure fonts after creating editors and fonts are loaded to avoid rendering issues
            document.fonts.ready.then(() => {
                monaco.editor.remeasureFonts()
            })
        })
    })
</script>

<div class="grid h-64 grow border-t border-b"
     class:grid-cols-2={layout===Layout.TwoColumns}
     class:grid-rows-2={layout===Layout.TwoRows}
>
    <div id="container-json"
         style:display={view===View.GoOnly ? "none" : ""}
         class:col-span-2={layout===Layout.TwoColumns && view===View.JSONOnly}
         class:row-span-2={layout===Layout.TwoRows && view===View.JSONOnly}
    >
        <div class="w-full bg-white/50 flex flex-row">
            <span class="py-1 px-4 select-none text-yellow-700 font-mono">JSON</span>
            <button on:click={pasteJSON}>{$_('Paste')}</button>
        </div>
        <div class="w-full h-32 grow" id="json-editor"></div>
    </div>

    <div id="container-go"
         style:display={view===View.JSONOnly ? "none" : ""}
         class:col-span-2={layout===Layout.TwoColumns && view===View.GoOnly}
         class:row-span-2={layout===Layout.TwoRows && view===View.GoOnly}
         class:border-l={layout===Layout.TwoColumns && view===View.JSONAndGo}
         class:border-t={layout===Layout.TwoRows && view===View.JSONAndGo}
    >
        <div class="w-full bg-white/50 flex flex-row">
            <span class="py-1 px-4 select-none text-purple-700 font-mono">Go</span>
            <button on:click={copyCode}>{$_('Copy')}</button>
        </div>
        <div class="w-full h-32 grow" id="go-editor"></div>
    </div>
</div>


<style>
    @font-face {
        font-family: "JetBrains Mono";
        font-style: normal;
        font-weight: 400;
        src: local(""), url("../assets/fonts/JetBrainsMono-Regular.woff2") format("woff2");;
    }

    #container-json, #container-go {
        @apply h-full max-h-full flex flex-col bg-white;
    }

    #container-json button, #container-go button {
        @apply text-sm px-5 py-1 ml-auto text-gray-800 select-none transition;
    }

    #container-json button:hover, #container-go button:hover {
        @apply bg-gray-300/25;
    }

    #container-json span, #container-go span {
        @apply transition duration-300;
    }

    #container-json:focus-within span {
        @apply text-yellow-500;
    }

    #container-go:focus-within span {
        @apply text-purple-500;
    }
</style>
