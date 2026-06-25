import { useMessageStore } from "@/store/message.ts";
import { Info, X } from "lucide-react";

export default function Message() {
    const { message, clear } = useMessageStore();

    if (!message) {
        return <></>;
    }

    return (
        <div
            className={
                "flex flex-row items-center justify-center" +
                " select-none mx-4 border bg-white/50 rounded shadow-sm space-x-1.5"
            }
        >
            <Info className="w-5 h-5 ml-2 mr-1 text-red-600" />
            <span>{message}</span>
            <button
                type={"button"}
                onClick={clear}
                className={"hover:bg-gray-200/50 py-1 px-1 rounded-r transition flex items-center justify-center"}
            >
                <X className="w-5 h-5" />
            </button>
        </div>
    );
}
