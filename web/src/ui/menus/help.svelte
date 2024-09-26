<script lang="ts">
    import {getContext} from "svelte";
    import {_} from "svelte-i18n";
    import {Menubar, Dialog} from "bits-ui";
    import About from "../dialogs/about.svelte";
    import Updates from "../dialogs/updates.svelte";
    import type {DialogState} from "../../state/dialog.svelte";

    const dialog = getContext<DialogState>("dialog");

    function openAbout() {
        dialog.dialog = About;
    }

    function openUpdates() {
        dialog.dialog = Updates;
    }
</script>

<Menubar.Menu>
    <Menubar.Trigger>{$_("help.title", {default: "Help"})}</Menubar.Trigger>
    <Menubar.Content>
        <Dialog.Trigger asChild>
            <Menubar.Item on:click={openUpdates}>
                {$_("help.updates", {default: "Check for updates"})}
            </Menubar.Item>
        </Dialog.Trigger>
        <Menubar.Separator/>
        <Dialogs.Trigger asChild>
            <Menubar.Item on:click={openAbout}>
                {$_("help.about", {default: "About"})}
            </Menubar.Item>
        </Dialogs.Trigger>
    </Menubar.Content>
</Menubar.Menu>
