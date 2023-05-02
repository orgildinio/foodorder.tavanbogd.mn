import{I as m,c as g}from"./queries.513f850b.js";import{a as x,$ as y,o as e,e as o,F as l,q as c,h as t,t as s,i as b}from"./entry.8ecb60d7.js";import{A as v,a as k}from"./AppleOutlined.79c31608.js";const E={name:"breakfast",props:["selected_type","selected_date","food_data1"],components:{AppleOutlined:v,AndroidOutlined:k},setup(){},data(){return{dayjs:y,activeKey:"1",IMAGE_URL:m,food_data:[]}},mounted(){},methods:{selectedCategory(){this.$apollo.query({query:g,fetchPolicy:"no-cache",variables:{ids:selected_type}}).then(({data:n})=>{})},getBreakFast(n,r){this.loader=!0,this.breakfast_data={},this.$apollo.query({query:g,fetchPolicy:"no-cache",variables:{ids:n,srch_date:r}}).then(({data:d})=>{this.food_data=d.gal_togoo_neg_main_menu,console.log("data.gal_togoo_neg_main_menu >>> ",d.gal_togoo_neg_main_menu)})}},watch:{}},q={class:"grid sm:grid-cols-1 md:grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 gap-5"},w={class:"text-center py-[10px] bg-green-600 rounded-t-[15px] font-bold text-[17px] text-gray-200"},A={class:"p-4"},B=t("div",{class:"pb-4 font-semibold text-gray-500"},"Гал тогоо 1",-1),F={class:"font-bold mb-[8px]"},L={class:"grid grid-cols-2 gap-3"},$={class:"h-[82px] rounded-[15px] relative"},G=["src"],I=t("div",{class:"bg-slate-900 opacity-60 rounded-lg absolute left-0 top-0 right-0 bottom-0"},null,-1),O={class:"absolute bottom-0 left-0 right-0 top-0"},T={class:"flex items-center justify-center text-center text-white w-full h-full text-sm font-semibold uppercase"},j={class:"p-3"},D={class:"grid grid-cols-2"},M=t("span",{class:"font-bold"},"Калори:",-1),R={class:"text-amber-600 font-semibold text-right"},S={class:"mt-1"},U=t("span",{class:"font-semibold"},"Орц:",-1),C={class:"text-right mt-3"},N=t("span",{class:"font-bold"},"Үлдэгдэл: ",-1),P={class:"text-red-600 font-semibold w-full"};function V(n,r,d,K,h,z){return e(),o("div",q,[(e(!0),o(l,null,c(d.food_data1,(i,p)=>(e(),o("div",{class:"border bg-gray-100 shadow-2xl rounded-[15px]",key:p},[t("div",w,s(i.set_name),1),t("div",A,[B,(e(!0),o(l,null,c(i.gal_togoo_neg_sub_menu,(_,f)=>(e(),o("div",{class:"border-b-2 border-gray-200 pb-3 mb-3",key:f},[t("div",F,s(_.food_type),1),t("div",L,[(e(!0),o(l,null,c(_.gal_togoo_neg_sub_menu_foods,(a,u)=>(e(),o("div",{class:"",key:u},[t("div",$,[t("img",{class:"object-cover w-full rounded-lg",style:{height:"82px!important"},src:h.IMAGE_URL+a.food_iamge},null,8,G),I,t("div",O,[t("div",T,s(a.food_name),1)])]),t("div",j,[t("div",D,[M,t("span",R,s(a.food_calorie)+" ккал",1)]),t("div",S,[U,b(" "+s(a.food_ingredients),1)]),t("div",C,[N,t("span",P,s(a.quantity_gt_neg),1)])])]))),128))])]))),128))])]))),128))])}const W=x(E,[["render",V]]);export{W as default};