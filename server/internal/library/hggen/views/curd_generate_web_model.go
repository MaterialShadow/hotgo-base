// Package views
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package views

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/library/dict"
	"hotgo/internal/model/input/sysin"
	"hotgo/utility/convert"
)

type StateItem struct {
	Name         string
	DefaultValue interface{}
	Dc           string
}

func (l *gCurd) webModelTplData(ctx context.Context, in *CurdPreviewInput) (data g.Map, err error) {
	data = make(g.Map)
	data["stateItems"] = l.generateWebModelStateItems(ctx, in)
	data["rules"] = l.generateWebModelRules(ctx, in)
	data["formSchema"] = l.generateWebModelFormSchema(ctx, in)
	if data["columns"], err = l.generateWebModelColumns(ctx, in); err != nil {
		return nil, err
	}

	// 根据表单生成情况，按需导包
	data["import"] = l.generateWebModelImport(ctx, in)
	return
}

func (l *gCurd) generateWebModelImport(ctx context.Context, in *CurdPreviewInput) string {
	importBuffer := bytes.NewBuffer(nil)

	importBuffer.WriteString("import { h, ref } from 'vue';\n")

	// 导入基础组件
	if len(in.options.Step.ImportModel.NaiveUI) > 0 {
		importBuffer.WriteString("import " + ImportWebMethod(in.options.Step.ImportModel.NaiveUI) + " from 'naive-ui';\n")
	}

	importBuffer.WriteString("import { cloneDeep } from 'lodash-es';\n")

	// 导入表单搜索
	if in.options.Step.HasSearchForm {
		importBuffer.WriteString("import { FormSchema } from '@/components/Form';\n")
	}

	// 导入字典选项
	if in.options.DictOps.Has {
		importBuffer.WriteString("import { Dicts } from '@/api/dict/dict';\n")
	}

	// 导入工具类
	if len(in.options.Step.ImportModel.UtilsIs) > 0 {
		importBuffer.WriteString("import " + ImportWebMethod(in.options.Step.ImportModel.UtilsIs) + " from '@/utils/is';\n")
	}

	if len(in.options.Step.ImportModel.UtilsUrl) > 0 {
		importBuffer.WriteString("import " + ImportWebMethod(in.options.Step.ImportModel.UtilsUrl) + " from '@/utils/urlUtils';\n")
	}

	if len(in.options.Step.ImportModel.UtilsDate) > 0 {
		importBuffer.WriteString("import " + ImportWebMethod(in.options.Step.ImportModel.UtilsDate) + " from '@/utils/dateUtil';\n")
	}

	if in.options.Step.HasRulesValidator {
		importBuffer.WriteString("import { validate } from '@/utils/validateUtil';\n")
	}

	if len(in.options.Step.ImportModel.UtilsHotGo) > 0 {
		importBuffer.WriteString("import " + ImportWebMethod(in.options.Step.ImportModel.UtilsHotGo) + " from '@/utils/hotgo';\n")
	}

	if len(in.options.Step.ImportModel.UtilsIndex) > 0 {
		importBuffer.WriteString("import " + ImportWebMethod(in.options.Step.ImportModel.UtilsIndex) + " from '@/utils';\n")
	}

	// 导入api
	var importApiMethod []string
	if in.options.Step.HasSwitch {
		importApiMethod = append(importApiMethod, "Switch")
	}
	if in.options.Step.IsTreeTable {
		importApiMethod = append(importApiMethod, "TreeOption")
	}
	if len(importApiMethod) > 0 {
		importBuffer.WriteString("import " + ImportWebMethod(importApiMethod) + " from '" + in.options.ImportWebApi + "';\n")
	}

	if in.options.Step.HasSwitch {
		importBuffer.WriteString("import { usePermission } from '@/hooks/web/usePermission';\n")
		importBuffer.WriteString("const { hasPermission } = usePermission();\n")
		importBuffer.WriteString("const $message = window['$message'];\n")
	}
	return importBuffer.String()
}

