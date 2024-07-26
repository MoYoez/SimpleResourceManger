<script setup>
import { ref } from "vue";
import axios from "axios";

const isBackendServiceClosed = ref(false)

axios.get("/").catch(() => {
  isBackendServiceClosed.value = true
})

</script>

<template>
  <el-dialog
      v-model="isBackendServiceClosed"
      title="您疑似没有打开后端服务器?"
      width="500"
  >
    <span>需要打开一下后端服务器一下/ 这样的话可以使用后端的各项数据</span>
    <template #footer>
      <div class="dialog-footer">
        <el-button type="primary" @click="isBackendServiceClosed = false">
          好滴
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<style scoped>

</style>
