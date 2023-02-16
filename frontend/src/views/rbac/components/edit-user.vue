<template>
  <dialog-panel title="用户修改" confirm-text="保存" :visible="visible" @cancel="cancel" @confirm="save" :confirm-loading="loading" width="420px">
    <el-form :model="userForm" ref="userForm" label-width="90px" size="small" :rules="userRules">
      <el-form-item label="用户名称" prop="username">
        <el-input v-model="userForm.username" placeholder="请填写用户名" />
      </el-form-item>
      <el-form-item label="邮箱地址" prop="email">
        <el-input v-model="userForm.email" placeholder="请填写邮箱，登陆使用" />
      </el-form-item>
      <el-form-item label="手机号" prop="mobile">
        <el-input v-model="userForm.mobile" placeholder="请填写手机号码" />
      </el-form-item>
      <el-form-item label="角色" prop="role_id">
        <el-select v-model="userForm.role_id" placeholder="请选择" style="width: 100%;" clearable>
          <el-option v-for="item in roles" :key="item.id" :label="item.role_name" :value="item.id" />
        </el-select>
      </el-form-item>
      <el-form-item label="状态" prop="state">
        <el-switch v-model="userForm.state" :active-value="1" :inactive-value="0" />
      </el-form-item>
      <el-form-item label="内部账号" prop="is_internal">
        <el-radio-group v-model="userForm.is_internal">
          <el-radio :label="1">是</el-radio>
          <el-radio :label="0">否</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="投放账户" prop="market_accounts" v-if="userForm.is_internal === 0" :rules="{required: true, message:'请关联投放账户'}">
        <el-select v-model="userForm.market_accounts" multiple clearable filterable style="width: 100%;">
          <el-option :key="item.id" :label="item.account_name" :value="item.id" v-for="item in accounts"
            v-show="item.account_type === Vars.AccountTypeMarket" />
        </el-select>
      </el-form-item>
      <el-form-item label="变现账户" prop="ads_accounts" v-if="userForm.is_internal === 0" :rules="{required: true, message:'请关联变现账户'}">
        <el-select v-model="userForm.ads_accounts" multiple clearable filterable style="width: 100%;">
          <el-option :key="item.id" :label="item.account_name" :value="item.id" v-for="item in accounts"
            v-show="item.account_type === Vars.AccountTypeAds" />
        </el-select>
      </el-form-item>
      <el-form-item label="登录密码" prop="pass">
        <el-input v-model="userForm.pass" placeholder="字母开头，数字特殊字符 [@.&!#?,%$] 的 6 - 18 位" />
      </el-form-item>
    </el-form>
  </dialog-panel>
</template>

<script>
import DialogPanel from "@c/DialogPanel"
import { getUserInfo, userUpdate } from "@a/user"
import { validPass, validEmail } from "@/utils/validate"
import { allAccounts } from "@a/account"
import Vars from "@/vars.js"

export default {
  components: {
    DialogPanel,
  },
  props: {
    roles: {
      default: () => [],
      type: Array,
    },
  },
  data() {
    var checkPass = (rule, value, callback) => {
      if (value === "") {
        callback()
      } else {
        if (!validPass(value)) {
          callback(new Error("密码格式不符合要求"))
        } else {
          callback()
        }
      }
    }
    var checkEmail = (rule, value, callback) => {
      if (!validEmail(value)) {
        callback(new Error("邮箱格式不正确"))
      } else {
        callback()
      }
    }
    return {
      Vars,
      visible: false,
      loading: false,
      remoteLoading: false,
      userForm: {
        id: 0,
        role_id: 0,
        is_internal: 1,
        username: "",
        email: "",
        mobile: "",
        state: 1,
        pass: "",
        market_accounts: [],
        ads_accounts: [],
      },
      accounts: [],
      userRules: {
        username: { required: true, message: "请填写用户名称" },
        email: [{ required: true, message: "请填写邮箱" }, { validator: checkEmail }],
        role_id: { required: true, message: "请选择角色" },
        pass: { validator: checkPass },
        is_internal: { required: true, message: "请选择用户类型" },
      },
    }
  },
  methods: {
    initUpdate(user_id) {
      Promise.all([this.getAllAccounts()])
        .then(() => {
          getUserInfo(user_id)
            .then((res) => {
              this.userForm = res.data
              this.visible = true
            })
            .catch(() => {})
        })
        .catch(() => {})
    },
    getAllAccounts() {
      return new Promise((resolve, reject) => {
        if (this.accounts.length > 0) {
          resolve()
        } else {
          allAccounts()
            .then((res) => {
              this.accounts = res.data
              resolve()
            })
            .catch((err) => {
              reject(err)
            })
        }
      })
    },
    cancel() {
      this.$refs.userForm.resetFields()
      this.visible = false
    },
    save() {
      this.$refs.userForm.validate((v) => {
        if (v) {
          this.loading = true
          userUpdate(this.userForm)
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
  },
}
</script>
