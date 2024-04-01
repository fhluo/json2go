<script lang="ts">
import { createConfigState } from "$lib/api.svelte";
import { setContext } from "svelte";
import { createEditorsState } from "../../json2go-web2/src/state/editors.svelte.js";
import { createUIState } from "../../json2go-web2/src/state/ui.svelte.js";

setContext("ui", createUIState());
const config = setContext("config", createConfigState());
setContext(
	"editors",
	createEditorsState({
		language: config.locale,
		fontSize: config.fontSize,
		goElement: undefined,
		jsonElement: undefined,
	}),
);
</script>

<main>

</main>

<style>

</style>
