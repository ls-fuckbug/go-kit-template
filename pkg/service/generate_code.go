package service

import (
	"fmt"
	"ls.com/kit/pkg/entity"
	"unicode"
)

// Generate 定义处理函数
func Generate(funcName string, codeType string) (code string) {
	if len(funcName) == 0 || len(codeType) == 0 {
		return "请输入函数名并选择代码类型"
	} else {
		switch codeType {
		case "backend":
			code = generateBackendTemplate([]string{funcName})
		case "service":
			code = generateServiceTemplate([]string{funcName})
		default:
			code = "不支持的codeType"
		}
		return
	}
}

// 首字母小写
func uncapitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	r := []rune(s)
	r[0] = unicode.ToLower(r[0])
	return string(r)
}

func generateBackendTemplate(funcName []string) (code string) {
	code += "$"
	for _, name := range funcName {
		greatHumpName := name
		code += fmt.Sprintf(entity.TemplateServiceImpl, greatHumpName, greatHumpName,
			greatHumpName, greatHumpName)
	}
	code += "$"
	for _, name := range funcName {
		greatHumpName := name
		code += fmt.Sprintf(entity.TemplateServiceInterface, greatHumpName, greatHumpName,
			greatHumpName, greatHumpName)
	}
	code += "$"
	for _, name := range funcName {
		greatHumpName := name
		code += fmt.Sprintf(entity.BackendTemplateEndpointMake, greatHumpName, greatHumpName,
			greatHumpName)
	}
	code += "$"
	for _, name := range funcName {
		greatHumpName := name
		code += fmt.Sprintf(entity.TemplateEndpointMakeCall, greatHumpName, greatHumpName)
	}
	code += "$"
	for _, name := range funcName {
		greatHumpName := name
		smallHumpName := uncapitalize(greatHumpName)
		code += fmt.Sprintf(entity.BackendTemplateTransportHandler, smallHumpName, greatHumpName,
			greatHumpName, greatHumpName)
	}
	code += "$"
	for _, name := range funcName {
		greatHumpName := name
		smallHumpName := uncapitalize(greatHumpName)
		code += fmt.Sprintf(entity.BackendTemplateTransportHandle, smallHumpName)
	}
	code += "$"
	for _, name := range funcName {
		greatHumpName := name
		code += fmt.Sprintf(entity.BackendTemplateTransportImpl, greatHumpName, greatHumpName,
			greatHumpName, greatHumpName)
	}

	return
}

func generateServiceTemplate(funcName []string) (code string) {
	code += "$"
	for _, name := range funcName {
		greatHumpName := name
		code += fmt.Sprintf(entity.TemplateServiceImpl, greatHumpName, greatHumpName,
			greatHumpName, greatHumpName)
	}
	code += "$"
	for _, name := range funcName {
		greatHumpName := name
		code += fmt.Sprintf(entity.TemplateServiceInterface, greatHumpName, greatHumpName,
			greatHumpName, greatHumpName)
	}
	code += "$"
	for _, name := range funcName {
		greatHumpName := name
		code += fmt.Sprintf(entity.ServiceTemplateEndpointMake, greatHumpName, greatHumpName,
			greatHumpName)
	}
	code += ""
	for _, name := range funcName {
		greatHumpName := name
		code += fmt.Sprintf(entity.TemplateEndpointMakeCall, greatHumpName, greatHumpName)
	}
	code += ""
	for _, name := range funcName {
		greatHumpName := name
		smallHumpName := uncapitalize(greatHumpName)
		code += fmt.Sprintf(entity.ServiceTemplateTransportHandler, smallHumpName, greatHumpName,
			greatHumpName, greatHumpName)
	}
	code += "$"
	for _, name := range funcName {
		greatHumpName := name
		smallHumpName := uncapitalize(greatHumpName)
		code += fmt.Sprintf(entity.ServiceTemplateTransportHandle, smallHumpName, smallHumpName)
	}
	code += "$"
	for _, name := range funcName {
		greatHumpName := name
		smallHumpName := uncapitalize(greatHumpName)
		code += fmt.Sprintf(entity.ServiceTemplateTransportImpl, greatHumpName, greatHumpName,
			greatHumpName, smallHumpName, greatHumpName, greatHumpName, greatHumpName,
			greatHumpName, greatHumpName)
	}
	return
}
