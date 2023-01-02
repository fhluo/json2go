<script lang="ts">
    import {Button, ContentDialog, TextBox} from "fluent-svelte";
    import {_} from "svelte-i18n";
    import {GetAllCapsWords, SetAllCapsWords} from "../wailsjs/go/main/App";

    export let open = false;

    export let allCapsWords = [] as string[]

    let gotAllCapsWords = false;

    GetAllCapsWords().then(result => {
        if (result != null) {
            allCapsWords = result
            gotAllCapsWords = true
        }
    })

    $: {
        if (gotAllCapsWords) {
            SetAllCapsWords(allCapsWords)
        }
    }

    let allCapsWord = ""

    function addAllCapsWord(): void {
        allCapsWord = allCapsWord.trim()
        if (allCapsWord === "") {
            return
        }

        // if allCapsWord contains a comma, split it and add each word separately
        allCapsWords = Array.from(new Set(allCapsWords.concat(allCapsWord.split(',').map(word => word.trim()).filter(word => word !== ""))))
        allCapsWord = ""
    }

    function onBeforeInput(event: InputEvent): void {
        // user can only enter letters, space and ','
        if (event.data !== null && !event.data.match(/[a-zA-Z', ]/)) {
            event.preventDefault()
        }
    }
</script>

<ContentDialog bind:open={open} title={$_('Settings')}>
    <div class="flex flex-col space-y-2">
        <span class="select-none font-semibold mr-3">{$_('All-caps words')}</span>
        {#if allCapsWords?.length > 0}
            <div class="flex flex-row flex-wrap space-x-1.5 pb-1.5">
                {#each allCapsWords as word}
                    <button class="px-3 leading-loose tracking-wide rounded border shadow-sm ml-1 mt-1.5 bg-white hover:bg-gray-50 transition cursor-default"
                            on:dblclick={() => allCapsWords = allCapsWords.filter(w => w !== word)}
                            on:click={() => allCapsWord = word}>{word}</button>
                {/each}
            </div>
        {/if}
        <div class="flex flex-row space-x-1.5">
            <TextBox bind:value={allCapsWord} on:beforeinput={onBeforeInput}></TextBox>
            <Button on:click={addAllCapsWord} class="min-w-fit">{$_('Add')}</Button>
            <Button on:click={() => allCapsWords = allCapsWords.filter(w => w !== allCapsWord)}
                    class="min-w-fit">{$_('Remove')}</Button>
        </div>
        <p class="text-sm text-gray-500 leading-relaxed py-2 px-1">{$_('tip.allCaps', {default: "Tip: Double click a word to remove it. To add multiple words, separate words with commas."})}</p>
    </div>
    <div class="flex justify-end">
        <Button variant="accent" on:click={() => {open = false}} class="mr-2">{$_('OK')}</Button>
    </div>
</ContentDialog>

<style>
</style>