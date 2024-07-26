<script setup>
import { getRes } from "../api/maintain"
import {ref} from "vue";
import Edit from "../components/maintain-mint/edit.vue"
import Add from "../components/maintain-mint/add.vue";
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
getRes().then(res => {dataExport.value = res.data["list"]
})

const ExportDelete = (row) => {
  axios.delete("/root/maintain",{
    data:{"m_id":(row.m_id)}
  }).then(() => {
    ElMessage.success("删除成功!")
    getRes().then(res => {
      dataExport.value = res.data["list"]
    })
  }).catch(err => {
    ElMessage.error(err.response.data.message)
  })
}

function MaintainerIsChecker(cellValue) {
  if (cellValue) {
    return "是"
  } else {
    return "否"
  }
}

function MaintainProcess(row,column,cellValue) {
  switch (cellValue) {
    case 0:
      return "未处理"
    case 1:
      return "正在处理"
    case 2:
      return "已处理"
  }
}

function MaintTainerTimeUnix(row, column, cellValue) {
  return new Date(cellValue ).toLocaleDateString() +" " +new Date(cellValue ).toLocaleTimeString()
}

function BaseTimeHandler(cellValue) {
  return new Date(cellValue ).toLocaleDateString() +" " +new Date(cellValue ).toLocaleTimeString()
}
const filTags = (value,row) => {
  return row.m_process === parseInt(value)
}

function OnModifeyHeaderColor(row) {
  if (row.row.m_process === 2) {
    return "--el-table-tr-bg-color: var(--el-color-success-light-9);"
  }
}



</script>

<template>
  <div class="body">

    <h1>日常维护模块</h1>
    <br>
    <div class="queryBody">
      <el-button type="primary" @click="ExportEdit">编辑</el-button>
      <el-button type="primary"  @click="exportAdd">新增</el-button>
    </div>
    <el-table :data="dataExport" :row-style="OnModifeyHeaderColor" >
      <el-table-column type="expand">
        <template #default="props">
          <h1>维护内容</h1>
          <p>{{ props.row.m_content }}</p>
        </template>
        </el-table-column>
      <el-table-column prop="m_id" label="维护ID"  />
      <el-table-column prop="m_process" label="维护状况"  :formatter="MaintainProcess" :filters="[{ text: '未处理', value: '0' },
        { text: '正在处理', value: '1' },{text:'已经处理',value: '2' },]" :filter-method="filTags" />

      <el-table-column prop="m_time" label="维护时间" :formatter="MaintTainerTimeUnix" />
      <el-table-column prop="m_name" label="维护人员"  />
      <el-table-column prop="v_verify" label="维护核实" >

        <template #default="scope">
          <el-popover effect='light' trigger="hover" placement="top" width="auto">
            <template #default v-if="scope.row.v_verify">
              <div>核实时间: {{ BaseTimeHandler(scope.row.v_time) }}</div>
              <div>核实人物: {{ scope.row.v_user }}</div>
            </template>
            <template #default v-else>
              <div>未经核实的维护项目</div>
            </template>
            <template #reference>
              <el-tag>{{ MaintainerIsChecker(scope.row.v_verify) }}</el-tag>
            </template>
          </el-popover>
        </template>
      </el-table-column>
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
