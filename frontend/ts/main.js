import { createApp } from "vue";
import App from "../index.vue"

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

import UnBackendServerRunning from "../components/unlogined/unlogin.vue"

import routers from "./router.js";

const app = createApp(App);

app.use(ElementPlus)
app.use(routers)
app.components = {
    UnBackendServerRunning
}
app.mount('#app')
