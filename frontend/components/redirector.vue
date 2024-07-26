<script setup>
import {getWhoami} from "../api/whoami";
import {useRouter} from "vue-router";
import {ElMessage} from "element-plus";

const routersHelper = useRouter();

getWhoami().then(res =>{
  if (res.data["user"] !== "" ) {
    ElMessage.success("已经登陆~正在跳转")
    setTimeout(() => {
      routersHelper.push("/admin/main")
    }, 1000)
  }})
</script>

