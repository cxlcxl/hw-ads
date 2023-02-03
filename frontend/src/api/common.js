import request from "@/utils/request";

export function regions() {
  return request({
    url: "/regions",
    method: "get",
  });
}

export function overseasAreas() {
  return request({
    url: "/region/area",
    method: "get",
  });
}

export function overseasCountries(params) {
  return request({
    url: "/region/country",
    method: "get",
    params,
  });
}

export function regionCreate(data) {
  return request({
    url: "/region",
    method: "post",
    data,
  });
}

export function regionAreaSet(data) {
  return request({
    url: "/region/area-set",
    method: "post",
    data,
  });
}

export function settingsCron() {
  return request({
    url: "/settings/cron",
    method: "get",
  });
}

export function settingsCronInfo(id) {
  return request({
    url: "/settings/cron/" + id,
    method: "get",
  });
}

export function settingsCronUpdate(data, id) {
  return request({
    url: "/settings/cron/" + id,
    method: "post",
    data,
  });
}

export function settingsCronSchedule(data) {
  return request({
    url: "/settings/cron/schedule",
    method: "post",
    data,
  });
}
