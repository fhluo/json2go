import { Button } from "@/components/ui/button.tsx";
import {
	DialogClose,
	DialogContent,
	DialogFooter,
	DialogHeader,
	DialogTitle,
} from "@/components/ui/dialog.tsx";
import {
	Tooltip,
	TooltipContent,
	TooltipProvider,
	TooltipTrigger,
} from "@/components/ui/tooltip.tsx";
import { openHomePage } from "@/lib/api.ts";
import { useEffect, useState } from "react";
import { useTranslation } from "react-i18next";
import { GetVersion } from "@api/app/services/version";

export default function About() {
	const { t } = useTranslation();

	const [version, setVersion] = useState("");
	useEffect(() => {
		void GetVersion().then((value) => setVersion(value));
	}, []);

	return (
		<DialogContent className="select-none">
			<DialogHeader>
				<DialogTitle>{t("about.title", "About")}</DialogTitle>
			</DialogHeader>
			<div className="flex flex-col items-center justify-center space-y-2 text-gray-900">
				<TooltipProvider>
					<Tooltip>
						<TooltipTrigger
							render={
								<p className="text-lg font-semibold" onDoubleClick={openHomePage}>
									json2go
								</p>
							}
						/>
						<TooltipContent className="bg-transparent text-slate-600">
							<p>
								{t("about.tip", "Double-click to visit json2go on Github.com")}
							</p>
						</TooltipContent>
					</Tooltip>
				</TooltipProvider>
				<p className="text-sm leading-relaxed text-muted-foreground">
					{t("about.description", "Generate Go type definitions from JSON")}
				</p>
				<div className="leading-relaxed text-sm">
					<p>
						<span className="select-none">
							{t("about.license", "License: ")}
						</span>
						MIT
					</p>
					<p>
						<span className="select-none">
							{t("about.version", "Version: ")}
						</span>
						{version}
					</p>
				</div>
				<p className="text-sm">Copyright © 2022 fhluo</p>
			</div>
			<DialogFooter>
				<DialogClose
					render={<Button variant="secondary">{t("OK")}</Button>}
				/>
			</DialogFooter>
		</DialogContent>
	);
}
