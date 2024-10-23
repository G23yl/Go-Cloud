<script setup lang="ts">
import Title from "@/components/Title/Title.vue"
import Main from "@/components/main/Main.vue"
import { ref } from "vue"

const title = ref("主页")
</script>

<template>
  <el-container>
    <el-header>
      <Title :title="title" />
    </el-header>
    <el-main>
      <Main />
    </el-main>
  </el-container>
</template>

<style scoped></style>
