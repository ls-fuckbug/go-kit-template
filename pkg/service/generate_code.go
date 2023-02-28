package service

import (
	"fmt"
	"ls.com/kit/pkg/entity"
)

// Generate 定义处理函数
func Generate(funcName string, codeType string) (code string) {
	if len(funcName) == 0 || len(codeType) == 0 {
		return "请输入函数名并选择代码类型"
	} else {
		switch codeType {
		case "backend":
			code = generateBackendTemplate(funcName)
		case "service":
			code = generateServiceTemplate(funcName)
		default:
			code = "不支持的codeType"
		}
		return
	}
}

func generateBackendTemplate(funcName string) (code string) {
	code = fmt.Sprintf(entity.BackendTemplateService, funcName)
	return
}

func generateServiceTemplate(funcName string) (code string) {
	code = fmt.Sprintf(entity.ServiceTemplateService, funcName)
	return
}
