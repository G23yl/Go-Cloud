<script setup lang="ts">
import Menu from "@/components/menu/Menu.vue"
</script>

<template>
  <div class="layout">
    <el-container>
      <el-aside width="220px" style="margin: 20px">
        <Menu />
      </el-aside>
      <RouterView />
    </el-container>
  </div>
</template>

<style scoped lang="scss">
.layout {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  overflow: hidden;
  background-color: #eee;
}
</style>
