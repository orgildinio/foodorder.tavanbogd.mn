import Vue from 'vue'
import axios from 'axios'
import lodash from 'lodash'

// import VcRibbon from 'vc-ribbon'
import iView from 'iview';
import locale from './mn-MN';
// import VcRibbon from 'vc-ribbon/src/components/vc-ribbon'




// import GanttTask from 'GanttTask'



Vue.config.productionTip = false;

window.Vue = Vue;
window._ = lodash;
window.axios = axios;
Vue.config.productionTip = false;
Vue.use(iView, { locale });



// Vue.use(VcRibbon)
// Vue.use(VcRibbon)

