import request from "@/utils/request";

export function campaignList(params) {
  return request({
    url: "/marketing/campaign/list",
    method: "get",
    params,
  });
}

export function campaignCreate(data) {
  return request({
    url: "/marketing/campaign/create",
    method: "post",
    data,
  });
}

export function campaignResources() {
  return request({
    url: "/marketing/campaign/resources",
    method: "get",
  });
}

export function campaignInfo(campaign_id) {
  return request({
    url: "/marketing/campaign/" + campaign_id,
    method: "get",
  });
}

export function campaignBindApp(data) {
  return request({
    url: "/marketing/campaign/bind-app",
    method: "post",
    data,
  });
}
