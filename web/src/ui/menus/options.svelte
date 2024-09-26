<script lang="ts">
    import { getContext } from "svelte";
    import { _ } from "svelte-i18n";
    import { Menubar } from "bits-ui";
    import type { ConfigState } from "src/lib/api.svelte";

    const config = getContext<ConfigState>("config");

    function toggleValidJSON() {
        config.validJSON = !config.validJSON;
    }

    function toggleRealTime() {
        config.realTime = !config.realTime;
    }
</script>

<Menubar.Menu>
    <Menubar.Trigger>{$_("options.title", { default: "Options" })}</Menubar.Trigger>
    <Menubar.Content>
        <Menubar.CheckboxItem checked={config.validJSON} on:checkedChange={toggleValidJSON}>
            {$_("options.valid", { default: "Validate JSON before generation" })}
        </Menubar.CheckboxItem>
        <Menubar.CheckboxItem checked={config.realTime} on:checkedChange={toggleRealTime}>
            {$_("options.realTime", { default: "Generate in real time" })}
        </Menubar.CheckboxItem>
    </Menubar.Content>
</Menubar.Menu>
