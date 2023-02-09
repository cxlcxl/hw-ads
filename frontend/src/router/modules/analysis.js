import Layout from "@/layout";

const analysisRouter = {
  path: "/report",
  component: Layout,
  redirect: "/report/comprehensive",
  meta: { title: "报表统计", icon: "el-icon-data-line" },
  children: [
    {
      path: "comprehensive",
      name: "Comprehensive",
      component: () => import("@v/analysis/comprehensive"),
      meta: { title: "综合报表", auth: "report/comprehensive" },
    },
    {
      path: "ads",
      name: "AnalysisAds",
      component: () => import("@v/analysis/ads"),
      meta: { title: "变现报表", auth: "report/ads" },
    },
  ],
};

export default analysisRouter;
