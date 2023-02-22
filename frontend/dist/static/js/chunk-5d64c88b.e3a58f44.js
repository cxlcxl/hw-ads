(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-5d64c88b"],{"06f2":function(e,t,n){"use strict";n.r(t);var a=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("el-row",[n("el-col",{staticClass:"search-container",attrs:{span:24}},[n("el-form",{ref:"_search",attrs:{model:e.search,inline:"",size:"small"}},[n("el-form-item",[n("el-input",{staticClass:"w200",attrs:{clearable:"",placeholder:"配置名"},model:{value:e.search._desc,callback:function(t){e.$set(e.search,"_desc",t)},expression:"search._desc"}})],1),n("el-form-item",[n("el-input",{staticClass:"w200",attrs:{clearable:"",placeholder:"配置代码"},model:{value:e.search._k,callback:function(t){e.$set(e.search,"_k",t)},expression:"search._k"}})],1),n("el-form-item",[n("el-select",{staticClass:"w110",attrs:{placeholder:"状态"},model:{value:e.search.state,callback:function(t){e.$set(e.search,"state",t)},expression:"search.state"}},[n("el-option",{key:1,attrs:{label:"正常",value:1}}),n("el-option",{key:0,attrs:{label:"停用",value:0}})],1)],1),n("el-form-item",{attrs:{label:""}},[n("el-button",{staticClass:"item",attrs:{type:"primary",icon:"el-icon-search"},on:{click:e.doSearch}},[e._v("查询")])],1)],1)],1),n("el-col",{attrs:{span:24}},[n("el-button",{attrs:{type:"primary",icon:"el-icon-plus",size:"mini"},on:{click:e.add}},[e._v("添加配置")])],1),n("el-col",{attrs:{span:24}},[n("el-table",{directives:[{name:"loading",rawName:"v-loading",value:e.loadings.pageLoading,expression:"loadings.pageLoading"}],staticStyle:{"margin-top":"15px"},attrs:{data:e.confList.list,"highlight-current-row":"",stripe:"",border:"",size:"mini"},on:{"sort-change":e.sortChange}},[n("el-table-column",{attrs:{prop:"desc",label:"配置名",width:"180","show-overflow-tooltip":""}}),n("el-table-column",{attrs:{prop:"key",label:"配置代码",width:"240",sortable:""}}),n("el-table-column",{attrs:{prop:"val",label:"值"}}),n("el-table-column",{attrs:{prop:"bak1",label:"扩展字段1"}}),n("el-table-column",{attrs:{prop:"bak2",label:"扩展字段2"}}),n("el-table-column",{attrs:{prop:"remark",label:"备注"}}),n("el-table-column",{attrs:{label:"状态",width:"80",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(1===t.row.state?"正常":"停用")+" ")]}}])}),n("el-table-column",{attrs:{align:"center",label:"操作",fixed:"right",width:"100"},scopedSlots:e._u([{key:"default",fn:function(t){return[n("el-button-group",{staticClass:"table-operate"},[n("el-button",{attrs:{type:"primary",plain:""},nativeOn:{click:function(n){return n.preventDefault(),e.editRow(t.row)}}},[e._v("编辑")]),n("el-button",{attrs:{type:"primary",plain:""},nativeOn:{click:function(n){return n.preventDefault(),e.copyRow(t.row)}}},[e._v("复制")])],1)]}}])})],1)],1),n("el-col",{staticStyle:{"margin-top":"15px"},attrs:{span:24}},[n("page",{ref:"page",attrs:{page:e.search.page,total:e.confList.total},on:{"current-change":e.handlePage,"size-change":e.handlePageSize}}),n("conf-create",{ref:"confCreate",on:{success:e.getConfList}}),n("conf-update",{ref:"confUpdate",on:{success:e.getConfList}})],1)],1)},o=[],r=(n("2ce8"),n("4221"),n("2934")),l=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("dialog-panel",{attrs:{title:"配置添加","confirm-text":"添加",visible:e.visible,"confirm-loading":e.loading,width:"500px"},on:{cancel:e.cancel,confirm:e.add}},[n("el-form",{ref:"confForm",attrs:{model:e.confForm,"label-width":"100px",size:"small"}},[n("el-form-item",{attrs:{label:"配置名称",prop:"desc",rules:{required:!0,message:"请填写配置名称"}}},[n("el-input",{model:{value:e.confForm.desc,callback:function(t){e.$set(e.confForm,"desc",t)},expression:"confForm.desc"}})],1),n("el-form-item",{attrs:{label:"配置代码",prop:"key",rules:{required:!0,message:"请填写配置代码"}}},[n("el-input",{attrs:{placeholder:"仅支持字符开头的大小写字母数字下划线组合 [50 位以内]"},model:{value:e.confForm.key,callback:function(t){e.$set(e.confForm,"key",t)},expression:"confForm.key"}})],1),n("el-form-item",{attrs:{label:"配置值",prop:"val",rules:{required:!0,message:"请填写配置值"}}},[n("el-input",{model:{value:e.confForm.val,callback:function(t){e.$set(e.confForm,"val",t)},expression:"confForm.val"}})],1),n("el-form-item",{attrs:{label:"扩展1",prop:"bak1"}},[n("el-input",{model:{value:e.confForm.bak1,callback:function(t){e.$set(e.confForm,"bak1",t)},expression:"confForm.bak1"}})],1),n("el-form-item",{attrs:{label:"扩展2",prop:"bak2"}},[n("el-input",{model:{value:e.confForm.bak2,callback:function(t){e.$set(e.confForm,"bak2",t)},expression:"confForm.bak2"}})],1),n("el-form-item",{attrs:{label:"配置描述",prop:"remark"}},[n("el-input",{model:{value:e.confForm.remark,callback:function(t){e.$set(e.confForm,"remark",t)},expression:"confForm.remark"}})],1)],1)],1)},i=[],c=n("d4fd"),s={components:{DialogPanel:c["a"]},data:function(){return{visible:!1,loading:!1,confForm:{key:"",desc:"",bak1:"",bak2:"",remark:"",val:""}}},methods:{cancel:function(){this.$refs.confForm.resetFields(),this.visible=!1},add:function(){var e=this;this.$refs.confForm.validate((function(t){if(!t)return!1;e.loading=!0,Object(r["a"])(e.confForm).then((function(t){e.$message.success("创建成功"),e.$emit("success"),e.loading=!1,e.cancel()})).catch((function(t){e.loading=!1,console.log(t)}))}))}}},u=s,f=n("cba8"),d=Object(f["a"])(u,l,i,!1,null,null,null),m=d.exports,p=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("dialog-panel",{attrs:{title:"配置修改","confirm-text":"保存",visible:e.visible,"confirm-loading":e.loading,width:"500px"},on:{cancel:e.cancel,confirm:e.save}},[n("el-form",{ref:"confForm",attrs:{model:e.confForm,"label-width":"100px",size:"small"}},[n("el-form-item",{attrs:{label:"配置名称",prop:"desc",rules:{required:!0,message:"请填写配置名称"}}},[n("el-input",{model:{value:e.confForm.desc,callback:function(t){e.$set(e.confForm,"desc",t)},expression:"confForm.desc"}})],1),n("el-form-item",{attrs:{label:"配置代码",prop:"key",rules:{required:!0,message:"请填写配置代码"}}},[n("el-input",{attrs:{placeholder:"仅支持字符开头的大小写字母数字下划线组合 [50 位以内]"},model:{value:e.confForm.key,callback:function(t){e.$set(e.confForm,"key",t)},expression:"confForm.key"}})],1),n("el-form-item",{attrs:{label:"配置值",prop:"val",rules:{required:!0,message:"请填写配置值"}}},[n("el-input",{model:{value:e.confForm.val,callback:function(t){e.$set(e.confForm,"val",t)},expression:"confForm.val"}})],1),n("el-form-item",{attrs:{label:"状态",prop:"state"}},[n("el-switch",{attrs:{"active-value":1,"inactive-value":0},model:{value:e.confForm.state,callback:function(t){e.$set(e.confForm,"state",t)},expression:"confForm.state"}})],1),n("el-form-item",{attrs:{label:"扩展1",prop:"bak1"}},[n("el-input",{model:{value:e.confForm.bak1,callback:function(t){e.$set(e.confForm,"bak1",t)},expression:"confForm.bak1"}})],1),n("el-form-item",{attrs:{label:"扩展2",prop:"bak2"}},[n("el-input",{model:{value:e.confForm.bak2,callback:function(t){e.$set(e.confForm,"bak2",t)},expression:"confForm.bak2"}})],1),n("el-form-item",{attrs:{label:"配置描述",prop:"remark"}},[n("el-input",{model:{value:e.confForm.remark,callback:function(t){e.$set(e.confForm,"remark",t)},expression:"confForm.remark"}})],1)],1)],1)},b=[],g={components:{DialogPanel:c["a"]},data:function(){return{visible:!1,loading:!1,confForm:{id:0,key:"",val:"",desc:"",state:1,bak1:"",bak2:"",remark:""}}},methods:{initPage:function(e){var t=this;Object(r["b"])(e).then((function(e){t.confForm=e.data,t.visible=!0})).catch((function(){}))},cancel:function(){this.$refs.confForm.resetFields(),this.visible=!1},save:function(){var e=this;this.$refs.confForm.validate((function(t){if(!t)return!1;e.loading=!0,Object(r["c"])(e.confForm).then((function(t){e.$message.success("修改成功"),e.$emit("success"),e.loading=!1,e.cancel()})).catch((function(t){e.loading=!1,console.log(t)}))}))}}},h=g,v=Object(f["a"])(h,p,b,!1,null,null,null),k=v.exports,F=n("fc23"),x={name:"Config",components:{Page:F["a"],ConfCreate:m,ConfUpdate:k},data:function(){return{loadings:{pageLoading:!1},confList:{total:0,list:[]},search:{_k:"",_desc:"",state:1,page:1,page_size:10}}},computed:{},mounted:function(){this.getConfList()},methods:{getConfList:function(){var e=this;this.loadings.pageLoading=!0,Object(r["m"])(this.search).then((function(t){e.confList=t.data,e.loadings.pageLoading=!1})).catch((function(){e.loadings.pageLoading=!1}))},add:function(){this.$refs.confCreate.visible=!0},copyRow:function(e){this.$refs.confCreate.confForm={desc:e.desc,key:e.key,val:e.val,state:e.state,bak1:e.bak1,bak2:e.bak2,remark:e.remark},this.$refs.confCreate.visible=!0},editRow:function(e){this.$refs.confUpdate.initPage(e.id)},doSearch:function(){this.search.page=1,this.getConfList()},handlePage:function(e){this.search.page=e,this.getConfList()},handlePageSize:function(e){this.search.page_size=e,this.getConfList()},sortChange:function(e){e.column;var t=e.prop,n=e.order}}},y=x,w=Object(f["a"])(y,a,o,!1,null,null,null);t["default"]=w.exports},2934:function(e,t,n){"use strict";n.d(t,"h",(function(){return o})),n.d(t,"p",(function(){return r})),n.d(t,"d",(function(){return l})),n.d(t,"e",(function(){return i})),n.d(t,"g",(function(){return c})),n.d(t,"f",(function(){return s})),n.d(t,"i",(function(){return u})),n.d(t,"j",(function(){return f})),n.d(t,"l",(function(){return d})),n.d(t,"k",(function(){return m})),n.d(t,"m",(function(){return p})),n.d(t,"a",(function(){return b})),n.d(t,"b",(function(){return g})),n.d(t,"c",(function(){return h})),n.d(t,"o",(function(){return v})),n.d(t,"n",(function(){return k}));var a=n("b775");function o(){return Object(a["a"])({url:"/regions",method:"get"})}function r(){return Object(a["a"])({url:"/settings/version",method:"get"})}function l(){return Object(a["a"])({url:"/region/area",method:"get"})}function i(e){return Object(a["a"])({url:"/region/country",method:"get",params:e})}function c(e){return Object(a["a"])({url:"/region",method:"post",data:e})}function s(e){return Object(a["a"])({url:"/region/area-set",method:"post",data:e})}function u(){return Object(a["a"])({url:"/settings/cron",method:"get"})}function f(e){return Object(a["a"])({url:"/settings/cron/"+e,method:"get"})}function d(e,t){return Object(a["a"])({url:"/settings/cron/"+t,method:"post",data:e})}function m(e){return Object(a["a"])({url:"/settings/cron/schedule",method:"post",data:e})}function p(e){return Object(a["a"])({url:"/settings/configs",method:"get",params:e})}function b(e){return Object(a["a"])({url:"/settings/config",method:"post",data:e})}function g(e){return Object(a["a"])({url:"/settings/config/"+e,method:"get"})}function h(e,t){return Object(a["a"])({url:"/settings/config/"+t,method:"post",data:e})}function v(e){return Object(a["a"])({url:"/settings/log",method:"post",data:e})}var k="https://hiads.mobgi.cc/api/settings/log/"},4221:function(e,t,n){"use strict";var a=n("4b8d"),o=n("4f7e"),r=n("c7b3"),l=n("3978"),i=n("b821"),c=n("3e87"),s=n("e001"),u=n("837a");o("search",(function(e,t,n){return[function(t){var n=l(this),o=void 0==t?void 0:s(t,e);return o?a(o,t,n):new RegExp(t)[e](c(n))},function(e){var a=r(this),o=c(e),l=n(t,a,o);if(l.done)return l.value;var s=a.lastIndex;i(s,0)||(a.lastIndex=0);var f=u(a,o);return i(a.lastIndex,s)||(a.lastIndex=s),null===f?-1:f.index}]}))},"80f8":function(e,t,n){},b821:function(e,t){e.exports=Object.is||function(e,t){return e===t?0!==e||1/e===1/t:e!=e&&t!=t}},b9c3:function(e,t,n){"use strict";n("80f8")},d4fd:function(e,t,n){"use strict";var a=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("el-dialog",{directives:[{name:"el-drag-dialog",rawName:"v-el-drag-dialog"}],attrs:{title:e.title,visible:e.visible,width:e.width,"append-to-body":"","modal-append-to-body":"","close-on-click-modal":!1,"close-on-press-escape":!1,"show-close":!1},on:{"update:visible":function(t){e.visible=t}}},[e._t("default"),n("div",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[n("el-button",{attrs:{loading:e.confirmLoading,icon:"el-icon-close"},on:{click:e.handleCancel}},[e._v("取消")]),n("el-button",{directives:[{name:"show",rawName:"v-show",value:e.confirmText,expression:"confirmText"}],attrs:{type:"primary",icon:"el-icon-check",loading:e.confirmLoading},on:{click:e.handleConfirm}},[e._v(e._s(e.confirmText))]),e._t("operate")],2)],2)},o=[],r={props:{confirmText:{default:"",type:String},title:{default:"",type:String},width:{default:"600px",type:String},confirmLoading:{default:!1,type:Boolean},visible:{default:!1,type:Boolean}},methods:{handleCancel:function(){this.$emit("cancel")},handleConfirm:function(){this.$emit("confirm")}}},l=r,i=(n("b9c3"),n("cba8")),c=Object(i["a"])(l,a,o,!1,null,null,null);t["a"]=c.exports},fc23:function(e,t,n){"use strict";var a=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("el-pagination",{attrs:{background:"","current-page":e.page,"page-size":e.limit,total:e.total,"hide-on-single-page":"",layout:"prev, pager, next, jumper, total, sizes","prev-text":"上页","next-text":"下页","page-sizes":[10,15,20,30,40,50,100]},on:{"current-change":e.handlePage,"size-change":e.handleSizeChange}})},o=[],r=(n("4582"),{props:{page:{type:Number,default:1},total:{type:Number,default:0},limit:{type:Number,default:10}},methods:{handlePage:function(e){this.$emit("current-change",e)},handleSizeChange:function(e){this.$emit("size-change",e)}}}),l=r,i=n("cba8"),c=Object(i["a"])(l,a,o,!1,null,null,null);t["a"]=c.exports}}]);