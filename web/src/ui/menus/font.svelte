<script lang="ts">
    import { getContext } from "svelte";
    import { _ } from "svelte-i18n";
    import { Menubar } from "bits-ui";
    import type { ConfigState } from "src/lib/api.svelte";

    const config = getContext<ConfigState>("config");
</script>

<Menubar.Menu>
    <Menubar.Trigger>{$_("font.title", { default: "Font" })}</Menubar.Trigger>
    <Menubar.Content>
        <Menubar.Item on:click={config.increaseFontSize}>
            {$_("font.increase", { default: "Increase size" })}
        </Menubar.Item>
        <Menubar.Item on:click={config.decreaseFontSize}>
            {$_("font.decrease", { default: "Decrease size" })}
        </Menubar.Item>
        <Menubar.Item on:click={config.resetFontSize}>
            {$_("font.reset", { default: "Reset size" })}
        </Menubar.Item>
    </Menubar.Content>
</Menubar.Menu>
