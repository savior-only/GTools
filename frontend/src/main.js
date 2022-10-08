import { createApp } from 'vue'
import App from './App.vue'
import naive from 'naive-ui'
import router from './router/index'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

const app = createApp(App)
app.use(naive)
app.use(ElementPlus)
app.use(router)
app.mount('#app')
