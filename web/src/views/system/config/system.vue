<template>
  <div>
    <n-grid cols="24 300:1 600:24" :x-gap="24">
      <n-grid-item span="6">
        <n-card :bordered="false" size="small" class="proCard">
          <n-thing
            class="thing-cell"
            v-for="item in typeTabList"
            :key="item.key"
            :class="{ 'thing-cell-on': type === item.key }"
            @click="switchType(item)"
          >
            <template #header>{{ item.name }}</template>
            <template #description>{{ item.desc }}</template>
          </n-thing>
        </n-card>
      </n-grid-item>
      <n-grid-item span="18">
        <n-card :bordered="false" size="small" :title="typeTitle" class="proCard">
          <BasicSetting v-if="type === 1" />
          <ThemeSetting v-if="type === 2" />
          <RevealSetting v-if="type === 3" />
          <EmailSetting v-if="type === 4" />
          <SmsSetting v-if="type === 5" />
          <LoginSetting v-if="type === 6" />
          <CashSetting v-if="type === 7" />
          <UploadSetting v-if="type === 8" />
          <GeoSetting v-if="type === 9" />
          <PaySetting v-if="type === 10" />
          <WechatSetting v-if="type === 11" />
        </n-card>
      </n-grid-item>
    </n-grid>
  </div>
</template>
<script lang="ts">
  import { defineAsyncComponent, defineComponent, reactive, toRefs } from 'vue';
  /** 异步加载的组件，用到的时候再加载组件 */
  const BasicSetting = defineAsyncComponent(() => {
      return import('./BasicSetting.vue');
    }),
    RevealSetting = defineAsyncComponent(() => {
      return import('./RevealSetting.vue');
    }),
    ThemeSetting = defineAsyncComponent(() => {
      return import('./ThemeSetting.vue');
    }),
    UploadSetting = defineAsyncComponent(() => {
      return import('./UploadSetting.vue');
    }),
    LoginSetting = defineAsyncComponent(() => {
      return import('./LoginSetting.vue');
    });
  const typeTabList = [
    {
      name: '基本设置',
      desc: '系统常规设置',
      key: 1,
    },
    {
      name: '登录注册',
      desc: '登录注册配置',
      key: 6,
    },
    {
      name: '云存储',
      desc: '配置上传文件驱动',
      key: 8,
    },
  ];
  export default defineComponent({
    components: {
      BasicSetting,
      RevealSetting,
      ThemeSetting,
      UploadSetting,
      LoginSetting,
    },
    setup() {
      const state = reactive({
        type: 1,
        typeTitle: '基本设置',
      });

      function switchType(e) {
        state.type = e.key;
        state.typeTitle = e.name;
      }

      return {
        ...toRefs(state),
        switchType,
        typeTabList,
      };
    },
  });
</script>
<style lang="less" scoped>
  .thing-cell {
    margin: 0 -16px 10px;
    padding: 5px 16px;

    &:hover {
      background: #f3f3f3;
      cursor: pointer;
    }
  }

  .thing-cell-on {
    background: #f0faff;
    color: #2d8cf0;

    ::v-deep(.n-thing-main .n-thing-header .n-thing-header__title) {
      color: #2d8cf0;
    }

    &:hover {
      background: #f0faff;
    }
  }
</style>
