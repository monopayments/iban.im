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
import Logout from "./components/Logout";
import Main from "./components/Main";
import Single from "./components/Single";
import SingleIban from "./components/SingleIban";

Vue.use(Router);

const routes =  [
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
            {
                path: 'logout',
                name: 'dashboard.logout',
                component: Logout,
                meta: {
                    bodyClass: 'logout'
                },
            },
        ]
    },
    {
        path: '/',
        name: 'home',
        component: Home,
        meta: {
            bodyClass: 'guest'
        },
        children: [
            {
                path: '/',
                name: 'home.main',
                component: Main,
                meta: {
                    bodyClass: 'guest'
                }
            },
            {
                path: '/login',
                name: 'home.login',
                component: Login,
                meta: {
                    bodyClass: 'guest'
                }
            },
            {
                path: '/register',
                name: 'home.register',
                component: Register,
                meta: {
                    bodyClass: 'guest'
                }
            },
            {
                path: '/:username',
                name: 'home.single',
                component: Single,
                meta: {
                    bodyClass: 'single guest'
                },
            },
            {
                path: '/:username/:alias',
                name: 'home.alias',
                component: SingleIban,
                meta: {
                    bodyClass: 'single alias guest'
                }
            }
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