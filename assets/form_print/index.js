import "./bootstrap"
import {i18n} from '@lambda-platform/locale';

import App from './app';





new Vue({

    extends: App,
    el:'#app',
    i18n,
});
