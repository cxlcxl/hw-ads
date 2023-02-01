<template>
  <DialogPanel :visible="visible" title="展示字段选择" width="460px" confirm-text="确认" @confirm="confirm" @cancel="cancel">
    <el-select v-model="selects" multiple placeholder="请选择" style="width: 100%;">
      <el-option v-for="item in Columns" :key="item.key" :label="item.label" :value="item.key">
      </el-option>
    </el-select>
  </DialogPanel>
</template>

<script>
import DialogPanel from "@/components/DialogPanel"
import { setReportComprehensiveColumns } from "@/api/report"

export default {
  name: "SelectColumns",
  components: { DialogPanel },
  props: {
    Columns: {
      required: true,
      type: Array,
    },
  },
  data() {
    return {
      visible: false,
      selects: [],
    }
  },
  methods: {
    setDefault() {
      this.Columns.forEach((v) => {
        if (v.show && !this.selects.includes(v.key)) {
          this.selects.push(v.key)
        }
      })

      this.visible = true
    },
    confirm() {
      setReportComprehensiveColumns({ columns: this.selects })
        .then((res) => {
          this.$emit("confirm", this.selects)
          this.visible = false
        })
        .catch((err) => {
          this.$message.error("字段设置失败：" + err)
        })
    },
    cancel() {
      this.visible = false
    },
  },
}
</script>