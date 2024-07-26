<script setup lang="ts">
import {UpdateReg,getRes} from "../../api/user";
import {reactive} from "vue";
import {ref} from "vue"
import {ElMessage} from "element-plus";


const sexOptions = [
  {"label":"男","value":0},
  {"label":"女","value":1}
]

const valid = ref(true)

const form = reactive({
  user_id: "",
  user_name: "",
  user_password:"",
  user_email: "",
  user_authority: 0,
  user_sex: 0,
})





function Updater() {
  getRes().then(res => {
    for (const user of res.data.list) {
      if (user.user_id == form.user_id) {
        ElMessage.error("用户ID已经存在")
        return
      }
    }
  })
  if (form.user_name === "" || form.user_id === "" || form.user_password === "" || form.user_email === "") {
    ElMessage.error("请填写完整信息")
    return
  }
  // updater
  UpdateReg(form).then(res => {
    ElMessage.success("新建账户成功")
    valid.value = !valid.value
  }).catch(err => {
    ElMessage.error(err.response.data.message)
  })
}



const Postptions = [
  {
    value: 0,
    label: 'Root 用户组',
  },
  {
    value: 1,
    label: 'Admin 用户组',
  },
  {
    value: 2,
    label: '普通用户组',
  },
]

</script>

<template>
  <div class="item-edit">
    <el-dialog
        v-model="valid"
        title="新增"
        width="600"
        destroy-on-close
    >

      <el-form :model="form" label-widh="auto" style="max-width: 600px">
        <el-form-item label="用户名">
          <el-input v-model="form.user_id" placeholder="请输入用户ID"></el-input>
        </el-form-item>
        <el-form-item label="用户名称">
          <el-input v-model="form.user_name" placeholder="请输入用户名称"></el-input>
        </el-form-item>
        <el-form-item label="用户密码">
          <el-input v-model="form.user_password" placeholder="请输入密码"></el-input>
        </el-form-item>
        <el-form-item label="用户邮箱">
          <el-input v-model="form.user_email" placeholder="请输入邮箱"></el-input>
        </el-form-item>
        <el-form-item label="用户性别">
          <el-select
              v-model="form.user_sex"
              placeholder="Select"
              size="large"
              style="width: 240px"
          >
            <el-option
                v-for="item in sexOptions"
                :key="item.value"
                :label="item.label"
                :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="用户权限组">
          <el-select
              v-model="form.user_authority"
              placeholder="Select"
              size="large"
              style="width: 240px"
          >
            <el-option
                v-for="item in Postptions"
                :key="item.value"
                :label="item.label"
                :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="Updater()">提交</el-button>
        </el-form-item>
      </el-form>

    </el-dialog>
  </div>
</template>

<style scoped>

</style>
