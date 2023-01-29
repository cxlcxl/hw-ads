<template>
  <el-row class="comprehensive">
    <el-col :span="24" class="search-container">
      <el-form ref="_search" :model="search" inline size="small">
        <el-form-item>
          <el-date-picker v-model="search.date_range" :picker-options="pickerOptions" class="w220" :clearable="false" value-format="yyyyMMdd"
            type="daterange" start-placeholder="开始日期" end-placeholder="结束日期" />
        </el-form-item>
        <el-form-item>
          <el-select v-model="search.dimensions" placeholder="数据维度" class="w200" multiple collapse-tags>
            <el-option :key="k" :label="v" :value="k" v-for="(v,k) in requestDimensions" />
          </el-select>
        </el-form-item>
        <el-form-item label="">
          <el-button type="primary" icon="el-icon-search" class="item" @click="doSearch">查询</el-button>
        </el-form-item>
      </el-form>
    </el-col>
    <!-- <el-col :span="24">
      <el-button type="primary" icon="el-icon-plus" size="mini" @click="add">添加用户</el-button>
    </el-col> -->
    <el-col :span="24">
      <el-table v-loading="loadings.pageLoading" :data="adsList.list" highlight-current-row stripe border size="mini">
        <el-table-column prop="stat_day" label="日期" width="120" align="center" />
        <el-table-column prop="email" label="收入" align="right" />
        <el-table-column prop="username" label="eCPM" align="right" />
        <el-table-column prop="mobile" label="请求" align="right" />
        <el-table-column prop="mobile" label="填充" align="right" />
        <el-table-column prop="mobile" label="填充率" align="right" />
        <el-table-column prop="mobile" label="曝光" align="right" />
        <el-table-column prop="mobile" label="点击" align="right" />
        <el-table-column prop="mobile" label="点击率" align="right" />
        <el-table-column prop="mobile" label="展示率" align="right" />
        <el-table-column prop="mobile" label="ARPU" align="right" />
      </el-table>
    </el-col>
    <el-col :span="24" style="margin-top: 15px;">
      <page ref="page" :page="search.page" :total="adsList.total" @current-change="handlePage" @size-change="handlePageSize" />
    </el-col>
  </el-row>
</template>

<script>
import Page from "@c/Page"
import { parseTime } from "@/utils"
import { requestDimensions } from "./data"
const nowDate = new Date()

export default {
  components: { Page },
  filters: {
    timeFormat(t) {
      return parseTime(t)
    },
  },
  data() {
    return {
      requestDimensions,
      loadings: {
        pageLoading: false,
      },
      search: {
        date_range: [],
        dimensions: [],
      },
      adsList: {},
      pickerOptions: {
        shortcuts: [
          {
            text: "近 7 天",
            onClick(picker) {
              const start = new Date()
              start.setTime(start.getTime() - 3600 * 1000 * 24 * 7)
              picker.$emit("pick", [start, nowDate])
            },
          },
          {
            text: "近 30 天",
            onClick(picker) {
              const start = new Date()
              start.setTime(start.getTime() - 3600 * 1000 * 24 * 30)
              picker.$emit("pick", [start, nowDate])
            },
          },
          {
            text: "近 3 个月",
            onClick(picker) {
              const start = new Date()
              start.setTime(start.getTime() - 3600 * 1000 * 24 * 90)
              picker.$emit("pick", [start, nowDate])
            },
          },
        ],
      },
    }
  },
  created() {
    this.setDefaultSearchDate()
  },
  mounted() {
    this.getAdsList()
  },
  methods: {
    getAdsList() {},
    handlePage() {},
    handlePageSize() {},
    add() {},
    doSearch() {},
    setDefaultSearchDate() {
      let s = new Date()
      let f = "{y}{m}{d}"
      this.$set(this.search, "date_range", [parseTime(s.getTime() - 3600 * 1000 * 24 * 7, f), parseTime(new Date(), f)])
    },
  },
}
</script>

<style lang="less" scoped>
</style>