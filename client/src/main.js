import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from './router'
import mdiVue from 'mdi-vue/v3'
import * as mdijs from '@mdi/js'
import store from './store/vuex'
import { cryption } from './functions'

let app = createApp(App)
app.config.globalProperties.cryption = cryption

app
    .use(mdiVue, {
        icons: mdijs
    })
    .use(store)
    .use(router)
    .mount('#app')
