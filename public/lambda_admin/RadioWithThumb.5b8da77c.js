import{a as u,H as p,r as o,o as t,c as r,w as n,f as c,e as f,m as _,h as b,F as h}from"./entry.82d5962d.js";const v={mixins:[p],computed:{options(){return this.meta.options.length>=1?this.meta.options:this.relation_data}}},g=["src"];function k(e,s,B,$,w,l){const m=o("a-radio"),d=o("a-radio-group"),i=o("lambda-form-item");return t(),r(i,{label:e.label,name:e.model.component,meta:e.meta},{default:n(()=>[c(d,{value:e.model.form[e.model.component],"onUpdate:value":s[0]||(s[0]=a=>e.model.form[e.model.component]=a)},{default:n(()=>[(t(!0),f(h,null,_(l.options,a=>(t(),r(m,{value:a.value,key:a.index,disabled:e.disabled},{default:n(()=>[b("img",{src:a.thumb,alt:""},null,8,g)]),_:2},1032,["value","disabled"]))),128))]),_:1},8,["value"])]),_:1},8,["label","name","meta"])}const F=u(v,[["render",k]]);export{F as default};