<script lang="ts">
    import {
        Generate,
        GetConfig,
        OpenJSONFile,
        ReadClipboard,
        SaveGoSourceFile,
        SetConfig,
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
        TextBlock,
        TextBox
    } from "fluent-svelte";
    import "./i18n"
    import {_, locale, locales} from "svelte-i18n";
    import {EventsEmit} from "../wailsjs/runtime";
    import * as monaco from 'monaco-editor';
    import {editor} from 'monaco-editor';
    import {onMount} from "svelte";
    import './worker';
    import IStandaloneCodeEditor = editor.IStandaloneCodeEditor;

    let jsonEditor: IStandaloneCodeEditor
    let goEditor: IStandaloneCodeEditor

    const defaultFontSize = 16
    let fontSize = defaultFontSize
    GetConfig("font_size").then(result => {
        fontSize = result
    })

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

    onMount(() => {
        jsonEditor = createEditor(document.getElementById('json-editor'), 'json', '')
        goEditor = createEditor(document.getElementById('go-editor'), 'go', '')
        // remeasure fonts after creating editors and fonts are loaded to avoid rendering issues
        document.fonts.ready.then(() => {
            monaco.editor.remeasureFonts()
        })
    })

    let allCaps: string = "ID,URL"

    function generate(): void {
        Generate(jsonEditor.getValue(), allCaps.split(",").map((v) => v.trim())).then(result => {
            goEditor.setValue(result)
        })
    }

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

    let openSettingsDialog = false;

    GetConfig("locale").then(result => {
        if (result !== "") {
            $locale = result
        }
    })

    $: SetConfig("locale", $locale)
    $: {
        jsonEditor?.updateOptions({fontSize})
        goEditor?.updateOptions({fontSize})
        SetConfig("font_size", fontSize)
    }

</script>

<main class="w-screen h-screen flex flex-col">
    <ContentDialog bind:open={openSettingsDialog} title={$_('Settings')}>
        <div class="space-y-1.5 flex flex-col">
            <TextBlock class="select-none text-lg">{$_('All caps')}</TextBlock>
            <TextBox id="allCaps" bind:value={allCaps}></TextBox>
        </div>
        <Button slot="footer" on:click={() => {openSettingsDialog = false}}>{$_('Close')}</Button>
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
    </MenuBar>

    <!-- use columns-2 will cause the editor to be rendered incorrectly, so use grid instead -->
    <div class="grid grid-cols-2 h-64 grow border-t border-b">
        <div id="container-json">
            <div class="w-full bg-white/50 flex flex-row">
                <span class="py-1 px-4 select-none text-yellow-700 font-mono">JSON</span>
                <button on:click={pasteJSON}>{$_('Paste')}</button>
            </div>
            <div class="w-full h-32 grow" id="json-editor"></div>
        </div>
        <div class="border-l" id="container-go">
            <div class="w-full bg-white/50 flex flex-row">
                <span class="py-1 px-4 select-none text-purple-700 font-mono">Go</span>
                <button on:click={copyCode}>{$_('Copy')}</button>
            </div>
            <div class="w-full h-32 grow" id="go-editor"></div>
        </div>
    </div>

    <div class="flex flex-row px-4 py-2">
        <Button variant="accent" on:click={generate} class="ml-auto mr-2">{$_('Generate')}</Button>
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
