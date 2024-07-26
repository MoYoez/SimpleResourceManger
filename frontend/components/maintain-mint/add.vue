<script setup lang="ts">

import {ref, watch} from "vue"
import {ElMessage} from "element-plus";
import {AddRes, getRes, InterfaceData} from "../../api/maintain";
import {onMounted, reactive} from "vue";
import {getWhoami} from "../../api/whoami";

const getValidData = ref<InterfaceData[]>()

const isProcessOk = ref(false)
const GetUserRole =ref()
const GetUserVerifiedName = ref("")

const valid =ref(true)

const form = reactive({
  m_id: "",
  m_time: 0,
  m_content: "",
  m_process: 0,
  m_name: GetUserVerifiedName,
  v_verify: false,
  v_time: 0,
  v_user: GetUserVerifiedName,
})

function Updater() {
  if (form.m_id === "") {
    ElMessage.error("维护ID不能为空")
    return
  }
  if (form.m_name === "" || form.m_content === "" || form.m_time === 0) {
    ElMessage.error("请填写完整信息")
    return
  }
  if (isProcessOk.value && (form.v_time === 0 || form.v_user === "")) {
    ElMessage.error("请填写完整信息")
    return
  }
  AddRes(form).then(() => {
    ElMessage.success("添加成功")
    valid.value = !valid.value
  }).catch(err => {
    ElMessage.error(err.response.data.message)
  })

}

onMounted(() => {
  getRes().then(res => {
    getValidData.value = res.data["list"]
  })
  // getUserRole
  getWhoami().then((res) => {
      GetUserVerifiedName.value = res.data.user
      GetUserRole.value = res.data.auth
  })
})

watch(() => form.m_process, (newVal) => {
  isProcessOk.value = newVal === 2 && GetUserRole.value === 0;
})


const StatusOptions = [
  {value: 0, label: "未处理"},
  {value: 1, label: "正在处理"},
  {value: 2, label: "已处理"},
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
        <el-form-item label="维护ID">
          <el-input v-model="form.m_id"> </el-input>
        </el-form-item>
        <el-form-item label="维护状态">
          <el-select v-model="form.m_process" placeholder="请选择维护状态">
            <el-option
                v-for="item in StatusOptions"
                :key="item.value"
                :label="item.label"
                :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="维护时间">
          <el-date-picker
              v-model="form.m_time"
              type="datetime"
              placeholder="Pick a Date"
              format="YYYY/MM/DD hh:mm:ss"
              value-format="x"
          />
        </el-form-item>
        <el-form-item label="维护内容">
          <el-input v-model="form.m_content" :autosize="{ minRows: 3, maxRows: 6 }" placeholder="请输入维护内容"></el-input>
        </el-form-item>
        <el-form-item label="维护人员">
          <el-input v-model="form.m_name"></el-input>
        </el-form-item>
        <el-form-item v-if="isProcessOk" label="是否核实">
          <el-switch v-model="form.v_verify"></el-switch>
        </el-form-item>
        <el-form-item v-if="isProcessOk" label="核实时间">
          <el-date-picker
              v-model="form.v_time"
              type="datetime"
              placeholder="Pick a Date"
              format="YYYY/MM/DD hh:mm:ss"
              value-format="x"
          />
        </el-form-item>
        <el-form-item v-if="isProcessOk" label="核实人员">
          <el-input v-model="form.v_user"></el-input>
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
