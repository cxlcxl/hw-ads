<template>
  <DialogPanel :visible="visible" title="展示字段选择" width="460px" confirm-text="确认" @confirm="confirm" @cancel="cancel">
    <el-select v-model="selected" multiple placeholder="请选择" style="width: 100%;">
      <el-option v-for="item in Columns" :key="item.key" :label="item.label" :value="item.key">
      </el-option>
    </el-select>
  </DialogPanel>
</template>

<script>
import DialogPanel from "@/components/DialogPanel"

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
      selected: [],
    }
  },
  methods: {
    setDefault(v) {
      if (Array.isArray(v)) {
        if (v.length === 0) {
          this.Columns.forEach((v) => {
            this.selected.push(v.key)
          })
        } else {
          this.selected = v
        }
      }
      this.visible = true
    },
    confirm() {
      this.$emit("confirm", this.selected)
      this.visible = false
    },
    cancel() {
      this.visible = false
    },
  },
}
</script>