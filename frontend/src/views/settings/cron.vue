<template>
  <el-row>
    <el-col :span="24" style="font-weight: 500; margin-bottom: 10px; color: #909399;">
      <p class="text-error">* 后台总调度任务每 5 分钟会刷新写入规则，修改规则后一般 6 分钟内生效</p>
      <p>“下一次调度日期” 按需设置「无日期属性的调度任务可以无视」</p>
    </el-col>
    <el-col :span="24">
      <el-button type="primary" icon="el-icon-refresh" size="mini" @click="getCronList">刷新列表</el-button>
    </el-col>
    <el-col :span="24">
      <el-table v-loading="loadings.pageLoading" :data="cronList.list" highlight-current-row stripe border size="mini" style="margin-top: 15px">
        <el-table-column prop="api_module" label="调度模块" width="268">
          <template slot-scope="scope">{{cronList.api_modules[scope.row.api_module]}}</template>
        </el-table-column>
        <el-table-column prop="stat_day" label="下一次调度日期" width="120">
          <template slot-scope="scope">{{scope.row.stat_day|timeFormat("{y}-{m}-{d}")}}</template>
        </el-table-column>
        <el-table-column prop="job_schedule" label="调度周期" width="110" />
        <el-table-column prop="pause_rule" label="停止调度规则" width="120">
          <template slot-scope="scope">
            <span :class="{'text-error': scope.row.pause_rule === -1, 'text-success': scope.row.pause_rule !== -1}">
              {{cronList.pause_rules[scope.row.pause_rule]}}
            </span>
          </template>
        </el-table-column>
        <el-table-column prop="version" label="版本" width="70" align="center" />
        <el-table-column prop="remark" label="备注" />
        <el-table-column prop="order_by" label="排序" width="70" align="center" />
        <el-table-column prop="last_schedule" label="最后调度完成时间" width="140">
          <template slot-scope="scope">{{scope.row.last_schedule|timeFormat}}</template>
        </el-table-column>
        <el-table-column align="center" label="操作" fixed="right" width="110">
          <template slot-scope="scope">
            <el-button-group class="table-operate">
              <el-button type="primary" plain @click.native.prevent="editRow(scope.row)">编辑</el-button>
              <el-button plain @click.native.prevent="scheduleRow(scope.row)" v-show="cronList.handler_jobs.includes(scope.row.api_module)">调度
              </el-button>
            </el-button-group>
          </template>
        </el-table-column>
      </el-table>
    </el-col>

    <cron-update ref="cronUpdate" :api-modules="cronList.api_modules" :pause-rules="cronList.pause_rules" @success="getCronList" />
    <cron-schedule ref="cronSchedule" :api-modules="cronList.api_modules" :pause-rules="cronList.handler_pause_rules" />
  </el-row>
</template>

<script>
import { settingsCron } from "@a/common"
import { parseTime } from "@/utils"
import CronUpdate from "./components/edit.vue"
import CronSchedule from "./components/schedule.vue"

export default {
  name: "Cron",
  components: { CronUpdate, CronSchedule },
  data() {
    return {
      loadings: {
        pageLoading: false,
      },
      cronList: {
        list: [],
        pause_rules: {},
        handler_pause_rules: {},
        api_modules: {},
        handler_jobs: [],
      },
    }
  },
  computed: {},
  mounted() {
    this.getCronList()
  },
  filters: {
    timeFormat(t, fmt) {
      return parseTime(t, fmt)
    },
  },
  methods: {
    getCronList() {
      this.loadings.pageLoading = true
      settingsCron()
        .then((res) => {
          this.cronList.list = res.data.list
          this.cronList.pause_rules = res.data.pause_rules
          this.cronList.api_modules = res.data.api_modules
          this.cronList.handler_jobs = res.data.handler_jobs
          this.cronList.handler_pause_rules = res.data.handler_pause_rules
          this.loadings.pageLoading = false
        })
        .catch(() => {
          this.loadings.pageLoading = false
        })
    },
    editRow(row) {
      this.$refs.cronUpdate.setDefault(row.id)
    },
    scheduleRow(row) {
      this.$refs.cronSchedule.setDefault(row.api_module)
    },
  },
}
</script>
