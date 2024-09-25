<script lang="ts">
    import {getContext} from "svelte";
    import {type ConfigState, openHomePage} from "$lib/api.svelte";
    import {_} from "svelte-i18n"
    import {Dialog, Button, Tooltip} from "bits-ui";

    let config = getContext<ConfigState>("config")
</script>

<Dialog.Content class="select-none">
    <Dialog.Header>
        <Dialog.Title>{$_("about.title", {default: "About"})}</Dialog.Title>
    </Dialog.Header>
    <div class="flex flex-col items-center justify-center space-y-2 text-gray-900">
        <Tooltip.Root>
            <Tooltip.Trigger asChild>
                <p class="text-lg font-semibold" ondblclick={openHomePage}>
                    json2go
                </p>
            </Tooltip.Trigger>
            <Tooltip.Content class="bg-transparent text-slate-600">
                <p>
                    {$_("about.tip", {default: "Double-click to visit json2go on Github.com"})}
                </p>
            </Tooltip.Content>
        </Tooltip.Root>
        <p class="text-sm leading-relaxed text-muted-foreground">
            {$_("about.description", {default: "Generate Go type definitions from JSON"})}
        </p>
        <div class="leading-relaxed text-sm">
            <p>
                <span class="select-none">
                    {$_("about.license", {default: "License: "})}
                </span>
                MIT
            </p>
            <p>
                <span class="select-none">
                    {$_("about.version", {default: "Version: "})}
                </span>
                {config.version}
            </p>
        </div>
        <p class="text-sm">Copyright © 2022 fhluo</p>
    </div>
    <Dialog.Footer>
        <Button variant="secondary">{$_("OK")}</Button>
    </Dialog.Footer>
</Dialog.Content>
