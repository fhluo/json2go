// @ts-ignore
import * as monaco from 'monaco-editor';
import JSONWorker from 'monaco-editor/esm/vs/language/json/json.worker?worker';
import EditorWorker from 'monaco-editor/esm/vs/editor/editor.worker?worker';

self.MonacoEnvironment = {
    getWorker: function (workerId, label) {
        switch (label) {
            case 'json':
                return new JSONWorker()
            default:
                return new EditorWorker()
        }
    }
};