import Layout from "@/layout";

const settingsRouter = {
  path: "/settings",
  component: Layout,
  redirect: "/settings/list",
  meta: { title: "系统设置", icon: "el-icon-setting" },
  children: [
    {
      path: "cron",
      name: "Cron",
      component: () => import("@v/settings/cron"),
      meta: { title: "任务调度", auth: "settings/cron" },
    },
    {
      path: "config",
      name: "Config",
      component: () => import("@v/settings/config"),
      meta: { title: "系统配置", auth: "settings/config" },
    },
    {
      path: "version",
      name: "Version",
      component: () => import("@v/settings/version"),
      meta: { title: "版本更新", auth: "settings/version" },
    },
    {
      path: "log",
      name: "ToolLog",
      component: () => import("@v/settings/tool-log"),
      meta: { title: "日志下载", auth: "settings/log" },
    },
    {
      path: "profile",
      name: "Profile",
      component: () => import("@v/profile/index"),
      meta: { title: "个人信息" },
    },
  ],
};

export default settingsRouter;
