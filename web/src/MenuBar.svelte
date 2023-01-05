<script lang="ts">
    import {MenuBar, MenuBarItem, MenuFlyoutDivider, MenuFlyoutItem} from "fluent-svelte";
    import {_, locale, locales} from "svelte-i18n";
    import {BrowserOpenURL, EventsEmit} from "../wailsjs/runtime";
    import SettingsDialog from "./SettingsDialog.svelte";
    import AboutDialog from "./AboutDialog.svelte";
    import {OpenJSONFile, SaveGoSourceFile} from "../wailsjs/go/main/App";
    import {Editors, Layout} from "./base.js";
    import {editor} from "monaco-editor/esm/vs/editor/editor.api";
    import IStandaloneCodeEditor = editor.IStandaloneCodeEditor;

    let openSettingsDialog = false
    let openAboutDialog = false

    export let layout: Layout
    export let editors: Editors

    export let defaultFontSize = 16
    export let fontSize: number

    export let jsonEditor: IStandaloneCodeEditor
    export let goEditor: IStandaloneCodeEditor

    export let allCapsWords = [] as string[]

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
</script>

<SettingsDialog bind:open={openSettingsDialog} bind:allCapsWords></SettingsDialog>
<AboutDialog bind:open={openAboutDialog}></AboutDialog>
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
            <MenuFlyoutItem cascading>
                {$_('Editors')}
                <svelte:fragment slot="flyout">
                    <MenuFlyoutItem variant="radio" bind:group={editors} name="editors"
                                    value={Editors.Both}>{$_(Editors.Both)}</MenuFlyoutItem>
                    <MenuFlyoutItem variant="radio" bind:group={editors} name="editors"
                                    value={Editors.JSON}>{$_(Editors.JSON)}</MenuFlyoutItem>
                    <MenuFlyoutItem variant="radio" bind:group={editors} name="editors"
                                    value={Editors.Go}>{$_(Editors.Go)}</MenuFlyoutItem>
                </svelte:fragment>
            </MenuFlyoutItem>
            <MenuFlyoutItem cascading>
                {$_('Layout')}
                <svelte:fragment slot="flyout">
                    <MenuFlyoutItem variant="radio" bind:group={layout} name="layout"
                                    value={Layout.TwoColumns}>{$_(Layout.TwoColumns)}</MenuFlyoutItem>
                    <MenuFlyoutItem variant="radio" bind:group={layout} name="layout"
                                    value={Layout.TwoRows}>{$_(Layout.TwoRows)}</MenuFlyoutItem>
                </svelte:fragment>
            </MenuFlyoutItem>
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
        {$_('Language')}
        <svelte:fragment slot="flyout">
            {#each $locales as _locale}
                <MenuFlyoutItem variant="radio" bind:group={$locale} name="locale" value={_locale}
                                checked={$locale===_locale}>{$_(_locale)}</MenuFlyoutItem>
            {/each}
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

<style>
</style>