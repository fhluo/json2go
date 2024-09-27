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
      "@": path.resolve(__dirname, "src"),
      "@state": path.resolve(__dirname, "src/state"),
      "@lib": path.resolve(__dirname, "src/lib"),
      "@ui": path.resolve(__dirname, "src/ui"),
      "@api": path.resolve(__dirname, "bindings/github.com/fhluo/json2go"),
    }
  },
  build: {
    outDir: "../app/web/dist"
  },
  server: {
    port: 9245,
    strictPort: true,
  }
})
