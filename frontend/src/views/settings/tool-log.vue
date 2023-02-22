<template>
  <el-row>
    <el-col :span="24" class="search-container" v-loading="loadings">
      <el-form ref="_search" :model="search" inline size="small">
        <el-form-item>
          <el-date-picker v-model="search.d" value-format="yyyy-MM-dd" placeholder="下载日期，最多下载 15 天前的数据" class="w300"/>
        </el-form-item>
        <el-form-item label="">
          <el-button icon="el-icon-download" @click="downloadLog">下载日志文件</el-button>
        </el-form-item>
      </el-form>
    </el-col>
  </el-row>
</template>

<script>
import { toolLogs, toolLogDownload } from '@a/common'

export default {
  name: "ToolLog",
  data() {
    return {
      loadings: false,
      search: {
        d: ''
      },
      logList: {
        list: [],
        total: 0
      }
    }
  },
  methods: {
    downloadLog() {
      this.loadings = true
      toolLogs(this.search).then(res => {
        this.loadings = false
        window.open(toolLogDownload+res.data)
      }).catch(() => {
        this.loadings = false
      })
    }
  }
}
</script>
