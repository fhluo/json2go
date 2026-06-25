import { MenubarContent, MenubarItem, MenubarMenu, MenubarTrigger } from "@/components/ui/menubar.tsx";
import { useEditorsStore } from "@/store/editors.ts";
import { useEffect, useState } from "react";
import { useTranslation } from "react-i18next";
import { Examples } from "@api/app/services";
import type { Example } from "@api/internal/examples/models";

export default function ExamplesMenu() {
    const { t } = useTranslation();
    const setJSON = useEditorsStore((state) => state.setJSON);

    const [exampleList, setExampleList] = useState([] as Example[]);
    useEffect(() => {
        void Examples.All().then((value) => {
            setExampleList(value ?? []);
        });
    }, []);

    return (
        <MenubarMenu>
            <MenubarTrigger>{t("examples.title", "Examples")}</MenubarTrigger>
            <MenubarContent>
                {exampleList.map((example) => (
                    <MenubarItem key={example.title} onClick={() => setJSON(example.content)}>
                        {example.title}
                    </MenubarItem>
                ))}
            </MenubarContent>
        </MenubarMenu>
    );
}
