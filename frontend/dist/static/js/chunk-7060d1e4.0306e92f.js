(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-7060d1e4"],{"1c31":function(t,e,n){"use strict";n.d(e,"b",(function(){return r})),n.d(e,"a",(function(){return c})),n.d(e,"c",(function(){return o}));var a=n("b775");function r(t){return Object(a["a"])({url:"/report/comprehensive",method:"post",data:t})}function c(t){return Object(a["a"])({url:"/report/ads",method:"post",data:t})}function o(t){return Object(a["a"])({url:"/report/column",method:"post",data:t})}},2934:function(t,e,n){"use strict";n.d(e,"h",(function(){return r})),n.d(e,"p",(function(){return c})),n.d(e,"d",(function(){return o})),n.d(e,"e",(function(){return i})),n.d(e,"g",(function(){return s})),n.d(e,"f",(function(){return l})),n.d(e,"i",(function(){return u})),n.d(e,"j",(function(){return d})),n.d(e,"l",(function(){return f})),n.d(e,"k",(function(){return p})),n.d(e,"m",(function(){return m})),n.d(e,"a",(function(){return h})),n.d(e,"b",(function(){return g})),n.d(e,"c",(function(){return b})),n.d(e,"o",(function(){return v})),n.d(e,"n",(function(){return _}));var a=n("b775");function r(){return Object(a["a"])({url:"/regions",method:"get"})}function c(){return Object(a["a"])({url:"/settings/version",method:"get"})}function o(){return Object(a["a"])({url:"/region/area",method:"get"})}function i(t){return Object(a["a"])({url:"/region/country",method:"get",params:t})}function s(t){return Object(a["a"])({url:"/region",method:"post",data:t})}function l(t){return Object(a["a"])({url:"/region/area-set",method:"post",data:t})}function u(){return Object(a["a"])({url:"/settings/cron",method:"get"})}function d(t){return Object(a["a"])({url:"/settings/cron/"+t,method:"get"})}function f(t,e){return Object(a["a"])({url:"/settings/cron/"+e,method:"post",data:t})}function p(t){return Object(a["a"])({url:"/settings/cron/schedule",method:"post",data:t})}function m(t){return Object(a["a"])({url:"/settings/configs",method:"get",params:t})}function h(t){return Object(a["a"])({url:"/settings/config",method:"post",data:t})}function g(t){return Object(a["a"])({url:"/settings/config/"+t,method:"get"})}function b(t,e){return Object(a["a"])({url:"/settings/config/"+e,method:"post",data:t})}function v(t){return Object(a["a"])({url:"/settings/log",method:"post",data:t})}var _="https://hiads.mobgi.cc/api/settings/log/"},4221:function(t,e,n){"use strict";var a=n("4b8d"),r=n("4f7e"),c=n("c7b3"),o=n("3978"),i=n("b821"),s=n("3e87"),l=n("e001"),u=n("837a");r("search",(function(t,e,n){return[function(e){var n=o(this),r=void 0==e?void 0:l(e,t);return r?a(r,e,n):new RegExp(e)[t](s(n))},function(t){var a=c(this),r=s(t),o=n(e,a,r);if(o.done)return o.value;var l=a.lastIndex;i(l,0)||(a.lastIndex=0);var d=u(a,r);return i(a.lastIndex,l)||(a.lastIndex=l),null===d?-1:d.index}]}))},5723:function(t,e,n){"use strict";n.d(e,"e",(function(){return r})),n.d(e,"b",(function(){return c})),n.d(e,"d",(function(){return o})),n.d(e,"c",(function(){return i})),n.d(e,"a",(function(){return s})),n.d(e,"j",(function(){return l})),n.d(e,"k",(function(){return u})),n.d(e,"g",(function(){return d})),n.d(e,"f",(function(){return f})),n.d(e,"i",(function(){return p})),n.d(e,"h",(function(){return m}));var a=n("b775");function r(t){return Object(a["a"])({url:"/account/update",method:"post",data:t})}function c(t){return Object(a["a"])({url:"/account/create",method:"post",data:t})}function o(t){return Object(a["a"])({url:"/account/list",method:"get",params:t})}function i(t){return Object(a["a"])({url:"/account/"+t,method:"get"})}function s(t){return Object(a["a"])({url:"/account/auth",method:"get",params:{id:t}})}function l(t){return Object(a["a"])({url:"/account/refresh/"+t,method:"post"})}function u(t){return Object(a["a"])({url:"/account/search",method:"get",params:{account_name:t}})}function d(){return Object(a["a"])({url:"/account/default",method:"get"})}function f(){return Object(a["a"])({url:"/account/all",method:"get"})}function p(t){return Object(a["a"])({url:"/account/parents",method:"get",params:t})}function m(t){return Object(a["a"])({url:"/account/token",method:"post",data:t})}},"5c07":function(t,e,n){"use strict";var a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("DialogPanel",{attrs:{visible:t.visible,title:"展示字段选择",width:"460px","confirm-text":"确认"},on:{confirm:t.confirm,cancel:t.cancel}},[n("el-select",{staticStyle:{width:"100%"},attrs:{multiple:"",placeholder:"请选择"},model:{value:t.selects,callback:function(e){t.selects=e},expression:"selects"}},t._l(t.Columns,(function(t){return n("el-option",{key:t.key,attrs:{label:t.label,value:t.key}})})),1)],1)},r=[],c=(n("8300"),n("60fe"),n("068b"),n("3bae"),n("d4fd")),o=n("1c31"),i={name:"SelectColumns",components:{DialogPanel:c["a"]},props:{Columns:{required:!0,type:Array},ModuleName:{required:!0,type:String}},data:function(){return{visible:!1,selects:[]}},methods:{setDefault:function(){var t=this;this.Columns.forEach((function(e){e.show&&!t.selects.includes(e.key)&&t.selects.push(e.key)})),this.visible=!0},confirm:function(){var t=this;Object(o["c"])({columns:this.selects,module:this.ModuleName}).then((function(e){t.$emit("confirm",t.selects),t.visible=!1})).catch((function(e){t.$message.error("字段设置失败："+e)}))},cancel:function(){this.visible=!1}}},s=i,l=n("cba8"),u=Object(l["a"])(s,a,r,!1,null,null,null);e["a"]=u.exports},"80f8":function(t,e,n){},9663:function(t,e,n){"use strict";n.r(e);var a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("el-row",{staticClass:"comprehensive"},[n("el-form",{ref:"_search",attrs:{model:t.search,inline:"",size:"small"}},[n("el-col",{staticClass:"search-container",attrs:{span:24}},[n("el-form-item",{attrs:{label:"日期"}},[n("el-date-picker",{staticClass:"w240",attrs:{"picker-options":t.pickerOptions,clearable:!1,"value-format":"yyyy-MM-dd",type:"daterange","start-placeholder":"开始日期","end-placeholder":"结束日期"},model:{value:t.search.date_range,callback:function(e){t.$set(t.search,"date_range",e)},expression:"search.date_range"}})],1),n("el-form-item",{attrs:{label:"维度"}},[n("el-select",{staticClass:"w220",attrs:{placeholder:"数据维度",multiple:"","collapse-tags":""},model:{value:t.search.dimensions,callback:function(e){t.$set(t.search,"dimensions",e)},expression:"search.dimensions"}},t._l(t.reportDimensions,(function(t,e){return n("el-option",{key:e,attrs:{label:t,value:e}})})),1)],1),n("el-form-item",{attrs:{label:""}},[n("el-button",{staticClass:"item",attrs:{type:"primary",icon:"el-icon-search"},on:{click:t.doSearch}},[t._v("查询")])],1),n("el-form-item",{staticStyle:{float:"right"},attrs:{label:""}},[n("el-button",{staticClass:"item",attrs:{type:"danger",icon:"el-icon-s-tools",circle:""},on:{click:t.selectColumns}})],1)],1),n("el-col",{staticClass:"search-container",attrs:{span:24}},[t.search.dimensions.includes("account_id")?n("el-form-item",{attrs:{label:"账户"}},[n("el-select",{staticClass:"w260",attrs:{placeholder:"账户选择",multiple:"","collapse-tags":"",clearable:"",filterable:""},model:{value:t.search.account_ids,callback:function(e){t.$set(t.search,"account_ids",e)},expression:"search.account_ids"}},t._l(t.accounts,(function(e){return n("el-option",{directives:[{name:"show",rawName:"v-show",value:e.account_type===t.Vars.AccountTypeAds,expression:"item.account_type === Vars.AccountTypeAds"}],key:e.id,attrs:{label:e.account_name,value:e.id}})})),1)],1):t._e(),t.search.dimensions.includes("app_id")?n("el-form-item",{attrs:{label:"应用"}},[n("el-select",{staticClass:"w260",attrs:{placeholder:"应用选择",multiple:"","collapse-tags":"",clearable:"",filterable:""},model:{value:t.search.app_ids,callback:function(e){t.$set(t.search,"app_ids",e)},expression:"search.app_ids"}},t._l(t.apps,(function(t){return n("el-option",{key:t.app_id,attrs:{label:t.app_name,value:t.app_id}})})),1)],1):t._e(),t.search.dimensions.includes("area_id")?n("el-form-item",{attrs:{label:"地区"}},[n("el-select",{staticClass:"w260",attrs:{placeholder:"地区选择",multiple:"","collapse-tags":"",clearable:"",filterable:""},model:{value:t.search.areas,callback:function(e){t.$set(t.search,"areas",e)},expression:"search.areas"}},t._l(t.regions,(function(t){return n("el-option",{key:t.id,attrs:{label:t.c_name,value:t.id}})})),1)],1):t._e(),t.search.dimensions.includes("country")?n("el-form-item",{attrs:{label:"国家"}},[n("el-cascader",{staticClass:"w300",attrs:{options:t.regions,props:{multiple:!0,value:"c_code",label:"c_name"},"collapse-tags":"",clearable:""},model:{value:t.search.countries,callback:function(e){t.$set(t.search,"countries",e)},expression:"search.countries"}})],1):t._e()],1)],1),n("el-col",{attrs:{span:24}},[n("el-table",{directives:[{name:"loading",rawName:"v-loading",value:t.loadings.pageLoading,expression:"loadings.pageLoading"}],attrs:{data:t.reportList.list,"highlight-current-row":"",stripe:"",border:"",size:"mini","show-summary":""}},t._l(t.reportList.columns,(function(e){return e.show?n("el-table-column",{key:e.key,attrs:{label:e.label,align:e.align,fixed:e.fix,"min-width":e.min,"show-overflow-tooltip":e.fix,prop:e.key},scopedSlots:t._u([{key:"default",fn:function(n){return[t._v(" "+t._s(e.prefix+n.row[e.key]+e.suffix)+" ")]}}],null,!0)}):t._e()})),1)],1),n("el-col",{staticStyle:{"margin-top":"15px"},attrs:{span:24}},[n("page",{ref:"page",attrs:{page:t.search.page,total:t.reportList.total,limit:t.search.page_size},on:{"current-change":t.handlePage,"size-change":t.handlePageSize}})],1),n("select-columns",{ref:"column",attrs:{Columns:t.reportList.columns,"module-name":"ads"},on:{confirm:t.confirm}})],1)},r=[],c=(n("8300"),n("ea5b"),n("e551"),n("2ce8"),n("4221"),n("60fe"),n("fc23")),o=n("ed08"),i=n("1c31"),s=n("2934"),l=n("5723"),u=n("b562"),d=n("5c07"),f=n("cf0b"),p=new Date,m={name:"Ads",components:{Page:c["a"],SelectColumns:d["a"]},filters:{timeFormat:function(t){return Object(o["c"])(t)}},data:function(){return{Vars:f["a"],loadings:{pageLoading:!1},reportDimensions:{},search:{date_range:[],dimensions:[],account_ids:[],app_ids:[],countries:[],show_columns:[],page:1,page_size:15},accounts:[],apps:[],appRels:{},regions:[],reportList:{list:[],total:0,columns:[],summaries:{}},pickerOptions:{shortcuts:[{text:"近 7 天",onClick:function(t){var e=new Date;e.setTime(e.getTime()-6048e5),t.$emit("pick",[e,p])}},{text:"本月",onClick:function(t){var e=new Date;t.$emit("pick",[new Date(e.setDate(1)),p])}},{text:"上月",onClick:function(t){var e=new Date((new Date).setDate(1));e.setTime(e.getTime()-864e5);var n=new Date(e-0),a=new Date(n.setDate(1));t.$emit("pick",[a,e])}},{text:"近 30 天",onClick:function(t){var e=new Date;e.setTime(e.getTime()-2592e6),t.$emit("pick",[e,p])}}]}}},created:function(){this.setDefaultSearchDate()},mounted:function(){this.getReportList()},methods:{getReportList:function(){var t=this;this.loadings.pageLoading=!0,Promise.all([this.getAllAccounts(),this.getAllApps(),this.getRegions()]).then((function(e){Object(i["a"])(t.search).then((function(e){t.reportList.columns=e.data.columns,t.reportList.list=e.data.list,t.reportList.total=e.data.total,t.reportDimensions=e.data.dimensions,t.loadings.pageLoading=!1})).catch((function(e){console.log(e),t.loadings.pageLoading=!1}))})).catch((function(e){console.log(e),t.loadings.pageLoading=!1}))},getRegions:function(){var t=this;return new Promise((function(e,n){if(t.regions.length>0)return e();Object(s["h"])().then((function(n){t.regions=n.data,e()})).catch((function(t){n(t)}))}))},getAllApps:function(){var t=this;return new Promise((function(e,n){if(t.apps.length>0)return e();Object(u["a"])().then((function(n){t.apps=n.data,e()})).catch((function(t){n(t)}))}))},getAllAccounts:function(){var t=this;return new Promise((function(e,n){t.accounts.length>0?e():Object(l["f"])().then((function(n){t.accounts=n.data,e()})).catch((function(t){n(t)}))}))},handlePage:function(t){this.search.page=t,this.getReportList()},handlePageSize:function(t){this.search.page_size=t,this.getReportList()},doSearch:function(){this.search.page=1,this.getReportList()},setDefaultSearchDate:function(){var t=new Date,e="{y}-{m}-{d}";this.$set(this.search,"date_range",[Object(o["c"])(t.getTime()-6048e5,e),Object(o["c"])(new Date,e)])},selectColumns:function(){this.$refs.column.setDefault()},confirm:function(t){this.search.show_columns=t,this.getReportList()},getSummaries:function(t){var e=this,n=t.columns,a=[];return n.forEach((function(t,n){if(0===n)a[n]="合计";else switch(t.property){case"earnings":a[n]=e.reportList.summaries.earnings;break}})),a}}},h=m,g=n("cba8"),b=Object(g["a"])(h,a,r,!1,null,"7d917aad",null);e["default"]=b.exports},b562:function(t,e,n){"use strict";n.d(e,"e",(function(){return r})),n.d(e,"c",(function(){return c})),n.d(e,"g",(function(){return o})),n.d(e,"d",(function(){return i})),n.d(e,"a",(function(){return s})),n.d(e,"h",(function(){return l})),n.d(e,"f",(function(){return u})),n.d(e,"b",(function(){return d}));var a=n("b775");function r(t){return Object(a["a"])({url:"/app/list",method:"post",data:t})}function c(t){return Object(a["a"])({url:"/app/create",method:"post",data:t})}function o(t){return Object(a["a"])({url:"/app/update",method:"post",data:t})}function i(t){return Object(a["a"])({url:"/app/"+t,method:"get"})}function s(){return Object(a["a"])({url:"/app/all",method:"get"})}function l(t){return Object(a["a"])({url:"/app/campaign-list",method:"get",params:t})}function u(t){return Object(a["a"])({url:"/app/pull",method:"post",data:t})}function d(){return Object(a["a"])({url:"/app/relation",method:"get"})}},b821:function(t,e){t.exports=Object.is||function(t,e){return t===e?0!==t||1/t===1/e:t!=t&&e!=e}},b9c3:function(t,e,n){"use strict";n("80f8")},cf0b:function(t,e,n){"use strict";var a={AccountTypeMarket:1,AccountTypeAds:2,ReportGranularity:[{name:"按日期",key:"date"},{name:"按整体",key:"all"}]};e["a"]=a},d4fd:function(t,e,n){"use strict";var a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("el-dialog",{directives:[{name:"el-drag-dialog",rawName:"v-el-drag-dialog"}],attrs:{title:t.title,visible:t.visible,width:t.width,"append-to-body":"","modal-append-to-body":"","close-on-click-modal":!1,"close-on-press-escape":!1,"show-close":!1},on:{"update:visible":function(e){t.visible=e}}},[t._t("default"),n("div",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[n("el-button",{attrs:{loading:t.confirmLoading,icon:"el-icon-close"},on:{click:t.handleCancel}},[t._v("取消")]),n("el-button",{directives:[{name:"show",rawName:"v-show",value:t.confirmText,expression:"confirmText"}],attrs:{type:"primary",icon:"el-icon-check",loading:t.confirmLoading},on:{click:t.handleConfirm}},[t._v(t._s(t.confirmText))]),t._t("operate")],2)],2)},r=[],c={props:{confirmText:{default:"",type:String},title:{default:"",type:String},width:{default:"600px",type:String},confirmLoading:{default:!1,type:Boolean},visible:{default:!1,type:Boolean}},methods:{handleCancel:function(){this.$emit("cancel")},handleConfirm:function(){this.$emit("confirm")}}},o=c,i=(n("b9c3"),n("cba8")),s=Object(i["a"])(o,a,r,!1,null,null,null);e["a"]=s.exports},fc23:function(t,e,n){"use strict";var a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("el-pagination",{attrs:{background:"","current-page":t.page,"page-size":t.limit,total:t.total,"hide-on-single-page":"",layout:"prev, pager, next, jumper, total, sizes","prev-text":"上页","next-text":"下页","page-sizes":[10,15,20,30,40,50,100]},on:{"current-change":t.handlePage,"size-change":t.handleSizeChange}})},r=[],c=(n("4582"),{props:{page:{type:Number,default:1},total:{type:Number,default:0},limit:{type:Number,default:10}},methods:{handlePage:function(t){this.$emit("current-change",t)},handleSizeChange:function(t){this.$emit("size-change",t)}}}),o=c,i=n("cba8"),s=Object(i["a"])(o,a,r,!1,null,null,null);e["a"]=s.exports}}]);