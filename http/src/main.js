import { createApp } from 'vue'
import PrimeVue from 'primevue/config';
import App from './App.vue'

import Button from "primevue/button"


// import './style.css'
// import 'primevue/resources/themes/aura-light-green/theme.css'
import 'primevue/resources/themes/aura-dark-green/theme.css'
import './sass/style.sass'



const app = createApp(App);


app.use(PrimeVue);


app.component(Button);

app.mount('#app')
