import { createApp } from 'vue'
import {createRouter, createWebHistory} from 'vue-router'
import PrimeVue from 'primevue/config';
import App from './App.vue'

// PrimeVue components import
import Button from "primevue/button"
import Badge from "primevue/badge"
import Breadcrumb from "primevue/breadcrumb"
import Menubar from "primevue/menubar"
import Column from 'primevue/column'
import DataTable from 'primevue/datatable'
import TabView from "primevue/tabview"
import TabPanel from "primevue/tabpanel"
import InputSwitch from 'primevue/inputswitch';
import InputText from "primevue/inputtext"
import InputNumber from "primevue/inputnumber"



// import './style.css'
// import 'primevue/resources/themes/aura-light-green/theme.css'
import 'primevue/resources/themes/aura-dark-green/theme.css'
import 'primeicons/primeicons.css'
import 'primeflex/primeflex.css'
// import 'primeflex/themes/primeone-light.css'
// import 'primeflex/themes/primeone-dark.css'
import './sass/style.sass'


import Index from './Index.vue'
import Modules from './pages/Modules.vue'
import ModulesId from './pages/ModulesId.vue'
import HomeView from './components/HelloWorld.vue'
import AboutView from './components/About.vue'

const routes = [
    { path: '/', name: 'index', component: Index },
    { path: '/modules', name: 'modules.index', component: Modules },
    { path: '/modules/:id', name: 'modules.id', component: ModulesId },
    { path: '/route1', name: 'route1', component: HomeView },
    { path: '/route2', name: 'route2', component: AboutView },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

const app = createApp(App)


app.use(router)
app.use(PrimeVue)


app.component('Badge', Badge)
app.component('Button', Button)
app.component('Breadcrumb', Breadcrumb)
app.component('Column', Column)
app.component('DataTable', DataTable)
app.component('TabView', TabView)
app.component('TabPanel', TabPanel)
app.component('Menubar', Menubar)
app.component('InputSwitch', InputSwitch)
app.component('InputText', InputText)
app.component('InputNumber', InputNumber)


app.mount('#app')
