import{a as s,b as a}from"./entry.e5ab3232.js";import{Y as n,Z as d,c,V as i}from"./vue.8575f319.js";import"./common.77f0e8e0.js";import"./ant.8a0ad9d7.js";import"./moment.8b5e7d95.js";import"./ag.2d2e5170.js";import"./lodash.007e3bbb.js";import"./numeral.183422ec.js";import"./cryptoJs.c870fd31.js";const m={components:{},computed:{url(){return a}},methods:{onSuccess(){},onError(){},getData(){this.editMode=!0,this.$refs.form.editModel(this.$route.params.id)}},data(){return{editMode:!1,formId:null}}},u={class:"form-preview-page"};function _(r,l,p,f,o,e){const t=i("dataform");return n(),d("div",u,[c(t,{ref:"form",schemaID:r.$route.params.schema_id,editMode:o.editMode,hideTitle:!0,onSuccess:e.onSuccess,do_render:!0,onReady:e.getData,url:e.url,permissions:{c:!0,r:!1,u:!1,d:!1},user_condition:null,onError:e.onError},null,8,["schemaID","editMode","onSuccess","onReady","url","onError"])])}const I=s(m,[["render",_]]);export{I as default};