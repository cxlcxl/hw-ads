(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-5bb72a10"],{2934:function(e,t,n){"use strict";n.d(t,"h",(function(){return r})),n.d(t,"d",(function(){return i})),n.d(t,"e",(function(){return o})),n.d(t,"g",(function(){return l})),n.d(t,"f",(function(){return s})),n.d(t,"i",(function(){return c})),n.d(t,"j",(function(){return u})),n.d(t,"l",(function(){return d})),n.d(t,"k",(function(){return m})),n.d(t,"m",(function(){return g})),n.d(t,"a",(function(){return f})),n.d(t,"b",(function(){return p})),n.d(t,"c",(function(){return h}));var a=n("b775");function r(){return Object(a["a"])({url:"/regions",method:"get"})}function i(){return Object(a["a"])({url:"/region/area",method:"get"})}function o(e){return Object(a["a"])({url:"/region/country",method:"get",params:e})}function l(e){return Object(a["a"])({url:"/region",method:"post",data:e})}function s(e){return Object(a["a"])({url:"/region/area-set",method:"post",data:e})}function c(){return Object(a["a"])({url:"/settings/cron",method:"get"})}function u(e){return Object(a["a"])({url:"/settings/cron/"+e,method:"get"})}function d(e,t){return Object(a["a"])({url:"/settings/cron/"+t,method:"post",data:e})}function m(e){return Object(a["a"])({url:"/settings/cron/schedule",method:"post",data:e})}function g(e){return Object(a["a"])({url:"/settings/configs",method:"get",params:e})}function f(e){return Object(a["a"])({url:"/settings/config",method:"post",data:e})}function p(e){return Object(a["a"])({url:"/settings/config/"+e,method:"get"})}function h(e,t){return Object(a["a"])({url:"/settings/config/"+t,method:"post",data:e})}},4221:function(e,t,n){"use strict";var a=n("4b8d"),r=n("4f7e"),i=n("c7b3"),o=n("3978"),l=n("b821"),s=n("3e87"),c=n("e001"),u=n("837a");r("search",(function(e,t,n){return[function(t){var n=o(this),r=void 0==t?void 0:c(t,e);return r?a(r,t,n):new RegExp(t)[e](s(n))},function(e){var a=i(this),r=s(e),o=n(t,a,r);if(o.done)return o.value;var c=a.lastIndex;l(c,0)||(a.lastIndex=0);var d=u(a,r);return l(a.lastIndex,c)||(a.lastIndex=c),null===d?-1:d.index}]}))},"80f8":function(e,t,n){},b821:function(e,t){e.exports=Object.is||function(e,t){return e===t?0!==e||1/e===1/t:e!=e&&t!=t}},b9c3:function(e,t,n){"use strict";n("80f8")},cf27:function(e,t,n){"use strict";n.r(t);var a=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("el-row",{attrs:{gutter:10}},[n("el-col",{staticClass:"search-container",attrs:{span:24}},[n("el-form",{ref:"_search",attrs:{model:e.search,inline:"",size:"small"}},[n("el-form-item",[n("el-input",{staticClass:"w240",attrs:{clearable:"",placeholder:"输入搜索：国家名称/代码"},model:{value:e.search.k,callback:function(t){e.$set(e.search,"k",t)},expression:"search.k"}})],1),n("el-form-item",{attrs:{label:""}},[n("el-button",{staticClass:"item",attrs:{type:"primary",icon:"el-icon-search"},on:{click:e.doSearch}},[e._v("查询")])],1)],1)],1),n("el-col",{attrs:{span:24}},[n("el-button",{attrs:{type:"primary",icon:"el-icon-plus",size:"mini"},on:{click:e.add}},[e._v("添加区域/国家")]),n("span",{staticClass:"notice text-error"},[e._v("国家信息因涉及投放定向设置，所以不可修改只可新增「"),n("a",{attrs:{href:"https://developer.huawei.com/consumer/cn/doc/distribution/promotion/marketing-api-tool-targeting7-0000001286343134#ZH-CN_TOPIC_0000001286343134__li1079145116117",target:"_blank"}},[e._v("国家数据来源")]),e._v("」")])],1),n("el-col",{staticStyle:{"margin-top":"15px"},attrs:{span:24}},[n("div",{staticClass:"region-area-country"},[n("div",{staticClass:"region-area"},[n("el-menu",{staticClass:"region-area-menu",attrs:{"default-active":e.search.area_id,width:"160px"},on:{select:e.handleSelect}},[n("el-menu-item",{attrs:{index:"-1"}},[e._v("全部国家")]),n("el-menu-item",{attrs:{index:"-2"}},[e._v("未分地区")]),e._l(e.areas,(function(t){return n("el-menu-item",{key:t.id,attrs:{index:t.id.toString()}},[e._v(e._s(t.name))])}))],2)],1),n("div",{staticClass:"region-country"},[n("el-table",{directives:[{name:"loading",rawName:"v-loading",value:e.loading,expression:"loading"}],attrs:{data:e.countries.list,size:"mini",border:"",stripe:""}},[n("el-table-column",{attrs:{align:"center",label:"操作",width:"100"},scopedSlots:e._u([{key:"default",fn:function(t){return[n("el-button-group",{staticClass:"table-operate"},[n("el-button",{attrs:{type:"primary",plain:""},nativeOn:{click:function(n){return n.preventDefault(),e.editRow(t.row)}}},[e._v("设置地区")])],1)]}}])}),n("el-table-column",{attrs:{label:"国家代码",prop:"c_code",align:"center",width:"80"}}),n("el-table-column",{attrs:{label:"国家名称",prop:"c_name"}}),n("el-table-column",{attrs:{label:"地区",prop:"area_name",width:"130"}})],1),n("div",{staticStyle:{"margin-top":"10px"}},[n("page",{ref:"page",attrs:{page:e.search.page,total:e.countries.total,limit:e.search.page_size},on:{"current-change":e.handlePage,"size-change":e.handlePageSize}})],1)],1)])]),n("region-create",{ref:"regionCreate",attrs:{areas:e.areas},on:{success:e.getRegionList}}),n("region-area-set",{ref:"regionAreaSet",attrs:{areas:e.areas},on:{success:e.getRegionList}})],1)},r=[],i=(n("8300"),n("ea5b"),n("e551"),n("2ce8"),n("4221"),n("7241"),n("2934")),o=n("fc23"),l=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("dialog-panel",{attrs:{title:"添加地区/国家","confirm-text":"添加",visible:e.visible,"confirm-loading":e.loading,width:"468px"},on:{cancel:e.cancel,confirm:e.add}},[n("el-form",{ref:"regionForm",attrs:{model:e.regionForm,"label-width":"90px",size:"small"}},[n("el-tabs",{attrs:{type:"border-card"},model:{value:e.regionForm.t,callback:function(t){e.$set(e.regionForm,"t",t)},expression:"regionForm.t"}},[n("el-tab-pane",{attrs:{label:"地区信息",name:"area"}},[n("el-form-item",{attrs:{label:"地区名称",prop:"area_name",rules:{required:!0,message:"请填写地区名称"}}},[n("el-input",{attrs:{placeholder:"请填写地区名称"},model:{value:e.regionForm.area_name,callback:function(t){e.$set(e.regionForm,"area_name",t)},expression:"regionForm.area_name"}})],1)],1),n("el-tab-pane",{attrs:{label:"国家信息",name:"country"}},["country"===e.regionForm.t?n("el-form-item",{attrs:{label:"所属地区",prop:"area_id",rules:{required:!0,message:"请选择所属地区"}}},[n("el-select",{attrs:{placeholder:"请选择所属地区"},model:{value:e.regionForm.area_id,callback:function(t){e.$set(e.regionForm,"area_id",t)},expression:"regionForm.area_id"}},e._l(e.areas,(function(e){return n("el-option",{key:e.id,attrs:{value:Number(e.id),label:e.name}})})),1)],1):e._e(),"country"===e.regionForm.t?n("el-form-item",{attrs:{label:"国家名称",prop:"c_name",rules:{required:!0,message:"请填写国家名称"}}},[n("el-input",{attrs:{placeholder:"请填写国家名称"},model:{value:e.regionForm.c_name,callback:function(t){e.$set(e.regionForm,"c_name",t)},expression:"regionForm.c_name"}})],1):e._e(),"country"===e.regionForm.t?n("el-form-item",{attrs:{label:"国家代码",prop:"c_code",rules:{required:!0,message:"请填写国家代码"}}},[n("el-input",{attrs:{placeholder:"请填写国家代码"},model:{value:e.regionForm.c_code,callback:function(t){e.$set(e.regionForm,"c_code",t)},expression:"regionForm.c_code"}})],1):e._e()],1)],1)],1)],1)},s=[],c=n("d4fd"),u={components:{DialogPanel:c["a"]},props:{areas:{required:!0,type:Array}},data:function(){return{visible:!1,loading:!1,remoteLoading:!1,regionForm:{t:"area",area_id:0,c_code:"",c_name:"",area_name:""}}},methods:{setDefault:function(){this.visible=!0},cancel:function(){this.$refs.regionForm.resetFields(),this.visible=!1},add:function(){var e=this;this.$refs.regionForm.validate((function(t){if(!t)return!1;e.loading=!0,Object(i["g"])(e.regionForm).then((function(t){e.$message.success("创建成功"),e.$emit("success"),e.loading=!1,e.cancel()})).catch((function(t){e.loading=!1,console.log(t)}))}))}}},d=u,m=n("cba8"),g=Object(m["a"])(d,l,s,!1,null,null,null),f=g.exports,p=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("dialog-panel",{attrs:{title:"添加地区/国家","confirm-text":"保存",visible:e.visible,"confirm-loading":e.loading,width:"368px"},on:{cancel:e.cancel,confirm:e.save}},[n("el-form",{ref:"regionForm",attrs:{model:e.regionForm,"label-width":"90px",size:"small"}},[n("el-form-item",{attrs:{label:"国家名称"}},[n("el-input",{attrs:{value:e.regionForm.c_name,disabled:""}})],1),n("el-form-item",{attrs:{label:"所属地区",prop:"area_id",rules:{required:!0,message:"请选择所属地区"}}},[n("el-select",{attrs:{placeholder:"请选择所属地区"},model:{value:e.regionForm.area_id,callback:function(t){e.$set(e.regionForm,"area_id",t)},expression:"regionForm.area_id"}},e._l(e.areas,(function(e){return n("el-option",{key:e.id,attrs:{value:Number(e.id),label:e.name}})})),1)],1)],1)],1)},h=[],b={components:{DialogPanel:c["a"]},props:{areas:{required:!0,type:Array}},data:function(){return{visible:!1,loading:!1,regionForm:{c_name:"",c_code:"",area_id:0}}},methods:{setDefault:function(e,t){this.regionForm.c_code=e,this.regionForm.c_name=t,this.visible=!0},cancel:function(){this.$refs.regionForm.resetFields(),this.visible=!1},save:function(){var e=this;this.$refs.regionForm.validate((function(t){if(!t)return!1;e.loading=!0,Object(i["f"])(e.regionForm).then((function(t){e.$message.success("设置成功"),e.$emit("success"),e.loading=!1,e.cancel()})).catch((function(t){e.loading=!1,console.log(t)}))}))}}},v=b,_=Object(m["a"])(v,p,h,!1,null,null,null),x=_.exports,F={name:"Region",components:{Page:o["a"],RegionCreate:f,RegionAreaSet:x},data:function(){return{loading:!1,search:{area_id:"-1",k:"",page:1,page_size:10},areas:[],countries:{total:0,list:[]}}},mounted:function(){this.getRegionList()},methods:{getRegionList:function(){var e=this;this.loading=!0,Promise.all([this.getRegionAreas()]).then((function(){Object(i["e"])(e.search).then((function(t){e.loading=!1,e.countries.list=t.data.list,e.countries.total=t.data.total})).catch((function(){e.loading=!1}))})).catch((function(t){e.loading=!1}))},getRegionAreas:function(){var e=this;return new Promise((function(t,n){Object(i["d"])().then((function(n){e.areas=n.data,t()})).catch((function(e){n()}))}))},add:function(){this.$refs.regionCreate.setDefault()},editRow:function(e){this.$refs.regionAreaSet.setDefault(e.c_code,e.c_name)},handlePage:function(e){this.search.page=e,this.getRegionList()},handlePageSize:function(e){this.search.page_size=e,this.doSearch()},handleSelect:function(e){this.search.area_id=e,this.doSearch()},doSearch:function(){this.search.page=1,this.getRegionList()},setRegion:function(e){console.log(e.name)}}},y=F,w=(n("d859"),Object(m["a"])(y,a,r,!1,null,"06fad778",null));t["default"]=w.exports},d4fd:function(e,t,n){"use strict";var a=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("el-dialog",{directives:[{name:"el-drag-dialog",rawName:"v-el-drag-dialog"}],attrs:{title:e.title,visible:e.visible,width:e.width,"append-to-body":"","modal-append-to-body":"","close-on-click-modal":!1,"close-on-press-escape":!1,"show-close":!1},on:{"update:visible":function(t){e.visible=t}}},[e._t("default"),n("div",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[n("el-button",{attrs:{loading:e.confirmLoading,icon:"el-icon-close"},on:{click:e.handleCancel}},[e._v("取消")]),n("el-button",{directives:[{name:"show",rawName:"v-show",value:e.confirmText,expression:"confirmText"}],attrs:{type:"primary",icon:"el-icon-check",loading:e.confirmLoading},on:{click:e.handleConfirm}},[e._v(e._s(e.confirmText))]),e._t("operate")],2)],2)},r=[],i={props:{confirmText:{default:"",type:String},title:{default:"",type:String},width:{default:"600px",type:String},confirmLoading:{default:!1,type:Boolean},visible:{default:!1,type:Boolean}},methods:{handleCancel:function(){this.$emit("cancel")},handleConfirm:function(){this.$emit("confirm")}}},o=i,l=(n("b9c3"),n("cba8")),s=Object(l["a"])(o,a,r,!1,null,null,null);t["a"]=s.exports},d859:function(e,t,n){"use strict";n("f1b1")},f1b1:function(e,t,n){},fc23:function(e,t,n){"use strict";var a=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("el-pagination",{attrs:{background:"","current-page":e.page,"page-size":e.limit,total:e.total,"hide-on-single-page":"",layout:"prev, pager, next, jumper, total, sizes","prev-text":"上页","next-text":"下页","page-sizes":[10,15,20,30,40,50,100]},on:{"current-change":e.handlePage,"size-change":e.handleSizeChange}})},r=[],i=(n("4582"),{props:{page:{type:Number,default:1},total:{type:Number,default:0},limit:{type:Number,default:10}},methods:{handlePage:function(e){this.$emit("current-change",e)},handleSizeChange:function(e){this.$emit("size-change",e)}}}),o=i,l=n("cba8"),s=Object(l["a"])(o,a,r,!1,null,null,null);t["a"]=s.exports}}]);