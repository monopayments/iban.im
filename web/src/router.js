import Vue from 'vue'
import Router from 'vue-router'
import Home from './components/Home'
import Login from "./components/Login";
import Register from "./components/Register";
import VueBodyClass from 'vue-body-class';
import Ibans from "./components/Ibans";
import Profile from "./components/Profile";
import Security from "./components/Security";
import Logout from "./components/Logout";
import Main from "./components/Main";
import Single from "./components/Single";
import SingleIban from "./components/SingleIban";
import store from "./store";


Vue.use(Router);

const routes =  [
    {
        path: '/',
        name: 'home',
        component: Home,
        meta: {
            bodyClass: 'guest home'
        },
        children: [
            {
                path: '/',
                name: 'home.main',
                component: Main,
                meta: {
                    bodyClass: 'guest home',
                    public: true,
                }
            },
            {
                path: '/dashboard',
                name: 'home.profile',
                component: Profile,
                meta: {
                    bodyClass: 'dashboard',
                    requiresAuth: true
                },
            },
            {
                path: '/dashboard/security',
                name: 'home.security',
                component: Security,
                meta: {
                    bodyClass: 'security',
                    requiresAuth: true
                },
            },
            {
                path: '/dashboard/ibans',
                name: 'home.ibans',
                component: Ibans,
                meta: {
                    bodyClass: 'ibans',
                    requiresAuth: true
                },
            },
            {
                path: '/dashboard/logout',
                name: 'home.logout',
                component: Logout,
                meta: {
                    bodyClass: 'logout',
                    requiresAuth: true
                },
            },
            {
                path: '/login',
                name: 'home.login',
                component: Login,
                meta: {
                    bodyClass: 'guest',
                    public: true,
                }
            },
            {
                path: '/register',
                name: 'home.register',
                component: Register,
                meta: {
                    bodyClass: 'guest',
                    public: true,
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

//router.beforeEach((to, from, next) => { vueBodyClass.guard(to, next) });
router.beforeEach(async(to,from,next) => {
    if (to.path === '/logout'){
        localStorage.removeItem('user');
        next('/login');
    }
    const token = localStorage.getItem("user");
    if(token && !store.state.logged) {
        store.commit('SET_HEADER', token);
        await store.dispatch('getUser');
    }
    if(token && !store.state.logged){
        next('/login');
    }
    if (to.matched.some(record => record.meta.requiresAuth) && !store.state.logged) {
        next('/login');
    }

    if (to.matched.some(record => record.meta.public) && store.state.logged) {
        next('/dashboard');
    }
    vueBodyClass.guard(to, next);
});


export default router;