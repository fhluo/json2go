<script lang="ts">
    import {Generate, ReadClipboard, WriteClipboard} from '../wailsjs/go/main/App.js'
    import Editor from "./Editor.svelte";
    import "fluent-svelte/theme.css";
    import {Button, MenuBar, MenuBarItem, MenuFlyoutDivider, MenuFlyoutItem, TextBlock, TextBox} from "fluent-svelte";
    import "./i18n"
    import {_, locale, locales} from "svelte-i18n";
    import {EventsEmit} from "../wailsjs/runtime";


    let input: { text: string }
    let output: { text: string }

    let allCaps: string = "ID"

    function generate(): void {
        Generate(input.text, allCaps.split(",").map((v) => v.trim())).then(result => {
            output.text = result
        })
    }

    function pasteJSON() {
        ReadClipboard().then(result => {
            input.text = result
            generate()
        })
    }

    function copyCode(): void {
        WriteClipboard(output.text)
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
                <MenuFlyoutDivider/>
                <MenuFlyoutItem>{$_('Settings')}</MenuFlyoutItem>
                <MenuFlyoutDivider/>
                <MenuFlyoutItem on:click={()=> EventsEmit("exit")}>{$_('Exit')}</MenuFlyoutItem>
            </svelte:fragment>
        </MenuBarItem>
        <MenuBarItem>{$_('Edit')}</MenuBarItem>
        <MenuBarItem>
            {$_('Language')}
            <svelte:fragment slot="flyout">
                {#each $locales as _locale}
                    <MenuFlyoutItem variant="radio" bind:group={$locale} value={_locale}>{$_(_locale)}</MenuFlyoutItem>
                {/each}
            </svelte:fragment>
        </MenuBarItem>
    </MenuBar>

    <div class="flex flex-col w-full h-full px-6 py-4 space-y-4 bg-white/50">
        <div class="columns-2 gap-6 mt-4 h-full max-h-full mx-2">
            <div class="code h-full max-h-full flex flex-col overflow-clip">
                <div class="w-full bg-white/50 border-b border-t flex flex-row">
                    <span class="py-1 px-4 select-none text-yellow-600 font-mono">JSON</span>
                    <button on:click={pasteJSON} class="text-sm">{$_('Paste')}</button>
                </div>
                <Editor class="w-full h-full border-b" lang="json" bind:editor={input}/>
            </div>
            <div class="code h-full max-h-full flex flex-col overflow-clip">
                <div class="w-full bg-white/50 border-b border-t flex flex-row">
                    <span class="py-1 px-4 select-none text-purple-600 font-mono">Go</span>
                    <button on:click={copyCode} class="text-sm">{$_('Copy')}</button>
                </div>
                <Editor class="w-full h-full border-b" lang="go" bind:editor={output}/>
            </div>
        </div>

        <div class="flex flex-col space-y-2 mx-2 border rounded-lg px-4 py-3 bg-white/50">
            <div class="flex flex-row items-center gap-x-5">
                <div class="space-x-1.5 flex-grow flex flex-row items-center">
                    <TextBlock class="select-none">{$_('All caps')}</TextBlock>
                    <TextBox id="allCaps" class="flex-grow" bind:value={allCaps}></TextBox>
                </div>
                <Button variant="accent" on:click={generate}>{$_('Generate')}</Button>
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
</style>
