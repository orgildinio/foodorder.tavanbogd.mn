"use strict";(self.webpackChunk=self.webpackChunk||[]).push([[4036],{17366:(e,t,l)=>{l.r(t),l.d(t,{default:()=>i});var a=l(36346);const o={props:["model","rule","label","meta","disabled","relation_data","do_render","showInformationModal"],computed:{lang:function(){var e=this,t=["dataNotFound"];return t.reduce((function(l,a,o){return l[a]=e.$t("dataForm."+t[o]),l}),{})},options:function(){var e=this;return(0,a.JY)(this.meta)&&(0,a.JY)(this.meta.options)&&this.meta.options.length>=1?this.searchval?this.filterOption(this.meta.options).filter((function(t){return t.label.toLowerCase().includes(e.searchval.toLowerCase())})):this.filterOption(this.meta.options):this.searchval?this.filterOption(this.relation_data).filter((function(t){return t.label.toLowerCase().includes(e.searchval.toLowerCase())})):this.filterOption(this.relation_data)}},data:function(){return{value:null,ignoreChange:!1,addAble:!1,clearAble:!1,modal_show:!1,searchval:null}},methods:{isValid:a.JY,filterOption:function(e){var t=this;if(e){if(this.$props.meta.relation.parentFieldOfForm){if(this.$props.model.form[this.$props.meta.relation.parentFieldOfForm]){var l=e.filter((function(e){return e.parent_value==t.$props.model.form[t.$props.meta.relation.parentFieldOfForm]}));return this.initialValue(l),l}return e||[]}return this.initialValue(e),e||[]}return[]},searchChange:function(e){this.searchval=e},initialValue:function(e){var t=this;if(!this.ignoreChange)if(this.model.form[this.model.component])if(1==this.meta.relation.multiple){var l=this.model.form[this.model.component].toString().split(","),a=e.filter((function(e){return l.includes(e.value.toString())}));a.length>=1&&(this.value=a)}else{var o=e.filter((function(e){return e.value==t.model.form[t.model.component]}));this.value=o.length>=1?o[0]:null}else this.value=null},addFromUrl:function(){var e=this;if(window.init.microserviceSettings){var t=window.init.microserviceSettings.findIndex((function(t){return t.project_id==e.meta.relation.addFromMicroservice}));return t>=0?window.init.microserviceSettings[t].production_url:""}return""},showAddModal:function(){this.modal_show=!0},clearState:function(){this.value=null,this.clearAble=!1,Vue.set(this.model.form,this.model.component,null)},closeModal:function(){this.modal_show=!1},onSuccess:function(e){var t=this.meta.relation.fields.map((function(t){return e[t]}));t=t.join(", ");var l={value:e[this.meta.relation.key],label:t};""!==this.meta.relation.parentFieldOfTable&&""!==this.meta.relation.parentFieldOfForm&&(l.parent_value=e[this.meta.relation.parentFieldOfTable].toString()),this.relation_data.push(l),this.closeModal()},onError:function(e){},showInfoModal:function(){this.model.form[this.model.component]?window.showInformationModal("".concat(this.meta.info_url).concat(this.model.form[this.model.component]),this.meta.placeHolder):this.$Message.success(this.lang.dataNotFound)}},watch:{do_render:function(e){if(e){if(this.model.form[this.model.component]){var t=this.model.form[this.model.component];if(1==this.meta.relation.multiple){var l=t.toString().split(",");this.value=this.options.filter((function(e){return l.includes(e.value.toString())}))}else{var a=this.options.filter((function(e){return e.value==t}));this.value=a.length>=1?a[0]:null}this.clearAble=!0}}else this.value=null,this.clearAble=!1,this.ignoreChange=!1,Vue.set(this.model.form,this.model.component,this.meta.default?this.meta.default:void 0)},value:function(e){e&&(1==this.meta.relation.multiple?Vue.set(this.model.form,this.model.component,e.map((function(e){return e.value})).join(",")):""==e.value||null===e.value?Vue.set(this.model.form,this.model.component,null):isNaN(e.value)?Vue.set(this.model.form,this.model.component,e.value):Vue.set(this.model.form,this.model.component,1*e.value),this.clearAble=!0)}},created:function(){this.meta.relation.addAble&&this.meta.relation.addFrom&&(this.addAble=!0)}};const i=(0,l(51900).Z)(o,(function(){var e=this,t=e.$createElement,l=e._self._c||t;return l("FormItem",{attrs:{label:e.label,prop:e.rule}},[e.meta.relation.multiple?l("multiselect",{class:e.meta.info_url?"with-info-caller":"",attrs:{multiple:!0,disabled:e.meta.disabled,"track-by":"value",searchable:!0,placeholder:e.meta&&null!==e.meta.placeHolder?e.meta.placeHolder:e.label?e.label:"",label:"label",options:e.options},on:{"search-change":e.searchChange},scopedSlots:e._u([{key:"caret",fn:function(t){t.toggle;return[l("div",{staticClass:"caret-container"},[l("div",{class:e.addAble?"multiselect__select addable-caret":"multiselect__select"}),e._v(" "),e.addAble?l("Button",{attrs:{type:"success",shape:"circle",size:"small",icon:"md-add"},on:{click:e.showAddModal}}):e._e()],1)]}}]),model:{value:e.value,callback:function(t){e.value=t},expression:"value"}}):l("multiselect",{class:e.meta.info_url?"with-info-caller":"",attrs:{disabled:e.meta.disabled,options:e.options,"track-by":"value",searchable:!0,"allow-empty":!0,placeholder:e.meta&&null!==e.meta.placeHolder?e.meta.placeHolder:e.label?e.label:"",label:"label"},on:{"search-change":e.searchChange},scopedSlots:e._u([{key:"singleLabel",fn:function(t){var l=t.option;return[e._v("\n            "+e._s(l.label)+"\n        ")]}},{key:"caret",fn:function(t){t.toggle;return[l("div",{staticClass:"caret-container"},[l("div",{class:e.addAble?"multiselect__select addable-caret":"multiselect__select"}),e._v(" "),e.addAble?l("Button",{attrs:{type:"success",shape:"circle",size:"small",icon:"md-add"},on:{click:e.showAddModal}}):e._e()],1)]}},{key:"clear",fn:function(t){t.search;return[l("div",{staticClass:"clear-container"},[e.clearAble?l("Button",{attrs:{shape:"circle",size:"small",icon:"md-close"},on:{click:e.clearState}}):e._e()],1)]}}],null,!1,1374386606),model:{value:e.value,callback:function(t){e.value=t},expression:"value"}}),e._v(" "),e.meta.info_url?l("div",{staticClass:"info-caller"},[l("Button",{attrs:{shape:"circle",type:"primary",icon:"ios-help-circle",size:"small"},on:{click:e.showInfoModal}})],1):e._e(),e._v(" "),e.addAble?l("Modal",{attrs:{"min-width":200,"min-height":100,draggable:!0,"footer-hide":!0,title:e.label,width:"800",height:"70%"},model:{value:e.modal_show,callback:function(t){e.modal_show=t},expression:"modal_show"}},[e.modal_show?l("section",{staticClass:"add-modal"},[l("div",{staticClass:"add-body"},[l("dataform",{ref:"form",attrs:{schemaID:e.meta.relation.addFrom,editMode:!1,onSuccess:e.onSuccess,url:e.addFromUrl(),do_render:e.modal_show,onError:e.onError}})],1)]):e._e()]):e._e()],1)}),[],!1,null,null,null).exports}}]);
//# sourceMappingURL=form-field-Select.15a71726f69e2787.js.map