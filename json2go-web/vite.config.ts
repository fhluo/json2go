import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'
import path from "path"

// https://vitejs.dev/config/
export default defineConfig({
  appType: "mpa",
  plugins: [svelte()],
  resolve: {
    alias: {
      "$lib": path.resolve("./src/lib"),
    },
  },
  build: {
    outDir: "../json2go-wails/web/dist"
  }
})
