import request from "@/utils/request";

export function regions() {
  return request({
    url: "/regions",
    method: "get",
  });
}
