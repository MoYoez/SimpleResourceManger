<script setup lang="ts">
import {AddRes, getRes} from "../../api/res";
import {reactive} from "vue";
import {ref} from "vue"
import {ElMessage} from "element-plus";


const valid = ref(true)

const form = reactive({
  res_id: "",
  res_number: 0,
  res_type: "",
  res_name: "",
  res_model: "",
  res_date: new Date().getTime(),
  res_place: "",
})



function Updater() {
  if (form.res_date === 0 || form.res_id === "" || form.res_model === "" || form.res_name === "" || form.res_place === "" || form.res_type === "") {
    ElMessage.error("请填写完整信息")
    return
  }
  AddRes(form).then(() => {
    ElMessage.success("新增成功!")
  }).catch(err => {
    ElMessage.error(err.response.data.message)
  })
}
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
        <el-form-item label="资源ID">
          <el-input v-model="form.res_id" placeholder="请输入资源ID"></el-input>
        </el-form-item>
        <el-form-item label="资源数量">
          <el-input-number v-model="form.res_number" placeholder="请输入资源数量"></el-input-number>
        </el-form-item>
        <el-form-item label="资源类型">
          <el-input v-model="form.res_type" placeholder="请输入资源类型"></el-input>
        </el-form-item>
        <el-form-item label="资源名称">
          <el-input v-model="form.res_name" placeholder="请输入资源名称"></el-input>
        </el-form-item>
        <el-form-item label="资源型号">
          <el-input v-model="form.res_model" placeholder="请输入资源型号"></el-input>
        </el-form-item>
        <el-form-item label="登记日期">
          <el-date-picker
              v-model="form.res_date"
              type="datetime"
              placeholder="Pick a Date"
              format="YYYY/MM/DD hh:mm:ss"
              value-format="x"
          />
        </el-form-item>
        <el-form-item label="存放地点">
          <el-input v-model="form.res_place" placeholder="请输入存放地点"></el-input>
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