func (l *gCurd) generateWebModelStateItems(ctx context.Context, in *CurdPreviewInput) (items []*StateItem) {
	for _, field := range in.masterFields {
		var value = field.DefaultValue
		if value == nil {
			value = "null"
		}
		if value == "" {
			value = `''`
		}

		// 选项组件默认值调整
		if gconv.Int(value) == 0 && IsSelectFormMode(field.FormMode) {
			value = "null"
		}

		if field.Name == "status" {
			value = 1
		}
		if field.FormMode == FormModeSwitch {
			value = 2
		}
		if field.FormMode == FormModeInputDynamic {
			value = "[]"
		}
		items = append(items, &StateItem{
			Name:         field.TsName,
			DefaultValue: value,
			Dc:           field.Dc,
		})

		// 查询用户摘要
		if field.IsList && in.options.Step.HasHookMemberSummary && IsMemberSummaryField(field.Name) {
			items = append(items, &StateItem{
				Name:         field.TsName + "Summa?: null | MemberSumma",
				DefaultValue: "null",
				Dc:           field.Dc + "摘要信息",
			})
		}
	}
	return
}

func (l *gCurd) generateWebModelDictOptions(ctx context.Context, in *CurdPreviewInput) error {
	type DictType struct {
		Id   int64  `json:"id"`
		Type string `json:"type"`
	}

	var (
		dictTypeIds         []int64
		dictTypeList        []*DictType
		builtinDictTypeIds  []int64
		builtinDictTypeList []*DictType
	)

	for _, field := range in.masterFields {
		if field.DictType > 0 {
			dictTypeIds = append(dictTypeIds, field.DictType)
		}

		if field.DictType < 0 {
			builtinDictTypeIds = append(builtinDictTypeIds, field.DictType)
		}
	}

	dictTypeIds = convert.UniqueSlice(dictTypeIds)
	builtinDictTypeIds = convert.UniqueSlice(builtinDictTypeIds)

	if len(dictTypeIds) == 0 && len(builtinDictTypeIds) == 0 {
		return nil
	}

	if len(dictTypeIds) > 0 {
		err := g.Model("sys_dict_type").Ctx(ctx).
			Fields("id", "type").
			WhereIn("id", dictTypeIds).
			Scan(&dictTypeList)
		if err != nil {
			return err
		}
	}

	if len(builtinDictTypeIds) > 0 {
		for _, id := range builtinDictTypeIds {
			typ, err := dict.GetTypeById(ctx, id)
			if err != nil && !errors.Is(err, dict.NotExistKeyError) {
				return err
			}
			if len(typ) > 0 {
				row := new(DictType)
				row.Id = id
				row.Type = typ
				builtinDictTypeList = append(builtinDictTypeList, row)
			}
		}
	}

	if len(dictTypeList) == 0 && len(builtinDictTypeList) == 0 {
		return nil
	}

	if len(builtinDictTypeList) > 0 {
		dictTypeList = append(dictTypeList, builtinDictTypeList...)
	}

	in.options.DictOps.Has = true

	// 导入选项包
	in.options.Step.ImportModel.UtilsHotGo = append(in.options.Step.ImportModel.UtilsHotGo, "Option")

	for _, v := range dictTypeList {
		// 字段映射字典
		for _, field := range in.masterFields {
			if field.DictType != 0 && v.Id == field.DictType {
				in.options.dictMap[field.TsName] = v.Type
				in.options.DictOps.Schemas = append(in.options.DictOps.Schemas, &OptionsSchemasField{
					Field: field.TsName,
					Type:  v.Type,
				})
			}
		}
		in.options.DictOps.Types = append(in.options.DictOps.Types, v.Type)
	}
	return nil
}

func (l *gCurd) generateWebModelRules(ctx context.Context, in *CurdPreviewInput) string {
	buffer := bytes.NewBuffer(nil)
	buffer.WriteString("export const rules = {\n")
	for _, field := range in.masterFields {
		if !field.IsEdit || (!field.Required && (field.FormRole == "" || field.FormRole == FormRoleNone)) {
			continue
		}

		in.options.Step.HasRules = true
		if field.FormRole == "" || field.FormRole == FormRoleNone || field.FormRole == "required" {
			buffer.WriteString(fmt.Sprintf("  %s: {\n    required: %v,\n    trigger: ['blur', 'input'],\n    type: '%s',\n    message: '请输入%s',\n  },\n", field.TsName, field.Required, field.TsType, field.Dc))
		} else {
			in.options.Step.HasRulesValidator = true
			buffer.WriteString(fmt.Sprintf("  %s: {\n    required: %v,\n    trigger: ['blur', 'input'],\n    type: '%s',\n    validator: validate.%v,\n  },\n", field.TsName, field.Required, field.TsType, field.FormRole))
		}
	}
	buffer.WriteString("};\n")
	return buffer.String()
}

