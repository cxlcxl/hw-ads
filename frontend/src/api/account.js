import request from "@/utils/request";

export function accountUpdate(data) {
  return request({
    url: "/account/update",
    method: "post",
    data,
  });
}

export function accountCreate(data) {
  return request({
    url: "/account/create",
    method: "post",
    data,
  });
}

export function accountList(params) {
  return request({
    url: "/account/list",
    method: "get",
    params,
  });
}

export function accountInfo(account_id) {
  return request({
    url: "/account/" + account_id,
    method: "get",
  });
}

export function accountAuth(account_id) {
  return request({
    url: "/account/auth",
    method: "get",
    params: { id: account_id },
  });
}

export function refreshAuth(id) {
  return request({
    url: "/account/refresh/" + id,
    method: "post",
  });
}

export function searchAccounts(accountName) {
  return request({
    url: "/account/search",
    method: "get",
    params: { account_name: accountName },
  });
}

export function defaultAccounts() {
  return request({
    url: "/account/default",
    method: "get",
  });
}

export function allAccounts() {
  return request({
    url: "/account/all",
    method: "get",
  });
}

export function parentAccounts(params) {
  return request({
    url: "/account/parents",
    method: "get",
    params,
  });
}

export function getAccessToken(data) {
  return request({ url: "/account/token", method: "post", data });
}
