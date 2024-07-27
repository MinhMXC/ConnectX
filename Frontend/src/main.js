import './assets/main.css';

import { createApp } from 'vue';
import PrimeVue from 'primevue/config';
import Aura from '@primevue/themes/aura';
import App from './App.vue';
import VueCookies from 'vue-cookies';
import router from "@/router.js";

const app = createApp(App);
app.use(PrimeVue, {
    theme: {
        preset: Aura
    }
});
app.use(VueCookies, {
    expires: '30d',
    sameSite: 'Strict'
});
app.use(router);

app.mount('#app');
