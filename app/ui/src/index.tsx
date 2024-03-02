import React from "react";
import { createRoot } from "react-dom/client";
import App from "./app.tsx";
import "./index.css";
import "./lib/i18n.ts";

const root = createRoot(document.getElementById("root")!);
root.render(
	<React.StrictMode>
		<App />
	</React.StrictMode>,
);
