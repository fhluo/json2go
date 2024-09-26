export interface DialogState {
    dialog: any;
}

export function createDialogState() {
    let dialog = $state(null);
    return {
        get dialog() {
            return dialog;
        },
        set dialog(value) {
            dialog = value;
        },
    };
}
