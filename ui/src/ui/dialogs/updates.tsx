import { Button } from "@/components/ui/button.tsx";
import {
	DialogClose,
	DialogContent,
	DialogFooter,
	DialogHeader,
	DialogTitle,
} from "@/components/ui/dialog.tsx";
import { openRelease } from "@/lib/api.ts";
import { useEffect, useRef, useState } from "react";
import { useTranslation } from "react-i18next";
import { CheckForUpdate } from "@api/app/services/version";

export default function Updates() {
	const { t } = useTranslation();

	const [version, setVersion] = useState("");
	const [latestVersion, setLatestVersion] = useState("");
	const [hasUpdate, setHasUpdate] = useState(false);
	const [message, setMessage] = useState(t("updates.checking", "Checking..."));
	const isChecking = useRef(true);

	useEffect(() => {
		void CheckForUpdate().then((info) => {
			setVersion(info.currentVersion);
			setLatestVersion(info.latestVersion);
			setHasUpdate(info.hasUpdate);
			isChecking.current = false;
		});
	}, []);

	useEffect(() => {
		if (isChecking.current) {
			return;
		}

		if (!version || !latestVersion) {
			setMessage(t("updates.failed", "Unable to check for updates."));
			return;
		}

		if (hasUpdate) {
			setMessage(
				`${t(
					"updates.available",
					"A new version is available: ",
				)}v${latestVersion}.`,
			);
		} else {
			setMessage(t("updates.none", "You are using the latest version."));
		}
	}, [version, latestVersion, hasUpdate]);

	return (
		<DialogContent className="select-none">
			<DialogHeader>
				<DialogTitle>{t("updates.title", "Check for updates")}</DialogTitle>
			</DialogHeader>
			<p className="text-gray-900 leading-relaxed">{message}</p>
			<DialogFooter>
				{version && latestVersion && hasUpdate && (
					<Button
						variant="secondary"
						onClick={() => openRelease(latestVersion)}
					>
						{t("Download")}
					</Button>
				)}
				<DialogClose
					render={<Button variant="secondary">{t("OK")}</Button>}
				/>
			</DialogFooter>
		</DialogContent>
	);
}
