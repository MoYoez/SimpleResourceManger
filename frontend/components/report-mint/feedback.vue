<script setup>
import {reactive, ref, watch,onMounted} from 'vue'
import {getWhoami} from "../../api/whoami";
import { getRes,ModifyRes} from "../../api/report";
import {ElMessage} from "element-plus";

function Updater() {
  if (form.report_id === "") {
    ElMessage.error("请选择一个上报ID")
    return
  }
  if (form.report_feedback_content === "") {
    ElMessage.error("请填写反馈内容")
    return
  }
  ModifyRes(form).then(res => {
      ElMessage.success("更新成功")
  }).catch(err => {
    ElMessage.error(err.response.data.message)
  })
}

const form = reactive({
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
})



onMounted(()=> {
  getWhoami().then((res) => {
    GetUserVerifiedName.value = res.data.user
    GetUserRole.value = res.data.auth
    form.report_user = GetUserVerifiedName.value
    form.report_feedback_user = GetUserVerifiedName.value
  })
  getRes().then(res => {
    validData.value = res.data
    if (parseInt(GetUserRole.value) >=2) {
      ValidDataList.value = res.data.list.filter(item => item.report_user === GetUserVerifiedName.value).filter(item => parseInt(item.report_process) === 2)
    } else {
      ValidDataList.value = res.data.list
    }
  })
})

 const validData = ref();
 const ValidDataList = ref([]);
 const newRefer = ref();

const valid = ref(true)
const GetUserVerifiedName = ref("");
const GetUserRole = ref("");

watch(()=> form.report_id,(newVal)=> {
  newRefer.value = (validData.value?.list ?? []).filter(item => item.report_id === newVal)[0]
  form.report_feedback_user = newRefer.value?.report_feedback_user
  form.report_feedback_logs = newRefer.value?.report_feedback_logs
  form.report_feedback_content = newRefer.value?.report_feedback_content
  form.report_feedback_rate = newRefer.value?.report_feedback_rate
  form.report_feedback_time = new Date().getTime()
  form.report_process = newRefer.value?.report_process
  form.report_emergency = newRefer.value?.report_emergency
  form.report_time = newRefer.value?.report_time
  form.report_content = newRefer.value?.report_content
  form.report_refer_res = newRefer.value?.report_refer_res
  form.report_user = newRefer.value?.report_user
})

</script>

<template>
  <el-dialog
      v-model="valid"
      title="反馈"
      width="600"
      destroy-on-close>

    <el-form :model="form">
        <el-form-item label="上报ID">
          <el-select
              v-model="form.report_id"
              placeholder="Select"
              size="large"
              style="width: 240px"
          >
            <el-option
                v-for="item in ValidDataList"
                :key="item.report_id"
                :label="item.report_id"
                :value="item.report_id"
            />
          </el-select>
        </el-form-item>
          <el-form-item label="反馈评价">
            <el-rate v-model="form.report_feedback_rate" />
          </el-form-item>
          <el-form-item label="反馈内容">
            <el-input v-model="form.report_feedback_content" />
          </el-form-item>
      <el-button type="primary" @click="Updater">提交</el-button>
    </el-form>

  </el-dialog>

</template>


<style scoped>

</style>
