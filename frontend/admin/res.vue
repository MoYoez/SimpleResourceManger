<script setup>
import { getResResource } from "../api/res"
import {ref} from "vue";
import Edit from "../components/res-mint/edit.vue"
import Add from "../components/res-mint/add.vue";
import axios from "axios";
import {ElMessage} from "element-plus";

const isshow1 =ref(false)
const isshow2 =ref(false)

const ExportEdit = () => {
  isshow1.value = !isshow1.value
}

const exportAdd = () => {
  isshow2.value = !isshow2.value
}

const dataExport = ref(null)
getResResource().then(res => {dataExport.value = res.data["list"]})

const ExportDelete = (row) => {
axios.delete("/root/res",{
  data:{"res_id":(row.res_id)}
}).then(() => {
    ElMessage.success("删除成功!")
    getResResource().then(res => {
      dataExport.value = res.data["list"]
    })
}).catch(err => {
  ElMessage.error(err.response.data["message"])
})
}

function Refresher() {
  getResResource().then(res => {
    dataExport.value = res.data["list"]
  })
}
setInterval(Refresher,2500)

function MaintTainerTimeUnix(row,path,cellValue) {
  return new Date(cellValue).toLocaleDateString() +" " +new Date(cellValue).toLocaleTimeString()
}

</script>

<template>
<div class="body">

  <h1>资产管理模块</h1>
  <br>
  <div class="queryBody">
    <el-button type="primary" @click="ExportEdit">编辑</el-button>
    <el-button type="primary"  @click="exportAdd">新增</el-button>
    </div>
  <el-table :data="dataExport" style="" >
    <el-table-column prop="res_id" label="资源ID"  />
    <el-table-column prop="res_number" label="资源数量" />
    <el-table-column prop="res_type" label="资源类型" />
    <el-table-column prop="res_name" label="资源名称"  />
    <el-table-column prop="res_model" label="资源型号"  />
    <el-table-column prop="res_date" label="登记日期"  :formatter="MaintTainerTimeUnix" />
    <el-table-column prop="res_place" label="存放地点"  />
    <el-table-column prop="" label="操作">
      <template #default="{row}">
        <el-popconfirm title="确认删除?" @confirm="ExportDelete(row)">
          <template #reference>
            <el-button link type="primary" size="small">删除</el-button>
          </template>
        </el-popconfirm>
      </template>
    </el-table-column>
  </el-table>
  <div v-if="isshow1">
    <edit />
  </div>
  <div v-if="isshow2">
    <add />
  </div>
</div>


</template>

<style scoped>
.queryBody {
  margin: 20px;
  position: relative;
  font-weight: 700;
  font-size: 18px;
}

</style>
