(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-21fd14fc"],{"1c31":function(e,t,n){"use strict";n.d(t,"b",(function(){return r})),n.d(t,"a",(function(){return o})),n.d(t,"c",(function(){return i}));var a=n("b775");function r(e){return Object(a["a"])({url:"/report/comprehensive",method:"post",data:e})}function o(e){return Object(a["a"])({url:"/report/ads",method:"post",data:e})}function i(e){return Object(a["a"])({url:"/report/column",method:"post",data:e})}},2934:function(e,t,n){"use strict";n.d(t,"h",(function(){return r})),n.d(t,"n",(function(){return o})),n.d(t,"d",(function(){return i})),n.d(t,"e",(function(){return c})),n.d(t,"g",(function(){return s})),n.d(t,"f",(function(){return l})),n.d(t,"i",(function(){return u})),n.d(t,"j",(function(){return d})),n.d(t,"l",(function(){return p})),n.d(t,"k",(function(){return f})),n.d(t,"m",(function(){return h})),n.d(t,"a",(function(){return m})),n.d(t,"b",(function(){return g})),n.d(t,"c",(function(){return b}));var a=n("b775");function r(){return Object(a["a"])({url:"/regions",method:"get"})}function o(){return Object(a["a"])({url:"/settings/version",method:"get"})}function i(){return Object(a["a"])({url:"/region/area",method:"get"})}function c(e){return Object(a["a"])({url:"/region/country",method:"get",params:e})}function s(e){return Object(a["a"])({url:"/region",method:"post",data:e})}function l(e){return Object(a["a"])({url:"/region/area-set",method:"post",data:e})}function u(){return Object(a["a"])({url:"/settings/cron",method:"get"})}function d(e){return Object(a["a"])({url:"/settings/cron/"+e,method:"get"})}function p(e,t){return Object(a["a"])({url:"/settings/cron/"+t,method:"post",data:e})}function f(e){return Object(a["a"])({url:"/settings/cron/schedule",method:"post",data:e})}function h(e){return Object(a["a"])({url:"/settings/configs",method:"get",params:e})}function m(e){return Object(a["a"])({url:"/settings/config",method:"post",data:e})}function g(e){return Object(a["a"])({url:"/settings/config/"+e,method:"get"})}function b(e,t){return Object(a["a"])({url:"/settings/config/"+t,method:"post",data:e})}},"3f5f":function(e,t,n){"use strict";n.r(t);var a,r=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("el-row",{staticClass:"comprehensive"},[n("el-form",{ref:"_search",attrs:{model:e.search,inline:"",size:"small"}},[n("el-col",{staticClass:"search-container",attrs:{span:24}},[n("el-form-item",{attrs:{label:"粒度"}},[n("el-select",{staticClass:"w100",attrs:{placeholder:"数据粒度"},model:{value:e.search.granularity,callback:function(t){e.$set(e.search,"granularity",t)},expression:"search.granularity"}},e._l(e.Vars.ReportGranularity,(function(e){return n("el-option",{key:e.key,attrs:{label:e.name,value:e.key}})})),1)],1),n("el-form-item",{attrs:{label:"日期"}},[n("el-date-picker",{staticClass:"w240",attrs:{"picker-options":e.pickerOptions,clearable:!1,"value-format":"yyyy-MM-dd",type:"daterange","start-placeholder":"开始日期","end-placeholder":"结束日期"},model:{value:e.search.date_range,callback:function(t){e.$set(e.search,"date_range",t)},expression:"search.date_range"}})],1),n("el-form-item",{attrs:{label:"维度"}},[n("el-select",{staticClass:"w220",attrs:{placeholder:"数据维度",multiple:"","collapse-tags":""},model:{value:e.search.dimensions,callback:function(t){e.$set(e.search,"dimensions",t)},expression:"search.dimensions"}},e._l(e.reportDimensions,(function(e,t){return n("el-option",{key:t,attrs:{label:e,value:t}})})),1)],1),n("el-form-item",{attrs:{label:""}},[n("el-button",{staticClass:"item",attrs:{type:"primary",icon:"el-icon-search"},on:{click:e.doSearch}},[e._v("查询")]),n("el-button",{directives:[{name:"permission",rawName:"v-permission",value:"report/comprehensive-download",expression:"'report/comprehensive-download'"}],staticClass:"item",attrs:{icon:"el-icon-download"},on:{click:e.download}},[e._v("下载数据")])],1),n("el-form-item",{staticStyle:{float:"right"},attrs:{label:""}},[n("el-button",{staticClass:"item",attrs:{type:"danger",icon:"el-icon-s-tools",circle:""},on:{click:e.selectColumns}})],1)],1),n("el-col",{staticClass:"search-container",attrs:{span:24}},[e.search.dimensions.includes("account_id")?n("el-form-item",{attrs:{label:"账户"}},[n("el-select",{staticClass:"w260",attrs:{placeholder:"账户选择",multiple:"","collapse-tags":"",clearable:"",filterable:""},on:{change:e.conditionApp},model:{value:e.search.account_ids,callback:function(t){e.$set(e.search,"account_ids",t)},expression:"search.account_ids"}},e._l(e.accounts,(function(t){return n("el-option",{directives:[{name:"show",rawName:"v-show",value:t.account_type===e.Vars.AccountTypeMarket,expression:"item.account_type === Vars.AccountTypeMarket"}],key:t.id,attrs:{label:t.account_name,value:t.id}})})),1)],1):e._e(),e.search.dimensions.includes("app_id")?n("el-form-item",{attrs:{label:"应用"}},[n("el-select",{staticClass:"w260",attrs:{placeholder:"应用选择",multiple:"","collapse-tags":"",clearable:"",filterable:""},model:{value:e.search.app_ids,callback:function(t){e.$set(e.search,"app_ids",t)},expression:"search.app_ids"}},e._l(e.apps,(function(t){return n("el-option",{directives:[{name:"show",rawName:"v-show",value:t.show||0===e.search.account_ids.length,expression:"item.show||search.account_ids.length===0"}],key:t.app_id,attrs:{label:t.app_name,value:t.app_id}})})),1)],1):e._e(),e.search.dimensions.includes("area_id")?n("el-form-item",{attrs:{label:"地区"}},[n("el-select",{staticClass:"w260",attrs:{placeholder:"地区选择",multiple:"","collapse-tags":"",clearable:"",filterable:""},model:{value:e.search.areas,callback:function(t){e.$set(e.search,"areas",t)},expression:"search.areas"}},e._l(e.regions,(function(e){return n("el-option",{key:e.id,attrs:{label:e.c_name,value:e.id}})})),1)],1):e._e(),e.search.dimensions.includes("country")?n("el-form-item",{attrs:{label:"国家"}},[n("el-cascader",{staticClass:"w300",attrs:{options:e.regions,props:{multiple:!0,value:"c_code",label:"c_name"},"collapse-tags":"",clearable:""},model:{value:e.search.countries,callback:function(t){e.$set(e.search,"countries",t)},expression:"search.countries"}})],1):e._e()],1)],1),n("el-col",{attrs:{span:24}},[n("el-table",{directives:[{name:"loading",rawName:"v-loading",value:e.loadings.pageLoading,expression:"loadings.pageLoading"}],attrs:{data:e.reportList.list,"highlight-current-row":"",stripe:"",border:"",size:"mini","show-summary":"","summary-method":e.getSummaries},on:{"sort-change":e.sortable}},e._l(e.reportList.columns,(function(t){return t.show?n("el-table-column",{key:t.key,attrs:{label:t.label,align:t.align,fixed:t.fix,"min-width":t.min,"show-overflow-tooltip":t.fix,sortable:e._f("filterSort")(t.sort),prop:t.key},scopedSlots:e._u([{key:"default",fn:function(n){return[e._v(" "+e._s(t.prefix+n.row[t.key]+t.suffix)+" ")]}}],null,!0)}):e._e()})),1)],1),n("el-col",{staticStyle:{"margin-top":"15px"},attrs:{span:24}},[n("page",{ref:"page",attrs:{page:e.search.page,total:e.reportList.total,limit:e.search.page_size},on:{"current-change":e.handlePage,"size-change":e.handlePageSize}})],1),n("select-columns",{ref:"column",attrs:{Columns:e.reportList.columns,"module-name":"comprehensive"},on:{confirm:e.confirm}})],1)},o=[],i=n("28b6"),c=(n("8300"),n("ea5b"),n("e551"),n("2ce8"),n("4221"),n("c76c"),n("9a2f"),n("60fe"),n("654b"),n("068b"),n("3bae"),n("fc23")),s=n("ed08"),l=n("1c31"),u=n("2934"),d=n("5723"),p=n("b562"),f=n("5c07"),h=n("cf0b"),m=new Date,g={custom:"custom",1:!0,0:!1},b=(a={name:"Comprehensive",components:{Page:c["a"],SelectColumns:f["a"]},filters:{timeFormat:function(e){return Object(s["c"])(e)}},data:function(){return{Vars:h["a"],loadings:{pageLoading:!1},reportDimensions:{},search:{date_range:[],dimensions:[],account_ids:[],app_ids:[],areas:[],countries:[],show_columns:[],granularity:"date",order:"",by:"",download:0,page:1,page_size:15},accounts:[],apps:[],appRels:{},regions:[],reportList:{list:[],total:0,columns:[],summaries:{}},pickerOptions:{shortcuts:[{text:"近 7 天",onClick:function(e){var t=new Date;t.setTime(t.getTime()-6048e5),e.$emit("pick",[t,m])}},{text:"本月",onClick:function(e){var t=new Date;e.$emit("pick",[new Date(t.setDate(1)),m])}},{text:"上月",onClick:function(e){var t=new Date((new Date).setDate(1));t.setTime(t.getTime()-864e5);var n=new Date(t-0),a=new Date(n.setDate(1));e.$emit("pick",[a,t])}},{text:"近 30 天",onClick:function(e){var t=new Date;t.setTime(t.getTime()-2592e6),e.$emit("pick",[t,m])}}]}}}},Object(i["a"])(a,"filters",{filterSort:function(e){return g[e]}}),Object(i["a"])(a,"created",(function(){this.setDefaultSearchDate()})),Object(i["a"])(a,"mounted",(function(){this.getReportList()})),Object(i["a"])(a,"methods",{getReportList:function(){var e=this;this.loadings.pageLoading=!0,Promise.all([this.getAllAccounts(),this.getAllApps(),this.getRegions(),this.getAppRels()]).then((function(t){Object(l["b"])(e.search).then((function(t){e.reportList.columns=t.data.columns,e.reportList.list=t.data.list,e.reportList.total=t.data.total,e.reportList.summaries=t.data.summaries,e.reportDimensions=t.data.dimensions,e.loadings.pageLoading=!1})).catch((function(t){console.log(t),e.loadings.pageLoading=!1}))})).catch((function(t){console.log(t),e.loadings.pageLoading=!1}))},getRegions:function(){var e=this;return new Promise((function(t,n){if(e.regions.length>0)return t();Object(u["h"])().then((function(n){e.regions=n.data,t()})).catch((function(e){n(e)}))}))},getAllApps:function(){var e=this;return new Promise((function(t,n){if(e.apps.length>0)return t();Object(p["a"])().then((function(n){e.apps=n.data,t()})).catch((function(e){n(e)}))}))},getAppRels:function(){var e=this;return new Promise((function(t,n){if(Object.keys(e.appRels).length>0)return t();Object(p["b"])().then((function(n){e.appRels=n.data,t()})).catch((function(e){n(e)}))}))},getAllAccounts:function(){var e=this;return new Promise((function(t,n){e.accounts.length>0?t():Object(d["f"])().then((function(n){e.accounts=n.data,t()})).catch((function(e){n(e)}))}))},handlePage:function(e){this.search.page=e,this.getReportList()},handlePageSize:function(e){this.search.page_size=e,this.getReportList()},doSearch:function(){this.search.page=1,this.getReportList()},download:function(){var e=this;this.loadings.pageLoading=!0,this.search.download=1,Object(l["b"])(this.search).then((function(t){e.search.download=0,e.loadings.downloadLoading=!1,Promise.all([n.e("chunk-4c7447ea"),n.e("chunk-2133cd4f")]).then(n.bind(null,"4bf8")).then((function(n){var a=[];t.data.columns.map((function(e){e.show&&a.push(e.label)}));var r=n.formatJson(t.data);if(0===r.length)return e.$message.info("没有筛选到需导出的数据"),void(e.loadings.pageLoading=!1);n.export_json_to_excel({header:a,data:r,filename:"综合报表数据"}),e.loadings.pageLoading=!1})).catch((function(t){console.log(t),e.loadings.pageLoading=!1}))})).catch((function(){e.search.download=0,e.loadings.pageLoading=!1}))},setDefaultSearchDate:function(){var e=new Date,t="{y}-{m}-{d}";this.$set(this.search,"date_range",[Object(s["c"])(e.getTime()-6048e5,t),Object(s["c"])(new Date,t)])},selectColumns:function(){this.$refs.column.setDefault()},confirm:function(e){this.search.show_columns=e,this.getReportList()},sortable:function(e){var t=e.column,n=e.prop,a=e.order;"custom"===t.sortable&&(this.search.order=n,this.search.by=a,this.getReportList())},getSummaries:function(e){var t=this,n=e.columns,a=[];return n.forEach((function(e,n){if(0===n)a[n]="合计";else switch(e.property){case"cost":a[n]=t.reportList.summaries.cost;break;case"earnings":a[n]=t.reportList.summaries.earnings;break;case"roi":a[n]=t.reportList.summaries.roi+"%";break;case"ad_show_count":a[n]=t.reportList.summaries.ad_show_count;break;case"ad_click_count":a[n]=t.reportList.summaries.ad_click_count;break}})),a},conditionApp:function(e){if(0===e.length)return!0;for(var t=[],n=0;n<e.length;n++)t=t.concat(this.appRels[e[n]]);for(var a=0;a<this.apps.length;a++)t.includes(this.apps[a].app_id)?this.apps[a].show=!0:this.apps[a].show=!1}}),a),v=b,_=n("cba8"),w=Object(_["a"])(v,r,o,!1,null,"14302f7f",null);t["default"]=w.exports},4221:function(e,t,n){"use strict";var a=n("4b8d"),r=n("4f7e"),o=n("c7b3"),i=n("3978"),c=n("b821"),s=n("3e87"),l=n("e001"),u=n("837a");r("search",(function(e,t,n){return[function(t){var n=i(this),r=void 0==t?void 0:l(t,e);return r?a(r,t,n):new RegExp(t)[e](s(n))},function(e){var a=o(this),r=s(e),i=n(t,a,r);if(i.done)return i.value;var l=a.lastIndex;c(l,0)||(a.lastIndex=0);var d=u(a,r);return c(a.lastIndex,l)||(a.lastIndex=l),null===d?-1:d.index}]}))},5723:function(e,t,n){"use strict";n.d(t,"e",(function(){return r})),n.d(t,"b",(function(){return o})),n.d(t,"d",(function(){return i})),n.d(t,"c",(function(){return c})),n.d(t,"a",(function(){return s})),n.d(t,"j",(function(){return l})),n.d(t,"k",(function(){return u})),n.d(t,"g",(function(){return d})),n.d(t,"f",(function(){return p})),n.d(t,"i",(function(){return f})),n.d(t,"h",(function(){return h}));var a=n("b775");function r(e){return Object(a["a"])({url:"/account/update",method:"post",data:e})}function o(e){return Object(a["a"])({url:"/account/create",method:"post",data:e})}function i(e){return Object(a["a"])({url:"/account/list",method:"get",params:e})}function c(e){return Object(a["a"])({url:"/account/"+e,method:"get"})}function s(e){return Object(a["a"])({url:"/account/auth",method:"get",params:{id:e}})}function l(e){return Object(a["a"])({url:"/account/refresh/"+e,method:"post"})}function u(e){return Object(a["a"])({url:"/account/search",method:"get",params:{account_name:e}})}function d(){return Object(a["a"])({url:"/account/default",method:"get"})}function p(){return Object(a["a"])({url:"/account/all",method:"get"})}function f(e){return Object(a["a"])({url:"/account/parents",method:"get",params:e})}function h(e){return Object(a["a"])({url:"/account/token",method:"post",data:e})}},"5c07":function(e,t,n){"use strict";var a=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("DialogPanel",{attrs:{visible:e.visible,title:"展示字段选择",width:"460px","confirm-text":"确认"},on:{confirm:e.confirm,cancel:e.cancel}},[n("el-select",{staticStyle:{width:"100%"},attrs:{multiple:"",placeholder:"请选择"},model:{value:e.selects,callback:function(t){e.selects=t},expression:"selects"}},e._l(e.Columns,(function(e){return n("el-option",{key:e.key,attrs:{label:e.label,value:e.key}})})),1)],1)},r=[],o=(n("8300"),n("60fe"),n("068b"),n("3bae"),n("d4fd")),i=n("1c31"),c={name:"SelectColumns",components:{DialogPanel:o["a"]},props:{Columns:{required:!0,type:Array},ModuleName:{required:!0,type:String}},data:function(){return{visible:!1,selects:[]}},methods:{setDefault:function(){var e=this;this.Columns.forEach((function(t){t.show&&!e.selects.includes(t.key)&&e.selects.push(t.key)})),this.visible=!0},confirm:function(){var e=this;Object(i["c"])({columns:this.selects,module:this.ModuleName}).then((function(t){e.$emit("confirm",e.selects),e.visible=!1})).catch((function(t){e.$message.error("字段设置失败："+t)}))},cancel:function(){this.visible=!1}}},s=c,l=n("cba8"),u=Object(l["a"])(s,a,r,!1,null,null,null);t["a"]=u.exports},"80f8":function(e,t,n){},b562:function(e,t,n){"use strict";n.d(t,"e",(function(){return r})),n.d(t,"c",(function(){return o})),n.d(t,"g",(function(){return i})),n.d(t,"d",(function(){return c})),n.d(t,"a",(function(){return s})),n.d(t,"h",(function(){return l})),n.d(t,"f",(function(){return u})),n.d(t,"b",(function(){return d}));var a=n("b775");function r(e){return Object(a["a"])({url:"/app/list",method:"post",data:e})}function o(e){return Object(a["a"])({url:"/app/create",method:"post",data:e})}function i(e){return Object(a["a"])({url:"/app/update",method:"post",data:e})}function c(e){return Object(a["a"])({url:"/app/"+e,method:"get"})}function s(){return Object(a["a"])({url:"/app/all",method:"get"})}function l(e){return Object(a["a"])({url:"/app/campaign-list",method:"get",params:e})}function u(e){return Object(a["a"])({url:"/app/pull",method:"post",data:e})}function d(){return Object(a["a"])({url:"/app/relation",method:"get"})}},b821:function(e,t){e.exports=Object.is||function(e,t){return e===t?0!==e||1/e===1/t:e!=e&&t!=t}},b9c3:function(e,t,n){"use strict";n("80f8")},cf0b:function(e,t,n){"use strict";var a={AccountTypeMarket:1,AccountTypeAds:2,ReportGranularity:[{name:"按日期",key:"date"},{name:"按整体",key:"all"}]};t["a"]=a},d4fd:function(e,t,n){"use strict";var a=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("el-dialog",{directives:[{name:"el-drag-dialog",rawName:"v-el-drag-dialog"}],attrs:{title:e.title,visible:e.visible,width:e.width,"append-to-body":"","modal-append-to-body":"","close-on-click-modal":!1,"close-on-press-escape":!1,"show-close":!1},on:{"update:visible":function(t){e.visible=t}}},[e._t("default"),n("div",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[n("el-button",{attrs:{loading:e.confirmLoading,icon:"el-icon-close"},on:{click:e.handleCancel}},[e._v("取消")]),n("el-button",{directives:[{name:"show",rawName:"v-show",value:e.confirmText,expression:"confirmText"}],attrs:{type:"primary",icon:"el-icon-check",loading:e.confirmLoading},on:{click:e.handleConfirm}},[e._v(e._s(e.confirmText))]),e._t("operate")],2)],2)},r=[],o={props:{confirmText:{default:"",type:String},title:{default:"",type:String},width:{default:"600px",type:String},confirmLoading:{default:!1,type:Boolean},visible:{default:!1,type:Boolean}},methods:{handleCancel:function(){this.$emit("cancel")},handleConfirm:function(){this.$emit("confirm")}}},i=o,c=(n("b9c3"),n("cba8")),s=Object(c["a"])(i,a,r,!1,null,null,null);t["a"]=s.exports},fc23:function(e,t,n){"use strict";var a=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("el-pagination",{attrs:{background:"","current-page":e.page,"page-size":e.limit,total:e.total,"hide-on-single-page":"",layout:"prev, pager, next, jumper, total, sizes","prev-text":"上页","next-text":"下页","page-sizes":[10,15,20,30,40,50,100]},on:{"current-change":e.handlePage,"size-change":e.handleSizeChange}})},r=[],o=(n("4582"),{props:{page:{type:Number,default:1},total:{type:Number,default:0},limit:{type:Number,default:10}},methods:{handlePage:function(e){this.$emit("current-change",e)},handleSizeChange:function(e){this.$emit("size-change",e)}}}),i=o,c=n("cba8"),s=Object(c["a"])(i,a,r,!1,null,null,null);t["a"]=s.exports}}]);