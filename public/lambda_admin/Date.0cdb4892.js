import{m as t}from"./_mixin.aa4b64aa.js";import{a as d,r as a,o as r,c as s,w as p,f as i}from"./entry.8ecb60d7.js";const c={mixins:[t]};function f(e,o,u,b,_,v){const l=a("a-date-picker"),m=a("lambda-form-item");return r(),s(m,{label:e.label,name:e.model.component,meta:e.meta},{default:p(()=>[i(l,{value:e.model.form[e.model.component],"onUpdate:value":o[0]||(o[0]=n=>e.model.form[e.model.component]=n),mode:"date",placeholder:e.placeholder,disabled:e.disabled,style:{width:"100%"},valueFormat:"YYYY-MM-DD"},null,8,["value","placeholder","disabled"])]),_:1},8,["label","name","meta"])}const Y=d(c,[["render",f]]);export{Y as default};