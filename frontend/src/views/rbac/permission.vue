<template>
  <el-row>
    <el-col :span="24" style="margin-bottom: 10px">
      <el-button type="primary" icon="el-icon-plus" @click="add({id: 0,p_name: '一级权限'})">添加一级权限
      </el-button>
      <el-button type="primary" icon="el-icon-refresh" @click="getPermissions">刷新列表</el-button>
    </el-col>
    <el-col :span="24">
      <el-table :data="list" style="width: 100%;margin-bottom: 20px;" row-key="id" border default-expand-all stripe
        :tree-props="{children: 'children', hasChildren: 'hasChildren'}" v-loading="dataLoading" size="mini">
        <el-table-column prop="p_name" label="权限名称" width="260"></el-table-column>
        <el-table-column prop="permission" label="路由地址" width="300"></el-table-column>
        <el-table-column prop="method" label="操作类型「W 写 / R 读 / * 全部」"></el-table-column>
        <el-table-column align="center" label="操作" fixed="right" width="150">
          <template slot-scope="scope">
            <el-button-group class="table-operate">
              <el-button type="primary" plain @click.native.prevent="add(scope.row)">添加
              </el-button>
              <el-button type="primary" plain @click.native.prevent="editRow(scope.row)">编辑
              </el-button>
              <el-button type="danger" plain @click.native.prevent="destroyRow(scope.row)">删除
              </el-button>
            </el-button-group>
          </template>
        </el-table-column>
      </el-table>

      <permission-create ref="permissionCreate" @success="getPermissions" />
      <permission-update ref="permissionUpdate" @success="getPermissions" />
    </el-col>
  </el-row>
</template>

<script>
import { permissionList, permissionDestroy } from "@a/role"
import PermissionCreate from "./components/add-permission"
import PermissionUpdate from "./components/edit-permission"

export default {
  name: "Permission",
  components: {
    PermissionCreate,
    PermissionUpdate,
  },
  data() {
    return {
      dataLoading: false,
      list: [],
    }
  },
  mounted() {
    this.getPermissions()
  },
  methods: {
    getPermissions() {
      this.dataLoading = true
      permissionList()
        .then((res) => {
          this.list = res.data
          this.dataLoading = false
        })
        .catch((err) => {
          this.dataLoading = false
        })
    },
    editRow(row) {
      this.$refs.permissionUpdate.permissionForm = {
        id: row.id,
        p_name: row.p_name,
        permission: row.permission,
        method: row.method,
        pid: row.pid,
      }
      this.$refs.permissionUpdate.visible = true
    },
    add(row) {
      this.$refs.permissionCreate.initCreate(row)
    },
    destroyRow(row) {
      this.$confirm("此操作会同步删除已分配的权限, 是否继续?", "警告", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "error",
      })
        .then(() => {
          this.dataLoading = true
          permissionDestroy(row.id)
            .then(() => {
              this.$message.success("删除成功")
              this.getPermissions()
            })
            .catch((err) => {
              this.dataLoading = false
            })
        })
        .catch(() => {})
    },
  },
}
</script>

<style scoped>
</style>