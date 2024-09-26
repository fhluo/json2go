<script lang="ts">
    import { getContext } from "svelte";
    import { Cross1Icon, InfoCircledIcon } from "@radix-ui/react-icons";
    import type { MessageState } from "../state/message.svelte";

    const messages = getContext<MessageState>("messages");

    function clearMessage() {
        messages.clear();
    }
</script>

{#if messages.message}
    <div class="flex flex-row items-center justify-center select-none mx-4 border bg-white/50 rounded shadow-sm space-x-1.5">
        <InfoCircledIcon class="w-5 h-5 ml-2 mr-1 text-red-600" />
        <span>{messages.message}</span>
        <button type="button" on:click={clearMessage} class="hover:bg-gray-200/50 py-1 px-1 rounded-r transition flex items-center justify-center">
            <Cross1Icon class="w-5 h-5" />
        </button>
    </div>
{/if}
