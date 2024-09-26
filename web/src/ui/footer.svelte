<script lang="ts">
    import {getContext} from "svelte";
    import {_} from "svelte-i18n";
    import {Button} from "bits-ui";
    import Message from "./message.svelte";
    import type {EditorsState} from "../state/editors.svelte";
    import type {MessageState} from "../state/message.svelte";

    const editors = getContext<EditorsState>("editors");
    const message = getContext<MessageState>("message");

    function handleClick() {
        message.clear();
        editors.json2go();
    }
</script>

<div class="flex flex-row px-4 py-2 justify-end items-center h-12">
    <Message/>
    <Button size="sm" on:click={handleClick} class="mr-2">
        {$_("Generate")}
    </Button>
</div>
