(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-39acfab5"],{"0906":function(e,t,o){"use strict";o.r(t);var r=function(){var e=this,t=e.$createElement,o=e._self._c||t;return o("el-row",[o("el-col",{staticClass:"search-container",attrs:{span:24}},[o("el-form",{ref:"_search",attrs:{model:e.search,inline:"",size:"small"}},[o("el-form-item",[o("el-input",{staticClass:"w150",attrs:{clearable:"",placeholder:"角色名称"},model:{value:e.search.role_name,callback:function(t){e.$set(e.search,"role_name",t)},expression:"search.role_name"}})],1),o("el-form-item",[o("el-select",{staticClass:"w110",attrs:{placeholder:"用户状态"},model:{value:e.search.state,callback:function(t){e.$set(e.search,"state",t)},expression:"search.state"}},[o("el-option",{key:1,attrs:{label:"正常",value:1}}),o("el-option",{key:0,attrs:{label:"停用",value:0}})],1)],1),o("el-form-item",{attrs:{label:""}},[o("el-button",{staticClass:"item",attrs:{type:"primary",icon:"el-icon-search"},on:{click:e.getRoleList}},[e._v("查询")])],1)],1)],1),o("el-col",{staticStyle:{"margin-bottom":"15px"},attrs:{span:24}},[o("el-button",{attrs:{icon:"el-icon-plus",size:"mini",type:"primary"},on:{click:e.addRole}},[e._v("新增角色")])],1),o("el-col",{attrs:{span:24}},[o("el-table",{directives:[{name:"loading",rawName:"v-loading",value:e.loadings.pageLoading,expression:"loadings.pageLoading"}],ref:"user",attrs:{data:e.roles.list,"highlight-current-row":"",stripe:"",border:"",size:"mini"}},[o("el-table-column",{attrs:{prop:"id",label:"ID",width:"110"}}),o("el-table-column",{attrs:{prop:"role_name",label:"角色名称"}}),o("el-table-column",{attrs:{align:"center",label:"操作",fixed:"right",width:"130"},scopedSlots:e._u([{key:"default",fn:function(t){return[o("el-button-group",{staticClass:"table-operate"},[o("el-button",{attrs:{type:"primary",plain:""},nativeOn:{click:function(o){return o.preventDefault(),e.editRow(t.row)}}},[e._v("编辑")])],1)]}}])})],1)],1),o("role-create",{ref:"roleCreate",attrs:{permissions:e.permissions},on:{success:e.getRoleList}}),o("role-update",{ref:"roleUpdate",attrs:{permissions:e.permissions},on:{success:e.getRoleList}})],1)},s=[],n=(o("2ce8"),o("4221"),o("cc5e")),i=function(){var e=this,t=e.$createElement,o=e._self._c||t;return o("dialog-panel",{attrs:{title:"角色创建","confirm-text":"添加",visible:e.visible,"confirm-loading":e.loading},on:{cancel:e.cancel,confirm:e.add}},[o("el-form",{ref:"roleForm",attrs:{model:e.roleForm,"label-width":"100px",size:"small"}},[o("el-form-item",{attrs:{label:"角色名称",prop:"role_name",rules:{required:!0,message:"请填写角色名称"}}},[o("el-input",{model:{value:e.roleForm.role_name,callback:function(t){e.$set(e.roleForm,"role_name",t)},expression:"roleForm.role_name"}})],1),o("el-form-item",{attrs:{label:"选择权限",prop:"permissions"}},[o("el-tree",{ref:"permission_tree",attrs:{data:e.permissions,"show-checkbox":"","node-key":"id",props:e.defaultProps},on:{"check-change":e.handleCheck}})],1)],1)],1)},a=[],l=(o("068b"),o("3bae"),o("d4fd")),c=o("ed08"),d={components:{DialogPanel:l["a"]},props:{permissions:Array,roleType:Object},data:function(){return{visible:!1,loading:!1,roleForm:{role_name:"",sys:0,permissions:[]},defaultProps:{children:"children",label:"p_name"}}},methods:{cancel:function(){this.$refs.roleForm.resetFields(),this.$set(this.roleForm,"permissions",[]),this.$refs.permission_tree.setCheckedKeys([]),this.visible=!1},add:function(){var e=this;this.$refs.roleForm.validate((function(t){if(!t)return!1;e.loading=!0,Object(n["e"])(e.roleForm).then((function(t){e.$message.success("创建成功"),e.$emit("success"),e.loading=!1,e.cancel()})).catch((function(t){e.loading=!1,console.log(t)}))}))},handleCheck:function(e,t,o){t&&!this.roleForm.permissions.includes(e.permission)?this.roleForm.permissions.push(e.permission):this.roleForm.permissions=Object(c["d"])(this.roleForm.permissions,e.permission)}}},u=d,m=o("cba8"),p=Object(m["a"])(u,i,a,!1,null,null,null),f=p.exports,h=function(){var e=this,t=e.$createElement,o=e._self._c||t;return o("dialog-panel",{attrs:{title:"角色修改","confirm-text":"保存",visible:e.visible,"confirm-loading":e.loading},on:{cancel:e.cancel,confirm:e.save}},[o("el-form",{ref:"roleForm",attrs:{model:e.roleForm,"label-width":"100px",size:"small"}},[o("el-form-item",{attrs:{label:"角色名称",prop:"role_name",rules:{required:!0,message:"请填写角色名称"}}},[o("el-input",{model:{value:e.roleForm.role_name,callback:function(t){e.$set(e.roleForm,"role_name",t)},expression:"roleForm.role_name"}})],1),o("el-form-item",{attrs:{label:"状态",prop:"state"}},[o("el-switch",{attrs:{"active-value":1,"inactive-value":0},model:{value:e.roleForm.state,callback:function(t){e.$set(e.roleForm,"state",t)},expression:"roleForm.state"}})],1),o("el-form-item",{attrs:{label:"选择权限",prop:"permissions",rules:{required:!0,message:"请选择权限"}}},[o("el-tree",{ref:"tree",attrs:{data:e.permissions,"show-checkbox":"","node-key":"permission",props:e.defaultProps,"default-checked-keys":e.roleForm.permissions},on:{"check-change":e.handleCheck}})],1)],1)],1)},g=[],b={components:{DialogPanel:l["a"]},props:{permissions:Array,roleType:Object},data:function(){return{visible:!1,loading:!1,roleForm:{id:0,role_name:"",state:1,sys:0,permissions:[]},defaultProps:{children:"children",label:"p_name"}}},methods:{initUpdate:function(e){var t=this;Object(n["g"])(e.id).then((function(e){t.roleForm.id=e.data.id,t.roleForm.role_name=e.data.role_name,t.roleForm.state=e.data.state,t.roleForm.sys=e.data.sys;var o=[];Array.isArray(e.data.permissions)&&(o=e.data.permissions),t.roleForm.permissions=o,t.visible=!0,t.$refs.tree.setCheckedNodes(o)})).catch((function(e){console.log(e)}))},cancel:function(){this.$refs.roleForm.resetFields(),this.visible=!1},save:function(){var e=this;this.$set(this.roleForm,"permissions",this.$refs.tree.getCheckedKeys()),this.$refs.roleForm.validate((function(t){if(!t)return!1;e.loading=!0,Object(n["i"])(e.roleForm).then((function(t){e.$message.success("修改成功"),e.$emit("success"),e.loading=!1,e.cancel()})).catch((function(t){e.loading=!1,console.log(t)}))}))},handleCheck:function(e,t,o){t&&!this.roleForm.permissions.includes(e.permission)?this.roleForm.permissions.push(e.permission):this.roleForm.permissions=Object(c["d"])(this.roleForm.permissions,e.permission)}}},v=b,F=Object(m["a"])(v,h,g,!1,null,null,null),_=F.exports,y={name:"RoleList",components:{RoleCreate:f,RoleUpdate:_},data:function(){return{loadings:{changeLoading:!1,pageLoading:!1},roles:{list:[]},search:{role_name:"",state:1},permissions:[]}},mounted:function(){this.getRoleList()},methods:{getRoleList:function(){var e=this;this.loadings.pageLoading=!0,Object(n["h"])(this.search).then((function(t){e.roles.list=t.data,e.loadings.pageLoading=!1})).catch((function(){e.loadings.pageLoading=!1}))},editRow:function(e){0===this.permissions.length&&this.getPermissions(),this.$refs.roleUpdate.initUpdate(e)},destroyRow:function(e){var t=this;this.$confirm("此操作会同步删除已分配的权限, 是否继续?","警告",{confirmButtonText:"确定",cancelButtonText:"取消",type:"error"}).then((function(){t.loadings.pageLoading=!0,Object(n["f"])(e.id).then((function(){t.$message.success("删除成功"),t.getRoleList()})).catch((function(e){t.loadings.pageLoading=!1}))})).catch((function(){}))},getPermissions:function(){var e=this;Object(n["c"])().then((function(t){e.permissions=t.data})).catch((function(e){}))},addRole:function(){0===this.permissions.length&&this.getPermissions(),this.$refs.roleCreate.visible=!0}}},k=y,x=Object(m["a"])(k,r,s,!1,null,null,null);t["default"]=x.exports},4221:function(e,t,o){"use strict";var r=o("4b8d"),s=o("4f7e"),n=o("c7b3"),i=o("3978"),a=o("b821"),l=o("3e87"),c=o("e001"),d=o("837a");s("search",(function(e,t,o){return[function(t){var o=i(this),s=void 0==t?void 0:c(t,e);return s?r(s,t,o):new RegExp(t)[e](l(o))},function(e){var r=n(this),s=l(e),i=o(t,r,s);if(i.done)return i.value;var c=r.lastIndex;a(c,0)||(r.lastIndex=0);var u=d(r,s);return a(r.lastIndex,c)||(r.lastIndex=c),null===u?-1:u.index}]}))},"80f8":function(e,t,o){},b821:function(e,t){e.exports=Object.is||function(e,t){return e===t?0!==e||1/e===1/t:e!=e&&t!=t}},b9c3:function(e,t,o){"use strict";o("80f8")},cc5e:function(e,t,o){"use strict";o.d(t,"h",(function(){return s})),o.d(t,"e",(function(){return n})),o.d(t,"i",(function(){return i})),o.d(t,"f",(function(){return a})),o.d(t,"g",(function(){return l})),o.d(t,"c",(function(){return c})),o.d(t,"a",(function(){return d})),o.d(t,"b",(function(){return u})),o.d(t,"d",(function(){return m}));var r=o("b775");function s(e){return Object(r["a"])({url:"/role/list",method:"get",params:e})}function n(e){return Object(r["a"])({url:"/role/create",method:"post",data:e})}function i(e){return Object(r["a"])({url:"/role/update",method:"post",data:e})}function a(e){return Object(r["a"])({url:"/role/destroy",method:"post",data:{id:e}})}function l(e){return Object(r["a"])({url:"/role/"+e,method:"get"})}function c(){return Object(r["a"])({url:"/permission/tree",method:"get"})}function d(e){return Object(r["a"])({url:"/permission/create",method:"post",data:e})}function u(e){return Object(r["a"])({url:"/permission/destroy",method:"post",data:{id:e}})}function m(e){return Object(r["a"])({url:"/permission/update",method:"post",data:e})}},d4fd:function(e,t,o){"use strict";var r=function(){var e=this,t=e.$createElement,o=e._self._c||t;return o("el-dialog",{directives:[{name:"el-drag-dialog",rawName:"v-el-drag-dialog"}],attrs:{title:e.title,visible:e.visible,width:e.width,"append-to-body":"","modal-append-to-body":"","close-on-click-modal":!1,"close-on-press-escape":!1,"show-close":!1},on:{"update:visible":function(t){e.visible=t}}},[e._t("default"),o("div",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[o("el-button",{attrs:{loading:e.confirmLoading,icon:"el-icon-close"},on:{click:e.handleCancel}},[e._v("取消")]),o("el-button",{directives:[{name:"show",rawName:"v-show",value:e.confirmText,expression:"confirmText"}],attrs:{type:"primary",icon:"el-icon-check",loading:e.confirmLoading},on:{click:e.handleConfirm}},[e._v(e._s(e.confirmText))]),e._t("operate")],2)],2)},s=[],n={props:{confirmText:{default:"",type:String},title:{default:"",type:String},width:{default:"600px",type:String},confirmLoading:{default:!1,type:Boolean},visible:{default:!1,type:Boolean}},methods:{handleCancel:function(){this.$emit("cancel")},handleConfirm:function(){this.$emit("confirm")}}},i=n,a=(o("b9c3"),o("cba8")),l=Object(a["a"])(i,r,s,!1,null,null,null);t["a"]=l.exports}}]);