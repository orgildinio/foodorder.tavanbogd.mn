import{m as p}from"./_mixin.4f22079c.js";import{a as d}from"./entry.f00561da.js";import{Z as a,a1 as r,a0 as i,W as l}from"./vue.65c51cdf.js";import"./common.72b693eb.js";import"./ant.770c87db.js";import"./moment.8b5e7d95.js";import"./ag.ab39fb93.js";import"./lodash.d251a0a5.js";import"./numeral.1f6cf9d6.js";import"./cryptoJs.f0105b17.js";const s={mixins:[p]};function u(e,m,f,b,v,$){const n=l("a-input-number"),t=l("lambda-form-item");return a(),r(t,{label:e.label,name:e.model.component,meta:e.meta},{default:i(()=>[e.meta.no_format?(a(),r(n,{key:0,type:"number",value:e.model.form[e.model.component],"onUpdate:value":m[0]||(m[0]=o=>e.model.form[e.model.component]=o),disabled:e.disabled,number:!0},null,8,["value","disabled"])):(a(),r(n,{key:1,value:e.model.form[e.model.component],"onUpdate:value":m[1]||(m[1]=o=>e.model.form[e.model.component]=o),formatter:o=>`${o}`.replace(/\B(?=(\d{3})+(?!\d))/g,","),parser:o=>o.replace(/\$\s?|(,*)/g,""),disabled:e.disabled},null,8,["value","formatter","parser","disabled"]))]),_:1},8,["label","name","meta"])}const Z=d(s,[["render",u]]);export{Z as default};