<script setup lang="ts">
import {ElMessage} from "element-plus";
import {getRes, ModifyUserInfo} from "../../api/user";
import {onMounted,reactive, ref,watch} from "vue";

const valid = ref(true)
const ValidData = ref()
const newRefer = ref()





const sexOptions = [
  {"label":"男","value":0},
  {"label":"女","value":1}
]

const form = reactive({
  user_id: "",
  user_name: "",
  user_password:"",
  user_email: "",
  user_authority: 0,
  user_sex: 0,
})


function Updater() {
  // updater
  ModifyUserInfo(form).then(res => {
    ElMessage.success("修改账户成功")
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

onMounted(() => {
  getRes().then(res => {
    ValidData.value = res.data.list
  })
})

watch(()=>form.user_id,(newVal) => {
  newRefer.value = ValidData.value.filter(items => items.user_id === newVal)[0]
  form.user_email = newRefer.value?.user_email ?? ""
  form.user_name = newRefer.value?.user_name ?? ""
  form.user_authority = newRefer.value?.user_authority ?? 0
  form.user_sex = newRefer.value?.user_sex ?? 0
})


</script>

<template>
  <div class="item-edit">
    <el-dialog
        v-model="valid"
        title="编辑"
        width="600"
        destroy-on-close
    >

      <el-form :model="form" label-widh="auto" style="max-width: 600px">
        <el-form-item label="用户名">
          <el-select v-model="form.user_id">
            <el-option v-for="items in ValidData" :key="items.user_id" :label="items.user_id" :value="items.user_id" />
          </el-select>
        </el-form-item>
        <el-form-item label="用户名称">
          <el-input v-model="form.user_name" placeholder="请输入用户名称"></el-input>
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
