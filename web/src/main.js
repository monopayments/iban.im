import Vue from 'vue';
import vuetify from './plugins/vuetify';
import App from "./App";
import router from './router';

Vue.config.productionTip = false;

new Vue({
  vuetify,
  router,
  render: h => h(App)
}).$mount('#app');
