<script lang="ts">
    import {Button, ContentDialog} from "fluent-svelte"
    import {_} from "svelte-i18n"
    import {GetVersion} from "../wailsjs/go/main/App"

    export let open = false

    let version: string
    GetVersion().then((v) => {
        version = v
    })
</script>

<ContentDialog bind:open={open} title={$_('about.title', {default: 'About'})}>
    <div class="flex flex-col space-y-2 items-center justify-center">
        <p class="text-lg font-semibold">JSON2Go</p>
        <p class="text-gray-900">{$_('about.about', {default: "Generate Go type definitions from JSON"})}</p>
        <div class="leading-relaxed">
            <p class="text-gray-900">
                <span class="font-semibold">{$_('about.license', {default: 'License: '})}</span>MIT
            </p>
            <p class="text-gray-900">
                <span class="font-semibold">{$_('about.version', {default: 'Version: '})}</span>{version}
            </p>
        </div>
        <p class="text-gray-900">Copyright © 2022 fhluo</p>
    </div>
    <div class="flex justify-end mt-3.5">
        <Button variant="accent" on:click={() => {open=false}} class="mr-2">{$_('OK')}</Button>
    </div>
</ContentDialog>
