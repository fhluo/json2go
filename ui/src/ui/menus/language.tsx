import {
	MenubarContent,
	MenubarMenu,
	MenubarRadioGroup,
	MenubarRadioItem,
	MenubarTrigger,
} from "@/components/ui/menubar.tsx";
import { languages } from "@/lib/i18n.ts";
import { useConfigStore } from "@/lib/api.ts";
import { useTranslation } from "react-i18next";

export default function Language() {
	const { t } = useTranslation();
	const language = useConfigStore((s) => s.language);
	const setLanguage = useConfigStore((s) => s.setLanguage);

	return (
		<MenubarMenu>
			<MenubarTrigger>{t("language.title", "Language")}</MenubarTrigger>
			<MenubarContent>
				<MenubarRadioGroup
					value={language}
					onValueChange={(value) => setLanguage(value)}
				>
					{languages.map((language) => (
						<MenubarRadioItem key={language} value={language}>
							{t(`language.${language}`)}
						</MenubarRadioItem>
					))}
				</MenubarRadioGroup>
			</MenubarContent>
		</MenubarMenu>
	);
}
