<script setup>
import {getWhoami} from "../../api/whoami";
import {useRouter} from "vue-router";
import {ElMessage} from "element-plus";

const routersHelper = useRouter();

getWhoami().catch(() => {
  ElMessage.error("你还没登陆呢!")
  setTimeout(() => {
    routersHelper.push("/login")
  }, 1000)
})

</script>

