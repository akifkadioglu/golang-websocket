import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home/View.vue'
import Chat from '../views/Chat/View.vue'
import Login from '../views/Login/View.vue'

import { routeNames } from './route_names'

import AppLayout from '../layouts/App.vue'
import AuthLayout from '../layouts/Auth.vue'

const routes = [
    {
        path: '/',
        name: routeNames.HOME,
        component: Home,
        meta: { layout: AppLayout },
    },
    {
        path: '/token/:token',
        meta: { layout: AppLayout },
        beforeEnter: (to, from, next) => {
            const token = to.params.token;
            localStorage.setItem("token", token);
            next('/');
        },
        component: () => null,
    },
    {
        path: '/login',
        name: routeNames.LOGIN,
        component: Login,
        meta: { layout: AuthLayout },
    },
    {
        path: '/chat/:name',
        name: routeNames.CHAT,
        component: Chat,
        meta: { layout: AppLayout },
    },
    {
        path: "/:catchAll(.*)",
        redirect: '/',
    },
]
export default createRouter({
    history: createWebHistory(),
    routes: routes,
})