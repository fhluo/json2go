<script lang="ts">
    import {getContext} from "svelte";
    import {_} from "svelte-i18n";
    import {Menubar} from "bits-ui";
    import {EventsEmit} from "../../../wailsjs/runtime";
    import Settings from "@/ui/dialogs/settings.svelte";
    import type {EditorsState} from "../../state/editors.svelte";

    const editors = getContext<EditorsState>("editors");
</script>

<Menubar.Menu>
    <Menubar.Trigger>{$_("file.title", {default: "File"})}</Menubar.Trigger>
    <Menubar.Content>
        <Menubar.Item on:click={editors.openJSONFile}>
            {$_("file.openJSON", {default: "Open JSON file"})}
        </Menubar.Item>
        <Menubar.Item on:click={editors.saveGoFile}>
            {$_("file.saveGo", {default: "Save Go source file"})}
        </Menubar.Item>
        <Menubar.Separator/>
        <Menubar.Item on:click={()=> {
            // Open the settings dialog
         }}>
            {$_("file.settings", {default: "Settings"})}
        </Menubar.Item>
        <Menubar.Separator/>
        <Menubar.Item on:click={() => EventsEmit("exit")}>
            {$_("file.exit", {default: "Exit"})}
        </Menubar.Item>
    </Menubar.Content>
</Menubar.Menu>
