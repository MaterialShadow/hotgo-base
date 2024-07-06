<template>
  <n-space :vertical="true">
    <n-divider>演示角色登录</n-divider>
    <n-space justify="center">
      <n-button
        v-for="item in accounts"
        :key="item.username"
        type="primary"
        @click="login(item.username, item.password)"
      >
        {{ item.label }}
      </n-button>
    </n-space>
    <n-space justify="center" class="mt-2">
      <n-text depth="3">SaaS系统多租户多应用设计</n-text>
    </n-space>
  </n-space>
</template>

<script lang="ts" setup>
  interface Emits {
    (e: 'login', param: { username: string; password: string }): void;
  }

  const emit = defineEmits<Emits>();

  const accounts = [
    {
      label: '超管',
      username: 'admin',
      password: '123456',
    },
    {
      label: '管理员',
      username: 'test',
      password: '123456',
    },
    {
      label: '租户',
      username: 'ameng',
      password: '123456',
    },
    {
      label: '商户',
      username: 'abai',
      password: '123456',
    },
    {
      label: '用户',
      username: 'asong',
      password: '123456',
    },
  ];

  function login(username: string, password: string) {
    emit('login', { username, password });
  }
</script>

<style scoped></style>
