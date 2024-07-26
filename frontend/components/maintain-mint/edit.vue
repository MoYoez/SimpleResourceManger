<script setup lang="ts">
import {AddRes, getRes, InterfaceData,UpdateRes} from "../../api/maintain";
import {onMounted, reactive, ref, watch} from "vue";
import {ElMessage} from "element-plus";
import {getWhoami} from "../../api/whoami";


const valid = ref(true)

const getValidData = ref<InterfaceData[]>()

const isProcessOk = ref(false)
const newRefer = ref<InterfaceData>()
const GetUserRole =ref()
const GetUserVerifiedName = ref("")

const form = reactive({
  m_id: "",
  m_time: 0,
  m_content: "",
  m_process: 0,
  m_name: "",
  v_verify: false,
  v_time: 0,
  v_user: "",
})


function Updater() {UpdateRes(form).then(() => {
    ElMessage.success("添加成功")
    valid.value = false
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

watch(() => form.m_id,(newReferValue)=> {
  newRefer.value = (getValidData?.value ?? []).filter(item => item.m_id === newReferValue)[0]
  form.m_name = newRefer.value?.m_name
  form.m_time = newRefer.value?.m_time
  form.m_content = newRefer.value?.m_content
  form.m_process = newRefer.value?.m_process
  form.v_verify = newRefer.value?.v_verify
  form.v_time = newRefer.value?.v_time
  form.v_user = newRefer.value?.v_user
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
        title="编辑"
        width="600"
        destroy-on-close
    >

      <el-form :model="form" label-widh="auto" style="max-width: 600px">
        <el-form-item label="维护ID">
          <el-select v-model="form.m_id" >
            <el-option
                v-for="item in getValidData"
                :key="item.m_id"
                :label="item.m_id"
                :value="item.m_id"
            />
          </el-select>
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
