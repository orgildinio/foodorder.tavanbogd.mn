import{_ as v,S as E,a as g,b as h,c as b}from"./SideMenu.fa5c568b.js";import{F as n,G as B,H,l as N,I as W,J as A,N as P,i as s,O as V,Q as D,R as G,V as O,W as m,a as R,E as j}from"./entry.6fd2a578.js";import{d as z,n as p,aj as F,j as I,w as L,o as J,s as Q,Z as l,a1 as r,a0 as o,a8 as f,W as u,c as a,$ as i,a2 as Y,k as Z}from"./vue.65c51cdf.js";import"./ant.770c87db.js";import"./SettingOutlined.193b4281.js";import"./RenderSubMenu.3f64804f.js";import"./ck.660397d8.js";import"./common.72b693eb.js";import"./moment.8b5e7d95.js";import"./ag.23ecb9ea.js";import"./lodash.d251a0a5.js";import"./numeral.1f6cf9d6.js";import"./cryptoJs.f0105b17.js";const q=z({name:"BasicLayout",components:{MultiTab:v,SideMenu:E,GlobalHeader:g,GlobalFooter:h,SettingDrawer:b},setup(){const e=p(!1);p([]);const c=F(),d=I(()=>!m.value||s.value?"0":n.value?"256px":"70px");return L(()=>n.value,t=>{e.value=!t}),J(()=>{navigator.userAgent.indexOf("Edge")>-1&&Q(()=>{e.value=!e.value,setTimeout(()=>{e.value=!e.value},16)}),n.value||(e.value=!0)}),{collapsed:e,contentPaddingLeft:d,toggle:()=>{e.value=!e.value,c.commit(D,!e.value),G()},paddingCalc:()=>{let t="";return n.value?t=O.value?"256px":"70px":t=s.value&&"0"||m.value&&"70px"||"0",t},drawerClose:()=>{e.value=!1},multiTab:B,device:H,layoutMode:N,contentWidth:W,fixedHeader:A,navTheme:P,isMobile:s,isSideMenu:V}}}),K={class:"bg-slate-60 dark:bg-slate-800",style:"height: 100%; "},U={class:"flex"};function X(e,c,d,y,x,T){const t=g,w=v,C=j,k=u("a-layout-content"),S=h,M=u("a-layout-footer"),$=b,_=u("a-layout");return l(),r(_,{class:f(["layout",e.device])},{default:o(()=>[a(_,{class:f([e.layoutMode,`content-width-${e.contentWidth}`]),style:{minHeight:"100vh"}},{default:o(()=>[a(t,{class:"module-page",theme:e.navTheme,collapsed:!0,moduleHeader:!0,device:e.device,onToggle:e.toggle},null,8,["theme","device","onToggle"]),a(k,null,{default:o(()=>[i("div",K,[e.multiTab?(l(),r(w,{key:0})):Y("",!0),a(Z,{name:"page-transition"},{default:o(()=>[i("section",null,[i("div",U,[(l(),r(C,{key:e.$route.fullPath}))])])]),_:1})])]),_:1}),a(M,{class:""},{default:o(()=>[a(S)]),_:1}),a($)]),_:1},8,["class"])]),_:1},8,["class"])}const me=R(q,[["render",X]]);export{me as default};