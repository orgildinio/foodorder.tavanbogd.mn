"use strict";(self.webpackChunk=self.webpackChunk||[]).push([[2943,1544,5311],{57936:(t,e,r)=>{function n(t){if(t.filterWithUser){if(t.filterWithUser&&t.filterWithUser.constructor===Array)t.filterWithUser.forEach((function(e){var r="".concat(e.tableField," = '").concat(window.init.user[e.userField],"'");""==t.filter||void 0===t.filter?t.filter=r:t.filter=t.filter+" AND "+r}));else{var e="".concat(t.filterWithUser.tableField," = '").concat(window.init.user[t.filterWithUser.userField],"'");""==t.filter||void 0===t.filter?t.filter=e:t.filter=t.filter+" AND "+e}t.filterWithUser=void 0}return t}r.d(e,{_:()=>n})},31702:(t,e,r)=>{r.r(e),r.d(e,{default:()=>l});var n=r(91317);const a=r.n(n)().extend({computed:{baseUrl:function(){return this.params.baseUrl?this.params.baseUrl:""}}});const l=(0,r(51900).Z)(a,(function(){var t=this,e=t._self._c;t._self._setupProxy;return e("img",{staticClass:"ag-grid-image",attrs:{src:"".concat(t.baseUrl).concat(t.params.value)}})}),[],!1,null,null,null).exports},67476:(t,e,r)=>{r.r(e),r.d(e,{default:()=>d});var n=r(91317),a=r.n(n),l=r(9669),i=r.n(l),o=r(99219),s=r(57936);function u(t){return function(t){if(Array.isArray(t))return c(t)}(t)||function(t){if("undefined"!=typeof Symbol&&null!=t[Symbol.iterator]||null!=t["@@iterator"])return Array.from(t)}(t)||function(t,e){if(!t)return;if("string"==typeof t)return c(t,e);var r=Object.prototype.toString.call(t).slice(8,-1);"Object"===r&&t.constructor&&(r=t.constructor.name);if("Map"===r||"Set"===r)return Array.from(t);if("Arguments"===r||/^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(r))return c(t,e)}(t)||function(){throw new TypeError("Invalid attempt to spread non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")}()}function c(t,e){(null==e||e>t.length)&&(e=t.length);for(var r=0,n=new Array(e);r<e;r++)n[r]=t[r];return n}const f=a().extend({data:function(){return{options:[],selected:"",currentValue:null}},created:function(){var t=this;if(!this.params.isClient){var e="/lambda/krud/".concat(this.params.schemaID,"/options");i().post(e,(0,s._)(this.params.column.filter.relation)).then((function(e){var r=e.data;t.options=r,t.loading=!1}))}},methods:{element:o.b,setMeta:function(t){return t.schemaID=this.params.schemaID,t},getValues:function(){var t=this;console.log("here"),console.log(this.params.api),this.params.api.forEachLeafNode((function(e){t.options.push(e.data[t.params.column.model])})),this.options=u(new Set(this.options))},setValue:function(t){this.selected=t,this.params.filterData(this.params.column.model,t,"equals")},valueChanged:function(t){this.params.isClient?this.params.filterData(this.params.column.model,t.target.value,"contains"):this.params.filterData(this.params.column.model,t.target.value)},onParentModelChanged:function(t){this.currentValue=t?t.filter:null}}});const d=(0,r(51900).Z)(f,(function(){var t=this,e=t._self._c;t._self._setupProxy;return e("div",{staticClass:"selectable-input-filter"},[e("input",{directives:[{name:"model",rawName:"v-model",value:t.selected,expression:"selected"}],ref:"sinput",domProps:{value:t.selected},on:{input:[function(e){e.target.composing||(t.selected=e.target.value)},t.valueChanged]}}),t._v(" "),e("div",{staticClass:"selectable-input-filter-arrow"},[e("Poptip",{on:{"on-popper-show":t.getValues}},[e("div",{staticClass:"arrow-down"}),t._v(" "),e("div",{attrs:{slot:"content"},slot:"content"},[e("ul",{staticClass:"selectable-input-filter-list"},t._l(t.options,(function(r){return e("li",{key:r.index,on:{click:function(e){return t.setValue(r)}}},[t._v("\n                        "+t._s(r)+"\n                    ")])})),0)])])],1)])}),[],!1,null,null,null).exports},98127:(t,e,r)=>{r.r(e),r.d(e,{default:()=>a});const n={props:["model","label"]};const a=(0,r(51900).Z)(n,(function(){var t=this,e=t._self._c;return e("FormItem",{attrs:{label:t.label}},[e("Input",{attrs:{type:"text"},model:{value:t.model.form[t.model.component],callback:function(e){t.$set(t.model.form,t.model.component,e)},expression:"model.form[model.component]"}})],1)}),[],!1,null,null,null).exports}}]);