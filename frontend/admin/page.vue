<script setup>
import {onMounted, ref,computed} from "vue";
import {useRouter} from "vue-router";
import Unlogin from "../components/unlogined/unlogin.vue";
import Status from "../components/status/status.vue";
import {getWhoami} from "../api/whoami";
import {SetAuth} from "../api/request";
import {ArrowRight} from "@element-plus/icons-vue";
import Self from "../components/self/main.vue"

const router = useRouter()

function getPostionName(path) {
  switch (path) {
    case "/admin/resadmin":
      return "资产管理"
    case "/admin/useradmin":
      return "用户管理"
    case "/admin/reportadmin":
      return "报修管理"
    case "/admin/maintainadmin":
      return "日常维护管理"
  }
}
const expName = computed(() => {
  return getPostionName(router.currentRoute.value.path)
})

const getName = ref("")

function logout() {
  SetAuth("")
  router.push("/login")
}

onMounted(() => {
  getName.value = getWhoami().then(res => {
    getName.value = res.data["user"]
  })
})
const editedtable = ref(false)

function ShowEdit() {
  editedtable.value = !editedtable.value
}

</script>

<template>
  <div class="body">
    <el-container>
      <el-aside width="200px">
        <el-header class="page-header">
          <div class="title">
            固定资产维护报修
          </div>
        </el-header>
        <el-menu class="menu-page" :default-active="$route.path" active-text-color="#fff" text-color="rgba(255,255,255,0.65)" router>
          <el-menu-item index="/admin/main">首页</el-menu-item>
          <br>
          <el-menu-item index="/admin/resadmin">资产管理</el-menu-item>
          <br>
          <el-menu-item index="/admin/useradmin">用户管理</el-menu-item>
          <br>
          <el-menu-item index="/admin/reportadmin">报修管理</el-menu-item>
          <br>
          <el-menu-item index="/admin/maintainadmin">日常维护管理</el-menu-item>
        </el-menu>
      </el-aside>
      <el-container>
        <el-header>
            <el-breadcrumb :separator-icon="ArrowRight">
              <el-breadcrumb-item to="/admin/main">首页</el-breadcrumb-item>
              <el-breadcrumb-item :to="router.currentRoute.value.path">{{expName}}</el-breadcrumb-item>
            </el-breadcrumb>
          <div class="user-icon-item">
            <el-dropdown>
              <span class="el-dropdown-link">
                <span>{{ getName }}</span>
              </span>
              <template #dropdown>
              <el-dropdown-menu slot="dropdown">
                <el-dropdown-item @click="logout">退出登录</el-dropdown-item>
                <el-dropdown-item @click="ShowEdit">修改信息</el-dropdown-item>
              </el-dropdown-menu>
              </template>
            </el-dropdown>
          </div>
        </el-header>
        <el-main>
          <router-view> </router-view>
        </el-main>
      </el-container>
    </el-container>
    <unlogin />
    <status />
    <div v-if="editedtable">
      <Self />
    </div>
  </div>
</template>

<style scoped>


.page-header {
  margin-top: 20px;

}

.el-aside {
  height: 100vh;
  background-color: #001529;
}
.el-menu {
  background-color: #001529;
  border: none;
}

.el-menu-item.is-active {
  background-color: #40a9ff !important;
  border-radius: 10px;
}

.el-menu-item:hover {
  background-color: #40a9ff !important;
  border-radius: 10px;
}

.el-menu-title:hover{
  color: white!important;
}


.page-header {
  color: white;
  font-weight: 700!important;
  background-color: #001529;
}

.menu-page {
  background-color: #001529;
}

.el-header {
  display: flex;
  align-items: center;
}
.user-icon-item {
  margin-left: auto;
  display: flex;
  align-items: center;
  margin-right: 20px;
}


</style>


