import {
    MenubarContent,
    MenubarMenu,
    MenubarRadioGroup,
    MenubarRadioItem,
    MenubarSub,
    MenubarSubContent,
    MenubarSubTrigger,
    MenubarTrigger
} from "@/components/ui/menubar.tsx"
import {Layout, View} from "@/lib/types.ts"
import {useTranslation} from "react-i18next"

interface Props {
    view: View
    setView: (view: View) => void

    layout: Layout
    setLayout: (layout: Layout) => void
}

export default function ({view, setView, layout, setLayout}: Props) {
    const {t} = useTranslation()

    return (
        <MenubarMenu>
            <MenubarTrigger>
                {t("view.title", "View")}
            </MenubarTrigger>
            <MenubarContent>
                <MenubarSub>
                    <MenubarSubTrigger>
                        {t("view.editors", "Editors")}
                    </MenubarSubTrigger>
                    <MenubarSubContent>
                        <MenubarRadioGroup value={view} onValueChange={value => setView(value as View)}>
                            <MenubarRadioItem value={View.JSONAndGo}>
                                {t("view.both", View.JSONAndGo)}
                            </MenubarRadioItem>
                            <MenubarRadioItem value={View.JSONOnly}>
                                {View.JSONOnly}
                            </MenubarRadioItem>
                            <MenubarRadioItem value={View.GoOnly}>
                                {View.GoOnly}
                            </MenubarRadioItem>
                        </MenubarRadioGroup>
                    </MenubarSubContent>
                </MenubarSub>
                <MenubarSub>
                    <MenubarSubTrigger>
                        {t("view.layout", "Layout")}
                    </MenubarSubTrigger>
                    <MenubarSubContent>
                        <MenubarRadioGroup value={layout} onValueChange={value => setLayout(value as Layout)}>
                            <MenubarRadioItem value={Layout.TwoColumns}>
                                {t("view.columns", Layout.TwoColumns)}
                            </MenubarRadioItem>
                            <MenubarRadioItem value={Layout.TwoRows}>
                                {t("view.rows", Layout.TwoRows)}
                            </MenubarRadioItem>
                        </MenubarRadioGroup>
                    </MenubarSubContent>
                </MenubarSub>
            </MenubarContent>
        </MenubarMenu>
    )
}
