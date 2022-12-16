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

    onMount(() => {
        jsonEditor = monaco.editor.create(document.getElementById('json-editor'), {
            value: "",
            language: 'json',
            fontFamily: 'Jetbrains Mono, monospace',
            fontSize: fontSize,
            minimap: {
                enabled: false
            },
            lineHeight: 25,
            automaticLayout: true,
        })

        goEditor = monaco.editor.create(document.getElementById('go-editor'), {
            value: "",
            language: 'go',
            fontFamily: 'Jetbrains Mono, monospace',
            fontSize: fontSize,
            minimap: {
                enabled: false
            },
            lineHeight: 25,
            automaticLayout: true,
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
                jsonEditor.setValue(result)
            }
        })
    }

    function saveGoSourceFile(): void {
        SaveGoSourceFile(goEditor.getValue())
    }

    function pasteJSON() {
        ReadClipboard().then(result => {
            jsonEditor.setValue(result)
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

    <MenuBar class="border-b">
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

    <div class="columns-2 gap-6 space-y-4 bg-white/25 h-64 grow">
        <div class="h-full max-h-full flex flex-col overflow-hidden border-r">
            <div class="w-full bg-white/50 border-b flex flex-row">
                <span class="py-1 px-4 select-none text-yellow-600 font-mono">JSON</span>
                <button on:click={pasteJSON}
                        class="text-sm px-5 py-1 ml-auto text-gray-800 select-none transition hover:bg-gray-300/25">{$_('Paste')}</button>
            </div>
            <div class="w-full h-full max-h-full" id="json-editor"></div>
        </div>
        <div class="h-full max-h-full flex flex-col overflow-hidden border-l">
            <div class="w-full bg-white/50 border-b flex flex-row">
                <span class="py-1 px-4 select-none text-purple-600 font-mono">Go</span>
                <button on:click={copyCode}
                        class="text-sm px-5 py-1 ml-auto text-gray-800 select-none transition hover:bg-gray-300/25">{$_('Copy')}</button>
            </div>
            <div class="w-full h-full max-h-full" id="go-editor"></div>
        </div>
    </div>

    <div class="flex flex-row px-4 py-2 border-t">
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

</style>
