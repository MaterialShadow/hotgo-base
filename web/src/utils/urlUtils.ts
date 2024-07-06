/**
 * 将对象添加当作参数拼接到URL上面
 * @param baseUrl 需要拼接的url
 * @param obj 参数对象
 * @returns {string} 拼接后的对象
 * 例子:
 *  let obj = {a: '3', b: '4'}
 *  setObjToUrlParams('www.baidu.com', obj)
 *  ==>www.baidu.com?a=3&b=4
 */
export function setObjToUrlParams(baseUrl: string, obj: object): string {
  let parameters = '';
  let url = '';
  for (const key in obj) {
    parameters += key + '=' + encodeURIComponent(obj[key]) + '&';
  }
  parameters = parameters.replace(/&$/, '');
  if (/\?$/.test(baseUrl)) {
    url = baseUrl + parameters;
  } else {
    url = baseUrl.replace(/\/?$/, '?') + parameters;
  }
  return url;
}

export function encodeParams(obj) {
  const arr = [];
  for (const p in obj) {
    // @ts-ignore
    arr.push(encodeURIComponent(p) + '=' + encodeURIComponent(obj[p]));
  }
  return arr.join('&');
}

/**
 * 获取文件后缀
 */
export function getFileExt(fileName: string) {
  if (fileName === undefined || fileName === '') {
    return ``;
  }
  return fileName.substring(fileName.lastIndexOf('.') + 1);
}

/**
 * 获取当访问的url，不含参数
 */
export function getNowUrl(): string {
  const w = window.location;
  return w.protocol + '//' + w.host + w.pathname;
}
