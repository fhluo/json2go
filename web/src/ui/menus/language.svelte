<script lang="ts">
    import {getContext} from "svelte";
    import {_, locales} from "svelte-i18n";
    import {Menubar} from "bits-ui";
    import type {ConfigState} from "src/lib/api.svelte";

    const config = getContext<ConfigState>("config");

    function setLanguage(language: string) {
        config.locale = language;
    }
</script>

<Menubar.Menu>
    <Menubar.Trigger>{$_("language.title", {default: "Language"})}</Menubar.Trigger>
    <Menubar.Content>
        <Menubar.RadioGroup value={config.locale} on:change={(event) => setLanguage(event.detail)}>
            {#each $locales as language}
                <Menubar.RadioItem key={language} value={language}>
                    {$_(`language.${language}`)}
                </Menubar.RadioItem>
            {/each}
        </Menubar.RadioGroup>
    </Menubar.Content>
</Menubar.Menu>
