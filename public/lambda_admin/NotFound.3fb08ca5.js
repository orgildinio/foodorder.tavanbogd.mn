import{m as s}from"./_mixin.abbd934f.js";import{a as i}from"./entry.c8558833.js";import{Y as l,a3 as p,$ as d,a0 as c,_ as m,V as u}from"./vue.091937b4.js";import"./ant.edd72bf4.js";import"./moment.8b5e7d95.js";import"./cryptoJs.ce41d5ee.js";import"./numeral.cf209493.js";import"./common.33f268d0.js";const f={mixins:[s],computed:{lang(){const e=["notFound"];return e.reduce((o,t,a)=>(o[t]=this.$t("dataForm."+e[a]),o),{})}},methods:{}};function _(e,o,t,a,b,r){const n=u("lambda-form-item");return l(),p(n,{label:e.label,name:e.model.component,meta:e.meta},{default:d(()=>[c("div",null,m(r.lang.notFound)+" - "+m(e.meta.formType),1)]),_:1},8,["label","name","meta"])}const y=i(f,[["render",_]]);export{y as default};