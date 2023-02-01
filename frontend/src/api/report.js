import request from "@/utils/request";

export function reportComprehensive(data) {
  return request({
    url: "/report/comprehensive",
    method: "post",
    data,
  });
}

export function setReportComprehensiveColumns(data) {
  return request({
    url: "/report/comprehensive/column",
    method: "post",
    data,
  });
}
