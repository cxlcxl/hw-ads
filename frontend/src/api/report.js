import request from "@/utils/request";

export function reportComprehensive(data) {
  return request({
    url: "/report/comprehensive",
    method: "post",
    data,
  });
}
