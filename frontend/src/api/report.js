import request from "@/utils/request";

export function reportComprehensive(data) {
  return request({
    url: "/report/comprehensive",
    method: "post",
    data,
  });
}

export function downloadComprehensive(data) {
  return request({
    url: "/report/comprehensive-download",
    method: "post",
    data,
  });
}

export function reportAds(data) {
  return request({
    url: "/report/ads",
    method: "post",
    data,
  });
}

export function setReportColumns(data) {
  return request({
    url: "/report/column",
    method: "post",
    data,
  });
}
