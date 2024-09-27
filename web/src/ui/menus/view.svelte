<script lang="ts">
    import {getContext} from "svelte";
    import {_} from "svelte-i18n";
    import {Menubar} from "bits-ui";
    import  {type UIState, Layout, View} from "../../state/ui.svelte";

    const ui = getContext<UIState>("ui");

    function setView(view: View) {
        ui.view = view;
    }

    function setLayout(layout: Layout) {
        ui.layout = layout;
    }
</script>

<Menubar.Menu>
    <Menubar.Trigger>{$_("view.title", {default: "View"})}</Menubar.Trigger>
    <Menubar.Content>
        <Menubar.Sub>
            <Menubar.SubTrigger>{$_("view.editors", {default: "Editors"})}</Menubar.SubTrigger>
            <Menubar.SubContent>
                <Menubar.RadioGroup value={ui.view} on:change={(event) => setView(event.detail)}>
                    <Menubar.RadioItem value={View.JSONAndGo}>
                        {$_("view.both", {default: "JSON and Go"})}
                    </Menubar.RadioItem>
                    <Menubar.RadioItem value={View.JSONOnly}>
                        {$_("view.jsonOnly", {default: "JSON"})}
                    </Menubar.RadioItem>
                    <Menubar.RadioItem value={View.GoOnly}>
                        {$_("view.goOnly", {default: "Go"})}
                    </Menubar.RadioItem>
                </Menubar.RadioGroup>
            </Menubar.SubContent>
        </Menubar.Sub>
        <Menubar.Sub>
            <Menubar.SubTrigger>{$_("view.layout", {default: "Layout"})}</Menubar.SubTrigger>
            <Menubar.SubContent>
                <Menubar.RadioGroup value={ui.layout} on:change={(event) => setLayout(event.detail)}>
                    <Menubar.RadioItem value={Layout.TwoColumns}>
                        {$_("view.columns", {default: "Two Columns"})}
                    </Menubar.RadioItem>
                    <Menubar.RadioItem value={Layout.TwoRows}>
                        {$_("view.rows", {default: "Two Rows"})}
                    </Menubar.RadioItem>
                </Menubar.RadioGroup>
            </Menubar.SubContent>
        </Menubar.Sub>
    </Menubar.Content>
</Menubar.Menu>