func (l *gCurd) generateWebModelFormSchema(ctx context.Context, in *CurdPreviewInput) string {
	buffer := bytes.NewBuffer(nil)
	buffer.WriteString("export const schemas = ref<FormSchema[]>([\n")

	// 主表
	l.generateWebModelFormSchemaEach(buffer, in.masterFields, in)

	// 关联表
	if len(in.options.Join) > 0 {
		for _, v := range in.options.Join {
			if !isEffectiveJoin(v) {
				continue
			}
			l.generateWebModelFormSchemaEach(buffer, v.Columns, in)
		}
	}

	buffer.WriteString("]);\n")
	return buffer.String()
}

func (l *gCurd) generateWebModelFormSchemaEach(buffer *bytes.Buffer, fields []*sysin.GenCodesColumnListModel, in *CurdPreviewInput) {
	for _, field := range fields {
		if !field.IsQuery {
			continue
		}
		in.options.Step.HasSearchForm = true

		// 查询用户摘要
		if field.IsQuery && in.options.Step.HasQueryMemberSummary && IsMemberSummaryField(field.Name) {
			buffer.WriteString(fmt.Sprintf("  {\n    field: '%s',\n    component: '%s',\n    label: '%s',\n    componentProps: {\n      placeholder: '请输入ID|用户名|姓名|手机号',\n      onUpdateValue: (e: any) => {\n        console.log(e);\n      },\n    },\n  },\n", field.TsName, "NInput", field.Dc))
			continue
		}

		var (
			defaultComponent = fmt.Sprintf("  {\n    field: '%s',\n    component: '%s',\n    label: '%s',\n    componentProps: {\n      placeholder: '请输入%s',\n      onUpdateValue: (e: any) => {\n        console.log(e);\n      },\n    },\n  },\n", field.TsName, "NInput", field.Dc, field.Dc)
			component        string
		)

		// 这里根据编辑表单组件来进行推断，如果没有则使用默认input，这可能会导致和查询条件所需参数不符的情况
		switch field.FormMode {
		case FormModeInput, FormModeInputTextarea, FormModeInputEditor:
			component = defaultComponent

		case FormModeInputNumber:
			component = fmt.Sprintf("  {\n    field: '%s',\n    component: '%s',\n    label: '%s',\n    componentProps: {\n      placeholder: '请输入%s',\n      onUpdateValue: (e: any) => {\n        console.log(e);\n      },\n    },\n  },\n", field.TsName, "NInputNumber", field.Dc, field.Dc)

		case FormModeDate:
			component = fmt.Sprintf("  {\n    field: '%s',\n    component: '%s',\n    label: '%s',\n    componentProps: {\n      type: '%s',\n      clearable: true,\n      shortcuts: %s,\n      onUpdateValue: (e: any) => {\n        console.log(e);\n      },\n    },\n  },\n", field.TsName, "NDatePicker", field.Dc, "date", "defShortcuts()")
			in.options.Step.ImportModel.UtilsDate = append(in.options.Step.ImportModel.UtilsDate, "defShortcuts")

		case FormModeDateRange:
			component = fmt.Sprintf("  {\n    field: '%s',\n    component: '%s',\n    label: '%s',\n    componentProps: {\n      type: '%s',\n      clearable: true,\n      shortcuts: %s,\n      onUpdateValue: (e: any) => {\n        console.log(e);\n      },\n    },\n  },\n", field.TsName, "NDatePicker", field.Dc, "daterange", "defRangeShortcuts()")
			in.options.Step.ImportModel.UtilsDate = append(in.options.Step.ImportModel.UtilsDate, "defRangeShortcuts")

		case FormModeTime:
			component = fmt.Sprintf("  {\n    field: '%s',\n    component: '%s',\n    label: '%s',\n    componentProps: {\n      type: '%s',\n      clearable: true,\n      shortcuts: %s,\n      onUpdateValue: (e: any) => {\n        console.log(e);\n      },\n    },\n  },\n", field.TsName, "NDatePicker", field.Dc, "datetime", "defShortcuts()")
			in.options.Step.ImportModel.UtilsDate = append(in.options.Step.ImportModel.UtilsDate, "defShortcuts")

		case FormModeTimeRange:
			component = fmt.Sprintf("  {\n    field: '%s',\n    component: '%s',\n    label: '%s',\n    componentProps: {\n      type: '%s',\n      clearable: true,\n      shortcuts: %s,\n      onUpdateValue: (e: any) => {\n        console.log(e);\n      },\n    },\n  },\n", field.TsName, "NDatePicker", field.Dc, "datetimerange", "defRangeShortcuts()")
			in.options.Step.ImportModel.UtilsDate = append(in.options.Step.ImportModel.UtilsDate, "defRangeShortcuts")

		case FormModeSwitch:
			fallthrough
		case FormModeRadio:
			component = fmt.Sprintf("  {\n    field: '%s',\n    component: '%s',\n    label: '%s',\n    giProps: {\n      //span: 24,\n    },\n    componentProps: {\n      options: [],\n      onUpdateChecked: (e: any) => {\n        console.log(e);\n      },\n    },\n  },\n", field.TsName, "NRadioGroup", field.Dc)

		case FormModeCheckbox:
			component = fmt.Sprintf("  {\n    field: '%s',\n    component: '%s',\n    label: '%s',\n    giProps: {\n      span: 1,\n    },\n    componentProps: {\n      placeholder: '请选择%s',\n      options: [],\n      onUpdateChecked: (e: any) => {\n        console.log(e);\n      },\n    },\n  },\n", field.TsName, "NCheckbox", field.Dc, field.Dc)

		case FormModeSelect:
			component = fmt.Sprintf("  {\n    field: '%s',\n    component: '%s',\n    label: '%s',\n    defaultValue: null,\n    componentProps: {\n      placeholder: '请选择%s',\n      options: [],\n      onUpdateValue: (e: any) => {\n        console.log(e);\n      },\n    },\n  },\n", field.TsName, "NSelect", field.Dc, field.Dc)

		case FormModeSelectMultiple:
			component = fmt.Sprintf("  {\n    field: '%s',\n    component: '%s',\n    label: '%s',\n    defaultValue: null,\n    componentProps: {\n      multiple: true,\n      placeholder: '请选择%s',\n      options: [],\n      onUpdateValue: (e: any) => {\n        console.log(e);\n      },\n    },\n  },\n", field.TsName, "NSelect", field.Dc, field.Dc)

		default:
			component = defaultComponent
		}

		buffer.WriteString(component)
	}
}

