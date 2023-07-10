import { createApp } from 'vue'
import App from './App.vue'
import router from './router';
import axios from 'axios';
import './assets/css/tailwind.css';

axios.defaults.baseURL = process.env.VUE_APP_BACKEND_PATH; 

const app = createApp(App);
app.use(router);
createApp(App).use(router).mount('#app')
