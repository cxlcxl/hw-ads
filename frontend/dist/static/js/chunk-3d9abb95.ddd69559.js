(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-3d9abb95"],{"1ba5":function(t,n,e){"use strict";e("46c8")},2934:function(t,n,e){"use strict";e.d(n,"h",(function(){return o})),e.d(n,"n",(function(){return i})),e.d(n,"d",(function(){return u})),e.d(n,"e",(function(){return c})),e.d(n,"g",(function(){return s})),e.d(n,"f",(function(){return a})),e.d(n,"i",(function(){return f})),e.d(n,"j",(function(){return d})),e.d(n,"l",(function(){return l})),e.d(n,"k",(function(){return g})),e.d(n,"m",(function(){return m})),e.d(n,"a",(function(){return b})),e.d(n,"b",(function(){return h})),e.d(n,"c",(function(){return p}));var r=e("b775");function o(){return Object(r["a"])({url:"/regions",method:"get"})}function i(){return Object(r["a"])({url:"/settings/version",method:"get"})}function u(){return Object(r["a"])({url:"/region/area",method:"get"})}function c(t){return Object(r["a"])({url:"/region/country",method:"get",params:t})}function s(t){return Object(r["a"])({url:"/region",method:"post",data:t})}function a(t){return Object(r["a"])({url:"/region/area-set",method:"post",data:t})}function f(){return Object(r["a"])({url:"/settings/cron",method:"get"})}function d(t){return Object(r["a"])({url:"/settings/cron/"+t,method:"get"})}function l(t,n){return Object(r["a"])({url:"/settings/cron/"+n,method:"post",data:t})}function g(t){return Object(r["a"])({url:"/settings/cron/schedule",method:"post",data:t})}function m(t){return Object(r["a"])({url:"/settings/configs",method:"get",params:t})}function b(t){return Object(r["a"])({url:"/settings/config",method:"post",data:t})}function h(t){return Object(r["a"])({url:"/settings/config/"+t,method:"get"})}function p(t,n){return Object(r["a"])({url:"/settings/config/"+n,method:"post",data:t})}},"36f8":function(t,n,e){"use strict";e.r(n);var r=function(){var t=this,n=t.$createElement,e=t._self._c||n;return e("el-row",[e("el-col",{staticStyle:{"margin-bottom":"20px"},attrs:{span:24}},[e("el-button",{attrs:{icon:"el-icon-refresh",size:"mini",plain:"",type:"primary"},on:{click:t.getVersions}})],1),e("el-col",{staticClass:"timeline-version",attrs:{span:24}},[e("el-timeline",t._l(t.versionList,(function(n){return e("el-timeline-item",{key:n.v,attrs:{timestamp:n.date,placement:"top",color:"#409eff"}},[e("el-card",[e("h3",[t._v("版本："+t._s(n.v))]),t._l(n.fixs,(function(n){return e("p",{staticClass:"v-text"},[t._v(t._s(n))])}))],2)],1)})),1)],1)],1)},o=[],i=e("2934"),u={data:function(){return{versionList:[{date:"2023-02-12",v:"v1.0.0",fixs:["更新了什么什么"]}]}},mounted:function(){this.getVersions()},methods:{getVersions:function(){var t=this;Object(i["n"])().then((function(n){t.versionList=n.data})).catch((function(){}))}}},c=u,s=(e("1ba5"),e("cba8")),a=Object(s["a"])(c,r,o,!1,null,null,null);n["default"]=a.exports},"46c8":function(t,n,e){}}]);