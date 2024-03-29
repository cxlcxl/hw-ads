<template>
  <el-row class="comprehensive">
    <el-form ref="_search" :model="search" inline size="small">
      <el-col :span="24" class="search-container">
        <el-form-item label="粒度">
          <el-select v-model="search.granularity" placeholder="数据粒度" class="w100">
            <el-option :key="item.key" :label="item.name" :value="item.key" v-for="item in Vars.ReportGranularity" />
          </el-select>
        </el-form-item>
        <el-form-item label="日期">
          <el-date-picker v-model="search.date_range" :picker-options="pickerOptions" class="w240" :clearable="false" value-format="yyyy-MM-dd"
            type="daterange" start-placeholder="开始日期" end-placeholder="结束日期" />
        </el-form-item>
        <el-form-item label="维度">
          <el-select v-model="search.dimensions" placeholder="数据维度" class="w220" multiple collapse-tags>
            <el-option :key="k" :label="v" :value="k" v-for="(v,k) in reportDimensions" />
          </el-select>
        </el-form-item>
        <el-form-item label="">
          <el-button type="primary" icon="el-icon-search" class="item" @click="doSearch">查询</el-button>
          <el-button icon="el-icon-download" class="item" @click="download" v-permission="'report/comprehensive-download'">下载数据</el-button>
        </el-form-item>
        <el-form-item label="" style="float: right;">
          <el-button type="danger" icon="el-icon-s-tools" class="item" @click="selectColumns" circle />
        </el-form-item>
      </el-col>
      <el-col :span="24" class="search-container">
        <el-form-item v-if="search.dimensions.includes('account_id')" label="账户">
          <el-select v-model="search.account_ids" placeholder="账户选择" class="w260" multiple collapse-tags clearable filterable @change="conditionApp">
            <el-option :key="item.id" :label="item.account_name" :value="item.id" v-for="item in accounts"
              v-show="item.account_type === Vars.AccountTypeMarket" />
          </el-select>
        </el-form-item>
        <el-form-item v-if="search.dimensions.includes('app_id')" label="应用">
          <el-select v-model="search.app_ids" placeholder="应用选择" class="w260" multiple collapse-tags clearable filterable>
            <el-option :key="item.app_id" :label="item.app_name" :value="item.app_id" v-for="item in apps"
              v-show="item.show||search.account_ids.length===0" />
          </el-select>
        </el-form-item>
        <el-form-item v-if="search.dimensions.includes('area_id')" label="地区">
          <el-select v-model="search.areas" placeholder="地区选择" class="w260" multiple collapse-tags clearable filterable>
            <el-option :key="item.id" :label="item.c_name" :value="item.id" v-for="item in regions" />
          </el-select>
        </el-form-item>
        <el-form-item v-if="search.dimensions.includes('country')" label="国家">
          <el-cascader :options="regions" :props="{ multiple: true, value:'c_code',label:'c_name' }" collapse-tags clearable class="w300"
            v-model="search.countries" />
        </el-form-item>
      </el-col>
    </el-form>
    <el-col :span="24">
      <el-table v-loading="loadings.pageLoading" :data="reportList.list" @sort-change="sortable" highlight-current-row stripe border size="mini"
        show-summary :summary-method="getSummaries">
        <el-table-column :label="item.label" :align="item.align" :fixed="item.fix" v-for="item in reportList.columns" :key="item.key" v-if="item.show"
          :min-width="item.min" :show-overflow-tooltip="item.fix" :sortable="item.sort|filterSort" :prop="item.key">
          <template slot-scope="scope">
            {{ item.prefix + scope.row[item.key] + item.suffix }}
          </template>
        </el-table-column>
      </el-table>
    </el-col>
    <el-col :span="24" style="margin-top: 15px;">
      <page ref="page" :page="search.page" :total="reportList.total" @current-change="handlePage" @size-change="handlePageSize"
        :limit="search.page_size" />
    </el-col>

    <select-columns ref="column" :Columns="reportList.columns" @confirm="confirm" module-name="comprehensive" />
  </el-row>
</template>

<script>
import Page from "@c/Page"
import { parseTime } from "@/utils"
import { reportComprehensive } from "@/api/report"
import { regions } from "@/api/common"
import { allAccounts } from "@/api/account"
import { allApp, appActRel } from "@/api/app"
import SelectColumns from "./components/columns"
import Vars from "@/vars.js"
const nowDate = new Date()
const sorts = { custom: "custom", 1: true, 0: false }

