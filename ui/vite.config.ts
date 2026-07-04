import path from "path";
import { defineConfig } from "vite";
import tailwindcss from "@tailwindcss/vite";
import react, { reactCompilerPreset } from "@vitejs/plugin-react";
import babel from "@rolldown/plugin-babel";

// https://vite.dev/config/
export default defineConfig({
    appType: "mpa",
    plugins: [tailwindcss(), react(), babel({ presets: [reactCompilerPreset()] })],
    resolve: {
        alias: {
            "@": path.resolve(__dirname, "./src"),
            "@api": path.resolve(__dirname, "./bindings/github.com/fhluo/json2go"),
        },
    },
    build: {
        outDir: "../app/ui/dist",
        emptyOutDir: true, // outDir is outside project root, Vite won't empty it by default
    },
    server: {
        host: "127.0.0.1",
        port: 9245,
        strictPort: true,
    },
});
