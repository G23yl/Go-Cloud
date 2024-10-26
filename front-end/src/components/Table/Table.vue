<script setup lang="ts">
import type { DIVAOData } from "@/types/types"

interface Props {
  data: DIVAOData[]
  icon: string[]
  size?: string
  color?: string
}
const { data, icon, size = "lg", color = "#74c0fc" } = defineProps<Props>()
</script>

<template>
  <div class="content">
    <el-table :data="data" style="width: 98%" border max-height="447px">
      <el-table-column label="名称" width="180">
        <template #default="scope">
          <div style="display: flex; align-items: center">
            <font-awesome-icon :icon="icon" :size="size" :style="{ color: color }" />
            <span style="margin-left: 10px">{{ scope.row.fileName }}</span>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="大小" width="180">
        <template #default="scope">
          <el-tag type="primary" round>{{ scope.row.fileSize }} KB</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="下载次数">
        <template #default="scope">
          <el-tag type="success" round>{{ scope.row.downloadNum }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="修改日期">
        <template #default="scope">
          <el-tag type="info" round>{{ scope.row.updateTime }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作">
        <template #default>
          <div>
            <el-tooltip content="删除" placement="top">
              <font-awesome-icon
                class="icon"
                :icon="['fas', 'trash-can']"
                style="color: #74c0fc"
                size="lg"
              />
            </el-tooltip>
            <el-tooltip content="下载" placement="top">
              <font-awesome-icon
                class="icon"
                :icon="['fas', 'cloud-arrow-down']"
                style="color: #74c0fc"
                size="lg"
              />
            </el-tooltip>
            <el-tooltip content="分享" placement="top">
              <font-awesome-icon
                class="icon"
                :icon="['fas', 'share-nodes']"
                style="color: #74c0fc"
                size="lg"
              />
            </el-tooltip>
            <el-tooltip content="重命名" placement="top">
              <font-awesome-icon
                class="icon"
                :icon="['fas', 'pen-to-square']"
                style="color: #74c0fc"
                size="lg"
              />
            </el-tooltip>
            <el-tooltip content="预览" placement="top">
              <font-awesome-icon
                class="icon"
                :icon="['fas', 'eye']"
                style="color: #74c0fc"
                size="lg"
              />
            </el-tooltip>
          </div>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<style scoped lang="scss">
.icon {
  cursor: pointer;
  margin-right: 10px;
  &:focus {
    outline: none;
  }
}
</style>
