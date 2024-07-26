<script setup lang="ts">
import {getResResource,UpdateRes} from "../../api/res";
import {onMounted, reactive, watch} from "vue";
import {ref} from "vue"
import {ElMessage} from "element-plus";

const getValidData = ref()
const newRefer = ref()

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

onMounted(() => {
  getResResource().then(res => {
    getValidData.value = res.data.list
  })
})

watch(()=> form.res_id,(newVal) => {
  newRefer.value = getValidData.value.filter(items => items.res_id === newVal)[0]
  form.res_date = newRefer.value.res_date
  form.res_model = newRefer.value.res_model
  form.res_name = newRefer.value.res_name
  form.res_number = newRefer.value.res_number
  form.res_place = newRefer.value.res_place
  form.res_type = newRefer.value.res_type
})



function Updater() {
  UpdateRes(form).then(res => {
    ElMessage.success("更新成功")
    valid.value = !valid.value
  }).catch(err => {
    ElMessage.error(err.response.data.message)
  })
}


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
      <el-form-item label="资源ID">
              <el-select
                  v-model="form.res_id"
                  placeholder="Select"
                  size="large"
                  style="width: 240px"
              >
                <el-option
                    v-for="item in getValidData"
                    :key="item.res_id"
                    :label="item.res_id"
                    :value="item.res_id"
                />
        </el-select>
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
