import{m as t}from"./_mixin.59e23610.js";import{a as n}from"./entry.0dfbcad4.js";import{S as m,Y as p,a8 as i,a4 as s,c as d}from"./vue.a1a1d817.js";import"./cryptoJs.37acf5e1.js";import"./ant.3cb96f1d.js";import"./common.db75f8be.js";import"./numeral.f32c122b.js";import"./moment.8b5e7d95.js";const f={mixins:[t]};function u(e,o,c,b,v,_){const a=m("a-time-picker"),l=m("lambda-form-item");return p(),i(l,{rule:e.rule,label:e.label,name:e.model.component,meta:e.meta},{default:s(()=>[d(a,{value:e.model.form[e.model.component],"onUpdate:value":o[0]||(o[0]=r=>e.model.form[e.model.component]=r),placeholder:e.placeholder,disabled:e.disabled,valueFormat:"HH:mm:ss"},null,8,["value","placeholder","disabled"])]),_:1},8,["rule","label","name","meta"])}const N=n(f,[["render",u]]);export{N as default};