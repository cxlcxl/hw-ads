<template>
  <dialog-panel title="添加地区/国家" confirm-text="保存" :visible="visible" @cancel="cancel" @confirm="save" :confirm-loading="loading" width="368px">
    <el-form :model="regionForm" ref="regionForm" label-width="90px" size="small">
      <el-form-item label="国家名称">
        <el-input :value="regionForm.c_name" disabled />
      </el-form-item>
      <el-form-item label="所属地区" prop="area_ids" :rules="{required: true,message:'请选择所属地区'}">
        <el-select v-model="regionForm.area_ids" placeholder="请选择所属地区" multiple>
          <el-option v-for="item in areas" :key="item.id" :value="Number(item.id)" :label="item.name" />
        </el-select>
      </el-form-item>
    </el-form>
  </dialog-panel>
</template>

<script>
import DialogPanel from "@c/DialogPanel"
import { regionAreaSet } from "@a/common"

export default {
  components: {
    DialogPanel,
  },
  props: {
    areas: {
      required: true,
      type: Array,
    },
  },
  data() {
    return {
      visible: false,
      loading: false,
      regionForm: {
        c_name: "",
        c_code: "",
        area_ids: [],
      },
    }
  },
  methods: {
    setDefault(c_code, c_name) {
      this.regionForm.c_code = c_code
      this.regionForm.c_name = c_name
      this.visible = true
    },
    cancel() {
      this.$refs.regionForm.resetFields()
      this.visible = false
    },
    save() {
      this.$refs.regionForm.validate((v) => {
        if (v) {
          this.loading = true
          regionAreaSet(this.regionForm)
            .then((res) => {
              this.$message.success("设置成功")
              this.$emit("success")
              this.loading = false
              this.cancel()
            })
            .catch((err) => {
              this.loading = false
              console.log(err)
            })
        } else {
          return false
        }
      })
    },
  },
}
</script>
