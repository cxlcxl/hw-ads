import request from "@/utils/request";

export function roleList(params) {
  return request({
    url: "/role/list",
    method: "get",
    params,
  });
}

export function rolePermissions(id) {
  return request({
    url: "/role/permissions",
    method: "get",
    params: { id },
  });
}

export function roleCreate(data) {
  return request({
    url: "/role/create",
    method: "post",
    data,
  });
}

export function roleUpdate(data) {
  return request({
    url: "/role/update",
    method: "post",
    data,
  });
}

export function roleDestroy(id) {
  return request({
    url: "/role/destroy",
    method: "post",
    data: {
      id,
    },
  });
}

export function roleInfo(id) {
  return request({
    url: "/role/" + id,
    method: "get",
  });
}

// 权限部分
export function permissionList() {
  return request({
    url: "/permission/tree",
    method: "get",
  });
}

export function permissionCreate(data) {
  return request({
    url: "/permission/create",
    method: "post",
    data,
  });
}

export function permissionDestroy(id) {
  return request({
    url: "/permission/destroy",
    method: "post",
    data: { id },
  });
}

export function permissionUpdate(data) {
  return request({
    url: "/permission/update",
    method: "post",
    data,
  });
}
