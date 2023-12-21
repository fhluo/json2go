import Message from "@/ui/message.tsx"
import {Button} from "@/components/ui/button.tsx"
import {useTranslation} from "react-i18next"

interface Props {
    message: string
    clearMessage: () => void
    generate: () => void
}

export default function ({message, clearMessage, generate}: Props) {
    const {t} = useTranslation()

    return (
        <div className="flex flex-row px-4 py-2 justify-end items-center h-12">
            {message && <Message message={message} clearMessage={clearMessage}/>}
            <Button size="sm" onClick={generate} className="mr-2">
                {t("Generate")}
            </Button>
        </div>
    )
}
