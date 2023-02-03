<template>
  <el-row :gutter="10">
    <el-col :span="24" class="search-container">
      <el-form ref="_search" :model="search" inline size="small">
        <el-form-item>
          <el-input v-model="search.k" class="w240" clearable placeholder="输入搜索：国家名称/代码" />
        </el-form-item>
        <el-form-item label="">
          <el-button type="primary" icon="el-icon-search" class="item" @click="doSearch">查询</el-button>
        </el-form-item>
      </el-form>
    </el-col>
    <el-col :span="24">
      <el-button type="primary" icon="el-icon-plus" size="mini" @click="add">添加区域/国家</el-button>
      <span class="notice text-error">国家信息因涉及投放定向设置，所以不可修改只可新增「<a
          href="https://developer.huawei.com/consumer/cn/doc/distribution/promotion/marketing-api-tool-targeting7-0000001286343134#ZH-CN_TOPIC_0000001286343134__li1079145116117"
          target="_blank">国家数据来源</a>」</span>
    </el-col>
    <el-col :span="24" style="margin-top: 15px;">
      <div class="region-area-country">
        <div class="region-area">
          <el-menu :default-active="search.area_id" class="region-area-menu" @select="handleSelect" width="160px">
            <el-menu-item index="-1">全部国家</el-menu-item>
            <el-menu-item index="-2">未分地区</el-menu-item>
            <el-menu-item :index="item.id.toString()" v-for="item in areas" :key="item.id">{{item.name}}</el-menu-item>
          </el-menu>
        </div>
        <div class="region-country">
          <el-table :data="countries.list" size="mini" border stripe v-loading="loading">
            <el-table-column align="center" label="操作" width="100">
              <template slot-scope="scope">
                <el-button-group class="table-operate">
                  <el-button type="primary" plain @click.native.prevent="editRow(scope.row)">设置地区</el-button>
                </el-button-group>
              </template>
            </el-table-column>
            <el-table-column label="国家代码" prop="c_code" align="center" width="80" />
            <el-table-column label="国家名称" prop="c_name" />
            <el-table-column label="地区" prop="area_name" width="130" />
          </el-table>

          <div style="margin-top: 10px;">
            <page ref="page" :page="search.page" :total="countries.total" @current-change="handlePage" @size-change="handlePageSize"
              :limit="search.page_size" />
          </div>
        </div>
      </div>
    </el-col>
    <region-create ref="regionCreate" :areas="areas" @success="getRegionList" />
    <region-area-set ref="regionAreaSet" :areas="areas" @success="getRegionList" />
  </el-row>
</template>

<script>
import { overseasAreas, overseasCountries } from "@/api/common"
import Page from "@/components/Page"
import RegionCreate from "./components/region-add.vue"
import RegionAreaSet from "./components/region-area-set.vue"

export default {
  // name:"Region",
  components: { Page, RegionCreate, RegionAreaSet },
  data() {
    return {
      loading: false,
      search: {
        area_id: "-1",
        k: "",
        page: 1,
        page_size: 10,
      },
      areas: [],
      countries: {
        total: 0,
        list: [],
      },
    }
  },
  mounted() {
    this.getRegionList()
  },
  methods: {
    getRegionList() {
      this.loading = true
      Promise.all([this.getRegionAreas()])
        .then(() => {
          overseasCountries(this.search)
            .then((res) => {
              this.loading = false
              this.countries.list = res.data.list
              this.countries.total = res.data.total
            })
            .catch(() => {
              this.loading = false
            })
        })
        .catch((err) => {
          this.loading = false
        })
    },
    getRegionAreas() {
      return new Promise((resolve, reject) => {
        overseasAreas()
          .then((res) => {
            this.areas = res.data
            resolve()
          })
          .catch((err) => {
            reject()
          })
      })
    },
    add() {
      this.$refs.regionCreate.setDefault()
    },
    editRow(row) {
      this.$refs.regionAreaSet.setDefault(row.c_code, row.c_name)
    },
    handlePage(v) {
      this.search.page = v
      this.getRegionList()
    },
    handlePageSize(v) {
      this.search.page_size = v
      this.doSearch()
    },
    handleSelect(v) {
      this.search.area_id = v
      this.doSearch()
    },
    doSearch() {
      this.search.page = 1
      this.getRegionList()
    },
    setRegion(v) {
      console.log(v.name)
    },
  },
}
</script>

<style lang="less" scoped>
.notice {
  font-size: 12px;
  margin-left: 15px;
  font-weight: 500;
}
.region-area-country {
  display: flex;
  flex-direction: row;

  .region-area {
    width: 180px;
    margin-right: 15px;

    .region-area-menu {
      .el-menu-item {
        height: 36px;
        line-height: 36px;
        font-weight: 500;
        color: #606266;
      }
      .is-active {
        color: #409eff;
      }
    }
  }
  .region-country {
    flex: 1;
  }
}
</style>