(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-3e254312"],{"0bb0":function(t,e,n){"use strict";n.d(e,"a",(function(){return a}));var a={account_id:"账户",app_id:"应用",country:"国家&地区"}},"1c31":function(t,e,n){"use strict";n.d(e,"b",(function(){return r})),n.d(e,"a",(function(){return o})),n.d(e,"c",(function(){return i}));var a=n("b775");function r(t){return Object(a["a"])({url:"/report/comprehensive",method:"post",data:t})}function o(t){return Object(a["a"])({url:"/report/ads",method:"post",data:t})}function i(t){return Object(a["a"])({url:"/report/column",method:"post",data:t})}},2934:function(t,e,n){"use strict";n.d(e,"h",(function(){return r})),n.d(e,"n",(function(){return o})),n.d(e,"d",(function(){return i})),n.d(e,"e",(function(){return c})),n.d(e,"g",(function(){return s})),n.d(e,"f",(function(){return l})),n.d(e,"i",(function(){return u})),n.d(e,"j",(function(){return d})),n.d(e,"l",(function(){return f})),n.d(e,"k",(function(){return p})),n.d(e,"m",(function(){return m})),n.d(e,"a",(function(){return h})),n.d(e,"b",(function(){return g})),n.d(e,"c",(function(){return b}));var a=n("b775");function r(){return Object(a["a"])({url:"/regions",method:"get"})}function o(){return Object(a["a"])({url:"/settings/version",method:"get"})}function i(){return Object(a["a"])({url:"/region/area",method:"get"})}function c(t){return Object(a["a"])({url:"/region/country",method:"get",params:t})}function s(t){return Object(a["a"])({url:"/region",method:"post",data:t})}function l(t){return Object(a["a"])({url:"/region/area-set",method:"post",data:t})}function u(){return Object(a["a"])({url:"/settings/cron",method:"get"})}function d(t){return Object(a["a"])({url:"/settings/cron/"+t,method:"get"})}function f(t,e){return Object(a["a"])({url:"/settings/cron/"+e,method:"post",data:t})}function p(t){return Object(a["a"])({url:"/settings/cron/schedule",method:"post",data:t})}function m(t){return Object(a["a"])({url:"/settings/configs",method:"get",params:t})}function h(t){return Object(a["a"])({url:"/settings/config",method:"post",data:t})}function g(t){return Object(a["a"])({url:"/settings/config/"+t,method:"get"})}function b(t,e){return Object(a["a"])({url:"/settings/config/"+e,method:"post",data:t})}},"3f5f":function(t,e,n){"use strict";n.r(e);var a,r=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("el-row",{staticClass:"comprehensive"},[n("el-form",{ref:"_search",attrs:{model:t.search,inline:"",size:"small"}},[n("el-col",{staticClass:"search-container",attrs:{span:24}},[n("el-form-item",{attrs:{label:"日期"}},[n("el-date-picker",{staticClass:"w240",attrs:{"picker-options":t.pickerOptions,clearable:!1,"value-format":"yyyy-MM-dd",type:"daterange","start-placeholder":"开始日期","end-placeholder":"结束日期"},model:{value:t.search.date_range,callback:function(e){t.$set(t.search,"date_range",e)},expression:"search.date_range"}})],1),n("el-form-item",{attrs:{label:"维度"}},[n("el-select",{staticClass:"w220",attrs:{placeholder:"数据维度",multiple:"","collapse-tags":""},model:{value:t.search.dimensions,callback:function(e){t.$set(t.search,"dimensions",e)},expression:"search.dimensions"}},t._l(t.requestDimensions,(function(t,e){return n("el-option",{key:e,attrs:{label:t,value:e}})})),1)],1),n("el-form-item",{attrs:{label:""}},[n("el-button",{staticClass:"item",attrs:{type:"primary",icon:"el-icon-search"},on:{click:t.doSearch}},[t._v("查询")]),n("el-button",{staticClass:"item",attrs:{icon:"el-icon-download"},on:{click:t.download}},[t._v("下载数据")])],1),n("el-form-item",{staticStyle:{float:"right"},attrs:{label:""}},[n("el-button",{staticClass:"item",attrs:{type:"danger",icon:"el-icon-s-tools",circle:""},on:{click:t.selectColumns}})],1)],1),n("el-col",{staticClass:"search-container",attrs:{span:24}},[t.search.dimensions.includes("account_id")?n("el-form-item",{attrs:{label:"账户"}},[n("el-select",{staticClass:"w260",attrs:{placeholder:"账户选择",multiple:"","collapse-tags":"",clearable:"",filterable:""},model:{value:t.search.account_ids,callback:function(e){t.$set(t.search,"account_ids",e)},expression:"search.account_ids"}},t._l(t.accounts,(function(t){return n("el-option",{directives:[{name:"show",rawName:"v-show",value:1===t.account_type,expression:"item.account_type === 1"}],key:t.id,attrs:{label:t.account_name,value:t.id}})})),1)],1):t._e(),t.search.dimensions.includes("app_id")?n("el-form-item",{attrs:{label:"应用"}},[n("el-select",{staticClass:"w260",attrs:{placeholder:"应用选择",multiple:"","collapse-tags":"",clearable:"",filterable:""},model:{value:t.search.app_ids,callback:function(e){t.$set(t.search,"app_ids",e)},expression:"search.app_ids"}},t._l(t.apps,(function(t){return n("el-option",{key:t.app_id,attrs:{label:t.app_name,value:t.app_id}})})),1)],1):t._e(),t.search.dimensions.includes("country")?n("el-form-item",{attrs:{label:"区域"}},[n("el-cascader",{staticClass:"w300",attrs:{options:t.regions,props:{multiple:!0,value:"c_code",label:"c_name"},"collapse-tags":"",clearable:""},model:{value:t.search.countries,callback:function(e){t.$set(t.search,"countries",e)},expression:"search.countries"}})],1):t._e()],1)],1),n("el-col",{attrs:{span:24}},[n("el-table",{directives:[{name:"loading",rawName:"v-loading",value:t.loadings.pageLoading,expression:"loadings.pageLoading"}],attrs:{data:t.reportList.list,"highlight-current-row":"",stripe:"",border:"",size:"mini","show-summary":"","summary-method":t.getSummaries},on:{"sort-change":t.sortable}},t._l(t.reportList.columns,(function(e){return e.show?n("el-table-column",{key:e.key,attrs:{label:e.label,align:e.align,fixed:e.fix,"min-width":e.min,"show-overflow-tooltip":e.fix,sortable:t._f("filterSort")(e.sort),prop:e.key},scopedSlots:t._u([{key:"default",fn:function(n){return[t._v(" "+t._s(e.prefix+n.row[e.key]+e.suffix)+" ")]}}],null,!0)}):t._e()})),1)],1),n("el-col",{staticStyle:{"margin-top":"15px"},attrs:{span:24}},[n("page",{ref:"page",attrs:{page:t.search.page,total:t.reportList.total,limit:t.search.page_size},on:{"current-change":t.handlePage,"size-change":t.handlePageSize}})],1),n("select-columns",{ref:"column",attrs:{Columns:t.reportList.columns,"module-name":"comprehensive"},on:{confirm:t.confirm}})],1)},o=[],i=n("28b6"),c=(n("8300"),n("ea5b"),n("e551"),n("2ce8"),n("4221"),n("9a2f"),n("60fe"),n("fc23")),s=n("ed08"),l=n("0bb0"),u=n("1c31"),d=n("2934"),f=n("5723"),p=n("b562"),m=n("5c07"),h=new Date,g={custom:"custom",1:!0,0:!1},b=(a={name:"Comprehensive",components:{Page:c["a"],SelectColumns:m["a"]},filters:{timeFormat:function(t){return Object(s["c"])(t)}},data:function(){return{requestDimensions:l["a"],loadings:{pageLoading:!1},search:{date_range:[],dimensions:[],account_ids:[],app_ids:[],countries:[],show_columns:[],order:"",by:"",download:0,page:1,page_size:15},accounts:[],apps:[],regions:[],reportList:{list:[],total:0,columns:[],summaries:{}},pickerOptions:{shortcuts:[{text:"近 7 天",onClick:function(t){var e=new Date;e.setTime(e.getTime()-6048e5),t.$emit("pick",[e,h])}},{text:"本月",onClick:function(t){var e=new Date;t.$emit("pick",[new Date(e.setDate(1)),h])}},{text:"上月",onClick:function(t){var e=new Date((new Date).setDate(1));e.setTime(e.getTime()-864e5);var n=new Date(e-0),a=new Date(n.setDate(1));t.$emit("pick",[a,e])}},{text:"近 30 天",onClick:function(t){var e=new Date;e.setTime(e.getTime()-2592e6),t.$emit("pick",[e,h])}}]}}}},Object(i["a"])(a,"filters",{filterSort:function(t){return g[t]}}),Object(i["a"])(a,"created",(function(){this.setDefaultSearchDate()})),Object(i["a"])(a,"mounted",(function(){this.getReportList()})),Object(i["a"])(a,"methods",{getReportList:function(){var t=this;this.loadings.pageLoading=!0,Promise.all([this.getAllAccounts(),this.getAllApps(),this.getRegions()]).then((function(e){Object(u["b"])(t.search).then((function(e){t.reportList.columns=e.data.columns,t.reportList.list=e.data.list,t.reportList.total=e.data.total,t.reportList.summaries=e.data.summaries,t.loadings.pageLoading=!1})).catch((function(e){console.log(e),t.loadings.pageLoading=!1}))})).catch((function(e){console.log(e),t.loadings.pageLoading=!1}))},getRegions:function(){var t=this;return new Promise((function(e,n){if(t.regions.length>0)return e();Object(d["h"])().then((function(n){t.regions=n.data,e()})).catch((function(t){n(t)}))}))},getAllApps:function(){var t=this;return new Promise((function(e,n){if(t.apps.length>0)return e();Object(p["a"])().then((function(n){t.apps=n.data,e()})).catch((function(t){n(t)}))}))},getAllAccounts:function(){var t=this;return new Promise((function(e,n){t.accounts.length>0?e():Object(f["f"])().then((function(n){t.accounts=n.data,e()})).catch((function(t){n(t)}))}))},handlePage:function(t){this.search.page=t,this.getReportList()},handlePageSize:function(t){this.search.page_size=t,this.getReportList()},doSearch:function(){this.search.page=1,this.getReportList()},download:function(){var t=this;this.loadings.pageLoading=!0,this.search.download=1,Object(u["b"])(this.search).then((function(e){t.search.download=0,t.loadings.downloadLoading=!1,Promise.all([n.e("chunk-4c7447ea"),n.e("chunk-2133cd4f")]).then(n.bind(null,"4bf8")).then((function(n){var a=[];e.data.columns.map((function(t){t.show&&a.push(t.label)}));var r=n.formatJson(e.data);if(0===r.length)return t.$message.info("没有筛选到需导出的数据"),void(t.loadings.pageLoading=!1);n.export_json_to_excel({header:a,data:r,filename:"综合报表数据"}),t.loadings.pageLoading=!1})).catch((function(e){console.log(e),t.loadings.pageLoading=!1}))})).catch((function(){t.search.download=0,t.loadings.pageLoading=!1}))},setDefaultSearchDate:function(){var t=new Date,e="{y}-{m}-{d}";this.$set(this.search,"date_range",[Object(s["c"])(t.getTime()-6048e5,e),Object(s["c"])(new Date,e)])},selectColumns:function(){this.$refs.column.setDefault()},confirm:function(t){this.search.show_columns=t,this.getReportList()},sortable:function(t){var e=t.column,n=t.prop,a=t.order;"custom"===e.sortable&&(this.search.order=n,this.search.by=a,this.getReportList())},getSummaries:function(t){var e=this,n=t.columns,a=[];return n.forEach((function(t,n){if(0===n)a[n]="合计";else switch(t.property){case"cost":a[n]=e.reportList.summaries.cost;break;case"earnings":a[n]=e.reportList.summaries.earnings;break;case"roi":a[n]=e.reportList.summaries.roi+"%";break}})),a}}),a),v=b,_=n("cba8"),w=Object(_["a"])(v,r,o,!1,null,"2ac045d6",null);e["default"]=w.exports},4221:function(t,e,n){"use strict";var a=n("4b8d"),r=n("4f7e"),o=n("c7b3"),i=n("3978"),c=n("b821"),s=n("3e87"),l=n("e001"),u=n("837a");r("search",(function(t,e,n){return[function(e){var n=i(this),r=void 0==e?void 0:l(e,t);return r?a(r,e,n):new RegExp(e)[t](s(n))},function(t){var a=o(this),r=s(t),i=n(e,a,r);if(i.done)return i.value;var l=a.lastIndex;c(l,0)||(a.lastIndex=0);var d=u(a,r);return c(a.lastIndex,l)||(a.lastIndex=l),null===d?-1:d.index}]}))},5723:function(t,e,n){"use strict";n.d(e,"e",(function(){return r})),n.d(e,"b",(function(){return o})),n.d(e,"d",(function(){return i})),n.d(e,"c",(function(){return c})),n.d(e,"a",(function(){return s})),n.d(e,"j",(function(){return l})),n.d(e,"k",(function(){return u})),n.d(e,"g",(function(){return d})),n.d(e,"f",(function(){return f})),n.d(e,"i",(function(){return p})),n.d(e,"h",(function(){return m}));var a=n("b775");function r(t){return Object(a["a"])({url:"/account/update",method:"post",data:t})}function o(t){return Object(a["a"])({url:"/account/create",method:"post",data:t})}function i(t){return Object(a["a"])({url:"/account/list",method:"get",params:t})}function c(t){return Object(a["a"])({url:"/account/"+t,method:"get"})}function s(t){return Object(a["a"])({url:"/account/auth",method:"get",params:{id:t}})}function l(t){return Object(a["a"])({url:"/account/refresh/"+t,method:"post"})}function u(t){return Object(a["a"])({url:"/account/search",method:"get",params:{account_name:t}})}function d(){return Object(a["a"])({url:"/account/default",method:"get"})}function f(){return Object(a["a"])({url:"/account/all",method:"get"})}function p(t){return Object(a["a"])({url:"/account/parents",method:"get",params:t})}function m(t){return Object(a["a"])({url:"/account/token",method:"post",data:t})}},"5c07":function(t,e,n){"use strict";var a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("DialogPanel",{attrs:{visible:t.visible,title:"展示字段选择",width:"460px","confirm-text":"确认"},on:{confirm:t.confirm,cancel:t.cancel}},[n("el-select",{staticStyle:{width:"100%"},attrs:{multiple:"",placeholder:"请选择"},model:{value:t.selects,callback:function(e){t.selects=e},expression:"selects"}},t._l(t.Columns,(function(t){return n("el-option",{key:t.key,attrs:{label:t.label,value:t.key}})})),1)],1)},r=[],o=(n("8300"),n("60fe"),n("068b"),n("3bae"),n("d4fd")),i=n("1c31"),c={name:"SelectColumns",components:{DialogPanel:o["a"]},props:{Columns:{required:!0,type:Array},ModuleName:{required:!0,type:String}},data:function(){return{visible:!1,selects:[]}},methods:{setDefault:function(){var t=this;this.Columns.forEach((function(e){e.show&&!t.selects.includes(e.key)&&t.selects.push(e.key)})),this.visible=!0},confirm:function(){var t=this;Object(i["c"])({columns:this.selects,module:this.ModuleName}).then((function(e){t.$emit("confirm",t.selects),t.visible=!1})).catch((function(e){t.$message.error("字段设置失败："+e)}))},cancel:function(){this.visible=!1}}},s=c,l=n("cba8"),u=Object(l["a"])(s,a,r,!1,null,null,null);e["a"]=u.exports},"80f8":function(t,e,n){},b562:function(t,e,n){"use strict";n.d(e,"d",(function(){return r})),n.d(e,"b",(function(){return o})),n.d(e,"f",(function(){return i})),n.d(e,"c",(function(){return c})),n.d(e,"a",(function(){return s})),n.d(e,"g",(function(){return l})),n.d(e,"e",(function(){return u}));var a=n("b775");function r(t){return Object(a["a"])({url:"/app/list",method:"get",params:t})}function o(t){return Object(a["a"])({url:"/app/create",method:"post",data:t})}function i(t){return Object(a["a"])({url:"/app/update",method:"post",data:t})}function c(t){return Object(a["a"])({url:"/app/"+t,method:"get"})}function s(){return Object(a["a"])({url:"/app/all",method:"get"})}function l(t){return Object(a["a"])({url:"/app/campaign-list",method:"get",params:t})}function u(t){return Object(a["a"])({url:"/app/pull",method:"post",data:t})}},b821:function(t,e){t.exports=Object.is||function(t,e){return t===e?0!==t||1/t===1/e:t!=t&&e!=e}},b9c3:function(t,e,n){"use strict";n("80f8")},d4fd:function(t,e,n){"use strict";var a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("el-dialog",{directives:[{name:"el-drag-dialog",rawName:"v-el-drag-dialog"}],attrs:{title:t.title,visible:t.visible,width:t.width,"append-to-body":"","modal-append-to-body":"","close-on-click-modal":!1,"close-on-press-escape":!1,"show-close":!1},on:{"update:visible":function(e){t.visible=e}}},[t._t("default"),n("div",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[n("el-button",{attrs:{loading:t.confirmLoading,icon:"el-icon-close"},on:{click:t.handleCancel}},[t._v("取消")]),n("el-button",{directives:[{name:"show",rawName:"v-show",value:t.confirmText,expression:"confirmText"}],attrs:{type:"primary",icon:"el-icon-check",loading:t.confirmLoading},on:{click:t.handleConfirm}},[t._v(t._s(t.confirmText))]),t._t("operate")],2)],2)},r=[],o={props:{confirmText:{default:"",type:String},title:{default:"",type:String},width:{default:"600px",type:String},confirmLoading:{default:!1,type:Boolean},visible:{default:!1,type:Boolean}},methods:{handleCancel:function(){this.$emit("cancel")},handleConfirm:function(){this.$emit("confirm")}}},i=o,c=(n("b9c3"),n("cba8")),s=Object(c["a"])(i,a,r,!1,null,null,null);e["a"]=s.exports},fc23:function(t,e,n){"use strict";var a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("el-pagination",{attrs:{background:"","current-page":t.page,"page-size":t.limit,total:t.total,"hide-on-single-page":"",layout:"prev, pager, next, jumper, total, sizes","prev-text":"上页","next-text":"下页","page-sizes":[10,15,20,30,40,50,100]},on:{"current-change":t.handlePage,"size-change":t.handleSizeChange}})},r=[],o=(n("4582"),{props:{page:{type:Number,default:1},total:{type:Number,default:0},limit:{type:Number,default:10}},methods:{handlePage:function(t){this.$emit("current-change",t)},handleSizeChange:function(t){this.$emit("size-change",t)}}}),i=o,c=n("cba8"),s=Object(c["a"])(i,a,r,!1,null,null,null);e["a"]=s.exports}}]);