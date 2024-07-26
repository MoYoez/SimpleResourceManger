<template>
  <h1>报修管理</h1>
  <br>
  <div class="queryBody">
      <el-button type="primary" @click="ExportEdit">编辑</el-button>
      <el-button type="primary"  @click="exportAdd">新增</el-button>
      <el-button type="primary"  @click="exportReport">反馈</el-button>
      <br>
  </div>
  <el-table :data="reports" style="width: 100%"  :row-style="rowStyle">
    <el-table-column type="expand">
                  <template #default="props" >
                    <div class="table-container">
                    <h1>详细信息</h1>
                    <br>
                    <h2>处理进度</h2>
                    <el-steps style="max-width: 600px;max-height: 300px" :active="props.row.report_process+1" align-center>
                      <el-step title="未开始"  />
                      <el-step title="正在处理"  />
                      <el-step title="已完成"/>
                    </el-steps>
                    <h2>处理详细信息</h2>
                    处理开始时间: {{ SimpleTimeUnix(props.row.report_time) }}
                    <br>
                    <br>
                    处理内容: {{ props.row.report_content }}
                    <h2>反馈与日志</h2>
                      <div class="item-container">
                        <el-card>
                          <template #header>
                            <div class="card-header">
                              <span>主用户反馈</span>
                            </div>
                          </template>
                          <div v-if="props.row.report_feedback_time !== 0">
                          <el-rate
                              v-model="props.row.report_feedback_rate"
                              disabled
                              show-score
                              text-color="#ff9900"
                              score-template="{value} points"
                          />
                          <br>
                          反馈时间： {{ SimpleTimeUnix(props.row.report_feedback_time) }}
                          <br>
                          反馈内容：{{props.row.report_feedback_content}}
                          </div>
                          <div v-else>
                            暂时还没有反馈哦
                          </div>
                        </el-card>
                        <br>
                      <el-card style="max-width: 400px;min-height: 400px">
                        <template #header>
                          <div class="card-header">
                            <span>维护日志</span>
                          </div>
                        </template>
                        <ul class="items">

                          <li v-if="props.row.report_feedback_user !== ''">
                            维护人: {{ props.row.report_feedback_user }}
                            <br>
                            <br>
                            内容:
                            <br>
                            <br>
                            {{props.row.report_feedback_logs }}
                            <br>
                          </li>
                          <li v-else>
                            目前还没有报修日志哦
                          </li>
                        </ul>
                      </el-card>
                      </div>
                    </div>
                  </template>
    </el-table-column>
    <el-table-column prop="report_id" label="上报ID" width="180" sortable />
    <el-table-column prop="report_user" label="上报发起人" width="180" />
    <el-table-column prop="report_refer_res" label="上报相关资源" width="180" />
    <el-table-column prop="report_time" label="上报时间" width="180" sortable :formatter="MaintTainerTimeUnix" />
    <el-table-column prop="report_emergency" label="是否加急" :formatter="FormatterStatusEmergency" sortable :filters="[
        {text: '已加急',value: true},
        {text: '未加急',value: false},
    ]" :filter-method="filterBool"/>
    <el-table-column prop="report_process" label="状况" :formatter="FormattedProcess" sortable :filters="[
        { text: '未开始', value: 0 },
        { text: '正在处理', value: 1 },
        { text: '已经完成', value: 2 },
      ]" :filter-method="filterTag"></el-table-column>
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
  <div v-if="isshow3">
    <Feedback />
  </div>
</template>

<script setup lang="ts">

import {getRes} from "../api/report";
import {onMounted, ref} from "vue";

import Add from "../components/report-mint/add.vue";
import Edit from "../components/report-mint/edit.vue";
import Feedback from "../components/report-mint/feedback.vue";

import axios from "axios";

import {ElMessage} from "element-plus";


const isshow1 =ref(false)
const isshow2 =ref(false)
const isshow3 =ref(false)

const reports = ref()

const ExportEdit = () => {
  isshow1.value = !isshow1.value
}

const exportAdd = () => {
  isshow2.value = !isshow2.value

}

const exportReport = () => {
  isshow3.value = !isshow3.value
}
function MaintTainerTimeUnix(row,path,cellValue) {
  return new Date(cellValue).toLocaleDateString() +" " +new Date(cellValue).toLocaleTimeString()
}

function SimpleTimeUnix(cellValue) {
  return new Date(cellValue).toLocaleDateString() +" " +new Date(cellValue).toLocaleTimeString()
}

onMounted(()=> {
  getRes().then(res => {
    reports.value = res.data["list"]
  })
})



function FormatterStatusEmergency(row,column,cellvalue) {
  if (cellvalue) {
    return "已加急"
  } else  {
    return "未加急"
  }
}

const filterTag = (value: number, row) => {
  return row.report_process === value
}

const filterBool = (value: Boolean, row) => {
  return row.report_emergency === value
}

function FormattedProcess(row,column,cellvalue,index) {
  cellvalue = parseInt(cellvalue)
  switch (cellvalue) {
    case (0) :
      return "未开始"
    case (1):
      return "正在处理"
    case (2):
      return "已完成"
  }
}

function rowStyle(row) {
  if (row.row.report_process === 2) {
    return "--el-table-tr-bg-color: var(--el-color-success-light-9);"
  }

  if (row.row.report_emergency) {
    return "--el-table-tr-bg-color: var(--el-color-warning-light-9);"
  }
}

const ExportDelete = (row) => {
  axios.delete("/root/report",{
    data:{"report_id":row.report_id}
  }).then(() => {
    ElMessage.success("删除成功!")
  }).catch(err => {
    ElMessage.error(err.response.data.message)
  })
}


</script>

<style scoped>

.table-container {
  height: 100%;
  min-height: 40vh;
  flex-direction: column;
  display: flex;
}

.item-container {
  display: flex;
  flex-direction: row;
  align-content: flex-start;
}

</style>
