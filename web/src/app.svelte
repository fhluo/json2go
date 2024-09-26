<script lang="ts">
    import { onMount } from "svelte";
    import { getContext, setContext } from "svelte";
    import { EventsEmit, EventsOn } from "../wailsjs/runtime";
    import { createConfigState } from "./lib/api.svelte";
    import { createEditorsState } from "./state/editors.svelte";
    import { createUIState } from "./state/ui.svelte";
    import Menubar from "./ui/menubar.svelte";
    import FileMenu from "./ui/menus/file.svelte";
    import OptionsMenu from "./ui/menus/options.svelte";
    import ViewMenu from "./ui/menus/view.svelte";
    import FontMenu from "./ui/menus/font.svelte";
    import LanguageMenu from "./ui/menus/language.svelte";
    import ExamplesMenu from "./ui/menus/examples.svelte";
    import HelpMenu from "./ui/menus/help.svelte";
    import Container from "./ui/container.svelte";
    import Footer from "./ui/footer.svelte";
    import GoContainer from "./ui/go.svelte";
    import JSONContainer from "./ui/json.svelte";
    import {createMessageState} from "./state/message.svelte";

    const config = createConfigState();
    const ui = createUIState();

    let goElement: HTMLElement;
    let jsonElement: HTMLElement;

    const editors = createEditorsState({
        language: config.locale,
        fontSize: config.fontSize,
        goElement: goElement,
        jsonElement: jsonElement,
    });
    const message = createMessageState()

    setContext("config", config);
    setContext("ui", ui);
    setContext("editors", editors);
    setContext("message", message);

   $effect(() => {
        config.init();
        editors.init();


       document.defaultView?.addEventListener("resize", () => {
            editors.jsonEditor?.layout();
            editors.goEditor?.layout();
            EventsEmit("resize");
        });

        EventsOn("error", (msg: string) => {
            message.message = msg;
        });
    });

    $effect(() => {
        if (config.locale) {
            editors.init();
        }
    });

    $effect(() =>{
        if (config.fontSize) {
            editors.jsonEditor?.updateOptions({ fontSize: value });
            editors.goEditor?.updateOptions({ fontSize: value });
        }
    })

    $effect(() => {
        if (editors.jsonEditor) {
            editors.jsonEditor.getModel()?.onDidChangeContent(() => {
                if (config.realTime) {
                    editors.json2go();
                }
            });
        }
    })
</script>

<main class="w-screen h-screen flex flex-col">
    <Menubar>
        <FileMenu />
        <OptionsMenu />
        <ViewMenu />
        <FontMenu />
        <LanguageMenu />
        <ExamplesMenu />
        <HelpMenu />
    </Menubar>
    <Container>
        <JSONContainer bind:goElement />
        <GoContainer bind:jsonElement/>
    </Container>
    <Footer />
</main>
