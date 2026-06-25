import { Button } from "@/components/ui/button.tsx";
import { useEditorsStore } from "@/store/editors.ts";
import { useTranslation } from "react-i18next";

export default function Footer() {
    const { t } = useTranslation();
    const generate = useEditorsStore((state) => state.generate);

    return (
        <div className="flex flex-row px-4 py-2 justify-end items-center h-12">
            <Button size="sm" onClick={generate} className="mr-2">
                {t("Generate")}
            </Button>
        </div>
    );
}
