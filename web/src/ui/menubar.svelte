<script lang="ts">
    import { getContext } from "svelte";
    import { Dialog } from "bits-ui";
    import { Menubar } from "bits-ui";
    import type { DialogState } from "../state/dialog.svelte";

    const dialog = getContext<DialogState>("dialog");
</script>

<Dialog>
    <Menubar class="rounded-none border-none bg-transparent">
        <slot></slot>
    </Menubar>
    {#if dialog}
        {dialog}
    {/if}
</Dialog>
