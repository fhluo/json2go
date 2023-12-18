import React from "react"
import {createRoot} from "react-dom/client"
import App from "./app.tsx"
import "./lib/i18n.ts"
import "./index.css"

const root = createRoot(document.getElementById("root")!)
root.render(
    <React.StrictMode>
        <App/>
    </React.StrictMode>,
)
