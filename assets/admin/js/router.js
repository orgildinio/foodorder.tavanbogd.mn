import Router from 'vue-router'

Vue.use(Router)

let base = '/';

if (location.pathname && location.pathname != '/') {
    base = location.pathname.split('/').slice(0, -1).join('/');
}

let routes = [
    {
        path: '/sync-data',
        component: () => import(/* webpackChunkName: "sync-data" */ './page/sync-data')
    },
    {
        path: '/dashboard',
        component: () => import(/* webpackChunkName: "dashboard" */ './page/dashboard')
    },
    {
        path: '/cons_plan',
        component: () => import(/* webpackChunkName: "cons_plan" */ './page/tables/cons_plan')
    },
    {
        path: '/year_round_trans',
        component: () => import(/* webpackChunkName: "year_round_trans" */ './page/tables/year_round_trans')
    },
    {
        path: '/tuluvluguu',
        component: () => import(/* webpackChunkName: "tuluvluguu" */ './page/tables/tuluvluguu')
    },
]

export default new Router({
    base,
    routes,
    history: true,
    linkActiveClass: 'active'
})
