<template>
  <el-row>
    <el-col :span="24" style="margin-bottom: 20px;">
      <el-button icon="el-icon-refresh" size="mini" plain @click="getVersions" type="primary" />
    </el-col>
    <el-col :span="24" class="timeline-version">
      <el-timeline>
        <el-timeline-item :timestamp="item.date" placement="top" v-for="item in versionList" :key="item.v" color="#409eff">
          <el-card>
            <h3>版本：{{item.v}}</h3>
            <p v-for="t in item.fixs" class="v-text">{{t}}</p>
          </el-card>
        </el-timeline-item>
      </el-timeline>
    </el-col>
  </el-row>

</template>

<script>
import { versionInfo } from "@a/common"

export default {
  // name: "Version",
  data() {
    return {
      versionList: [{ date: "2023-02-12", v: "v1.0.0", fixs: ["更新了什么什么"] }],
    }
  },
  mounted() {
    this.getVersions()
  },
  methods: {
    getVersions() {
      versionInfo()
        .then((res) => {
          this.versionList = res.data
        })
        .catch(() => {})
    },
  },
}
</script>

<style lang="less">
.timeline-version {
  .el-card__body {
    padding: 8px 20px;

    .v-text {
      font-size: 12px;
    }
  }
}
</style>