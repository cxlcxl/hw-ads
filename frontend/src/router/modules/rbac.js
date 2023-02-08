/** When your routing table is too long, you can split it into small modules **/

import Layout from "@/layout";

const userRouter = {
  path: "/rbac",
  component: Layout,
  redirect: "/rbac/user",
  meta: { title: "用户权限", icon: "el-icon-unlock" },
  children: [
    {
      path: "user",
      name: "UserList",
      component: () => import("@v/rbac/user"),
      meta: { title: "用户列表", auth: "user/list" },
    },
    {
      path: "role",
      name: "RoleList",
      component: () => import("@v/rbac/role"),
      meta: { title: "角色列表", auth: "role/list" },
    },
    {
      path: "permission",
      name: "Permission",
      component: () => import("@v/rbac/permission"),
      meta: { title: "权限设置", auth: "permission/list" },
    },
  ],
};

export default userRouter;
