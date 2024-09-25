export enum Layout {
    TwoColumns = "Two Columns",
    TwoRows = "Two Rows",
}

export enum View {
    JSONAndGo = "JSON and Go",
    JSONOnly = "JSON Only",
    GoOnly = "Go Only",
}

export interface UIState {
    layout: Layout,
    view: View,
    showTwoColumns: boolean,
    showTwoRows: boolean,
    hideJSONEditor: boolean,
    hideGoEditor: boolean,
}

export function createUIState(): UIState {
    let layout = $state(Layout.TwoColumns)
    let view = $state(View.JSONAndGo)

    return {
        get layout() {
            return layout
        },
        set layout(value) {
            layout = value
        },
        get view() {
            return view
        },
        set view(value) {
            view = value
        },
        get showTwoColumns() {
            return view === View.JSONAndGo && layout === Layout.TwoColumns
        },
        get showTwoRows() {
            return view === View.JSONAndGo && layout === Layout.TwoRows
        },
        get hideJSONEditor() {
            return view === View.GoOnly
        },
        get hideGoEditor() {
            return view === View.JSONOnly
        }
    }
}
