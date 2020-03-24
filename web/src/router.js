import Vue from 'vue'
import Router from 'vue-router'
import Home from './components/Home'
import Login from "./components/Login";
import Register from "./components/Register";
import Dashboard from "./components/Dashboard";
import VueBodyClass from 'vue-body-class';
import Ibans from "./components/Ibans";
import Profile from "./components/Profile";
import Security from "./components/Security";

Vue.use(Router);

const routes =  [
    {
        path: '/',
        name: 'home',
        component: Home,
        meta: {
            bodyClass: 'guest'
        }
    },
    {
        path: '/login',
        name: 'login',
        component: Login,
        meta: {
            bodyClass: 'guest'
        }
    },
    {
        path: '/register',
        name: 'register',
        component: Register,
        meta: {
            bodyClass: 'guest'
        }
    },
    {
        path: '/dashboard',
        name: 'dashboard',
        component: Dashboard,
        meta: {
            bodyClass: 'dashboard'
        },
        children: [
            {
                path: '/',
                name: 'dashboard.profile',
                component: Profile,
            },
            {
                path: 'security',
                name: 'dashboard.security',
                component: Security,
                meta: {
                    bodyClass: 'security'
                },
            },
            {
                path: 'ibans',
                name: 'dashboard.ibans',
                component: Ibans,
                meta: {
                    bodyClass: 'ibans'
                },
            },
        ]
    },
];

const router = new Router({
    mode: 'history',
    routes
});

const vueBodyClass = new VueBodyClass(routes);

router.beforeEach((to, from, next) => { vueBodyClass.guard(to, next) });

export default router;