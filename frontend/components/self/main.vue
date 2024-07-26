<script setup lang="ts">
import {GetUserInfo, PostUserInfoUpdater} from "../../api/self"
import {computed, onMounted, reactive, ref} from "vue";
import {ElMessage} from "element-plus";


const userInfoData = ref()

const form = reactive({
  user_name:"",
  user_password:"",
  retry_user_password:"",
  user_email:"",
  user_sex:""
})

const showref = ref(true)

const sexOptions = [
  {"label":"男","value":"0"},
  {"label":"女","value":"1"}
]

onMounted(() => {
  GetUserInfo().then((res) => {
    userInfoData.value = res.data
    form.user_name = userInfoData.value.user_name
    form.user_email = userInfoData.value.user_email
    form.user_sex = SexModifier(userInfoData.value.user_sex)
  })
})

function SexModifier(sex) {
  if (sex === 0) {
    return "男"
  } else {
    return "女"
  }
}

function updater() {
  if (form.user_password !== form.retry_user_password) {
    ElMessage.error("两次密码不一致")
    return
  }
  const formUpdater = {
    user_id: userInfoData.value.user_id,
    user_name: form.user_name,
    user_password: form.user_password,
    user_email: form.user_email,
    user_authority: parseInt(userInfoData.value.user_authority),
    user_sex: parseInt(form.user_sex),

  }
PostUserInfoUpdater(formUpdater).then(() => {
    ElMessage.success("修改成功")
    showref.value = false
  }).catch(err => {
    ElMessage.error(err.response.data["message"])
  })
}




</script>

<template>
  <el-dialog title="修改个人信息" v-model="showref">
    <el-form>
      <el-form-item label="用户名">
    <el-input v-model="form.user_name" placeholder="用户名" />
        </el-form-item>
      <el-form-item label="用户邮箱">
    <el-input v-model="form.user_email" placeholder="用户邮箱" />
        </el-form-item >
      <el-form-item label="用户性别">
        <el-select v-model="form.user_sex" placeholder="用户性别" >
          <el-option
              v-for="item in sexOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
          />
        </el-select>
      </el-form-item >
      <el-form-item label="用户密码">
    <el-input v-model="form.user_password" placeholder="用户密码" />
        </el-form-item>
      <div class="items-password" v-if="form.user_password !== ''">
      <el-form-item label="请在输入一次密码">
    <el-input v-model="form.retry_user_password" placeholder="请再一次输入密码" />
        </el-form-item>
      </div>
      <br>
    <el-button type="primary" @click="updater()">提交</el-button>
    </el-form>
  </el-dialog>
</template>

<style scoped>

</style>
