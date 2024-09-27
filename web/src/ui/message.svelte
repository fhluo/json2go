<script lang="ts">
    import { getContext } from "svelte";
    import { X, Info } from "lucide-svelte";
    import type { MessageState } from "../state/message.svelte";

    const message = getContext<MessageState>("message");

    function clearMessage() {
        message.clear();
    }
</script>

{#if message.message}
    <div class="flex flex-row items-center justify-center select-none mx-4 border bg-white/50 rounded shadow-sm space-x-1.5">
        <Info class="w-5 h-5 ml-2 mr-1 text-red-600" />
        <span>{message.message}</span>
        <button type="button" on:click={clearMessage} class="hover:bg-gray-200/50 py-1 px-1 rounded-r transition flex items-center justify-center">
            <X class="w-5 h-5" />
        </button>
    </div>
{/if}
