<script setup>
import { reactive,ref } from 'vue'
import { ElMessage } from 'element-plus'
import {Login} from "../api/login";
import {getWhoami} from "../api/whoami";
import Unlogin from "../components/unlogined/unlogin.vue";
import axios from "axios";
import Redirector from "../components/redirector.vue";
import {useRouter} from "vue-router";
import {SetAuth} from "../api/request";

const routerHelper = useRouter();

const form = reactive({
  user_name: '',
  user_password: '',
})

axios.get("/").then(() => {
  ElMessage.info("后端服务正在运行~")
})

function LoginUsage() {
  if (form.user_name === "" || form.user_password === "") {
    ElMessage.error("用户名或者密码不能为空")
    return
  }
 new Login(form).then(res => {
   if (res.data["status"] === true) {
     ElMessage.success("登陆成功!");
     SetAuth(res.data["auth"]);
     setTimeout(() => {
       routerHelper.push("/admin/main")
     }, 1000)
   } else {
      ElMessage.error("登陆失败,请检查用户名和密码")
   }
  })
}


function Whoami() {
  getWhoami().then(res => {
    ElMessage.info(res.data["user"])
  }).catch(err => {
    ElMessage.error("你还没登陆呢!")
  })
}


</script>

<template>
  <div class="body">
    <div class="title">
      固定资产日常维护报修系统
    </div>
    <div class="container-inputbox">
    <el-input v-model="form.user_name" placeholder="请输入用户名" type="text"></el-input>
      <br>
      <el-input v-model="form.user_password" placeholder="请输入密码" type="password" show-password ></el-input>
      <br>
      <el-button type="primary" @click="LoginUsage">登录</el-button>
      <br>
    </div>
  <unlogin />
  <redirector />
  </div>

</template>

<style scoped>
.body {
  display: flex;
  background: repeating-linear-gradient(45deg,#5bcefa,#5bcefa 50px,#fff 20px,#fff 650px);
  justify-content: center;
  align-items: center;
  flex-direction: column;
  height: 100vh;
  width: 100vw;
}
.title {
  margin-bottom: 10vh;
  font-size: 24px;
  font-weight: bold;
  color: #333;
}
.container-inputbox {
  display: flex;
  flex-direction: column;
  width: 300px;
  padding: 20px;
  background-color: #fff;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  border-radius: 25px;
  border: 2px dashed #aaa;
}

</style>
