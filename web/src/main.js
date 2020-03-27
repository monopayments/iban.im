import Vue from 'vue';
import vuetify from './plugins/vuetify';
import App from "./App";
import router from './router';
import store from './store'
// import axios from 'axios'

Vue.config.productionTip = false;

//axios.defaults.baseURL = "http://195.201.97.159:4880";


new Vue({
    vuetify,
    store,
    router,
    render: h => h(App),

    created() {
        // const token = localStorage.getItem('user');
        // if (token) {
        //     this.$store.commit('SET_TOKEN', token)
        // }
        // axios.interceptors.response.use(
        //     response => response,
        //     error => {
        //         console.log(error.response);
        //         if (error.response.status === 401) {
        //             this.$store.dispatch('logout');
        //             this.$router.push('/');
        //         }
        //         return Promise.reject(error)
        //     }
        // )
    }
}).$mount('#app');
