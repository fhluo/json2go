import {Cross1Icon, InfoCircledIcon} from "@radix-ui/react-icons"

interface Props {
    message: string
    clearMessage: () => void
}

export default function ({message, clearMessage}: Props) {
    return (
        <div className={
            "flex flex-row items-center justify-center" +
            "select-none mx-4 border bg-white/50 rounded shadow-sm space-x-1.5"
        }>
            <InfoCircledIcon className="w-5 h-5 ml-2 mr-1 text-red-600"/>
            <span>
                {message}
            </span>
            <button onClick={clearMessage} className={
                "hover:bg-gray-200/50 py-1 px-1 rounded-r transition flex items-center justify-center"
            }>
                <Cross1Icon className="w-5 h-5"/>
            </button>
        </div>
    )
}
