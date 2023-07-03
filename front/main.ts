import 'vite/modulepreload-polyfill';

import './assets/main.css';

import VueApolloComponents from '@vue/apollo-components';
import { createApp } from 'vue';
import { createPinia } from 'pinia';
import App from './App.vue';
import router from './router';
import apolloProvider from './apollo';

const app = createApp(App);
app.use(createPinia());
app.use(router);
app.use(apolloProvider);
app.use(VueApolloComponents);
app.mount('#app');
