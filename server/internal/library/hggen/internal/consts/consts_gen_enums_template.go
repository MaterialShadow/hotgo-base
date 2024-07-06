// Copyright GoFrame gf Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package consts

const TemplateGenEnums = `
// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ================================================================================

package {PackageName}

import (
	"github.com/gogf/gf/v2/util/gtag"
)

func init() {
	gtag.SetGlobalEnums({EnumsJson})
}
`
