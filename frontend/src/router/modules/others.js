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
      name: "AccountList",
      component: () => import("@v/account/list"),
      meta: { title: "账户列表" },
    },
    {
      path: "app",
      name: "AppList",
      component: () => import("@v/application/list"),
      meta: { title: "应用列表" },
    },
    {
      path: "region",
      name: "RegionList",
      component: () => import("@v/others/region"),
      meta: { title: "区域管理" },
    },
  ],
};

export default userRouter;
