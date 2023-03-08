import{_ as M,S as L,a as S,b as $,c as B}from"./SideMenu.95e05123.js";import{p as V,q as d,r as R,t as W,v as x,w as G,x as O,y as Y,i as r,z as j,B as q,C as F,D as I,E as y,a as J,o as K}from"./entry.e5ab3232.js";import{d as Q,m as b,ag as U,j as X,w as T,o as Z,n as ee,Y as n,a3 as l,$ as s,a5 as w,V as i,u as C,c as t,a2 as m,a0 as p,k as ae,a7 as k}from"./vue.8575f319.js";import"./ant.8a0ad9d7.js";import"./SettingOutlined.d48bd242.js";import"./ck.69e8a6ca.js";import"./common.77f0e8e0.js";import"./moment.8b5e7d95.js";import"./ag.2d2e5170.js";import"./lodash.007e3bbb.js";import"./numeral.183422ec.js";import"./cryptoJs.c870fd31.js";const oe=Q({name:"BasicLayout",components:{MultiTab:M,SideMenu:L,GlobalHeader:S,GlobalFooter:$,SettingDrawer:B},setup(){const e=V(),a=b(!1);b([]);const c=U(),v=X(()=>!y.value||r.value?"0":d.value?"256px":"70px");d.value?a.value=!1:a.value=!0,T(()=>d.value,o=>{a.value=!o}),T(()=>e.fullPath,o=>{r.value&&u()}),Z(()=>{navigator.userAgent.indexOf("Edge")>-1&&ee(()=>{a.value=!a.value,setTimeout(()=>{a.value=!a.value},16)})});const _=()=>{a.value=!a.value,c.commit(q,!a.value),F()},g=()=>{let o="";return d.value?o=I.value?"256px":"70px":o=r.value&&"0"||y.value&&"70px"||"0",o},u=()=>{a.value=!1};return{collapsed:a,contentPaddingLeft:v,toggle:_,paddingCalc:g,drawerClose:u,multiTab:R,device:W,layoutMode:x,contentWidth:G,fixedHeader:O,navTheme:Y,isMobile:r,isSideMenu:j}}}),te={class:"flex"};function ne(e,a,c,v,_,g){const u=i("side-menu"),o=i("a-drawer"),E=S,N=M,f=i("portal-target"),P=K,z=i("a-layout-content"),A=$,D=i("a-layout-footer"),H=B,h=i("a-layout");return n(),l(h,{class:w(["layout",e.device])},{default:s(()=>[C(r)?(n(),l(o,{key:0,placement:"left",wrapClassName:`drawer-sider ${e.navTheme}`,closable:!1,visible:e.collapsed,width:"275",onClose:e.drawerClose},{default:s(()=>[t(u,{mode:"inline",theme:e.navTheme,collapsed:!1,collapsible:!0,onToggle:e.toggle,device:e.device},null,8,["theme","onToggle","device"])]),_:1},8,["wrapClassName","visible","onClose"])):e.isSideMenu()?(n(),l(u,{key:1,mode:"inline",theme:e.navTheme,collapsed:e.collapsed,collapsible:!0,onToggle:e.toggle,device:e.device},null,8,["theme","collapsed","onToggle","device"])):m("",!0),t(h,{class:w([e.layoutMode,`content-width-${e.contentWidth}`]),style:k({paddingLeft:e.contentPaddingLeft,minHeight:"100vh"})},{default:s(()=>[t(E,{mode:e.layoutMode,theme:e.navTheme,collapsed:e.collapsed,device:e.device,onToggle:e.toggle},null,8,["mode","theme","collapsed","device","onToggle"]),t(z,null,{default:s(()=>[p("div",{class:"bg-slate-60 dark:bg-slate-800",style:k(`height: 100%; padding: ${e.fixedHeader?e.layoutMode==="levelmenu"?"20px":"87px":"20px"} 24px 20px;`)},[e.multiTab?(n(),l(N,{key:0})):m("",!0),t(ae,{name:"page-transition"},{default:s(()=>[p("section",null,[C(r)||e.layoutMode==="topmenu"?(n(),l(f,{key:0,name:"mobile-page-title"})):m("",!0),p("div",te,[t(f,{name:"level-menu"}),(n(),l(P,{key:e.$route.path}))])])]),_:1})],4)]),_:1}),t(D,{class:""},{default:s(()=>[t(A)]),_:1}),t(H)]),_:1},8,["class","style"])]),_:1},8,["class"])}const fe=J(oe,[["render",ne]]);export{fe as default};