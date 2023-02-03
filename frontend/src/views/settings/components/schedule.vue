<template>
  <DialogPanel :visible="visible" title="调度任务" width="430px" confirm-text="立即调度" @confirm="confirm" @cancel="cancel" :confirm-loading="loading">
    <el-form size="mini" label-width="120px" :rules="rules" :model="jobForm" ref="jobForm">
      <el-form-item label="调度模块">
        <el-input :value="ApiModules[jobForm.api_module]" disabled />
      </el-form-item>
      <el-form-item label="调度日期" prop="stat_day">
        <el-date-picker v-model="jobForm.stat_day" type="date" placeholder="选择日期" value-format="yyyy-MM-dd" :clearable="false"
          :picker-options="pickerOptions" />
      </el-form-item>
      <el-form-item label="停止调度规则" prop="pause_rule">
        <el-select v-model="jobForm.pause_rule">
          <el-option :key="k" :label="v" :value="Number(k)" v-for="(v,k) in PauseRules" />
        </el-select>
      </el-form-item>
    </el-form>
  </DialogPanel>
</template>

<script>
import DialogPanel from "@/components/DialogPanel"
import { settingsCronSchedule } from "@a/common"

export default {
  name: "SettingsCronSchedule",
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
      jobForm: {
        api_module: "",
        pause_rule: 0,
        stat_day: "",
      },
      rules: {
        stat_day: { required: true, message: "调度日期必填" },
        pause_rule: { required: true, message: "停止调度规则必选" },
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
    setDefault(api_module) {
      this.jobForm.api_module = api_module
      this.visible = true
    },
    confirm() {
      this.$refs.jobForm.validate((v) => {
        if (v) {
          this.$confirm("此调度可能需要较长时间, 是否继续?", "提示", {
            confirmButtonText: "确定",
            cancelButtonText: "取消",
            type: "warning",
          })
            .then(() => {
              this.loading = true
              settingsCronSchedule(this.jobForm)
                .then(() => {
                  this.$message.info("后台设置调度成功，刷新列表以最后调度完成时间为准")
                  this.loading = false
                  this.visible = false
                })
                .catch((err) => {
                  this.loading = false
                })
            })
            .catch(() => {})
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
