import 'vite/modulepreload-polyfill';

import './assets/main.css';

import VueApolloComponents from '@vue/apollo-components';
import { createApp } from 'vue';
import { createPinia } from 'pinia';
import App from './App.vue';
import router from './router';
import { library } from '@fortawesome/fontawesome-svg-core';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { faUserSecret, faCoffee, faArrowRight, faBars, faPlay } from '@fortawesome/free-solid-svg-icons';
import apolloProvider from './apollo';

library.add(faUserSecret, faCoffee, faArrowRight, faBars, faPlay);

const app = createApp(App);
app.use(createPinia());
app.use(router);
app.use(apolloProvider);
app.use(VueApolloComponents);
app.mount('#app');
app.component("FontAwesomeIcon", FontAwesomeIcon)
