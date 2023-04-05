import "../controlTemplate/bootstrap"
import App from "./App";
import router from "./router";
import {i18n} from '@lambda-platform/locale';
import VueApollo from 'vue-apollo'



import {
    store
} from '../controlTemplate/store/store'


Vue.prototype.$init = window.init;
Vue.prototype.$user = window.init.user;

Vue.use(VueApollo)

import {apolloClient} from "./vue-apollo.js"


const apolloProvider = new VueApollo({
    defaultClient: apolloClient,
})
new Vue({
    router,
    store,
    extends:App,
    apolloProvider,
    i18n,
    el: '#app',
});
