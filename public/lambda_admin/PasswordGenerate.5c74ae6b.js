import{a as u,H as f,r as a,o as _,c as g,w as s,f as t,i as v,t as b,h}from"./entry.82d5962d.js";const w={mixins:[f],computed:{lang(){const e=["Create_a_password"];return e.reduce((o,n,l)=>(o[n]=this.$t("dataForm."+e[l]),o),{})}},methods:{generatePass(){this.model.form[this.model.component]=Math.floor(1e5+Math.random()*9e5)}}};function C(e,o,n,l,$,r){const m=a("inline-svg"),p=a("a-tooltip"),i=a("a-input-password"),c=a("lambda-form-item");return _(),g(c,{rule:e.rule,label:e.label,name:e.model.component,meta:e.meta},{default:s(()=>[t(i,{value:e.model.form[e.model.component],"onUpdate:value":o[1]||(o[1]=d=>e.model.form[e.model.component]=d),password:"",placeholder:e.placeholder,disabled:e.disabled},{addonAfter:s(()=>[t(p,{slot:"append"},{title:s(()=>[v(b(r.lang.Create_a_password),1)]),default:s(()=>[h("span",{class:"svg-icon",onClick:o[0]||(o[0]=d=>r.generatePass())},[t(m,{src:"/assets/icons/duotone/Home/Key.svg"})])]),_:1})]),_:1},8,["value","placeholder","disabled"])]),_:1},8,["rule","label","name","meta"])}const B=u(w,[["render",C]]);export{B as default};