export default {
  name: "Comprehensive", // 写上 name 可以使 keep-alive 生效
  components: { Page, SelectColumns },
  filters: {
    timeFormat(t) {
      return parseTime(t)
    },
  },
  data() {
    return {
      Vars,
      loadings: {
        pageLoading: false,
      },
      reportDimensions: {},
      search: {
        date_range: [],
        dimensions: [],
        account_ids: [],
        app_ids: [],
        areas: [],
        countries: [],
        show_columns: [],
        granularity: "date",
        order: "",
        by: "",
        download: 0,
        page: 1,
        page_size: 15,
      },
      accounts: [],
      apps: [],
      appRels: {},
      regions: [],
      reportList: {
        list: [],
        total: 0,
        columns: [],
        summaries: {},
      },
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
            text: "本月",
            onClick(picker) {
              const start = new Date()
              picker.$emit("pick", [new Date(start.setDate(1)), nowDate])
            },
          },
          {
            text: "上月",
            onClick(picker) {
              const end = new Date(new Date().setDate(1)) // 本月第一天
              end.setTime(end.getTime() - 3600 * 1000 * 24 * 1) // -1 天 => 上月最后一天
              const s = new Date(end - 0)
              const start = new Date(s.setDate(1))
              picker.$emit("pick", [start, end])
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
        ],
      },
    }
  },
  filters: {
    filterSort(v) {
      return sorts[v]
    },
  },
  created() {
    this.setDefaultSearchDate()
  },
  mounted() {
    this.getReportList()
  },
  methods: {
    getReportList() {
      this.loadings.pageLoading = true
      Promise.all([this.getAllAccounts(), this.getAllApps(), this.getRegions(), this.getAppRels()])
        .then((res) => {
          reportComprehensive(this.search)
            .then((res) => {
              this.reportList.columns = res.data.columns
              this.reportList.list = res.data.list
              this.reportList.total = res.data.total
              this.reportList.summaries = res.data.summaries
              this.reportDimensions = res.data.dimensions
              this.loadings.pageLoading = false
            })
            .catch((err) => {
              console.log(err)
              this.loadings.pageLoading = false
            })
        })
        .catch((err) => {
          console.log(err)
          this.loadings.pageLoading = false
        })
    },
    getRegions() {
      return new Promise((resolve, reject) => {
        if (this.regions.length > 0) {
          return resolve()
        }
        regions()
          .then((res) => {
            this.regions = res.data
            resolve()
          })
          .catch((err) => {
            reject(err)
          })
      })
    },
    getAllApps() {
      return new Promise((resolve, reject) => {
        if (this.apps.length > 0) {
          return resolve()
        }
        allApp()
          .then((res) => {
            this.apps = res.data
            resolve()
          })
          .catch((err) => {
            reject(err)
          })
      })
    },
    getAppRels() {
      return new Promise((resolve, reject) => {
        if (Object.keys(this.appRels).length > 0) {
          return resolve()
        }
        appActRel()
          .then((res) => {
            this.appRels = res.data
            resolve()
          })
          .catch((err) => {
            reject(err)
          })
      })
    },
    getAllAccounts() {
      return new Promise((resolve, reject) => {
        if (this.accounts.length > 0) {
          resolve()
        } else {
          allAccounts()
            .then((res) => {
              this.accounts = res.data
              resolve()
            })
            .catch((err) => {
              reject(err)
            })
        }
      })
    },
    handlePage(v) {
      this.search.page = v
      this.getReportList()
    },
    handlePageSize(v) {
      this.search.page_size = v
      this.getReportList()
    },
    doSearch() {
      this.search.page = 1
      this.getReportList()
    },
    download() {
      this.loadings.pageLoading = true
      this.search.download = 1
      reportComprehensive(this.search)
        .then((res) => {
          this.search.download = 0
          this.loadings.downloadLoading = false
          import("@/vendor/Export2Excel")
            .then((excel) => {
              const tHeader = []
              res.data.columns.map((item) => {
                if (item.show) {
                  tHeader.push(item.label)
                }
              })
              const data = excel.formatJson(res.data)
              if (data.length === 0) {
                this.$message.info("没有筛选到需导出的数据")
                this.loadings.pageLoading = false
                return
              }
              excel.export_json_to_excel({ header: tHeader, data, filename: "综合报表数据" })
              this.loadings.pageLoading = false
            })
            .catch((err) => {
              console.log(err)
              this.loadings.pageLoading = false
            })
        })
        .catch(() => {
          this.search.download = 0
          this.loadings.pageLoading = false
        })
    },
    setDefaultSearchDate() {
      let s = new Date()
      let f = "{y}-{m}-{d}"
      this.$set(this.search, "date_range", [parseTime(s.getTime() - 3600 * 1000 * 24 * 7, f), parseTime(new Date(), f)])
    },
    selectColumns() {
      this.$refs.column.setDefault()
    },
    confirm(selected) {
      this.search.show_columns = selected
      this.getReportList()
    },
    sortable({ column, prop, order }) {
      if (column.sortable === "custom") {
        this.search.order = prop
        this.search.by = order
        this.getReportList()
      }
    },
    getSummaries(param) {
      const { columns } = param
      const sums = []
      columns.forEach((column, index) => {
        if (index === 0) {
          sums[index] = "合计"
        } else {
          switch (column.property) {
            case "cost":
              sums[index] = this.reportList.summaries.cost
              break
            case "earnings":
              sums[index] = this.reportList.summaries.earnings
              break
            case "roi":
              sums[index] = this.reportList.summaries.roi + "%"
              break
            case "ad_show_count":
              sums[index] = this.reportList.summaries.ad_show_count
              break
            case "ad_click_count":
              sums[index] = this.reportList.summaries.ad_click_count
              break
          }
        }
      })

      return sums
    },
    conditionApp(ids) {
      if (ids.length === 0) {
        return true
      }
      let appIds = []
      for (let i = 0; i < ids.length; i++) {
        appIds = appIds.concat(this.appRels[ids[i]])
      }
      for (let i = 0; i < this.apps.length; i++) {
        if (appIds.includes(this.apps[i].app_id)) {
          this.apps[i].show = true
        } else {
          this.apps[i].show = false
        }
      }
    },
  },
}
</script>

<style lang="less" scoped>
</style>