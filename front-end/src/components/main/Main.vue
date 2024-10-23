<script setup lang="ts"></script>

<template>
  <div class="main"></div>
</template>

<style scoped lang="scss">
@use "../../style/variables.scss" as *;

.main {
  margin-top: 20px;
  height: calc(100vh - 65px - 60px);
  background-color: $aside-bg-color;
  border-radius: $border-radius;
}
</style>
