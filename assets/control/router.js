import Router from 'vue-router'
import {routes} from "../controlTemplate/router"

Vue.use(Router)

let base = '/';
if (location.pathname && location.pathname != '/') {
    base = location.pathname.split('/').slice(0, -1).join('/');
}

// let AppRoutes = [
//     ...routes
// ];

let routesn = [
    ...routes,
    {
        path: '/prjctadv',
        component: () => import(/* webpackChunkName: "project_advertisement" */ './pages/project_advertisement')
    },
]

export default new Router({
    base,
    routes: routesn,
    history: true,
    linkActiveClass: 'active'
})
