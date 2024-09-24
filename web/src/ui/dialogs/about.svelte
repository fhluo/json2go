<script lang="ts">
    import {getContext} from "svelte";
    import {type ConfigState, openHomePage} from "$lib/api.svelte";
    import {_} from "svelte-i18n"
    import {DialogContent, DialogFooter, DialogHeader, DialogTitle} from "$lib/components/ui/dialog";
    import {Tooltip, TooltipContent, TooltipTrigger} from "$lib/components/ui/tooltip";
    import {Button} from "$lib/components/ui/button";

    let config = getContext<ConfigState>("config")
</script>

<DialogContent class="select-none">
    <DialogHeader>
        <DialogTitle>{$_("about.title", {default: "About"})}</DialogTitle>
    </DialogHeader>
    <div class="flex flex-col items-center justify-center space-y-2 text-gray-900">
        <Tooltip>
            <TooltipTrigger asChild>
                <p class="text-lg font-semibold" ondblclick={openHomePage}>
                    json2go
                </p>
            </TooltipTrigger>
            <TooltipContent class="bg-transparent text-slate-600">
                <p>
                    {$_("about.tip", {default: "Double-click to visit json2go on Github.com"})}
                </p>
            </TooltipContent>
        </Tooltip>
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
    <DialogFooter>
        <Button variant="secondary">{$_("OK")}</Button>
    </DialogFooter>
</DialogContent>
