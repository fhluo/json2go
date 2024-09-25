<script lang="ts">
    import {getContext} from "svelte";
    import {Button, Dialog} from "bits-ui";
    import {_} from "svelte-i18n";
    import type {ConfigState} from "src/lib/api.svelte";

    let config = getContext<ConfigState>("config");
    let input = $state("");

    function onBeforeInput(event: InputEvent) {
        // user can only enter letters, space and ','
        const value = (event.target as HTMLInputElement).value;
        if (value && !value.match(/[a-zA-Z', ]/)) {
            event.preventDefault();
        }
    }

    function addWord() {
        config.addAllCapsWords(input);
        input = "";
    }

    function removeWord() {
        config.removeAllCapsWords(input);
        input = "";
    }

    function handleKeyDown(event: KeyboardEvent) {
        if (event.key === "Enter") {
            addWord();
        }
    }
</script>

<Dialog.Content>
    <Dialog.Header>
        <Dialog.Title>{$_("settings.title", {default: "Settings"})}</Dialog.Title>
    </Dialog.Header>
    <div class="flex flex-col space-y-2">
        <span class="select-none font-semibold mr-3">
            {$_("settings.all-caps", {default: "All-caps words"})}
        </span>
        {#if config.allCapsWords?.length > 0}
            <div class="flex flex-row flex-wrap space-x-1.5 pb-1.5">
                {#each config.allCapsWords as word}
                    <Button size="sm" variant="outline" class="ml-1 mt-1.5 hover:bg-gray-50 transition cursor-default"
                            on:click={() => (input = word)} on:dblclick={() => config.removeAllCapsWords(input)}>
                        {word}
                    </Button>
                {/each}
            </div>
        {/if}
        <div>
            <input bind:value={input} on:beforeinput={onBeforeInput} on:keydown={handleKeyDown}/>
            <p class="text-sm text-gray-500 leading-relaxed py-2 px-1 select-none">
                {$_("settings.tip", {default: "Tip: Double click a word to remove it. To add multiple words, separate words with commas."})}
            </p>
        </div>
    </div>
    <Dialog.Footer>
        <Button variant="secondary" class="min-w-fit" on:click={() => addWord()}>
            {$_("settings.add", {default: "Add"})}
        </Button>
        <Button variant="secondary" class="min-w-fit" on:click={() => removeWord()}>
            {$_("settings.remove", {default: "Remove"})}
        </Button>
        <Dialog.Close asChild>
            <Button variant="secondary">{$_("OK")}</Button>
        </Dialog.Close>
    </Dialog.Footer>
</Dialog.Content>
