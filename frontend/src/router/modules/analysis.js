import Layout from "@/layout";

const analysisRouter = {
  path: "/analysis",
  component: Layout,
  redirect: "/analysis/list",
  meta: { title: "统计报表", icon: "el-icon-data-line" },
  children: [
    {
      path: "comprehensive",
      name: "AnalysisList",
      component: () => import("@v/analysis/comprehensive"),
      meta: { title: "综合报表" },
    },
    {
      path: "ads",
      name: "AnalysisAds",
      component: () => import("@v/analysis/ads"),
      meta: { title: "变现报表" },
    },
  ],
};

export default analysisRouter;
