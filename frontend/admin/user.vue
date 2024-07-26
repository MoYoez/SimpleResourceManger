<script setup>
import { getRes} from "../api/user"
import {ref} from "vue";
import Edit from "../components/user-mint/edit.vue"
import Add from "../components/user-mint/new.vue";
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

const querydata = ref("")
const dataExport = ref(null)

getRes().then(res => {
  dataExport.value = res.data["list"]
})


function returnGroup(row, column, cellValue, index) {
  switch (cellValue) {
    case 0:
      return "Root用户组"
    case 1:
      return "Admin 用户组"
    case 2:
      return "普通用户组"
  }
}

function SexModflier(row,column,cellvalue,index) {
  switch (cellvalue) {
    case 0:
      return "男"
    case 1:
      return "女"
  }
}

const ExportDelete = (row) => {
  axios.delete("/admin/register",{
    data:{"user_id":row.user_id}
  }).then(() => {
    ElMessage.success("删除成功!")
    getRes().then(res => {
      dataExport.value = res.data["list"]
    })
  }).catch(err => {
    ElMessage.error(err.response.data["message"])
  })
}

function Refresher() {
  getRes().then(res => {
    dataExport.value = res.data["list"]
  })
}

setInterval(Refresher,2500)

</script>

<template>
  <div class="body">

    <h1>用户管理模块</h1>
    <br>
    <div class="queryBody">
      <el-button type="primary" @click="ExportEdit">编辑</el-button>
      <el-button type="primary"  @click="exportAdd">新增</el-button>
    </div>
    <el-table :data="dataExport" style="" >
      <el-table-column prop="user_id" label="用户ID"  />
      <el-table-column prop="user_name" label="用户名" />
      <el-table-column prop="user_email" label="用户邮箱" />
      <el-table-column prop="user_sex" label="用户性别"  :formatter="SexModflier"/>

      <el-table-column prop="user_authority" label="用户组" :formatter="returnGroup" />
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
