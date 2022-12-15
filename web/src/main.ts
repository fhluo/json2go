import './style.css'
import App from './App.svelte'
import './worker'

const app = new App({
    target: document.getElementById('app')
})

export default app
