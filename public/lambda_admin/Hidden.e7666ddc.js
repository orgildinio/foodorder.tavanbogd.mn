import{a as n}from"./entry.0dfbcad4.js";import{S as l,Y as d,a8 as i,a4 as c,c as r}from"./vue.a1a1d817.js";import"./cryptoJs.37acf5e1.js";import"./ant.3cb96f1d.js";import"./common.db75f8be.js";import"./numeral.f32c122b.js";import"./moment.8b5e7d95.js";const u={props:["model","rule","label","meta"]};function s(_,a,e,f,b,v){const t=l("a-input"),o=l("a-form-item");return d(),i(o,{rules:e.rule,label:e.label,name:e.model.component,style:{visibility:"hidden"}},{default:c(()=>[r(t,{value:e.model.form[e.model.component],"onUpdate:value":a[0]||(a[0]=m=>e.model.form[e.model.component]=m),placeholder:e.meta&&e.meta.placeHolder!==null?e.meta.placeHolder:e.label,disabled:e.meta&&e.meta.disabled?e.meta.disabled:!1},null,8,["value","placeholder","disabled"])]),_:1},8,["rules","label","name"])}const p=n(u,[["render",s]]);export{p as default};