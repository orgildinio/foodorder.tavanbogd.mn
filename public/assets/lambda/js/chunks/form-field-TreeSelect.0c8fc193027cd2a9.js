"use strict";(self.webpackChunk=self.webpackChunk||[]).push([[84],{73634:(e,t,n)=>{n.r(t),n.d(t,{default:()=>o});var a=function e(t,n,a){return t.filter((function(e){return e.parent_value==n})).map((function(n){return n.children=e(t,n.value,a),n.children.length>=1&&(n.expand=!0),"".concat(n.value)=="".concat(a)?n.selected=!0:n.selected=!1,n.title=n.label,n}))};const l={props:["model","rule","label","meta","disabled","relation_data"],data:function(){return{treeData:[],loading:!1}},computed:{treeValue:function(){return this.model.form[this.model.component]}},watch:{relation_data:function(){this.options()},treeValue:function(){this.loading=!0,this.options()}},methods:{options:function(){var e=this;if(this.relation_data)if(this.relation_data.length>=1){var t=this.relation_data.filter((function(e){return void 0===e.parent_value}));t.map((function(t){return t.children=a(e.relation_data,t.value,e.model.form[e.model.component]),t.children.length>=1&&(t.expand=!0),t.value==e.model.form[e.model.component]?t.selected=!0:t.selected=!1,t.title=t.label,t})),Vue.set(this.$data,"treeData",t),setTimeout((function(){e.loading=!1}),10)}else this.treeData=[];else this.treeData=[]},selected:function(e){e.length>=1&&Vue.set(this.model.form,this.model.component,e[0].value)}}};const o=(0,n(51900).Z)(l,(function(){var e=this,t=e._self._c;return t("FormItem",{attrs:{label:e.label,prop:e.rule}},[t("div",{staticClass:"tree-select"},[e.loading?e._e():t("tree",{attrs:{data:e.treeData},on:{"on-select-change":e.selected}})],1)])}),[],!1,null,null,null).exports}}]);
//# sourceMappingURL=form-field-TreeSelect.0c8fc193027cd2a9.js.map