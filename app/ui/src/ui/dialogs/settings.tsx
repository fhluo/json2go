import {DialogClose, DialogContent, DialogFooter, DialogHeader, DialogTitle,} from "@/components/ui/dialog.tsx"
import {FormEvent, useEffect, useState} from "react"
import {Button} from "@/components/ui/button.tsx"
import {Input} from "@/components/ui/input.tsx"
import {useTranslation} from "react-i18next"
import {useAllCapsWordsStore} from "@/lib/store.ts"

function onBeforeInput(event: FormEvent<HTMLInputElement>) {
    // user can only enter letters, space and ','
    const value = event.currentTarget.value
    if (value && !value.match(/[a-zA-Z', ]/)) {
        event.preventDefault()
    }
}

export default function () {
    const {t} = useTranslation()

    const {words, init, add, remove} = useAllCapsWordsStore()
    useEffect(() => {
        init()
    }, [])

    const [input, setInput] = useState("")

    return (
        <DialogContent>
            <DialogHeader>
                <DialogTitle>{t("settings.title", "Settings")}</DialogTitle>
            </DialogHeader>
            <div className="flex flex-col space-y-2">
                <span className="select-none font-semibold mr-3">
                    {t("settings.all-caps", "All-caps words")}
                </span>
                {words?.length > 0 &&
                  <div className="flex flex-row flex-wrap space-x-1.5 pb-1.5">
                      {words.map(value =>
                          <Button size="sm" variant={"outline"}
                                  onClick={() => setInput(value)}
                                  onDoubleClick={() => remove(input)}
                                  className="ml-1 mt-1.5 hover:bg-gray-50 transition cursor-default">
                              {value}
                          </Button>
                      )}
                  </div>
                }
                <div>
                    <Input value={input}
                           onBeforeInput={onBeforeInput}
                           onChange={event => setInput(event.currentTarget.value)}
                           onKeyDown={event => {
                               if (event.key == "Enter") {
                                   add(input)
                                   setInput("")
                               }
                           }}/>
                    <p className="text-sm text-gray-500 leading-relaxed py-2 px-1 select-none">
                        {t(
                            "settings.tip",
                            "Tip: Double click a word to remove it. To add multiple words, separate words with commas."
                        )}
                    </p>
                </div>
            </div>
            <DialogFooter>
                <Button variant="secondary" className="min-w-fit" onClick={() => {
                    add(input)
                    setInput("")
                }}>
                    {t("settings.add", "Add")}
                </Button>
                <Button variant="secondary" className="min-w-fit" onClick={() => {
                    remove(input)
                    setInput("")
                }}>
                    {t("settings.remove", "Remove")}
                </Button>
                <DialogClose asChild>
                    <Button variant="secondary">{t("OK")}</Button>
                </DialogClose>
            </DialogFooter>
        </DialogContent>
    )
}
