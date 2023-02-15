<template>
  <DialogPanel :visible="visible" title="调度规则修改" width="430px" confirm-text="保存" @confirm="confirm" @cancel="cancel" :confirm-loading="loading">
    <el-form size="mini" label-width="120px" :rules="rules" :model="jobForm" ref="jobForm">
      <el-form-item label="调度模块">
        <el-input :value="ApiModules[jobForm.api_module]" disabled />
      </el-form-item>
      <el-form-item label="当前版本">
        <el-input :value="jobForm.version" disabled />
      </el-form-item>
      <el-form-item label="下次调度日期" prop="stat_day">
        <el-date-picker v-model="jobForm.stat_day" type="date" placeholder="选择日期" value-format="yyyy-MM-dd" :clearable="false"
          :picker-options="pickerOptions" />
      </el-form-item>
      <el-form-item label="停止调度规则" prop="pause_rule">
        <el-select v-model="jobForm.pause_rule">
          <el-option :key="k" :label="v" :value="Number(k)" v-for="(v,k) in PauseRules" />
        </el-select>
      </el-form-item>
      <el-form-item label="调度周期" prop="job_schedule">
        <el-input v-model="jobForm.job_schedule" placeholder="* * * * * [分/时/日/月/周]" />
        <p class="schedule">
          <span class="text-error">
            * 仅支持 Crontab 表达式和 <code>robfig/cron</code> 包的规则
            [<a href="https://cron.qqe2.com/" target="_blank">Crontab 参考</a>,<a href="https://pkg.go.dev/github.com/robfig/cron"
              target="_blank">包规则参考</a>]
          </span>
        </p>
      </el-form-item>
      <el-form-item label="排序" prop="order_by">
        <el-input-number v-model="jobForm.order_by" :min="1" :max="99" />
      </el-form-item>
      <el-form-item label="备注" prop="remark">
        <el-input v-model="jobForm.remark" type="textarea" :rows="3" />
      </el-form-item>
    </el-form>
  </DialogPanel>
</template>

<script>
import DialogPanel from "@/components/DialogPanel"
import { settingsCronInfo, settingsCronUpdate } from "@a/common"
import { parseTime } from "@/utils"

export default {
  name: "SettingsCronCreate",
  components: { DialogPanel },
  props: {
    ApiModules: {
      type: Object,
      required: true,
    },
    PauseRules: {
      type: Object,
      required: true,
    },
  },
  data() {
    return {
      visible: false,
      loading: false,
      job_id: 0,
      jobForm: {
        api_module: "",
        job_schedule: "",
        order_by: 99,
        pause_rule: 0,
        remark: "",
        stat_day: "",
        version: 1,
      },
      rules: {
        stat_day: { required: true, message: "调度日期必填" },
        job_schedule: { required: true, message: "调度周期必填" },
        pause_rule: { required: true, message: "停止调度规则必选" },
        order_by: { required: true, message: "排序必填" },
      },
      pickerOptions: {
        disabledDate(time) {
          return time.getTime() > Date.now()
        },
        shortcuts: [
          {
            text: "今天",
            onClick(picker) {
              picker.$emit("pick", new Date())
            },
          },
        ],
      },
    }
  },
  methods: {
    setDefault(id) {
      settingsCronInfo(id)
        .then((res) => {
          this.visible = true
          this.job_id = id
          this.jobForm = res.data
          this.jobForm.stat_day = parseTime(this.jobForm.stat_day, "{y}-{m}-{d}")
        })
        .catch((err) => {
          this.$message.error(err)
        })
    },
    confirm() {
      this.$refs.jobForm.validate((v) => {
        if (v) {
          this.loading = true
          settingsCronUpdate(this.jobForm, this.job_id)
            .then((res) => {
              this.loading = false
              this.$message.success("修改成功")
              this.$emit("success")
              this.cancel()
            })
            .catch((err) => {
              this.loading = false
            })
        } else {
          return false
        }
      })
    },
    cancel() {
      this.visible = false
    },
  },
}
</script>

<style lang="scss">
.schedule {
  font-size: 12px;
  font-weight: 600;
  line-height: 20px;
  margin-top: 5px;
}
</style>