import { createApp } from 'vue'
import { createMemoryHistory, createRouter } from 'vue-router'
import PrimeVue from 'primevue/config';
import App from './App.vue'

import Button from "primevue/button"


// import './style.css'
// import 'primevue/resources/themes/aura-light-green/theme.css'
import 'primevue/resources/themes/aura-dark-green/theme.css'
import 'primeicons/primeicons.css'
import './sass/style.sass'


import Home from './Home.vue'
import HomeView from './components/HelloWorld.vue'
import AboutView from './components/About.vue'

const routes = [
    { path: '/', name: 'home', component: Home },
    { path: '/route1', name: 'route1', component: HomeView },
    { path: '/route2', name: 'route2', component: AboutView },
]

const router = createRouter({
    history: createMemoryHistory(),
    routes,
})

const app = createApp(App);


app.use(router);
app.use(PrimeVue);


app.component(Button);

app.mount('#app')
