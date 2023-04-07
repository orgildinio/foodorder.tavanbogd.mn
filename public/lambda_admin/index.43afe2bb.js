import{H as w,T as A}from"./components.7e1849ad.js";import{S as H,g as S,a as T,s as k,K as R,P as C,L as N,M as E,i as B,b as P,d as L}from"./entry.9b01c438.js";import{S as V}from"./SettingOutlined.67f581ce.js";import{d as F,X as s,Y as t,a4 as f,a1 as r,a2 as a,_ as u,a0 as d,F as g,$ as v,Z as M,S as y,E as _,c as $}from"./vue.300029e2.js";import"./common.01c9107c.js";import"./ant.d96fdebf.js";import"./moment.8b5e7d95.js";import"./ag.1d467b62.js";import"./lodash.ccbc03fa.js";import"./numeral.6cca0570.js";import"./cryptoJs.b65e2552.js";const O=F({name:"HomeMenuRender",props:["item","cruds","permissions","title","children"],components:{SvgIcon:H,SettingOutlined:V},data(){return{}},mounted(){this.setActive()},methods:{getModuleItem(e){if(e.children)return e},findActivehild(e){let n=0;for(;!this.can(e.children[n]);)n++;return{...e.children[n],svg:e.svg,icon:e.icon}},setActive(){this.hasItems(this.item)&&(this.item.children.findIndex(n=>n.id===this.$route.params.menu_id)>=0?this.selectMain(this.item,this.getTitle(this.item)):this.item.children.forEach(n=>{n.children.findIndex(i=>i.id===this.$route.params.menu_id)>=0&&this.selectMain(this.item,this.getTitle(this.item))}))},getPath(e){return S(e)},findActiveMenu(e,n,l){e.forEach(i=>{i.children?i.children.length>=1?this.findActiveMenu(i.children,`${n}/${i.id}`,`${l}${i.id}`):this.setActiveMenu(i,n,l):this.setActiveMenu(i,n,l)})},setActiveMenu(e,n,l){(e.link_to=="iframe"||e.link_to=="crud")&&this.cleanPath==`${n}/${e.id}`&&l!=""&&(this.menu_open=[l])},can(e){return this.permissions[e.id]?!!this.permissions[e.id].show:!1},getTitle(e,n){if(this.title&&!n)return this.title;if(e.link_to=="crud"){let l=this.cruds.findIndex(i=>i.id==e.url);return l>=0?this.cruds[l].title:""}else return e.title},hasItems(e){return e&&e.children!==void 0?e.children.length>0:!1}}});const D=["href"],K={key:2,class:"module-menus"},U={key:0},z={key:2,class:"module-menus"},G={key:0},X={key:2,class:"module-menus"},Y={key:0};function Z(e,n,l,i,c,b){const h=y("inline-svg"),p=y("router-link"),I=y("HomeMenuRender",!0);return s(),t(g,null,[e.can(e.item)&&!e.hasItems(e.item)&&e.item.link_to==="link"?(s(),t("a",{key:0,class:"card p-4 sm:p-5 shadow-md bg-white dark:bg-slate-900 bg-white dark:bg-slate-900",href:e.item.url,target:"_blank"},[e.item.icon?(s(),t("i",{key:0,class:f(e.item.icon)},null,2)):r("",!0),e.item.svg?(s(),a(h,{key:1,class:"svg-icon",src:e.item.svg},null,8,["src"])):r("",!0),u("span",null,d(e.getTitle(e.item)),1),e.children?(s(),t("p",K,[(s(!0),t(g,null,v(e.children,(o,m)=>(s(),t("span",{key:o.index},[m>=1?(s(),t("span",U,", ")):r("",!0),_(d(e.getTitle(o,!0)),1)]))),128))])):r("",!0)],8,D)):e.can(e.item)&&!e.hasItems(e.item)&&e.item.link_to==="router-link"?(s(),a(p,{key:1,to:e.item.url,class:"card p-4 sm:p-5 shadow-md bg-white dark:bg-slate-900 bg-white dark:bg-slate-900"},{default:M(()=>[e.item.icon?(s(),t("i",{key:0,class:f(e.item.icon)},null,2)):r("",!0),e.item.svg?(s(),a(h,{key:1,class:"svg-icon",src:e.item.svg},null,8,["src"])):r("",!0),u("span",null,d(e.getTitle(e.item)),1),e.children?(s(),t("p",z,[(s(!0),t(g,null,v(e.children,(o,m)=>(s(),t("span",{key:o.index},[m>=1?(s(),t("span",G,", ")):r("",!0),_(d(e.getTitle(o,!0)),1)]))),128))])):r("",!0)]),_:1},8,["to"])):e.can(e.item)&&!e.hasItems(e.item)?(s(),a(p,{key:2,to:`/admin/p/${e.item.id}`,class:"card p-4 sm:p-5 shadow-md bg-white dark:bg-slate-900 bg-white dark:bg-slate-900"},{default:M(()=>[e.item.icon?(s(),t("i",{key:0,class:f(e.item.icon)},null,2)):r("",!0),e.item.svg?(s(),a(h,{key:1,class:"svg-icon",src:e.item.svg},null,8,["src"])):r("",!0),u("span",null,d(e.getTitle(e.item)),1),e.children?(s(),t("p",X,[(s(!0),t(g,null,v(e.children,(o,m)=>(s(),t("span",{key:m},[m>=1?(s(),t("span",Y,", ")):r("",!0),_(d(e.getTitle(o,!0)),1)]))),128))])):r("",!0)]),_:1},8,["to"])):r("",!0),e.can(e.item)&&e.hasItems(e.item)?(s(),a(I,{key:3,title:e.getTitle(e.item),children:e.item.children,item:e.findActivehild(e.item),cruds:e.cruds,permissions:e.permissions},null,8,["title","children","item","cruds","permissions"])):r("",!0)],64)}const j=T(O,[["render",Z],["__scopeId","data-v-e82381a3"]]),q={name:"index",components:{HomeMenuRender:j},setup(e,{emit:n}){const l=k.get(R),i=k.get(C),c=k.get(N),b=k.get(E);return{isMobile:B,menu:b,permissions:i,kruds:l,title:c.title,logo:c.logo,logoDark:c.logo_dark,base_url:P,darkMode:L}}},J={class:"grid grow grid-nogutter module-page container pb-16 px-4 mx-auto"},Q=["src","alt"],W=["src","alt"],x={class:"module-navigation"};function ee(e,n,l,i,c,b){const h=A,p=w,I=y("HomeMenuRender");return s(),t("div",J,[$(p,null,{default:M(()=>[$(h,null,{default:M(()=>[_(d(i.title),1)]),_:1})]),_:1}),u("div",null,[i.darkMode?(s(),t("img",{key:0,src:`${i.base_url}${i.logoDark}`,class:"logo-light",alt:i.title},null,8,Q)):r("",!0),i.darkMode?r("",!0):(s(),t("img",{key:1,src:`${i.base_url}${i.logo}`,class:"logo-dark",alt:i.title},null,8,W)),u("nav",x,[u("div",{class:f(i.isMobile?"grid grid-cols-1  gap-4":"grid grid-cols-3  gap-4")},[(s(!0),t(g,null,v(i.menu,o=>(s(),a(I,{key:o.index,item:o,cruds:i.kruds,permissions:i.permissions.permissions},null,8,["item","cruds","permissions"]))),128))],2)])])])}const ce=T(q,[["render",ee]]);export{ce as default};