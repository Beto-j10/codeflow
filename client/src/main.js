import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import fun from './modules/testBabel';

fun()

import App from './App.vue'

const app = createApp(App)

app.use(ElementPlus)

app.mount('#app')
