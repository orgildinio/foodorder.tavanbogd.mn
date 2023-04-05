import Vue from "vue";
import {i18n} from '@lambda-platform/locale';
import axios from "axios";
import router from "./router";
import iView from 'iview';
import lodash from 'lodash';
import Multiselect from 'vue-multiselect'
import locale from 'iview/src/locale/lang/mn-MN';
// import 'swiper/swiper.scss';
// import 'swiper/swiper-bundle.css';
// import Swiper, { Navigation, Pagination } from 'swiper';


// Swiper.use([Navigation, Pagination]);
Vue.use(iView, {locale});

window.Vue = Vue;
window.axios = axios;
window._ = lodash;
Vue.component('multiselect', Multiselect)
window.axios.defaults.headers.common = {
    "X-Requested-With": "XMLHttpRequest",
    "X-CSRF-TOKEN": document
        .querySelector('meta[name="csrf-token"]')
        .getAttribute("content"),
};
Vue.config.productionTip = false;

import(/* webpackChunkName: "auth-[request]" */ `./views/theme/${window.lambda.theme}/index`).then(theme => {
    new Vue({
        router,
        i18n,
        render: h => h(theme.default),
        // render: h => h(require(`./views/theme/${window.lambda.theme}/index`).default)
    }).$mount('#app');
});



