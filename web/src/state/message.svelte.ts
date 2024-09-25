export interface MessageState {
    message: string;
    clear: () => void;
}

export function createMessageState() {
    let message = $state("");
    return {
        get message() {
            return message;
        },
        set message(value) {
            message = value;
        },
        clear: () => {
            message = "";
        },
    };
}
