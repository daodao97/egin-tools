package gen

const SimpleApi = `
r.Handle("{{ .method }}", "{{ .path }}", func(ctx *gin.Context) {
	result, code, err := controller.{{ .entity }}{}.{{ .funcName }}(ctx)
	egin.Response(ctx, result, code, err)
}{{ .middleware }})
`

const ApiWithParam = `
r.Handle("{{ .method }}", "{{ .path }}", func () func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var params controller.{{ .paramsStruct }}
		errs := utils.Validated(ctx, &params)
		if errs != nil {
			egin.Fail(ctx, consts.ErrorParam, strings.Join(errs, "\n"))
			return
		}
		result, code, err := controller.{{ .entity }}{}.{{ .funcName }}(ctx, params)
		egin.Response(ctx, result, code, err)
	}
}(){{ .middleware }})
`

const RouteFile = `
// ****************************
// 该文件为系统生成, 请勿随意更改
// ****************************
package routes

import (
	"strings"

	"github.com/daodao97/egin"
	"github.com/daodao97/egin/consts"
	"github.com/daodao97/egin/middleware"
	"github.com/daodao97/egin/utils"
	"github.com/gin-gonic/gin"

	"skeleton/controller"
)

func Reg{{ .entity }}Router(r *gin.Engine) {
	{{ range $index, $value := .handles }} {{$value}} {{ end }}
}

{{ if .hasCustomValidateFuncs }}
func RegUserCustomValidateFunc() {
	utils.RegCustomValidateFuncs([]utils.CustomValidateFunc{
		{{ range $index, $value := .customValidateFuncs }} controller.{{$value}}, {{ end }}
	})
}
{{ end }}
`

const RouteExport = `
// ****************************
// 该文件为系统生成, 请勿随意更改
// ****************************
package config

import (
	"github.com/gin-gonic/gin"

	"skeleton/config/routes"
)

func RegRouter(r *gin.Engine) {
	{{ range $index, $value := .list }}
		{{ $value }}
	{{ end }}
}
`

const Entity = `
package model

import (
	"github.com/daodao97/egin/db"
)

{{ if .tableComment }} // {{ .tableComment }} {{ end }}
type {{ .entityName }}Entity struct {
	{{- range $index, $value := .fieldList }}
		{{ $value.Name }} {{ $value.Type }} {{ $.backquote }}json:"{{ $value.Field }}"{{ if $value.Comment }} comment:"{{ $value.Comment }}"{{ end }}{{ $.backquote }}
	{{- end }}
}

type {{ .entityName }}Model struct {
	db.BaseModel
}

var {{ .entityName }} {{ .entityName }}Model

func init() {
	{{ .entityName }} = *New{{ .entityName }}Model()
}

func New{{ .entityName }}Model() *{{ .entityName }}Model {
	return &{{ .entityName }}Model{
		BaseModel: db.BaseModel{
			Connection: "{{ .connection }}",
			Table:      "{{ .table }}",
			{{- if .fakeDel }}
			FakeDelete: true,
			FakeDelKey: "is_deleted",
			{{- end }}
		},
	}
}
`
