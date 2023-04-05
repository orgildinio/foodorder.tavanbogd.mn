import "./bootstrap"
import App from './App';
import router from "./router";
import VueApollo from 'vue-apollo'
import {i18n} from '@lambda-platform/locale';
Vue.prototype.$init = window.init;
Vue.prototype.$user = window.init.user;
Vue.prototype.$logo = window.logo;
//
Vue.use(VueApollo)
import {apolloClient} from "./vue-apollo.js"

const apolloProvider = new VueApollo({
    defaultClient: apolloClient,
})

new Vue({
    components: {
        App
    },
    router,
     apolloProvider,
    i18n,
    render: h => h(App),
}).$mount('#app');
