<script lang="ts">
    import { getContext } from "svelte";
    import { _ } from "svelte-i18n";
    import { Menubar } from "bits-ui";
    import type { EditorsState } from "../../state/editors.svelte";
    import {Examples} from "@api/app/services";
    import {Example} from "@api/internal/examples";

    let exampleList: Example[] = [];
    const editors = getContext<EditorsState>("editors");

    $effect(() => {
        Examples.All().then((examples) => {
            exampleList = examples;
        });
    });

    function setJSON(content: string) {
        editors.json = content;
    }
</script>

<Menubar.Menu>
    <Menubar.Trigger>{$_("examples.title", { default: "Examples" })}</Menubar.Trigger>
    <Menubar.Content>
        {#each exampleList as example}
            <Menubar.Item key={example.title} on:click={() => setJSON(example.content)}>
                {example.title}
            </Menubar.Item>
        {/each}
    </Menubar.Content>
</Menubar.Menu>
