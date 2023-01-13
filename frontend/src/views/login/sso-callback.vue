<template></template>

<script>
import { ssoLogin } from "@/api/user"

export default {
  name: "SsoCallBack",
  data() {
    return {}
  },
  mounted() {
    this.ssoCallback()
  },
  methods: {
    ssoCallback() {
      if (!this.$route.query.ticket || this.$route.query.ticket === "") {
        this.$message.error("凭证信息有误")
        return
      }
      this.$store
        .dispatch("user/ssoLogin", this.$route.query.ticket)
        .then(() => {
          this.$router.replace({ path: "/" })
        })
        .catch((err) => {
          this.loading = false
        })
    },
  },
}
</script>
