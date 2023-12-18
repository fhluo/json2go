<script lang="ts">
    import {Button, ContentDialog} from "fluent-svelte"
    import {_} from "svelte-i18n"
    import {GetLatestVersion, GetVersion} from "../../wailsjs/go/main/App"
    import {BrowserOpenURL} from "../../wailsjs/runtime"

    export let open = false

    let version: string
    GetVersion().then(v => version = v)

    let latestVersion: string

    let failed = false
    GetLatestVersion().then(v => {
        if (v) {
            latestVersion = v
        } else {
            failed = true
        }
    })
</script>

<ContentDialog bind:open={open} title={$_('checkForUpdates.title', {default: 'Check for updates'})}>
    {#if version && latestVersion}
        {#if version === latestVersion}
            <p class="text-gray-900 leading-relaxed">
                {$_('checkForUpdates.upToDate', {default: 'You are using the latest version.'})}
            </p>
        {:else}
            <p class="text-gray-900 leading-relaxed">
                {$_('checkForUpdates.newVersion', {default: 'A new version is available: '})}
                <button class="hover:underline hover:cursor-pointer text-blue-600 tracking-wide"
                        on:click={()=>BrowserOpenURL("https://github.com/fhluo/json2go/releases/tag/v"+latestVersion)}>
                    v{latestVersion}</button>
                {$_('.')}
            </p>
        {/if}
    {:else}
        {#if failed}
            <p class="text-gray-900 leading-relaxed">
                {$_('checkForUpdates.failed', {default: 'Unable to check for updates.'})}
            </p>
        {:else}
            <p class="text-gray-900 leading-relaxed">
                {$_('checkForUpdates.checking', {default: 'Checking...'})}
            </p>
        {/if}
    {/if}
    <div class="flex justify-end mt-3.5">
        <Button variant="accent" on:click={() => {open=false}} class="mr-2">{$_('OK')}</Button>
    </div>
</ContentDialog>
