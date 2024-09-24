import { mount } from "svelte";
import "../../json2go-web2/src/app.css";
import App from "./app.svelte";

const app = mount(App, {
	target: document.getElementById("app")!,
});

export default app;
