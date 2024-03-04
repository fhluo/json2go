import React from "react";
import { createRoot } from "react-dom/client";
import App from "./app.tsx";
import "./index.css";
import "./lib/i18n.ts";

const rootElement = document.getElementById("root");

if (rootElement) {
	const root = createRoot(rootElement);
	root.render(
		<React.StrictMode>
			<App />
		</React.StrictMode>,
	);
}
