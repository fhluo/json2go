import {
	useFontSizeStore,
	useLanguageStore,
	useRealTimeStore,
	useValidJSONStore,
} from "@/lib/store.ts";
import { useEditorsStore } from "@/store/editors.ts";
import { useMessageStore } from "@/store/message.ts";
import Container from "@/ui/container.tsx";
import Footer from "@/ui/footer.tsx";
import GoContainer from "@/ui/go.tsx";
import JSONContainer from "@/ui/json.tsx";
import Menubar from "@/ui/menubar.tsx";
import ExamplesMenu from "@/ui/menus/examples.tsx";
import FileMenu from "@/ui/menus/file.tsx";
import FontMenu from "@/ui/menus/font.tsx";
import HelpMenu from "@/ui/menus/help.tsx";
import LanguageMenu from "@/ui/menus/language.tsx";
import OptionsMenu from "@/ui/menus/options.tsx";
import ViewMenu from "@/ui/menus/view.tsx";
import { useEffect, useRef } from "react";
import { EventsEmit, EventsOn } from "../wailsjs/runtime";
import "./app.css";

function App() {
	const { fontSize, init: initFontSize } = useFontSizeStore();
	const { language, init: initLanguage } = useLanguageStore();

	const { init: initValidJSON } = useValidJSONStore();
	const { realTime, init: initRealTime } = useRealTimeStore();

	const setMessage = useMessageStore((state) => state.setMessage);

	const init = useRef(false);

	const {
		jsonEditor,
		goEditor,
		init: initEditors,
		generate,
	} = useEditorsStore();

	useEffect(() => {
		initLanguage();
		initFontSize();
		initValidJSON();
		initRealTime();

		document.defaultView?.addEventListener("resize", () => {
			jsonEditor?.layout();
			goEditor?.layout();
			EventsEmit("resize");
		});

		EventsOn("error", (message: string) => setMessage(message));

		init.current = true;
	}, []);

	useEffect(() => {
		if (language) {
			initEditors(language, fontSize, "json-editor", "go-editor");
		}
	}, [language]);

	useEffect(() => {
		if (init.current) {
			jsonEditor?.updateOptions({ fontSize });
			goEditor?.updateOptions({ fontSize });
		}
	}, [fontSize]);

	useEffect(() => {
		if (jsonEditor) {
			jsonEditor.getModel()?.onDidChangeContent(() => {
				if (realTime) {
					generate();
				}
			});
		}
	}, [jsonEditor]);

	return (
		<main className="w-screen h-screen flex flex-col">
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
				<JSONContainer />
				<GoContainer />
			</Container>
			<Footer />
		</main>
	);
}

export default App;
