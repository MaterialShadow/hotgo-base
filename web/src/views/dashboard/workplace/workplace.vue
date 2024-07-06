<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="工作台">
        <n-grid cols="2 s:1 m:1 l:2 xl:2 2xl:2" responsive="screen">
          <n-gi>
            <div class="flex items-center">
              <div>
                <p class="px-4 text-xl">工作台</p>
              </div>
            </div>
          </n-gi>
        </n-grid>
      </n-card>
    </div>
  </div>
</template>

<script lang="ts" setup></script>

<style lang="less" scoped>
  .project-card {
    margin-right: -6px;

    &-item {
      margin: -1px;
      width: 33.333333%;
    }
  }
</style>
