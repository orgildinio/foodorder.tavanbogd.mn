import{_ as C,S as L,a as S,b as k,c as $}from"./SideMenu.b1b69c06.js";import{u as D,n as r,o as R,p as V,q as W,r as G,t as O,v as Y,i as d,w as j,S as q,x as F,y as I,z as y,a as x}from"./entry.0dfbcad4.js";import{d as J,l as b,ai as K,f as Q,w as T,o as U,n as X,S as l,Y as n,a8 as s,a4 as i,c as t,a7 as p,_ as c,j as Z,$ as w,aa as M}from"./vue.a1a1d817.js";import"./ant.3cb96f1d.js";import"./ck.b590734d.js";import"./cryptoJs.37acf5e1.js";import"./common.db75f8be.js";import"./numeral.f32c122b.js";import"./moment.8b5e7d95.js";const ee=J({name:"BasicLayout",components:{MultiTab:C,SideMenu:L,GlobalHeader:S,GlobalFooter:k,SettingDrawer:$},setup(){const e=D(),a=b(!1);b([]);const m=K(),v=Q(()=>!y.value||d.value?"0":r.value?"256px":"70px");r.value?a.value=!1:a.value=!0,T(()=>r.value,o=>{a.value=!o}),T(()=>e.fullPath,o=>{d.value&&u()}),U(()=>{navigator.userAgent.indexOf("Edge")>-1&&X(()=>{a.value=!a.value,setTimeout(()=>{a.value=!a.value},16)})});const g=()=>{a.value=!a.value,m.commit(q,!a.value),F()},_=()=>{let o="";return r.value?o=I.value?"256px":"70px":o=d.value&&"0"||y.value&&"70px"||"0",o},u=()=>{a.value=!1};return{collapsed:a,contentPaddingLeft:v,toggle:g,paddingCalc:_,drawerClose:u,multiTab:R,device:V,layoutMode:W,contentWidth:G,fixedHeader:O,navTheme:Y,isMobile:d,isSideMenu:j}}}),ae={class:"flex"};function oe(e,a,m,v,g,_){const u=l("side-menu"),o=l("a-drawer"),N=S,P=C,f=l("portal-target"),B=l("NuxtPage"),E=l("a-layout-content"),z=k,A=l("a-layout-footer"),H=$,h=l("a-layout");return n(),s(h,{class:M(["layout",e.device])},{default:i(()=>[e.isMobile?(n(),s(o,{key:0,placement:"left",wrapClassName:`drawer-sider ${e.navTheme}`,closable:!1,visible:e.collapsed,width:"275",onClose:e.drawerClose},{default:i(()=>[t(u,{mode:"inline",theme:e.navTheme,collapsed:!1,collapsible:!0,onToggle:e.toggle,device:e.device},null,8,["theme","onToggle","device"])]),_:1},8,["wrapClassName","visible","onClose"])):e.isSideMenu()?(n(),s(u,{key:1,mode:"inline",theme:e.navTheme,collapsed:e.collapsed,collapsible:!0,onToggle:e.toggle,device:e.device},null,8,["theme","collapsed","onToggle","device"])):p("",!0),t(h,{class:M([e.layoutMode,`content-width-${e.contentWidth}`]),style:w({paddingLeft:e.contentPaddingLeft,minHeight:"100vh"})},{default:i(()=>[t(N,{mode:e.layoutMode,theme:e.navTheme,collapsed:e.collapsed,device:e.device,onToggle:e.toggle},null,8,["mode","theme","collapsed","device","onToggle"]),t(E,null,{default:i(()=>[c("div",{class:"bg-slate-60 dark:bg-slate-800",style:w(`height: 100%; padding: ${e.fixedHeader?e.layoutMode==="levelmenu"?"20px":"87px":"20px"} 24px 20px;`)},[e.multiTab?(n(),s(P,{key:0})):p("",!0),t(Z,{name:"page-transition"},{default:i(()=>[c("section",null,[e.isMobile||e.layoutMode==="topmenu"?(n(),s(f,{key:0,name:"mobile-page-title"})):p("",!0),c("div",ae,[t(f,{name:"level-menu"}),(n(),s(B,{key:e.$route.path}))])])]),_:1})],4)]),_:1}),t(A,{class:""},{default:i(()=>[t(z)]),_:1}),t(H)]),_:1},8,["class","style"])]),_:1},8,["class"])}const ce=x(ee,[["render",oe]]);export{ce as default};