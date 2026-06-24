import { Button } from "@/components/ui/button.tsx";
import { useEditorsStore } from "@/store/editors.ts";
import { useMessageStore } from "@/store/message.ts";
import Message from "@/ui/message.tsx";
import { useTranslation } from "react-i18next";

export default function () {
	const { t } = useTranslation();
	const clear = useMessageStore((state) => state.clear);
	const generate = useEditorsStore((state) => state.generate);

	return (
		<div className="flex flex-row px-4 py-2 justify-end items-center h-12">
			<Message />
			<Button
				size="sm"
				onClick={() => {
					clear();
					generate();
				}}
				className="mr-2"
			>
				{t("Generate")}
			</Button>
		</div>
	);
}
