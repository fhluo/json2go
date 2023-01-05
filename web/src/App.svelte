<script lang="ts">
    import {
        Generate,
        GetFontSize,
        GetLocale,
        ReadClipboard,
        SetFontSize,
        SetLocale,
        WriteClipboard
    } from '../wailsjs/go/main/App.js'
    import "fluent-svelte/theme.css";
    import {Button} from "fluent-svelte";
    import "./i18n"
    import {_, locale} from "svelte-i18n";
    import {EventsEmit, EventsOn} from "../wailsjs/runtime";
    import {editor} from "monaco-editor/esm/vs/editor/editor.api";
    import {onMount} from "svelte";
    import loader from "@monaco-editor/loader";
    import MenuBar from "./MenuBar.svelte";
    import {Editors, Layout} from "./base";
    import IStandaloneCodeEditor = editor.IStandaloneCodeEditor;

    let jsonEditor: IStandaloneCodeEditor
    let goEditor: IStandaloneCodeEditor

    const defaultFontSize = 16
    let fontSize = defaultFontSize
    GetFontSize().then(result => {
        fontSize = result
    })

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

    loader.config({paths: {vs: 'monaco-editor/min/vs'}})

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

    let showErrorInfo = false
    let errorMessage = ""
    let allCapsWords = [] as string[]

    let layout = Layout.TwoColumns
    let editors = Editors.Both

    function generate(): void {
        showErrorInfo = false
        errorMessage = ""
        Generate(jsonEditor.getValue(), allCapsWords).then(result => {
            goEditor.setValue(result)
        })
    }

    $: SetLocale($locale)
    $: {
        jsonEditor?.updateOptions({fontSize})
        goEditor?.updateOptions({fontSize})
        SetFontSize(fontSize)
    }

    document.defaultView.addEventListener('resize', () => {
        jsonEditor?.layout()
        goEditor?.layout()
        EventsEmit("resize")
    })

    EventsOn("error", (message: string) => {
        showErrorInfo = true
        errorMessage = message
    })
</script>

<main class="w-screen h-screen flex flex-col">
    <MenuBar bind:layout={layout} bind:editors={editors} bind:fontSize={fontSize} bind:jsonEditor={jsonEditor}
             bind:goEditor={goEditor} bind:allCapsWords></MenuBar>
    <!-- use columns-2 will cause the editor to be rendered incorrectly, so use grid instead -->
    <div class="grid h-64 grow border-t border-b" class:grid-cols-2={layout===Layout.TwoColumns}
         class:grid-rows-2={layout===Layout.TwoRows}>
        <div id="container-json" style:display={editors===Editors.Go ? "none" : ""}
             class:col-span-2={layout===Layout.TwoColumns && editors===Editors.JSON}
             class:row-span-2={layout===Layout.TwoRows && editors===Editors.JSON}>
            <div class="w-full bg-white/50 flex flex-row">
                <span class="py-1 px-4 select-none text-yellow-700 font-mono">JSON</span>
                <button on:click={pasteJSON}>{$_('Paste')}</button>
            </div>
            <div class="w-full h-32 grow" id="json-editor"></div>
        </div>
        <div id="container-go" style:display={editors===Editors.JSON ? "none" : ""}
             class:col-span-2={layout===Layout.TwoColumns && editors===Editors.Go}
             class:row-span-2={layout===Layout.TwoRows && editors===Editors.Go}
             class:border-l={layout===Layout.TwoColumns && editors===Editors.Both}
             class:border-t={layout===Layout.TwoRows && editors===Editors.Both}>
            <div class="w-full bg-white/50 flex flex-row">
                <span class="py-1 px-4 select-none text-purple-700 font-mono">Go</span>
                <button on:click={copyCode}>{$_('Copy')}</button>
            </div>
            <div class="w-full h-32 grow" id="go-editor"></div>
        </div>
    </div>

    <div class="flex flex-row px-4 py-2 justify-end items-center h-12">
        {#if showErrorInfo}
            <div class="select-none mx-4 border bg-white/50 rounded shadow-sm flex flex-row items-center justify-center space-x-1.5">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
                     stroke="currentColor" class="w-6 h-6 ml-2 mr-1 text-red-600">
                    <path stroke-linecap="round" stroke-linejoin="round"
                          d="M12 9v3.75m9-.75a9 9 0 11-18 0 9 9 0 0118 0zm-9 3.75h.008v.008H12v-.008z"/>
                </svg>
                <span>{errorMessage}</span>
                <button class="hover:bg-gray-200/50 py-1 px-1 rounded-r transition flex items-center justify-center"
                        on:click={()=>showErrorInfo=false}>
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1"
                         stroke="currentColor" class="w-6 h-6">
                        <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12"/>
                    </svg>
                </button>
            </div>
        {/if}
        <Button variant="accent" on:click={generate} class="mr-2">{$_('Generate')}</Button>
    </div>
</main>

<style>
    @font-face {
        font-family: "JetBrains Mono";
        font-style: normal;
        font-weight: 400;
        src: local(""), url("assets/fonts/JetBrainsMono-Regular.woff2") format("woff2");;
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
