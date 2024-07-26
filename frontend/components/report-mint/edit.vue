<script setup lang="ts">

import {onMounted, reactive, ref, watch} from "vue";
import {AddRes, getRes,ReportMain, ResResponse,ModifyRes} from "../../api/report";
import {ElMessage} from "element-plus";
import {getResResource, ResData} from "../../api/res";
import {getWhoami} from "../../api/whoami";

const valid = ref(true)
const GetUserVerifiedName = ref("")
const GetUserRole = ref("")
const getResResourceList = ref<ResData[]>()
const GetProcess = ref(false)
const newRefer = ref<ReportMain>()

const form = reactive<ReportMain>({
  report_id: "",
  report_user:"",
  report_refer_res: "",
  report_content: "",
  report_time: 0,
  report_emergency: false,
  report_process: 0,
  report_feedback_time: 0,
  report_feedback_rate: 0,
  report_feedback_content: "",
  report_feedback_logs: "",
  report_feedback_user: "",
});

const Postptions = [
  {
    value: 0,
    label: '未开始',
  },
  {
    value: 1,
    label: '正在进行',
  },
  {
    value: 2,
    label: '已经完成',
  },
]


const validData = ref<ResResponse>()
const ValidDataList = ref<ReportMain[]>()

onMounted(() => {
  getWhoami().then((res) => {
    GetUserVerifiedName.value = res.data.user
    GetUserRole.value = res.data.auth
    form.report_user = GetUserVerifiedName.value
    form.report_feedback_user = GetUserVerifiedName.value
  })
  getRes().then(res => {
    validData.value = res.data
    if (parseInt(GetUserRole.value) >=2) {
      ValidDataList.value = res.data.list.filter(item => item.report_user === GetUserVerifiedName.value)
      UserIsGuest.value = true
    } else {
      ValidDataList.value = res.data.list
    }
  })
  getResResource().then(res => {
    getResResourceList.value = res.data.list
  })
  form.report_time = new Date().getTime()
})

const UserIsGuest = ref(false)

watch(()=>form.report_process,(newVal)=> {
  GetProcess.value = newVal === 2;
})

watch(()=> form.report_id,(newVal)=> {
  newRefer.value = (validData.value?.list ?? []).filter(item => item.report_id === newVal)[0]
  form.report_feedback_user = newRefer.value?.report_feedback_user
  form.report_feedback_logs = newRefer.value?.report_feedback_logs
  form.report_feedback_content = newRefer.value?.report_feedback_content
  form.report_feedback_rate = newRefer.value?.report_feedback_rate
  form.report_feedback_time = newRefer.value?.report_feedback_time
  form.report_process = newRefer.value?.report_process
  form.report_emergency = newRefer.value?.report_emergency
  form.report_time = newRefer.value?.report_time
  form.report_content = newRefer.value?.report_content
  form.report_refer_res = newRefer.value?.report_refer_res
  form.report_user = newRefer.value?.report_user
})

function Updater() {
  ModifyRes(form).then((res) => {
    ElMessage.success("修改成功")
    valid.value = !valid.value
  }).catch(err => {
    ElMessage.error(err.response.data.message)
  })
}


</script>

<template>
  <el-dialog
      v-model="valid"
      title="编辑"
      width="600"
      destroy-on-close>

    <el-form :model="form">
      <el-form-item label="上报ID">
        <el-select v-model="form.report_id">
          <el-option v-for="items in ValidDataList" :key="items.report_id" :label="items.report_id" :value="items.report_id" />
        </el-select>
      </el-form-item>

      <el-form-item label="上报相关资源">
        <el-select v-model="form.report_refer_res" placeholder="请选择资源">
          <el-option v-for="items in getResResourceList" :key="items.res_id" :label="items.res_id" :value="items.res_id" />
        </el-select>
      </el-form-item>
      <el-form-item label="上报内容">
        <el-input v-model="form.report_content" placeholder="请输入上报内容" />
      </el-form-item>
      <el-form-item label="上报是否加急">
        <el-switch v-model="form.report_emergency" />
      </el-form-item>
      <div v-if="!UserIsGuest">
      <el-form-item label="上报用户">
        <el-input v-model="form.report_user" placeholder="请输入用户ID" />
      </el-form-item>
      <el-form-item label="上报时间">
        <el-date-picker
            v-model="form.report_time"
            type="datetime"
            placeholder="Pick a Date"
            format="YYYY/MM/DD hh:mm:ss"
            value-format="x"
        />
      </el-form-item>
      <el-form-item label="处理情况">
        <el-select
            v-model="form.report_process"
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
      <div v-if="GetProcess">
        <el-form-item label="反馈时间">
          <el-date-picker
              v-model="form.report_feedback_time"
              type="datetime"
              placeholder="Pick a Date"
              format="YYYY/MM/DD hh:mm:ss"
              value-format="x"
          />
        </el-form-item>
        <el-form-item label="反馈评价">
          <el-rate v-model="form.report_feedback_rate" />
        </el-form-item>
        <el-form-item label="反馈内容">
          <el-input v-model="form.report_feedback_content" />
        </el-form-item>
        <el-form-item label="维护日志">
          <el-input v-model="form.report_feedback_logs" />
        </el-form-item>
        <el-form-item label="维护人员">
          <el-input v-model="form.report_feedback_user" />
        </el-form-item>
      </div>
      </div>
      <el-button type="primary" @click="Updater">提交</el-button>
    </el-form>

  </el-dialog>

</template>

<style scoped>

</style>
