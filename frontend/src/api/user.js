import request from "@/utils/request";

export function ssoLoginPath() {
  return request({
    url: "/sso-login",
    method: "get",
  });
}

export function ssoLogin(data) {
  return request({
    url: "/sso-login",
    method: "post",
    data,
  });
}

export function login(data) {
  return request({
    url: "/login",
    method: "post",
    data,
  });
}

export function getInfo() {
  return request({
    url: "/profile",
    method: "get",
  });
}

export function getUserInfo(id) {
  return request({
    url: "/user/" + id,
    method: "get",
  });
}

export function logout() {
  return request({
    url: "/logout",
    method: "post",
  });
}

export function create(data) {
  return request({
    url: "/user/create",
    method: "post",
    data,
  });
}

export function infoUpdate(data) {
  return request({
    url: "/self/update",
    method: "post",
    data,
  });
}

export function userUpdate(data) {
  return request({
    url: "/user/update",
    method: "post",
    data,
  });
}

export function userClose(id) {
  return request({
    url: "/user/close",
    method: "post",
    data: { id },
  });
}

export function list(params) {
  return request({
    url: "/user/list",
    method: "get",
    params,
  });
}

export function resetPass(data) {
  return request({
    url: "/reset-pass",
    method: "post",
    data,
  });
}
