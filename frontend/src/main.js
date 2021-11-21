import { createApp } from 'vue'
import App from './App.vue'
import axios from 'axios';

if (process.env.NODE_ENV === 'development') {
    axios.defaults.baseURL = 'http://localhost:10001';
}

createApp(App).mount('#app')