func (l *gCurd) generateWebModelColumns(ctx context.Context, in *CurdPreviewInput) (string, error) {
	buffer := bytes.NewBuffer(nil)
	buffer.WriteString("export const columns = [\n")

	// 主表
	if err := l.generateWebModelColumnsEach(buffer, in, in.masterFields); err != nil {
		return "", err
	}

	// 关联表
	if len(in.options.Join) > 0 {
		for _, v := range in.options.Join {
			if !isEffectiveJoin(v) {
				continue
			}
			if err := l.generateWebModelColumnsEach(buffer, in, v.Columns); err != nil {
				return "", err
			}
		}
	}

	buffer.WriteString("];\n")
	return buffer.String(), nil
}

func (l *gCurd) generateWebModelColumnsEach(buffer *bytes.Buffer, in *CurdPreviewInput, fields []*sysin.GenCodesColumnListModel) (err error) {
	for _, field := range fields {
		if !field.IsList {
			continue
		}
		var (
			defaultComponent = fmt.Sprintf("  {\n    title: '%s',\n    key: '%s',\n    align: '%v',\n    width: %v,\n  },\n", field.Dc, field.TsName, field.Align, field.Width)
			component        string
		)

		// 查询用户摘要
		if in.options.Step.HasHookMemberSummary && IsMemberSummaryField(field.Name) {
			buffer.WriteString(fmt.Sprintf("  {\n    title: '%v',\n    key: '%v',\n    align: '%v',\n    width: %v,\n    render(row) {\n      return renderPopoverMemberSumma(row.%vSumma);\n    },\n  },\n", field.Dc, field.TsName, field.Align, field.Width, field.TsName))
			in.options.Step.ImportModel.UtilsIndex = append(in.options.Step.ImportModel.UtilsIndex, []string{"renderPopoverMemberSumma", "MemberSumma"}...)
			continue
		}

		// 这里根据编辑表单组件来进行推断，如果没有则使用默认input，这可能会导致和查询条件所需参数不符的情况
		switch field.FormMode {
		case FormModeDate:
			component = fmt.Sprintf("  {\n    title: '%s',\n    key: '%s',\n    align: '%v',\n    width: %v,\n    render(row) {\n      return formatToDate(row.%s);\n    },\n  },\n", field.Dc, field.TsName, field.Align, field.Width, field.TsName)
			in.options.Step.ImportModel.UtilsDate = append(in.options.Step.ImportModel.UtilsDate, "formatToDate")

		case FormModeRadio:
			fallthrough
		case FormModeSelect:
			if g.IsEmpty(in.options.dictMap[field.TsName]) {
				err = gerror.Newf("设置单选下拉框选项时，必须选择字典类型，字段名称:%v", field.Name)
				return
			}
			component = fmt.Sprintf("  {\n    title: '%s',\n    key: '%s',\n    align: '%v',\n    width: %v,\n    render(row) {\n      if (isNullObject(row.%s)) {\n        return ``;\n      }\n      return h(\n        NTag,\n        {\n          style: {\n            marginRight: '6px',\n          },\n          type: getOptionTag(options.value.%s, row.%s),\n          bordered: false,\n        },\n        {\n          default: () => getOptionLabel(options.value.%s, row.%s),\n        }\n      );\n    },\n  },\n", field.Dc, field.TsName, field.Align, field.Width, field.TsName, in.options.dictMap[field.TsName], field.TsName, in.options.dictMap[field.TsName], field.TsName)
			in.options.Step.ImportModel.NaiveUI = append(in.options.Step.ImportModel.NaiveUI, "NTag")
			in.options.Step.ImportModel.UtilsIs = append(in.options.Step.ImportModel.UtilsIs, "isNullObject")
			in.options.Step.ImportModel.UtilsHotGo = append(in.options.Step.ImportModel.UtilsHotGo, []string{"getOptionLabel", "getOptionTag"}...)

		case FormModeSelectMultiple:
			if g.IsEmpty(in.options.dictMap[field.TsName]) {
				err = gerror.Newf("设置多选下拉框选项时，必须选择字典类型，字段名称:%v", field.Name)
				return
			}
			component = fmt.Sprintf("  {\n    title: '%s',\n    key: '%s',\n    align: '%v',\n    width: %v,\n    render(row) {\n      if (isNullObject(row.%s) || !isArray(row.%s)) {\n        return ``;\n      }\n      return row.%s.map((tagKey) => {\n        return h(\n          NTag,\n          {\n            style: {\n              marginRight: '6px',\n            },\n            type: getOptionTag(options.value.%s, tagKey),\n            bordered: false,\n          },\n          {\n            default: () => getOptionLabel(options.value.%s, tagKey),\n          }\n        );\n      });\n    },\n  },\n", field.Dc, field.TsName, field.Align, field.Width, field.TsName, field.TsName, field.TsName, in.options.dictMap[field.TsName], in.options.dictMap[field.TsName])
			in.options.Step.ImportModel.NaiveUI = append(in.options.Step.ImportModel.NaiveUI, "NTag")
			in.options.Step.ImportModel.UtilsIs = append(in.options.Step.ImportModel.UtilsIs, "isNullObject")
			in.options.Step.ImportModel.UtilsHotGo = append(in.options.Step.ImportModel.UtilsHotGo, []string{"getOptionLabel", "getOptionTag"}...)

		case FormModeUploadImage:
			component = fmt.Sprintf("  {\n    title: '%s',\n    key: '%s',\n    align: '%v',\n    width: %v,\n    render(row) {\n      return h(%s, {\n        width: 32,\n        height: 32,\n        src: row.%s,\n        fallbackSrc: errorImg,\n        onError: errorImg,\n        style: {\n          width: '32px',\n          height: '32px',\n          'max-width': '100%%',\n          'max-height': '100%%',\n        },\n      });\n    },\n  },\n", field.Dc, field.TsName, field.Align, field.Width, "NImage", field.TsName)
			in.options.Step.ImportModel.NaiveUI = append(in.options.Step.ImportModel.NaiveUI, "NImage")
			in.options.Step.ImportModel.UtilsHotGo = append(in.options.Step.ImportModel.UtilsHotGo, "errorImg")

		case FormModeUploadImages:
			component = fmt.Sprintf("  {\n    title: '%s',\n    key: '%s',\n    align: '%v',\n    width: %v,\n    render(row) {\n      if (isNullObject(row.%s)) {\n        return ``;\n      }\n      return row.%s.map((image) => {\n        return h(%s, {\n          width: 32,\n          height: 32,\n          src: image,\n        onError: errorImg,\n          style: {\n            width: '32px',\n            height: '32px',\n            'max-width': '100%%',\n            'max-height': '100%%',\n            'margin-left': '2px',\n          },\n        });\n      });\n    },\n  },\n", field.Dc, field.TsName, field.Align, field.Width, field.TsName, field.TsName, "NImage")
			in.options.Step.ImportModel.NaiveUI = append(in.options.Step.ImportModel.NaiveUI, "NImage")
			in.options.Step.ImportModel.UtilsIs = append(in.options.Step.ImportModel.UtilsIs, "isArray")
			in.options.Step.ImportModel.UtilsHotGo = append(in.options.Step.ImportModel.UtilsHotGo, "errorImg")

		case FormModeUploadFile:
			component = fmt.Sprintf("  {\n    title: '%s',\n    key: '%s',\n    align: '%v',\n    width: %v,\n    render(row) {\n      if (row.%s === '') {\n        return ``;\n      }\n      return h(\n        %s,\n        {\n          size: 'small',\n        },\n        {\n          default: () => getFileExt(row.%s),\n        }\n      );\n    },\n  },\n", field.Dc, field.TsName, field.Align, field.Width, field.TsName, "NAvatar", field.TsName)
			in.options.Step.ImportModel.NaiveUI = append(in.options.Step.ImportModel.NaiveUI, "NAvatar")
			in.options.Step.ImportModel.UtilsUrl = append(in.options.Step.ImportModel.UtilsUrl, "getFileExt")

		case FormModeUploadFiles:
			component = fmt.Sprintf("  {\n    title: '%s',\n    key: '%s',\n    align: '%v',\n    width: %v,\n    render(row) {\n      if (isNullObject(row.%s)) {\n        return ``;\n      }\n      return row.%s.map((attachfile) => {\n        return h(\n          %s,\n          {\n            size: 'small',\n            style: {\n              'margin-left': '2px',\n            },\n          },\n          {\n            default: () => getFileExt(attachfile),\n          }\n        );\n      });\n    },\n  },\n", field.Dc, field.TsName, field.Align, field.Width, field.TsName, field.TsName, "NAvatar")
			in.options.Step.ImportModel.NaiveUI = append(in.options.Step.ImportModel.NaiveUI, "NAvatar")
			in.options.Step.ImportModel.UtilsIs = append(in.options.Step.ImportModel.UtilsIs, "isNullObject")
			in.options.Step.ImportModel.UtilsUrl = append(in.options.Step.ImportModel.UtilsUrl, "getFileExt")

		case FormModeSwitch:
			component = fmt.Sprintf("  {\n    title: '%s',\n    key: '%s',\n    align: '%v',\n    width: %v,\n    render(row) {\n      return h(%s, {\n        value: row.%s === 1,\n        checked: '开启',\n        unchecked: '关闭',\n        disabled: !hasPermission(['%s']),\n        onUpdateValue: function (e) {\n          console.log('onUpdateValue e:' + JSON.stringify(e));\n          row.%s = e ? 1 : 2;\n          Switch({ %s: row.%s, key: '%s', value: row.%s }).then((_res) => {\n            $message.success('操作成功');\n          });\n        },\n      });\n    },\n  },\n", field.Dc, field.TsName, field.Align, field.Width, "NSwitch", field.TsName, "/"+in.options.ApiPrefix+"/switch", field.TsName, in.pk.TsName, in.pk.TsName, convert.CamelCaseToUnderline(field.TsName), field.TsName)
			in.options.Step.ImportModel.NaiveUI = append(in.options.Step.ImportModel.NaiveUI, "NSwitch")

		case FormModeRate:
			component = fmt.Sprintf("  {\n    title: '%s',\n    key: '%s',\n    align: '%v',\n    width: %v,\n    render(row) {\n      return h(%s, {\n        allowHalf: true,\n        readonly: true,\n        defaultValue: row.%s,\n      });\n    },\n  },\n", field.Dc, field.TsName, field.Align, field.Width, "NRate", field.TsName)
			in.options.Step.ImportModel.NaiveUI = append(in.options.Step.ImportModel.NaiveUI, "NRate")

		default:
			component = defaultComponent
		}

		buffer.WriteString(component)
	}
	return
}
