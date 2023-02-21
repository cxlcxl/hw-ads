<template>
  <el-row>
    <el-col :span="24" class="search-container">
      <el-form ref="_search" :model="search" inline size="small">
        <el-form-item>
          <el-select v-model="search.account_ids" placeholder="投放账户选择" class="w260" multiple collapse-tags clearable filterable>
            <el-option :key="item.id" :label="item.account_name" :value="item.id" v-for="item in accounts" v-show="item.account_type === 1" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-input v-model="search.app_name" class="w150" clearable placeholder="应用名" />
        </el-form-item>
        <el-form-item>
          <el-input v-model="search.app_id" class="w120" clearable placeholder="APP ID" />
        </el-form-item>
        <!-- <el-form-item>
          <el-select v-model="search.channel" class="w100">
            <el-option label="全部渠道" :value="0" />
            <el-option v-for="(key, val) in appList.app_channel" :label="key" :value="Number(val)" :key="val" />
          </el-select>
        </el-form-item> -->
        <el-form-item label="">
          <el-button type="primary" icon="el-icon-search" class="item" @click="doSearch">查询</el-button>
        </el-form-item>
      </el-form>
    </el-col>
    <!-- <el-col :span="24">
      <el-button type="primary" icon="el-icon-plus" size="mini" @click="add">添加应用</el-button>
      <el-button type="primary" icon="el-icon-download" size="mini" @click="pullApp">拉取应用</el-button>
    </el-col> -->
    <el-col :span="24">
      <el-table v-loading="loading" :data="appList.data" highlight-current-row stripe border size="mini">
        <el-table-column prop="id" label="#" width="60" align="center">
          <template slot-scope="scope">
            <el-avatar shape="square" :size="30" :src="scope.row.icon_url"/>
          </template>
        </el-table-column>
        <el-table-column prop="app_name" label="应用名称" width="200" show-overflow-tooltip/>
        <el-table-column prop="app_id" label="AppID" width="100" align="center" />
        <!-- <el-table-column label="渠道" width="120" align="center">
          <template slot-scope="scope">{{appList.app_channel[scope.row.channel]}}</template>
        </el-table-column> -->
        <el-table-column label="关联投放账户" prop="account_ids" min-width="150" show-overflow-tooltip>
          <template slot-scope="scope">
            {{scope.row.account_ids|accountsFilter(accounts)}}
          </template>
        </el-table-column>
        <el-table-column label="应用包名" prop="pkg_name" min-width="150" show-overflow-tooltip />
        <el-table-column prop="created_at" label="添加时间" width="140" align="center">
          <template slot-scope="scope">{{scope.row.created_at|timeFormat}}</template>
        </el-table-column>
        <el-table-column align="center" label="操作" width="90">
          <template slot-scope="scope">
            <el-button-group class="table-operate" v-permission="'app/update'">
              <el-button type="primary" plain @click.native.prevent="editRow(scope.row.id)">编辑</el-button>
            </el-button-group>
          </template>
        </el-table-column>
      </el-table>
    </el-col>
    <el-col :span="24" style="margin-top: 15px;">
      <page ref="page" :page="search.page" :total="appList.total" @current-change="handlePage" @size-change="handlePageSize" />
    </el-col>

    <create-application ref="appCreate" :app-channel="appList.app_channel" @success="getApplicationList" />
    <update-application ref="appUpdate" :app-channel="appList.app_channel" @success="getApplicationList" />
    <pull-application ref="appPull" @success="getApplicationList" />
  </el-row>
</template>

<script>
import { appList } from "@a/app"
import { parseTime } from "@/utils"
import Page from "@c/Page"
import { allAccounts } from "@/api/account"
import CreateApplication from "./components/add-app"
import UpdateApplication from "./components/edit-app"
import PullApplication from "./components/pull-app"

export default {
  name: "App",
  components: { Page, CreateApplication, UpdateApplication, PullApplication },
  data() {
    return {
      loading: false,
      search: {
        app_name: "",
        app_id: "",
        channel: 0,
        account_ids: [],
        page: 1,
        page_size: 10,
      },
      accounts: [],
      appList: {
        total: 0,
        app_channel: {},
        data: [],
      },
    }
  },
  mounted() {
    this.getApplicationList()
  },
  filters: {
    timeFormat(timestamp) {
      return parseTime(timestamp)
    },
    accountsFilter(ids, acts) {
      let names = ""
      if (Array.isArray(ids) && ids.length > 0) {
        acts.forEach((item) => {
          if (ids.includes(item.id)) {
            names = names + item.account_name + "、"
          }
        })
      }
      return names
    },
  },
  methods: {
    add() {
      this.$refs.appCreate.initCreate()
    },
    pullApp() {
      this.$refs.appPull.initPull()
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
    doSearch() {
      this.search.page = 1
      this.getApplicationList()
    },
    getApplicationList() {
      this.loading = true
      Promise.all([this.getAllAccounts()])
        .then((res) => {
          appList(this.search)
            .then((res) => {
              const { list, total, app_channel } = res.data
              this.appList.data = list
              this.appList.total = total
              this.appList.app_channel = app_channel
              this.loading = false
            })
            .catch((err) => {
              this.loading = false
            })
        })
        .catch(() => {
          this.loading = false
        })
    },
    handlePage(p) {
      this.search.page = p
      this.getApplicationList()
    },
    handlePageSize(p) {
      this.search.page_size = p
      this.getApplicationList()
    },
    editRow(id) {
      this.$refs.appUpdate.initUpdate(id)
    },
  },
}
</script>

<style scoped>
</style>
