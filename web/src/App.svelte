<script lang="ts">
    import {Generate, GetFontSize} from '../wailsjs/go/main/App.js'
    import "fluent-svelte/theme.css"
    import {Button} from "fluent-svelte"
    import "./lib/i18n"
    import {_} from "svelte-i18n"
    import {EventsEmit, EventsOn} from "../wailsjs/runtime"
    import MenuBar from "./lib/MenuBar.svelte"
    import Editors from './lib/Editors.svelte'
    import {editor} from "monaco-editor";
    import {onMount} from "svelte";
    import {Layout, View} from "./lib/types";
    import IStandaloneCodeEditor = editor.IStandaloneCodeEditor;

    let showErrorInfo = false
    let errorMessage = ""
    let allCapsWords = [] as string[]

    let jsonEditor: IStandaloneCodeEditor
    let goEditor: IStandaloneCodeEditor

    function generate(): void {
        showErrorInfo = false
        errorMessage = ""
        Generate(jsonEditor.getValue(), allCapsWords).then(result => {
            goEditor.setValue(result)
        })
    }

    EventsOn("error", (message: string) => {
        showErrorInfo = true
        errorMessage = message
    })

    let optionsGenerateInRealTime = false

    document.defaultView.addEventListener('resize', () => {
        jsonEditor?.layout()
        goEditor?.layout()
        EventsEmit("resize")
    })

    onMount(() => {
        jsonEditor.getModel().onDidChangeContent(() => {
            if (optionsGenerateInRealTime) {
                generate()
            }
        })
    })

    let layout: Layout = Layout.TwoColumns
    let view: View = View.JSONAndGo
    let fontSize: number

    GetFontSize().then(result => {
        fontSize = result
    })
</script>

<main class="w-screen h-screen flex flex-col">
  <MenuBar bind:layout={layout} bind:view={view} bind:fontSize={fontSize} bind:jsonEditor={jsonEditor}
           bind:goEditor={goEditor} bind:allCapsWords bind:optionsGenerateInRealTime></MenuBar>

  <Editors bind:jsonEditor={jsonEditor} bind:goEditor={goEditor} bind:layout={layout} bind:fontSize={fontSize}
           bind:view={view}></Editors>

  <div class="flex flex-row px-4 py-2 justify-end items-center h-12">
    {#if showErrorInfo}
      <div class="select-none mx-4 border bg-white/50 rounded shadow-sm flex flex-row items-center justify-center space-x-1.5">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
             stroke="currentColor" class="w-6 h-6 ml-2 mr-1 text-red-600">
          <path stroke-linecap="round" stroke-linejoin="round"
                d="M12 9v3.75m9-.75a9 9 0 11-18 0 9 9 0 0118 0zm-9 3.75h.008v.008H12v-.008z"/>
        </svg>
        <span>{errorMessage}</span>
        <button class="hover:bg-gray-200/50 py-1 px-1 rounded-r transition flex items-center justify-center"
                on:click={()=>showErrorInfo=false}>
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1"
               stroke="currentColor" class="w-6 h-6">
            <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12"/>
          </svg>
        </button>
      </div>
    {/if}
    <Button variant="accent" on:click={generate} class="mr-2">{$_('Generate')}</Button>
  </div>
</main>

