<script lang="ts">
    import {Generate, ReadClipboard, WriteClipboard} from '../wailsjs/go/main/App.js'
    import "fluent-svelte/theme.css";
    import {Button, MenuBar, MenuBarItem, MenuFlyoutDivider, MenuFlyoutItem} from "fluent-svelte";
    import "./i18n"
    import {_, locale, locales} from "svelte-i18n";
    import {EventsEmit} from "../wailsjs/runtime";
    import * as monaco from 'monaco-editor';
    import {editor} from 'monaco-editor';
    import {onMount} from "svelte";
    import IStandaloneCodeEditor = editor.IStandaloneCodeEditor;

    let jsonEditor: IStandaloneCodeEditor
    let goEditor: IStandaloneCodeEditor

    onMount(() => {
        jsonEditor = monaco.editor.create(document.getElementById('json-editor'), {
            value: "",
            language: 'json'
        });

        goEditor = monaco.editor.create(document.getElementById('go-editor'), {
            value: "",
            language: 'go'
        });
    })


    let allCaps: string = "ID"

    function generate(): void {
        Generate(jsonEditor.getValue(), allCaps.split(",").map((v) => v.trim())).then(result => {
            goEditor.setValue(result)
        })
    }

    function pasteJSON() {
        ReadClipboard().then(result => {
            jsonEditor.setValue(result)
            generate()
        })
    }

    function copyCode(): void {
        WriteClipboard(goEditor.getValue())
    }

    let items = $locales.map(function (value) {
        return {
            name: value,
            value: value,
        }
    })


</script>

<main class="w-screen h-screen flex flex-col">
    <MenuBar class="border-b">
        <MenuBarItem>
            {$_('File')}
            <svelte:fragment slot="flyout">
                <MenuFlyoutItem>{$_('Load from JSON file')}</MenuFlyoutItem>
                <MenuFlyoutItem>{$_('Save to Go source file')}</MenuFlyoutItem>
                <MenuFlyoutDivider/>
                <MenuFlyoutItem>{$_('Settings')}</MenuFlyoutItem>
                <MenuFlyoutDivider/>
                <MenuFlyoutItem on:click={()=> EventsEmit("exit")}>{$_('Exit')}</MenuFlyoutItem>
            </svelte:fragment>
        </MenuBarItem>
        <MenuBarItem>
            {$_('Language')}
            <svelte:fragment slot="flyout">
                {#each $locales as _locale}
                    <MenuFlyoutItem variant="radio" bind:group={$locale} value={_locale}>{$_(_locale)}</MenuFlyoutItem>
                {/each}
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
</style>
