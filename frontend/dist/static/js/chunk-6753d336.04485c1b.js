(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-6753d336"],{4221:function(e,t,r){"use strict";var a=r("4b8d"),l=r("4f7e"),s=r("c7b3"),n=r("3978"),o=r("b821"),i=r("3e87"),c=r("e001"),u=r("837a");l("search",(function(e,t,r){return[function(t){var r=n(this),l=void 0==t?void 0:c(t,e);return l?a(l,t,r):new RegExp(t)[e](i(r))},function(e){var a=s(this),l=i(e),n=r(t,a,l);if(n.done)return n.value;var c=a.lastIndex;o(c,0)||(a.lastIndex=0);var d=u(a,l);return o(a.lastIndex,c)||(a.lastIndex=c),null===d?-1:d.index}]}))},"80f8":function(e,t,r){},b821:function(e,t){e.exports=Object.is||function(e,t){return e===t?0!==e||1/e===1/t:e!=e&&t!=t}},b9c3:function(e,t,r){"use strict";r("80f8")},cc5e:function(e,t,r){"use strict";r.d(t,"h",(function(){return l})),r.d(t,"e",(function(){return s})),r.d(t,"i",(function(){return n})),r.d(t,"f",(function(){return o})),r.d(t,"g",(function(){return i})),r.d(t,"c",(function(){return c})),r.d(t,"a",(function(){return u})),r.d(t,"b",(function(){return d})),r.d(t,"d",(function(){return m}));var a=r("b775");function l(e){return Object(a["a"])({url:"/role/list",method:"get",params:e})}function s(e){return Object(a["a"])({url:"/role/create",method:"post",data:e})}function n(e){return Object(a["a"])({url:"/role/update",method:"post",data:e})}function o(e){return Object(a["a"])({url:"/role/destroy",method:"post",data:{id:e}})}function i(e){return Object(a["a"])({url:"/role/"+e,method:"get"})}function c(){return Object(a["a"])({url:"/permission/tree",method:"get"})}function u(e){return Object(a["a"])({url:"/permission/create",method:"post",data:e})}function d(e){return Object(a["a"])({url:"/permission/destroy",method:"post",data:{id:e}})}function m(e){return Object(a["a"])({url:"/permission/update",method:"post",data:e})}},d4fd:function(e,t,r){"use strict";var a=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("el-dialog",{directives:[{name:"el-drag-dialog",rawName:"v-el-drag-dialog"}],attrs:{title:e.title,visible:e.visible,width:e.width,"append-to-body":"","modal-append-to-body":"","close-on-click-modal":!1,"close-on-press-escape":!1,"show-close":!1},on:{"update:visible":function(t){e.visible=t}}},[e._t("default"),r("div",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[r("el-button",{attrs:{loading:e.confirmLoading,icon:"el-icon-close"},on:{click:e.handleCancel}},[e._v("取消")]),r("el-button",{directives:[{name:"show",rawName:"v-show",value:e.confirmText,expression:"confirmText"}],attrs:{type:"primary",icon:"el-icon-check",loading:e.confirmLoading},on:{click:e.handleConfirm}},[e._v(e._s(e.confirmText))]),e._t("operate")],2)],2)},l=[],s={props:{confirmText:{default:"",type:String},title:{default:"",type:String},width:{default:"600px",type:String},confirmLoading:{default:!1,type:Boolean},visible:{default:!1,type:Boolean}},methods:{handleCancel:function(){this.$emit("cancel")},handleConfirm:function(){this.$emit("confirm")}}},n=s,o=(r("b9c3"),r("cba8")),i=Object(o["a"])(n,a,l,!1,null,null,null);t["a"]=i.exports},e093:function(e,t,r){"use strict";r.r(t);var a=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("el-row",[r("el-col",{staticClass:"search-container",attrs:{span:24}},[r("el-form",{ref:"_search",attrs:{model:e.search,inline:"",size:"small"}},[r("el-form-item",[r("el-input",{staticClass:"w150",attrs:{clearable:"",placeholder:"用户名"},model:{value:e.search.username,callback:function(t){e.$set(e.search,"username",t)},expression:"search.username"}})],1),r("el-form-item",[r("el-input",{staticClass:"w230",attrs:{clearable:"",placeholder:"邮箱"},model:{value:e.search.email,callback:function(t){e.$set(e.search,"email",t)},expression:"search.email"}})],1),r("el-form-item",[r("el-select",{staticClass:"w110",attrs:{placeholder:"用户状态"},model:{value:e.search.state,callback:function(t){e.$set(e.search,"state",t)},expression:"search.state"}},[r("el-option",{key:1,attrs:{label:"正常",value:1}}),r("el-option",{key:0,attrs:{label:"停用",value:0}})],1)],1),r("el-form-item",{attrs:{label:""}},[r("el-button",{staticClass:"item",attrs:{type:"primary",icon:"el-icon-search"},on:{click:e.doSearch}},[e._v("查询")])],1)],1)],1),r("el-col",{attrs:{span:24}},[r("el-button",{attrs:{type:"primary",icon:"el-icon-plus",size:"mini"},on:{click:e.add}},[e._v("添加用户")])],1),r("el-col",{attrs:{span:24}},[r("el-table",{directives:[{name:"loading",rawName:"v-loading",value:e.loadings.pageLoading,expression:"loadings.pageLoading"}],staticStyle:{"margin-top":"15px"},attrs:{data:e.userList.list,"highlight-current-row":"",stripe:"",border:"",size:"mini"}},[r("el-table-column",{attrs:{prop:"id",label:"ID",width:"80",align:"center"}}),r("el-table-column",{attrs:{prop:"email",label:"邮箱",width:"220"}}),r("el-table-column",{attrs:{prop:"username",label:"用户名",width:"180"}}),r("el-table-column",{attrs:{prop:"mobile",label:"手机号",width:"130"}}),r("el-table-column",{attrs:{label:"角色"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(e._s(e._f("userRolesFilter")(t.row.role_id,e.roles)))]}}])}),r("el-table-column",{attrs:{label:"状态",width:"90",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(1===t.row.state?"正常":"停用")+" ")]}}])}),r("el-table-column",{attrs:{prop:"created_at",label:"添加时间",width:"140",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(e._s(e._f("timeFormat")(t.row.created_at)))]}}])}),r("el-table-column",{attrs:{align:"center",label:"操作",fixed:"right",width:"100"},scopedSlots:e._u([{key:"default",fn:function(t){return[r("el-button-group",{staticClass:"table-operate"},[r("el-button",{attrs:{type:"primary",plain:""},nativeOn:{click:function(r){return r.preventDefault(),e.editRow(t.row)}}},[e._v("编辑")])],1)]}}])})],1)],1),r("el-col",{staticStyle:{"margin-top":"15px"},attrs:{span:24}},[r("page",{ref:"page",attrs:{page:e.search.page,total:e.userList.total},on:{"current-change":e.handlePage,"size-change":e.handlePageSize}}),r("user-create",{ref:"userCreate",attrs:{roles:e.roles},on:{success:e.getUserList}}),r("user-update",{ref:"userUpdate",attrs:{roles:e.roles},on:{success:e.getUserList}})],1)],1)},l=[],s=(r("8300"),r("ea5b"),r("e551"),r("2ce8"),r("4221"),r("c24f")),n=r("cc5e"),o=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("dialog-panel",{attrs:{title:"添加用户","confirm-text":"添加",visible:e.visible,"confirm-loading":e.loading,width:"388px"},on:{cancel:e.cancel,confirm:e.add}},[r("el-form",{ref:"userForm",attrs:{model:e.userForm,"label-width":"90px",size:"small",rules:e.userRules}},[r("el-form-item",{attrs:{label:"用户名称",prop:"username"}},[r("el-input",{attrs:{placeholder:"请填写用户名"},model:{value:e.userForm.username,callback:function(t){e.$set(e.userForm,"username",t)},expression:"userForm.username"}})],1),r("el-form-item",{attrs:{label:"邮箱地址",prop:"email"}},[r("el-input",{attrs:{placeholder:"请填写邮箱，登陆使用"},model:{value:e.userForm.email,callback:function(t){e.$set(e.userForm,"email",t)},expression:"userForm.email"}})],1),r("el-form-item",{attrs:{label:"手机号",prop:"mobile"}},[r("el-input",{attrs:{placeholder:"请填写手机号"},model:{value:e.userForm.mobile,callback:function(t){e.$set(e.userForm,"mobile",t)},expression:"userForm.mobile"}})],1),r("el-form-item",{attrs:{label:"角色",prop:"role_id"}},[r("el-select",{staticStyle:{width:"100%"},attrs:{placeholder:"请选择",clearable:""},model:{value:e.userForm.role_id,callback:function(t){e.$set(e.userForm,"role_id",t)},expression:"userForm.role_id"}},e._l(e.roles,(function(e){return r("el-option",{key:e.id,attrs:{label:e.role_name,value:e.id}})})),1)],1),r("el-form-item",{attrs:{label:"登录密码",prop:"pass"}},[r("el-input",{attrs:{placeholder:"字母开头，数字特殊字符 [@.&!#?,%$] 的 6 - 18 位"},model:{value:e.userForm.pass,callback:function(t){e.$set(e.userForm,"pass",t)},expression:"userForm.pass"}})],1)],1)],1)},i=[],c=(r("5f3e"),r("d4fd")),u=r("61f7"),d={components:{DialogPanel:c["a"]},props:{roles:{default:function(){return[]},type:Array}},data:function(){var e=function(e,t,r){if(""===t)return r();Object(u["f"])(t)?r():r(new Error("密码格式不符合要求"))},t=function(e,t,r){Object(u["d"])(t)?r():r(new Error("邮箱格式不正确"))};return{visible:!1,loading:!1,remoteLoading:!1,userForm:{role_id:"",username:"",email:"",mobile:"",pass:""},userRules:{username:{required:!0,message:"请填写用户名称"},email:[{required:!0,message:"请填写邮箱"},{validator:t}],role_id:{required:!0,message:"请选择角色"},pass:[{required:!0,message:"请填写登陆密码"},{validator:e}]}}},methods:{cancel:function(){this.$refs.userForm.resetFields(),this.visible=!1},add:function(){var e=this;this.$refs.userForm.validate((function(t){if(!t)return!1;e.loading=!0,Object(s["a"])(e.userForm).then((function(t){e.$message.success("创建成功"),e.$emit("success"),e.loading=!1,e.cancel()})).catch((function(t){e.loading=!1,console.log(t)}))}))}}},m=d,p=r("cba8"),f=Object(p["a"])(m,o,i,!1,null,null,null),h=f.exports,b=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("dialog-panel",{attrs:{title:"用户修改","confirm-text":"保存",visible:e.visible,"confirm-loading":e.loading,width:"388px"},on:{cancel:e.cancel,confirm:e.save}},[r("el-form",{ref:"userForm",attrs:{model:e.userForm,"label-width":"90px",size:"small",rules:e.userRules}},[r("el-form-item",{attrs:{label:"用户名称",prop:"username"}},[r("el-input",{attrs:{placeholder:"请填写用户名"},model:{value:e.userForm.username,callback:function(t){e.$set(e.userForm,"username",t)},expression:"userForm.username"}})],1),r("el-form-item",{attrs:{label:"邮箱地址",prop:"email"}},[r("el-input",{attrs:{placeholder:"请填写邮箱，登陆使用"},model:{value:e.userForm.email,callback:function(t){e.$set(e.userForm,"email",t)},expression:"userForm.email"}})],1),r("el-form-item",{attrs:{label:"手机号",prop:"mobile"}},[r("el-input",{attrs:{placeholder:"请填写手机号码"},model:{value:e.userForm.mobile,callback:function(t){e.$set(e.userForm,"mobile",t)},expression:"userForm.mobile"}})],1),r("el-form-item",{attrs:{label:"角色",prop:"role_id"}},[r("el-select",{staticStyle:{width:"100%"},attrs:{placeholder:"请选择",clearable:""},model:{value:e.userForm.role_id,callback:function(t){e.$set(e.userForm,"role_id",t)},expression:"userForm.role_id"}},e._l(e.roles,(function(e){return r("el-option",{key:e.id,attrs:{label:e.role_name,value:e.id}})})),1)],1),r("el-form-item",{attrs:{label:"状态",prop:"state"}},[r("el-switch",{attrs:{"active-value":1,"inactive-value":0},model:{value:e.userForm.state,callback:function(t){e.$set(e.userForm,"state",t)},expression:"userForm.state"}})],1),r("el-form-item",{attrs:{label:"登录密码",prop:"pass"}},[r("el-input",{attrs:{placeholder:"字母开头，数字特殊字符 [@.&!#?,%$] 的 6 - 18 位"},model:{value:e.userForm.pass,callback:function(t){e.$set(e.userForm,"pass",t)},expression:"userForm.pass"}})],1)],1)],1)},g=[],v={components:{DialogPanel:c["a"]},props:{roles:{default:function(){return[]},type:Array}},data:function(){var e=function(e,t,r){""===t||Object(u["f"])(t)?r():r(new Error("密码格式不符合要求"))},t=function(e,t,r){Object(u["d"])(t)?r():r(new Error("邮箱格式不正确"))};return{visible:!1,loading:!1,remoteLoading:!1,userForm:{id:0,role_id:0,username:"",email:"",mobile:"",state:1,pass:""},userRules:{username:{required:!0,message:"请填写用户名称"},email:[{required:!0,message:"请填写邮箱"},{validator:t}],role_id:{required:!0,message:"请选择角色"},pass:{validator:e}}}},methods:{initUpdate:function(e){var t=this;Object(s["c"])(e).then((function(e){t.userForm=e.data,t.visible=!0})).catch((function(){}))},cancel:function(){this.$refs.userForm.resetFields(),this.visible=!1},save:function(){var e=this;this.$refs.userForm.validate((function(t){if(!t)return!1;e.loading=!0,Object(s["k"])(e.userForm).then((function(t){e.$message.success("修改成功"),e.$emit("success"),e.loading=!1,e.cancel()})).catch((function(t){e.loading=!1,console.log(t)}))}))}}},_=v,F=Object(p["a"])(_,b,g,!1,null,null,null),x=F.exports,w=r("fc23"),y=r("ed08"),$={components:{UserCreate:h,UserUpdate:x,Page:w["a"]},data:function(){return{loadings:{pageLoading:!1},userList:{total:0,list:[]},roles:[],search:{email:"",username:"",state:1,page:1,page_size:10}}},computed:{},filters:{userRolesFilter:function(e,t){if(null===t||!t||0===t.length)return"";for(var r=0;r<t.length;r++)if(e===t[r].id)return t[r].role_name;return""},timeFormat:function(e){return Object(y["c"])(e)}},mounted:function(){this.getUserList()},methods:{getUserList:function(){var e=this;this.loadings.pageLoading=!0,Promise.all([this.userRoles()]).then((function(){Object(s["e"])(e.search).then((function(t){e.userList=t.data,e.loadings.pageLoading=!1})).catch((function(){e.loadings.pageLoading=!1}))})).catch((function(){e.loadings.pageLoading=!1}))},userRoles:function(){var e=this;if(this.roles.length>0)return!0;Object(n["h"])({state:1,role_name:""}).then((function(t){e.roles=t.data}))},add:function(){this.$refs.userCreate.visible=!0},editRow:function(e){this.$refs.userUpdate.initUpdate(e.id)},doSearch:function(){this.search.page=1,this.getUserList()},handlePage:function(e){this.search.page=e,this.getUserList()},handlePageSize:function(e){this.search.page_size=e,this.getUserList()}}},k=$,j=Object(p["a"])(k,a,l,!1,null,null,null);t["default"]=j.exports},fc23:function(e,t,r){"use strict";var a=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("el-pagination",{attrs:{background:"","current-page":e.page,"page-size":e.limit,total:e.total,"hide-on-single-page":"",layout:"prev, pager, next, jumper, total, sizes","prev-text":"上页","next-text":"下页","page-sizes":[10,15,20,30,40,50,100]},on:{"current-change":e.handlePage,"size-change":e.handleSizeChange}})},l=[],s=(r("4582"),{props:{page:{type:Number,default:1},total:{type:Number,default:0},limit:{type:Number,default:10}},methods:{handlePage:function(e){this.$emit("current-change",e)},handleSizeChange:function(e){this.$emit("size-change",e)}}}),n=s,o=r("cba8"),i=Object(o["a"])(n,a,l,!1,null,null,null);t["a"]=i.exports}}]);