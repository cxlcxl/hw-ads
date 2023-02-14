/** When your routing table is too long, you can split it into small modules **/

import Layout from "@/layout";

const userRouter = {
  path: "/other",
  component: Layout,
  redirect: "/other/account",
  meta: { title: "其他功能", icon: "el-icon-s-grid" },
  children: [
    {
      path: "account",
      name: "Account",
      component: () => import("@v/account/list"),
      meta: { title: "账户列表", auth: "account/list" },
    },
    {
      path: "app",
      name: "App",
      component: () => import("@v/application/list"),
      meta: { title: "应用列表", auth: "app/list" },
    },
    {
      path: "region",
      name: "Region",
      component: () => import("@v/others/region"),
      meta: { title: "区域管理", auth: "region" },
    },
  ],
};

export default userRouter;
