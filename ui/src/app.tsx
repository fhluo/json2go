import { useConfigStore } from "@/lib/api.ts";
import { useEditorsStore } from "@/store/editors.ts";
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
import { toast } from "sonner";
import { Toaster } from "@/components/ui/sonner";
import "./app.css";
import { Events } from "@wailsio/runtime";

function App() {
    const fontSize = useConfigStore((s) => s.fontSize);
    const language = useConfigStore((s) => s.language);
    const realTime = useConfigStore((s) => s.realTime);
    const initConfig = useConfigStore((s) => s.init);

    const init = useRef(false);

    const { jsonEditor, goEditor, init: initEditors, generate } = useEditorsStore();

    const realTimeRef = useRef(realTime);
    realTimeRef.current = realTime;
    const generateRef = useRef(generate);
    generateRef.current = generate;

    useEffect(() => {
        void initConfig();

        document.defaultView?.addEventListener("resize", () => {
            jsonEditor?.layout();
            goEditor?.layout();
            void Events.Emit("resize");
        });

        Events.On("error", (ev) => toast.error(ev.data));

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
                if (realTimeRef.current) {
                    generateRef.current();
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
            <Toaster position="top-right" />
        </main>
    );
}

export default App;
