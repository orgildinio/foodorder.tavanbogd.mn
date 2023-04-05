import "./sass/app.scss"

import windowVariables from './common/window_variables';
windowVariables();
import './esriconfig';

Vue.config.productionTip = false
import Vue from 'vue';

import Vuex from 'vuex';
import VueRouter from 'vue-router';



import plugins from './config/plugins';
import directives from './common/directives';

import configRouter from './config/router';
import vuexStore from './store/store';


import CollapseItem  from 'element-ui/lib/collapse-item'
import lang from 'element-ui/lib/locale/lang/en'
import FromItem  from 'element-ui/lib/form-item'
import Collapse  from 'element-ui/lib/collapse'
import Checkbox from 'element-ui/lib/checkbox'
import Tooltip  from 'element-ui/lib/tooltip'
import Popover from 'element-ui/lib/popover'
import Button  from 'element-ui/lib/button'
import Input  from 'element-ui/lib/input'
import Rate  from 'element-ui/lib/rate'
import locale from 'element-ui/lib/locale'
import Select from 'element-ui/lib/select'
import Option from 'element-ui/lib/option'
import Dialog from 'element-ui/lib/dialog'
import Slider from 'element-ui/lib/slider'
import Steps from 'element-ui/lib/steps'
import Form  from 'element-ui/lib/form'
import Step from 'element-ui/lib/step'
import Radio from 'element-ui/lib/radio'
import Switch from 'element-ui/lib/switch'

// configure language
locale.use(lang);

// import components
Vue.component(CollapseItem.name, CollapseItem);
Vue.component(FromItem.name, FromItem);
Vue.component(Collapse.name, Collapse);
Vue.component(Checkbox.name, Checkbox);
Vue.component(Tooltip.name, Tooltip);
Vue.component(Popover.name, Popover);
Vue.component(Button.name, Button);
Vue.component(Select.name, Select);
Vue.component(Option.name, Option);
Vue.component(Dialog.name, Dialog);
Vue.component(Slider.name, Slider);
Vue.component(Input.name, Input);
Vue.component(Steps.name, Steps);
Vue.component(Form.name, Form);
Vue.component(Step.name, Step);
Vue.component(Rate.name, Rate);
Vue.component(Radio.name, Radio);
Vue.component(Switch.name, Switch);



plugins(Vue, Vuex, VueRouter);

directives(Vue);


const store = new Vuex.Store(vuexStore);
const router = configRouter(VueRouter);
// main loyout
import Main from './components/Main.vue'

new Vue({
  router:router,
  store:store,
  render: h => h(Main),
}).$mount('#app');

