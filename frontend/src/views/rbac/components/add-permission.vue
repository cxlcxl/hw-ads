<template>
  <dialog-panel title="权限创建" confirm-text="添加" :visible="visible" @cancel="cancel" @confirm="add" :confirm-loading="loading" width="450px">
    <el-form :model="permissionForm" ref="permissionForm" label-width="100px" size="small">
      <el-form-item label="上级权限" prop="pid" :rules="{required: true, message: '请选择上级权限路由'}">
        <el-input :value="permissionForm.parent_name" disabled placeholder="由页面选择传递" />
      </el-form-item>
      <el-form-item label="权限名称" prop="p_name" :rules="{required: true, message: '请填写权限名称'}">
        <el-input v-model="permissionForm.p_name" />
      </el-form-item>
      <el-form-item label="路由地址" prop="permission" :rules="{required: true, message: '请填写路由地址'}">
        <el-input v-model="permissionForm.permission" placeholder="请严格填写，示例：user/list" />
      </el-form-item>
      <el-form-item label="请求类型" prop="method" :rules="{required: true, message: '请选择允许的请求类型'}">
        <el-radio-group v-model="permissionForm.method" size="mini">
          <el-radio-button label="*"></el-radio-button>
          <el-radio-button label="R"></el-radio-button>
          <el-radio-button label="W"></el-radio-button>
        </el-radio-group>
      </el-form-item>
    </el-form>
  </dialog-panel>
</template>

<script>
import DialogPanel from "@c/DialogPanel"
import { permissionCreate } from "@a/role"

export default {
  components: {
    DialogPanel,
  },
  data() {
    return {
      visible: false,
      loading: false,
      permissionForm: {
        parent_name: "",
        p_name: "",
        permission: "",
        method: "*",
        pid: 0,
      },
    }
  },
  methods: {
    initCreate(row) {
      this.$set(this.permissionForm, "pid", row.id)
      this.$set(this.permissionForm, "parent_name", row.p_name)
      this.visible = true
    },
    cancel() {
      this.$refs.permissionForm.resetFields()
      this.visible = false
    },
    add() {
      this.$refs.permissionForm.validate((v) => {
        if (v) {
          this.loading = true
          permissionCreate(this.permissionForm)
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
