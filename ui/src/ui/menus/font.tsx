import {
	MenubarContent,
	MenubarItem,
	MenubarMenu,
	MenubarTrigger,
} from "@/components/ui/menubar.tsx";
import { useConfigStore } from "@/lib/api.ts";
import { useTranslation } from "react-i18next";

export default function Font() {
	const { t } = useTranslation();
	const increase = useConfigStore((s) => s.increaseFontSize);
	const decrease = useConfigStore((s) => s.decreaseFontSize);
	const reset = useConfigStore((s) => s.resetFontSize);

	return (
		<MenubarMenu>
			<MenubarTrigger>{t("font.title", "Font")}</MenubarTrigger>
			<MenubarContent>
				<MenubarItem onClick={increase}>
					{t("font.increase", "Increase size")}
				</MenubarItem>
				<MenubarItem onClick={decrease}>
					{t("font.decrease", "Decrease size")}
				</MenubarItem>
				<MenubarItem onClick={reset}>
					{t("font.reset", "Reset size")}
				</MenubarItem>
			</MenubarContent>
		</MenubarMenu>
	);
}
