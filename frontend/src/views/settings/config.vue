<template>
  <el-row>
    <el-col :span="24" class="search-container">
      <el-form ref="_search" :model="search" inline size="small">
        <el-form-item>
          <el-input v-model="search._desc" class="w200" clearable placeholder="配置名" />
        </el-form-item>
        <el-form-item>
          <el-input v-model="search._k" class="w200" clearable placeholder="配置代码" />
        </el-form-item>
        <el-form-item>
          <el-select v-model="search.state" placeholder="状态" class="w110">
            <el-option :key="1" label="正常" :value="1" />
            <el-option :key="0" label="停用" :value="0" />
          </el-select>
        </el-form-item>
        <el-form-item label="">
          <el-button type="primary" icon="el-icon-search" class="item" @click="doSearch">查询</el-button>
        </el-form-item>
      </el-form>
    </el-col>
    <el-col :span="24">
      <el-button type="primary" icon="el-icon-plus" size="mini" @click="add">添加配置</el-button>
    </el-col>
    <el-col :span="24">
      <el-table v-loading="loadings.pageLoading" :data="confList.list" highlight-current-row stripe border size="mini" style="margin-top: 15px"
        @sort-change="sortChange">
        <el-table-column prop="desc" label="配置名" width="180" show-overflow-tooltip />
        <el-table-column prop="key" label="配置代码" width="240" sortable />
        <el-table-column prop="val" label="值" />
        <el-table-column prop="bak1" label="扩展字段1" />
        <el-table-column prop="bak2" label="扩展字段2" />
        <el-table-column prop="remark" label="备注" />
        <el-table-column label="状态" width="80" align="center">
          <template slot-scope="scope">
            {{ scope.row.state === 1 ? '正常' : '停用' }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="操作" fixed="right" width="100">
          <template slot-scope="scope">
            <el-button-group class="table-operate">
              <el-button type="primary" plain @click.native.prevent="editRow(scope.row)">编辑</el-button>
              <el-button type="primary" plain @click.native.prevent="copyRow(scope.row)">复制</el-button>
            </el-button-group>
          </template>
        </el-table-column>
      </el-table>
    </el-col>
    <el-col :span="24" style="margin-top: 15px;">
      <page ref="page" :page="search.page" :total="confList.total" @current-change="handlePage" @size-change="handlePageSize" />

      <conf-create ref="confCreate" @success="getConfList" />
      <conf-update ref="confUpdate" @success="getConfList" />
    </el-col>
  </el-row>
</template>

<script>
import { sysConfigs } from "@a/common"
import ConfCreate from "./components/add-conf"
import ConfUpdate from "./components/edit-conf"
import Page from "@c/Page"

export default {
  name: "Config",
  components: {
    Page,
    ConfCreate,
    ConfUpdate,
  },
  data() {
    return {
      loadings: {
        pageLoading: false,
      },
      confList: {
        total: 0,
        list: [],
      },
      search: { _k: "", _desc: "", state: 1, page: 1, page_size: 10 },
    }
  },
  computed: {},
  mounted() {
    this.getConfList()
  },
  methods: {
    getConfList() {
      this.loadings.pageLoading = true

      sysConfigs(this.search)
        .then((response) => {
          this.confList = response.data
          this.loadings.pageLoading = false
        })
        .catch(() => {
          this.loadings.pageLoading = false
        })
    },
    add() {
      this.$refs.confCreate.visible = true
    },
    copyRow(row) {
      this.$refs.confCreate.confForm = {
        desc: row.desc,
        key: row.key,
        val: row.val,
        state: row.state,
        bak1: row.bak1,
        bak2: row.bak2,
        remark: row.remark,
      }
      this.$refs.confCreate.visible = true
    },
    editRow(row) {
      this.$refs.confUpdate.initPage(row.id)
    },
    doSearch() {
      this.search.page = 1
      this.getConfList()
    },
    handlePage(p) {
      this.search.page = p
      this.getConfList()
    },
    handlePageSize(p) {
      this.search.page_size = p
      this.getConfList()
    },
    sortChange({ column, prop, order }) {
      let orderBy = { order: "", by: "" }
      if (order !== null) {
        orderBy = { order: prop, by: order === "ascending" ? "asc" : "desc" }
      }
    },
  },
}
</script>
