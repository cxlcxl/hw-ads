import request from "@/utils/request";

export function appList(params) {
  return request({
    url: "/app/list",
    method: "get",
    params,
  });
}

export function appCreate(data) {
  return request({
    url: "/app/create",
    method: "post",
    data,
  });
}

export function appUpdate(data) {
  return request({
    url: "/app/update",
    method: "post",
    data,
  });
}

export function appInfo(id) {
  return request({
    url: "/app/" + id,
    method: "get",
  });
}

export function allApp() {
  return request({
    url: "/app/all",
    method: "get",
  });
}

export function campaignAppList(params) {
  return request({
    url: "/app/campaign-list",
    method: "get",
    params,
  });
}

export function appPull(data) {
  return request({
    url: "/app/pull",
    method: "post",
    data,
  });
}
