import { ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { Option } from '@/utils/hotgo';
import { Dicts } from '@/api/dict/dict';

export interface State {
  id: number;
  pid?: number;
  title: string;
  name: string;
  path: string;
  label: string;
  icon: string;
  type: number;
  redirect: string;
  permissions: string;
  permissionName: string;
  component?: string;
  alwaysShow: number;
  activeMenu: string;
  isRoot: number;
  isFrame: number;
  frameSrc: string;
  keepAlive: number;
  hidden: number;
  affix: number;
  status: number;
  sort: number;
  disabled: boolean;
  children?: State[];
}

export const defaultState: State = {
  id: 0,
  pid: 0,
  title: '',
  name: '',
  path: '',
  label: '',
  icon: '',
  type: 1,
  redirect: '',
  permissions: '',
  permissionName: '',
  component: null,
  alwaysShow: 2,
  activeMenu: '',
  isRoot: 2,
  isFrame: 2,
  frameSrc: '',
  keepAlive: 2,
  hidden: 2,
  affix: 2,
  status: 1,
  sort: 10,
  disabled: false,
  children: null,
};

export function newState(state: State | null): State {
  if (state !== null) {
    return defaultValueCheck(cloneDeep(state));
  }
  return defaultValueCheck(cloneDeep(defaultState));
}

// 默认值校正，主要为了解决历史数据格式不规范问题
export function defaultValueCheck(state: State): State {
  if (state.pid < 1) {
    state.pid = null;
  }
  if (state.alwaysShow != 1) {
    state.alwaysShow = 2;
  }
  if (state.isRoot != 1) {
    state.isRoot = 2;
  }
  if (state.isFrame != 1) {
    state.isFrame = 2;
  }
  if (state.keepAlive != 1) {
    state.keepAlive = 2;
  }
  if (state.hidden != 1) {
    state.hidden = 2;
  }
  if (state.affix != 1) {
    state.affix = 2;
  }
  if (state.status != 1) {
    state.status = 2;
  }
  return state;
}

// 字典数据选项
export const options = ref({
  sys_menu_types: [] as Option[],
  sys_menu_component: [] as Option[],
  sys_normal_disable: [] as Option[],
  sys_switch: [] as Option[],
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['sys_menu_types', 'sys_menu_component', 'sys_normal_disable', 'sys_switch'],
  }).then((res) => {
    options.value = res;
  });
}
