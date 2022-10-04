"use strict";(self.webpackChunk=self.webpackChunk||[]).push([[9451],{55210:(t,a,s)=>{s.d(a,{Z:()=>i});var e=s(22964);const r={props:["title"],methods:{beforeMount:function(){"mn"!=this.selectedLang&&(0,e.loadLanguageAsync)(this.selectedLang)},switchLanguage:function(t){this.selectedLang=t,(0,e.loadLanguageAsync)(t)}},computed:{lang:function(){var t=this,a=["applications","additions","logistics","letter","fileDirectory","admin","systemAdministrator","personalInformation","settings","logOut","ui_builder","_form"];return a.reduce((function(s,e,r){return s[e]=t.$t("components."+a[r]),s}),{})}}};const i=(0,s(51900).Z)(r,(function(){var t=this,a=t._self._c;return a("subheader",[a("div",{attrs:{slot:"left"},slot:"left"},[a("h3",{staticClass:"title"},[t._v("\n            "+t._s(t.title)+"\n        ")])]),t._v(" "),a("div",{attrs:{slot:"right"},slot:"right"},[a("ul",[a("li",[a("Badge",{attrs:{dot:""}},[a("a",{staticClass:"demo-badge",attrs:{href:"#"}},[a("Icon",{attrs:{type:"ios-bell-outline"}})],1)])],1),t._v(" "),a("li",[a("Poptip",{attrs:{placement:"bottom-end"}},[a("a",{attrs:{href:"javascript:void(0)"}},[a("Icon",{attrs:{type:"ios-keypad"}})],1),t._v(" "),a("div",{staticClass:"header-apps",attrs:{slot:"content"},slot:"content"},[a("Tabs",{attrs:{value:"apps",size:"small"}},[a("TabPane",{attrs:{label:t.lang.applications,name:"apps"}},[a("Row",[a("Col",{attrs:{span:"8"}},[a("a",{attrs:{href:"/logistic"}},[a("img",{attrs:{src:"/images/apps/github.png",alt:""}}),t._v(" "),a("span",[t._v("CMS")])])]),t._v(" "),a("Col",{attrs:{span:"8"}},[a("a",{attrs:{href:"/logistic"}},[a("img",{attrs:{src:"/images/apps/slack.png",alt:""}}),t._v(" "),a("span",[t._v("HR")])])]),t._v(" "),a("Col",{attrs:{span:"8"}},[a("a",{attrs:{href:"/logistic"}},[a("img",{attrs:{src:"/images/apps/bitbucket.png",alt:""}}),t._v(" "),a("span",[t._v("CRM")])])]),t._v(" "),a("Col",{attrs:{span:"8"}},[a("a",{attrs:{href:"/logistic"}},[a("img",{attrs:{src:"/images/apps/dribbble.png",alt:""}}),t._v(" "),a("span",[t._v(t._s(t.lang.logistics))])])]),t._v(" "),a("Col",{attrs:{span:"8"}},[a("a",{attrs:{href:"/logistic"}},[a("img",{attrs:{src:"/images/apps/mail_chimp.png",alt:""}}),t._v(" "),a("span",[t._v(t._s(t.lang.letter))])])]),t._v(" "),a("Col",{attrs:{span:"8"}},[a("a",{attrs:{href:"/logistic"}},[a("img",{attrs:{src:"/images/apps/dropbox.png",alt:""}}),t._v(" "),a("span",[t._v(t._s(t.lang.fileDirectory))])])])],1)],1),t._v(" "),a("TabPane",{attrs:{label:t.lang.additions,name:"plugins"}},[a("Row",[a("Col",{attrs:{span:"8"}},[a("a",{attrs:{href:"/lambda/puzzle"}},[a("img",{attrs:{src:"/images/apps/github.png",alt:""}}),t._v(" "),a("span",[t._v(t._s(t.lang.ui_builder))])])]),t._v(" "),a("Col",{attrs:{span:"8"}},[a("a",{attrs:{href:"/logistic"}},[a("img",{attrs:{src:"/images/apps/mail_chimp.png",alt:""}}),t._v(" "),a("span",[t._v(t._s(t.lang._form))])])])],1)],1)],1)],1)])],1),t._v(" "),a("li",{staticClass:"avatar-item"},[a("Poptip",{attrs:{placement:"bottom-end"}},[a("a",{staticClass:"avatar",attrs:{href:"javascript:void(0)"}},[a("img",{staticClass:"avatar",attrs:{src:"/assets/lambda/images/avatar.png",alt:"avatar"}})]),t._v(" "),a("div",{staticClass:"header-profile",attrs:{slot:"content"},slot:"content"},[a("div",{staticClass:"header-profile-info"},[a("h3",[t._v(t._s(t.lang.admin))]),t._v(" "),a("small",[t._v(t._s(t.lang.systemAdministrator))])]),t._v(" "),a("ul",[a("li",[a("a",{attrs:{href:"#"}},[a("Icon",{attrs:{type:"ios-color-filter-outline"}}),t._v(" "),a("span",[t._v(t._s(t.lang.personalInformation))])],1)]),t._v(" "),a("li",[a("a",{attrs:{href:"#"}},[a("Icon",{attrs:{type:"ios-gear-outline"}}),t._v(" "),a("span",[t._v(t._s(t.lang.settings))])],1)]),t._v(" "),a("li",[a("a",{attrs:{href:"#"}},[a("Icon",{attrs:{type:"ios-locked-outline"}}),t._v(" "),a("span",[t._v(t._s(t.lang.logOut))])],1)])])])])],1)])])])}),[],!1,null,null,null).exports},93:(t,a,s)=>{s.r(a),s.d(a,{default:()=>r});const e={components:{"page-header":s(55210).Z}};const r=(0,s(51900).Z)(e,(function(){var t=this,a=t._self._c;return a("section",{staticClass:"page page-preview"},[a("div",{staticClass:"pz-grid-preview"},[t.$project?a("datagrid",{attrs:{projectID:t.$project.id,schemaID:t.$route.params.id,paginate:50}}):a("datagrid",{attrs:{schemaID:t.$route.params.id,paginate:50}})],1)])}),[],!1,null,null,null).exports},73391:(t,a,s)=>{s.r(a),s.d(a,{default:()=>r});const e={components:{"page-header":s(55210).Z}};const r=(0,s(51900).Z)(e,(function(){var t=this,a=t._self._c;return a("section",{staticClass:"page page-preview"},[a("div",{staticClass:"pz-grid-preview"},[t.$project?a("datagrid",{attrs:{projectID:t.$project.id,schemaID:t.$route.params.id,paginate:50}}):a("datagrid",{attrs:{schemaID:t.$route.params.id,paginate:50}})],1)])}),[],!1,null,null,null).exports}}]);