<script lang="ts">
    import {getContext} from "svelte";
    import {Button, Dialog} from "bits-ui";
    import {_} from "svelte-i18n";
    import type {ConfigState} from "src/lib/api.svelte";
    import {openRelease} from "src/lib/api.svelte";
    import {GetLatestVersion, GetVersion} from "../../../wailsjs/go/main/App";

    let config = getContext<ConfigState>("config");
    let version = $state("");
    let latestVersion = $state("");
    let message = $state("");
    let isChecking = $state(true);

    $effect(() => {
        Promise.all([GetVersion(), GetLatestVersion()]).then(
            ([version_, latestVersion_]) => {
                version = version_;
                latestVersion = latestVersion_;
                isChecking = false;
            },
        );
    })

    $effect(() => {
        if (isChecking) {
            message = $_("updates.checking", {default: "Checking..."});
        } else if (version && latestVersion && version === latestVersion) {
            // todo: update key updates.none to updates.latest in translations file
            message = $_("updates.latest", {default: "You are using the latest version."});
        } else if (version && latestVersion && version !== latestVersion) {
            // todo: update key updates.available to updates.new in translations file
            message = $_("updates.new", {default: "A new version is available: "}) + `v${latestVersion}`;
        } else {
            message = $_("updates.failed", {default: "Unable to check for updates."});
        }
    })

</script>

<Dialog.Content class="select-none">
    <Dialog.Header>
        <Dialog.Title>{$_("updates.title", {default: "Check for updates"})}</Dialog.Title>
    </Dialog.Header>
    <p class="text-gray-900 leading-relaxed">{message}</p>
    <Dialog.Footer>
        {#if version && latestVersion && version !== latestVersion}
            <Button variant="secondary" on:click={() => openRelease(latestVersion)}>
                {$_("Download")}
            </Button>
        {/if}
        <Dialog.Close asChild>
            <Button variant="secondary">{$_("OK")}</Button>
        </Dialog.Close>
    </Dialog.Footer>
</Dialog.Content>
