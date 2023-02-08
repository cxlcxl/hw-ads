<template>
  <dialog-panel title="应用信息修改" confirm-text="保存" :visible="visible" @cancel="cancel" @confirm="save" :confirm-loading="loading">
    <el-form :model="appForm" ref="appForm" label-width="120px" size="small" :rules="userRules">
      <el-form-item label="APP ID">
        <el-input v-model="appForm.app_id" placeholder="请填写APP ID" disabled />
      </el-form-item>
      <el-form-item label="应用名称" prop="app_name">
        <el-input v-model="appForm.app_name" placeholder="请填写应用名称" />
      </el-form-item>
      <el-form-item label="应用包名" prop="pkg_name">
        <el-input v-model="appForm.pkg_name" placeholder="请填写应用包名" />
      </el-form-item>
      <el-form-item label="渠道" prop="channel">
        <el-select v-model="appForm.channel" style="width: 100%;">
          <el-option v-for="(key, val) in appChannel" :label="key" :value="Number(val)" :key="val" />
        </el-select>
      </el-form-item>
      <el-form-item label="应用类型" prop="app_type">
        <el-select v-model="appForm.app_type" style="width: 100%;">
          <el-option v-for="(key, val) in appType" :label="key" :value="val" :key="val" />
        </el-select>
      </el-form-item>
      <el-form-item label="标签" prop="tags">
        <el-input v-model="appForm.tags" placeholder="多个请用英文 ',' 号隔开" />
      </el-form-item>
    </el-form>
  </dialog-panel>
</template>

<script>
import DialogPanel from "@c/DialogPanel"
import { appUpdate, appInfo } from "@a/app"

export default {
  components: {
    DialogPanel,
  },
  props: {
    appType: Object,
    appChannel: Object,
  },
  data() {
    return {
      visible: false,
      loading: false,
      remoteLoading: false,
      appForm: {
        id: 0,
        app_name: "",
        app_id: "",
        pkg_name: "",
        tags: "",
        channel: "",
        app_type: "",
      },
      userRules: {
        app_name: { required: true, message: "请填写应用名称" },
        channel: { required: true, message: "请选择渠道" },
        pkg_name: { required: true, message: "请填写应用包名" },
      },
    }
  },
  methods: {
    initUpdate(id) {
      appInfo(id)
        .then((res) => {
          this.visible = true
          this.appForm = res.data
        })
        .catch(() => {
          this.$message.error("请求错误")
        })
    },
    cancel() {
      this.$refs.appForm.resetFields()
      this.visible = false
    },
    save() {
      this.$refs.appForm.validate((v) => {
        if (v) {
          this.loading = true
          appUpdate(this.appForm)
            .then((res) => {
              this.$message.success("修改成功")
              this.$emit("success")
              this.loading = false
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
  },
}
</script>
