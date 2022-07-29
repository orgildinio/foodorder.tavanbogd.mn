import Vue from 'vue';
import iView from 'iview';
import locale from 'iview/dist/locale/mn-MN';
import axios from 'axios';

import moqupView from './moqup/moqup';


window.Vue = Vue;


window.axios = axios;
Vue.config.productionTip = false;
Vue.use(iView, { locale });

new Vue({
    el: '#moqup',
    extends: moqupView,
});
