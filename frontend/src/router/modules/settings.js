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
      meta: { title: "版本更新" },
    },
    {
      path: "tool-log",
      name: "ToolLog",
      component: () => import("@v/settings/tool-log"),
      meta: { title: "日志下载" },
    },
  ],
};

export default settingsRouter;
