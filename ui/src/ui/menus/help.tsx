import { DialogTrigger } from "@/components/ui/dialog.tsx";
import {
	MenubarContent,
	MenubarItem,
	MenubarMenu,
	MenubarSeparator,
	MenubarTrigger,
} from "@/components/ui/menubar.tsx";
import { openHomePage } from "@/lib/api.ts";
import { useDialogStore } from "@/store/dialog.ts";
import About from "@/ui/dialogs/about.tsx";
import Updates from "@/ui/dialogs/updates.tsx";
import { useTranslation } from "react-i18next";

export default function Help() {
	const { t } = useTranslation();
	const setDialog = useDialogStore((state) => state.setDialog);

	return (
		<MenubarMenu>
			<MenubarTrigger>{t("help.title", "Help")}</MenubarTrigger>
			<MenubarContent>
				<DialogTrigger
					render={
						<MenubarItem onClick={() => setDialog(<About />)}>
							{t("help.about", "About")}
						</MenubarItem>
					}
				/>
				<DialogTrigger
					render={
						<MenubarItem onClick={() => setDialog(<Updates />)}>
						{t("help.updates", "Check for updates")}
					</MenubarItem>
					}
				/>
				<MenubarSeparator />
				<MenubarItem onClick={openHomePage}>
					{t("help.repository", "Repository")}
				</MenubarItem>
			</MenubarContent>
		</MenubarMenu>
	);
}
