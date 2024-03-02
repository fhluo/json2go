import {
	MenubarContent,
	MenubarItem,
	MenubarMenu,
	MenubarTrigger,
} from "@/components/ui/menubar.tsx";
import { useEditorsStore } from "@/store/editors.ts";
import { useEffect, useState } from "react";
import { useTranslation } from "react-i18next";
import { GetExamples } from "../../../wailsjs/go/main/App";
import { examples } from "../../../wailsjs/go/models.ts";
import Example = examples.Example;

export default function () {
	const { t } = useTranslation();
	const setJSON = useEditorsStore((state) => state.setJSON);

	const [exampleList, setExampleList] = useState([] as Example[]);
	useEffect(() => {
		GetExamples().then((value) => {
			setExampleList(value);
		});
	}, []);

	return (
		<MenubarMenu>
			<MenubarTrigger>{t("examples.title", "Examples")}</MenubarTrigger>
			<MenubarContent>
				{exampleList.map((example) => (
					<MenubarItem
						key={example.title}
						onClick={() => setJSON(example.content)}
					>
						{example.title}
					</MenubarItem>
				))}
			</MenubarContent>
		</MenubarMenu>
	);
}
