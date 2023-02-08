<template>
  <dialog-panel title="角色修改" confirm-text="保存" :visible="visible" @cancel="cancel" @confirm="save" :confirm-loading="loading">
    <el-form :model="roleForm" ref="roleForm" label-width="100px" size="small">
      <el-form-item label="角色名称" prop="role_name" :rules="{required: true, message: '请填写角色名称'}">
        <el-input v-model="roleForm.role_name" />
      </el-form-item>
      <el-form-item label="状态" prop="state">
        <el-switch v-model="roleForm.state" :active-value="1" :inactive-value="0" />
      </el-form-item>
      <el-form-item label="选择权限" prop="permissions" :rules="{required: true, message: '请选择权限'}">
        <el-tree :data="permissions" show-checkbox node-key="permission" :props="defaultProps" @check-change="handleCheck" ref="tree"
          :default-checked-keys="roleForm.permissions" />
      </el-form-item>
    </el-form>
  </dialog-panel>
</template>

<script>
import DialogPanel from "@c/DialogPanel"
import { roleUpdate, roleInfo } from "@a/role"
import { removeArrayItem } from "@/utils"

export default {
  components: {
    DialogPanel,
  },
  props: {
    permissions: Array,
    roleType: Object,
  },
  data() {
    return {
      visible: false,
      loading: false,
      roleForm: {
        id: 0,
        role_name: "",
        state: 1,
        sys: 0,
        permissions: [],
      },
      defaultProps: {
        children: "children",
        label: "p_name",
      },
    }
  },
  methods: {
    // Tips：el-tree 默认选中有bug，需设置 setCheckedNodes
    initUpdate(row) {
      // this.visible = true
      roleInfo(row.id)
        .then((res) => {
          this.roleForm.id = res.data.id
          this.roleForm.role_name = res.data.role_name
          this.roleForm.state = res.data.state
          this.roleForm.sys = res.data.sys

          let ps = []
          if (Array.isArray(res.data.permissions)) {
            ps = res.data.permissions
          }
          this.roleForm.permissions = ps
          this.visible = true
          this.$refs.tree.setCheckedNodes(ps)
        })
        .catch((err) => {
          console.log(err)
        })
    },
    cancel() {
      this.$refs.roleForm.resetFields()
      this.visible = false
    },
    save() {
      this.$set(this.roleForm, "permissions", this.$refs.tree.getCheckedKeys())
      this.$refs.roleForm.validate((v) => {
        if (v) {
          this.loading = true
          roleUpdate(this.roleForm)
            .then((res) => {
              this.$message.success("修改成功")
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
    // 	共三个参数，依次为：传递给 data 属性的数组中该节点所对应的对象、节点本身是否被选中、节点的子树中是否有被选中的节点
    handleCheck(data, selfIsChecked, hasBrotherChecked) {
      if (selfIsChecked && !this.roleForm.permissions.includes(data.permission)) {
        this.roleForm.permissions.push(data.permission)
      } else {
        this.roleForm.permissions = removeArrayItem(this.roleForm.permissions, data.permission)
      }
    },
  },
}
</script>
