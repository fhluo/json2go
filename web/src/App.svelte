<script lang="ts">
    import {
        Generate,
        GetAllCapsWords,
        GetFontSize,
        GetLocale,
        OpenJSONFile,
        ReadClipboard,
        SaveGoSourceFile,
        SetAllCapsWords,
        SetFontSize,
        SetLocale,
        WriteClipboard
    } from '../wailsjs/go/main/App.js'
    import "fluent-svelte/theme.css";
    import {
        Button,
        ContentDialog,
        MenuBar,
        MenuBarItem,
        MenuFlyoutDivider,
        MenuFlyoutItem,
        TextBox
    } from "fluent-svelte";
    import "./i18n"
    import {_, locale, locales} from "svelte-i18n";
    import {EventsEmit, EventsOn} from "../wailsjs/runtime";
    import {editor} from "monaco-editor/esm/vs/editor/editor.api";
    import {onMount} from "svelte";
    import loader from "@monaco-editor/loader";
    import {BrowserOpenURL} from "../wailsjs/runtime/runtime.js";
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

    function openJSONFile(): void {
        OpenJSONFile().then(result => {
            if (result !== "") {
                jsonEditor.executeEdits("", [{
                    range: jsonEditor.getModel()!.getFullModelRange(),
                    text: result,
                }])
            }
        })
    }

    function saveGoSourceFile(): void {
        SaveGoSourceFile(goEditor.getValue())
    }

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

    let openSettingsDialog = false
    let openAboutDialog = false
    let showErrorInfo = false
    let showJSONContainer = true
    let showGoContainer = true
    let errorMessage = ""
    let allCapsWord = ""
    let allCapsWords = [] as string[]
    GetAllCapsWords().then(result => {
        if (result != null) {
            allCapsWords = result
        }
    })

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
    $: SetAllCapsWords(allCapsWords)

    document.defaultView.addEventListener('resize', () => {
        jsonEditor?.layout()
        goEditor?.layout()
        EventsEmit("resize")
    })

    EventsOn("error", (message: string) => {
        showErrorInfo = true
        errorMessage = message
    })

    function addAllCapsWord(): void {
        allCapsWord = allCapsWord.trim()
        if (allCapsWord === "") {
            return
        }

        // if allCapsWord contains a comma, split it and add each word separately
        allCapsWords = Array.from(new Set(allCapsWords.concat(allCapsWord.split(',').map(word => word.trim()).filter(word => word !== ""))))
        allCapsWord = ""
    }

    function onBeforeInput(event: InputEvent): void {
        // user can only enter letters, space and ','
        if (event.data !== null && !event.data.match(/[a-zA-Z', ]/)) {
            event.preventDefault()
        }
    }
</script>

<main class="w-screen h-screen flex flex-col">
    <ContentDialog bind:open={openSettingsDialog} title={$_('Settings')}>
        <div class="flex flex-col space-y-2">
            <span class="select-none font-semibold mr-3">{$_('All-caps words')}</span>
            {#if allCapsWords?.length > 0}
                <div class="flex flex-row flex-wrap space-x-1.5 pb-1.5">
                    {#each allCapsWords as word}
                        <button class="px-3 leading-loose tracking-wide rounded border shadow-sm ml-1 mt-1.5 bg-white hover:bg-gray-50 transition cursor-default"
                                on:dblclick={() => allCapsWords = allCapsWords.filter(w => w !== word)}
                                on:click={() => allCapsWord = word}>{word}</button>
                    {/each}
                </div>
            {/if}
            <div class="flex flex-row space-x-1.5">
                <TextBox bind:value={allCapsWord} on:beforeinput={onBeforeInput}></TextBox>
                <Button on:click={addAllCapsWord} class="min-w-fit">{$_('Add')}</Button>
                <Button on:click={() => allCapsWords = allCapsWords.filter(w => w !== allCapsWord)}
                        class="min-w-fit">{$_('Remove')}</Button>
            </div>
            <p class="text-sm text-gray-500 leading-relaxed py-2 px-1">{$_('tip.allCaps', {default: "Tip: Double click a word to remove it. To add multiple words, separate words with commas."})}</p>
        </div>
        <div class="flex justify-end">
            <Button variant="accent" on:click={() => {openSettingsDialog = false}} class="mr-2">{$_('OK')}</Button>
        </div>
    </ContentDialog>

    <ContentDialog bind:open={openAboutDialog} title={$_('about.title', {default: 'About'})}>
        <div class="flex flex-col space-y-2 items-center justify-center">
            <p class="text-lg font-semibold">JSON2Go</p>
            <p class="text-gray-900">{$_('about.about', {default: "Generate Go type definitions from JSON"})}</p>
            <div class="leading-relaxed">
                <p class="text-gray-900">
                    <span class="font-semibold">{$_('about.license', {default: 'License: '})}</span>MIT
                </p>
                <p class="text-gray-900">
                    <span class="font-semibold">{$_('about.version', {default: 'Version: '})}</span>0.1.0
                </p>
            </div>
            <p class="text-gray-900">Copyright © 2022 fhluo</p>
        </div>
        <div class="flex justify-end mt-3.5">
            <Button variant="accent" on:click={() => {openAboutDialog=false}} class="mr-2">{$_('OK')}</Button>
        </div>
    </ContentDialog>

    <MenuBar>
        <MenuBarItem>
            {$_('File')}
            <svelte:fragment slot="flyout">
                <MenuFlyoutItem on:click={openJSONFile}>{$_('Open JSON file')}</MenuFlyoutItem>
                <MenuFlyoutItem on:click={saveGoSourceFile}>{$_('Save Go source file')}</MenuFlyoutItem>
                <MenuFlyoutDivider/>
                <MenuFlyoutItem on:click={() => {openSettingsDialog = true}}>{$_('Settings')}</MenuFlyoutItem>
                <MenuFlyoutDivider/>
                <MenuFlyoutItem on:click={() => EventsEmit("exit")}>{$_('Exit')}</MenuFlyoutItem>
            </svelte:fragment>
        </MenuBarItem>
        <MenuBarItem>
            {$_('View')}
            <svelte:fragment slot="flyout">
                <MenuFlyoutItem variant="toggle" bind:checked={showJSONContainer}>{$_('JSON')}</MenuFlyoutItem>
                <MenuFlyoutItem variant="toggle" bind:checked={showGoContainer}>{$_('Go')}</MenuFlyoutItem>
            </svelte:fragment>
        </MenuBarItem>
        <MenuBarItem>
            {$_('Language')}
            <svelte:fragment slot="flyout">
                {#each $locales as _locale}
                    <MenuFlyoutItem variant="radio" bind:group={$locale} name="locale" value={_locale}
                                    checked={$locale===_locale}>{$_(_locale)}</MenuFlyoutItem>
                {/each}
            </svelte:fragment>
        </MenuBarItem>
        <MenuBarItem>
            {$_('Font')}
            <svelte:fragment slot="flyout">
                <MenuFlyoutItem on:click={()=>fontSize++}>{$_('Increase size')}</MenuFlyoutItem>
                <MenuFlyoutItem on:click={()=>fontSize--}>{$_('Decrease size')}</MenuFlyoutItem>
                <MenuFlyoutItem on:click={()=>fontSize=defaultFontSize}>{$_('Reset size')}</MenuFlyoutItem>
            </svelte:fragment>
        </MenuBarItem>
        <MenuBarItem>
            {$_('Help')}
            <svelte:fragment slot="flyout">
                <MenuFlyoutItem on:click={()=>BrowserOpenURL("https://github.com/fhluo/json2go")}>
                    {$_('Document')}
                </MenuFlyoutItem>
                <MenuFlyoutDivider/>
                <MenuFlyoutItem on:click={()=>openAboutDialog=true}>
                    {$_('about.title', {default: 'About'})}
                </MenuFlyoutItem>
            </svelte:fragment>
        </MenuBarItem>
    </MenuBar>

    <!-- use columns-2 will cause the editor to be rendered incorrectly, so use grid instead -->
    <div class="grid grid-cols-2 h-64 grow border-t border-b">
        <div id="container-json" class:col-span-2={showJSONContainer &&!showGoContainer}
             style:display={!showJSONContainer ? "none" : ""}>
            <div class="w-full bg-white/50 flex flex-row">
                <span class="py-1 px-4 select-none text-yellow-700 font-mono">JSON</span>
                <button on:click={pasteJSON}>{$_('Paste')}</button>
            </div>
            <div class="w-full h-32 grow" id="json-editor"></div>
        </div>
        <div id="container-go" class:col-span-2={showGoContainer && !showJSONContainer}
             style:display={!showGoContainer ? "none" : ""} class="border-l">
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
