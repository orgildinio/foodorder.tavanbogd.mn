import{_ as C,S as z,a as k,b as S,c as $}from"./SideMenu.c9b3c425.js";import{u as D,m as R,n as V,o as W,p as q,q as G,r as O,i as d,t as Y,S as j,v as F,w as u,x as I,y,a as U}from"./entry.c548f030.js";import{d as x,q as b,ai as J,j as K,w as T,o as Q,n as X,U as n,Y as l,a8 as s,a4 as i,c as t,a7 as m,_ as p,k as Z,$ as w,aa as M}from"./vue.9a5b07a3.js";import"./ant.9c0b679d.js";import"./ck.b19d9291.js";import"./cryptoJs.12d0d048.js";import"./common.a3ab4f3b.js";import"./moment.8b5e7d95.js";import"./ag.16cfe0bb.js";import"./lodash.2101602c.js";import"./numeral.08866ac3.js";const ee=x({name:"BasicLayout",components:{MultiTab:C,SideMenu:z,GlobalHeader:k,GlobalFooter:S,SettingDrawer:$},setup(){const e=D(),a=b(!1);b([]);const c=J(),g=K(()=>!y.value||d.value?"0":u.value?"256px":"70px");T(()=>u.value,o=>{a.value=!o}),T(()=>e.fullPath,o=>{d.value&&r()}),Q(()=>{navigator.userAgent.indexOf("Edge")>-1&&X(()=>{a.value=!a.value,setTimeout(()=>{a.value=!a.value},16)})});const v=()=>{a.value=!a.value,c.commit(j,!a.value),F()},_=()=>{let o="";return u.value?o=I.value?"256px":"70px":o=d.value&&"0"||y.value&&"70px"||"0",o},r=()=>{a.value=!1};return{collapsed:a,contentPaddingLeft:g,toggle:v,paddingCalc:_,drawerClose:r,multiTab:R,device:V,layoutMode:W,contentWidth:q,fixedHeader:G,navTheme:O,isMobile:d,isSideMenu:Y}}}),ae={class:"flex"};function oe(e,a,c,g,v,_){const r=n("side-menu"),o=n("a-drawer"),N=k,P=C,f=n("portal-target"),B=n("NuxtPage"),E=n("a-layout-content"),A=S,H=n("a-layout-footer"),L=$,h=n("a-layout");return l(),s(h,{class:M(["layout",e.device])},{default:i(()=>[e.isMobile?(l(),s(o,{key:0,placement:"left",wrapClassName:`drawer-sider ${e.navTheme}`,closable:!1,visible:e.collapsed,width:"275",onClose:e.drawerClose},{default:i(()=>[t(r,{mode:"inline",theme:e.navTheme,collapsed:!1,collapsible:!0,onToggle:e.toggle,device:e.device},null,8,["theme","onToggle","device"])]),_:1},8,["wrapClassName","visible","onClose"])):e.isSideMenu()?(l(),s(r,{key:1,mode:"inline",theme:e.navTheme,collapsed:e.collapsed,collapsible:!0,onToggle:e.toggle,device:e.device},null,8,["theme","collapsed","onToggle","device"])):m("",!0),t(h,{class:M([e.layoutMode,`content-width-${e.contentWidth}`]),style:w({paddingLeft:e.contentPaddingLeft,minHeight:"100vh"})},{default:i(()=>[t(N,{mode:e.layoutMode,theme:e.navTheme,collapsed:e.collapsed,device:e.device,onToggle:e.toggle},null,8,["mode","theme","collapsed","device","onToggle"]),t(E,null,{default:i(()=>[p("div",{class:"bg-slate-60 dark:bg-slate-800",style:w(`height: 100%; padding: ${e.fixedHeader?e.layoutMode==="levelmenu"?"20px":"87px":"20px"} 24px 20px;`)},[e.multiTab?(l(),s(P,{key:0})):m("",!0),t(Z,{name:"page-transition"},{default:i(()=>[p("section",null,[e.isMobile||e.layoutMode==="topmenu"?(l(),s(f,{key:0,name:"mobile-page-title"})):m("",!0),p("div",ae,[t(f,{name:"level-menu"}),(l(),s(B,{key:e.$route.path}))])])]),_:1})],4)]),_:1}),t(H,{class:""},{default:i(()=>[t(A)]),_:1}),t(L)]),_:1},8,["class","style"])]),_:1},8,["class"])}const ge=U(ee,[["render",oe]]);export{ge as default};