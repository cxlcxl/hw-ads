<template>
  <dialog-panel title="添加地区/国家" confirm-text="添加" :visible="visible" @cancel="cancel" @confirm="add" :confirm-loading="loading" width="468px">
    <el-form :model="regionForm" ref="regionForm" label-width="90px" size="small">
      <el-tabs type="border-card" v-model="regionForm.t">
        <el-tab-pane label="地区信息" name="area">
          <el-form-item label="地区名称" prop="area_name" :rules="{required: true,message:'请填写地区名称'}" v-if="regionForm.t === 'area'">
            <el-input v-model="regionForm.area_name" placeholder="请填写地区名称" />
          </el-form-item>
        </el-tab-pane>
        <el-tab-pane label="国家信息" name="country">
          <el-form-item label="所属地区" prop="area_ids" :rules="{required: true,message:'请选择所属地区'}" v-if="regionForm.t === 'country'">
            <el-select v-model="regionForm.area_ids" placeholder="请选择所属地区" multiple>
              <el-option v-for="item in areas" :key="item.id" :value="Number(item.id)" :label="item.name" />
            </el-select>
          </el-form-item>
          <el-form-item label="国家名称" prop="c_name" :rules="{required: true,message:'请填写国家名称'}" v-if="regionForm.t === 'country'">
            <el-input v-model="regionForm.c_name" placeholder="请填写国家名称" />
          </el-form-item>
          <el-form-item label="国家代码" prop="c_code" :rules="{required: true,message:'请填写国家代码'}" v-if="regionForm.t === 'country'">
            <el-input v-model="regionForm.c_code" placeholder="请填写国家代码" />
          </el-form-item>
        </el-tab-pane>
      </el-tabs>
    </el-form>
  </dialog-panel>
</template>

<script>
import DialogPanel from "@c/DialogPanel"
import { regionCreate } from "@a/common"

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
      remoteLoading: false,
      regionForm: {
        t: "area",
        area_ids: [],
        c_code: "",
        c_name: "",
        area_name: "",
      },
    }
  },
  methods: {
    setDefault() {
      this.visible = true
    },
    cancel() {
      this.$refs.regionForm.resetFields()
      this.visible = false
    },
    add() {
      this.$refs.regionForm.validate((v) => {
        if (v) {
          this.loading = true
          regionCreate(this.regionForm)
            .then((res) => {
              this.$message.success("创建成功")
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
