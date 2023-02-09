<template>
  <el-row class="comprehensive">
    <el-form ref="_search" :model="search" inline size="small">
      <el-col :span="24" class="search-container">
        <el-form-item label="日期">
          <el-date-picker v-model="search.date_range" :picker-options="pickerOptions" class="w240" :clearable="false" value-format="yyyy-MM-dd"
            type="daterange" start-placeholder="开始日期" end-placeholder="结束日期" />
        </el-form-item>
        <el-form-item label="维度">
          <el-select v-model="search.dimensions" placeholder="数据维度" class="w220" multiple collapse-tags>
            <el-option :key="k" :label="v" :value="k" v-for="(v,k) in requestDimensions" />
          </el-select>
        </el-form-item>
        <el-form-item label="">
          <el-button type="primary" icon="el-icon-search" class="item" @click="doSearch">查询</el-button>
        </el-form-item>
        <el-form-item label="" style="float: right;">
          <el-button type="danger" icon="el-icon-s-tools" class="item" @click="selectColumns" circle />
        </el-form-item>
      </el-col>
      <el-col :span="24" class="search-container">
        <el-form-item v-if="search.dimensions.includes('account_id')" label="账户">
          <el-select v-model="search.account_ids" placeholder="账户选择" class="w260" multiple collapse-tags clearable filterable>
            <el-option :key="item.id" :label="item.account_name" :value="item.id" v-for="item in accounts" v-show="item.account_type === 2" />
          </el-select>
        </el-form-item>
        <el-form-item v-if="search.dimensions.includes('app_id')" label="应用">
          <el-select v-model="search.app_ids" placeholder="应用选择" class="w260" multiple collapse-tags clearable filterable>
            <el-option :key="item.app_id" :label="item.app_name" :value="item.app_id" v-for="item in apps" />
          </el-select>
        </el-form-item>
        <el-form-item v-if="search.dimensions.includes('country')" label="区域">
          <el-cascader :options="regions" :props="{ multiple: true, value:'c_code',label:'c_name' }" collapse-tags clearable class="w300"
            v-model="search.countries" />
        </el-form-item>
      </el-col>
    </el-form>
    <el-col :span="24">
      <el-table v-loading="loadings.pageLoading" :data="reportList.list" highlight-current-row stripe border size="mini" show-summary>
        <el-table-column prop="stat_day" label="日期" width="90" align="center" fixed="left" />
        <el-table-column :label="item.label" :align="item.align" :fixed="item.fix" v-for="item in reportList.columns" :key="item.key" v-if="item.show"
          :min-width="item.min" :show-overflow-tooltip="item.fix" :prop="item.key">
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

    <select-columns ref="column" :Columns="reportList.columns" @confirm="confirm" module-name="ads" />
  </el-row>
</template>

<script>
import Page from "@c/Page"
import { parseTime } from "@/utils"
import { requestDimensions } from "./data"
import { reportAds } from "@/api/report"
import { regions } from "@/api/common"
import { allAccounts } from "@/api/account"
import { allApp } from "@/api/app"
import SelectColumns from "./components/columns"
const nowDate = new Date()

export default {
  // name: "Ads", // 写上 name 可以使 keep-alive 生效
  components: { Page, SelectColumns },
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
        account_ids: [],
        app_ids: [],
        countries: [],
        show_columns: [],
        page: 1,
        page_size: 15,
      },
      accounts: [],
      apps: [],
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
  created() {
    this.setDefaultSearchDate()
  },
  mounted() {
    this.getReportList()
  },
  methods: {
    getReportList() {
      this.loadings.pageLoading = true
      Promise.all([this.getAllAccounts(), this.getAllApps(), this.getRegions()])
        .then((res) => {
          reportAds(this.search)
            .then((res) => {
              this.reportList.columns = res.data.columns
              this.reportList.list = res.data.list
              this.reportList.total = res.data.total
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
    getSummaries(param) {
      const { columns } = param
      const sums = []
      columns.forEach((column, index) => {
        if (index === 0) {
          sums[index] = "合计"
        } else {
          switch (column.property) {
            case "earnings":
              sums[index] = this.reportList.summaries.earnings
              break
          }
        }
      })

      return sums
    },
  },
}
</script>

<style lang="less" scoped>
</style>