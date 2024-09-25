import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'
import tailwindcss from '@tailwindcss/vite'
import path from "path"

// https://vitejs.dev/config/
export default defineConfig({
  appType: "mpa",
  plugins: [tailwindcss(), svelte()],
  resolve: {
    alias: {
      "$lib": path.resolve("./src/lib"),
    },
  },
  build: {
    outDir: "../app/web/dist"
  },
  server: {
    port: 9245,
    strictPort: true,
  }
})
