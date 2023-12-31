import { createRouter, createWebHistory } from 'vue-router';

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'home',
            component: () => import('../views/HomeView.vue'),
        },
        {
            path: '/auth',
            children: [
                {
                    path: '',
                    name: 'login',
                    component: () => import('../views/LogInView.vue'),
                },
                {
                    path: 'new',
                    name: 'signup',
                    component: () => import('../views/SignUpView.vue'),
                },
            ],
        },
        {
            path: '/about',
            name: 'about',
            // route level code-splitting
            // this generates a separate chunk (About.[hash].js) for this route
            // which is lazy-loaded when the route is visited.
            component: () => import('../views/AboutView.vue'),
        },
        {
            path: '/stats',
            name: 'stats',
            component: () => import('../views/StatsView.vue'),
        },
        {
            path: '/profile',
            name: 'profile',
            component: () => import('../views/ProfileView.vue'),
        },
        {
            path: '/checkup',
            name: 'checkup',
            component: () => import('../views/CheckupView.vue'),
        },
        {
            path: '/collection/:songIdOrName?',
            name: 'collection',
            component: () => import('../views/CollectionView.vue'),
        },
        {
            path: '/play',
            redirect: '/collection',
        },
        {
            path: '/collection/:url',
            name: 'url',
            props: true,
            component: () => import('../views/SongView.vue'),
        },
        {
            path: '/play/:performanceId',
            props: true,
            component: () => import('../views/PostPerformanceView.vue'),
        },
    ],
});

export default router;
