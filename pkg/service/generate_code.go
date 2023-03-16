package service

import (
	"encoding/json"
	"fmt"
	"log"
	"ls.com/kit/pkg/entity"
	"strings"
	"unicode"
)

// Generate 定义处理函数
func Generate(funcNames string, codeType string) string {
	funcNameArr := strings.Split(funcNames, "_")
	if len(funcNames) == 0 || len(codeType) == 0 {
		return ""
	} else {
		var codeMap = make(map[string]string)
		switch codeType {
		case "backend":
			generateBackendTemplate(funcNameArr, codeMap)
		case "service":
			generateServiceTemplate(funcNameArr, codeMap)
		default:
			return ""
		}
		generateCommonTemplate(funcNameArr, codeMap)

		jsonStr, err := json.Marshal(codeMap)
		if err != nil {
			log.Println("err", err)
		}
		return string(jsonStr)
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

func generateBackendTemplate(funcName []string, codeMap map[string]string) {
	for _, name := range funcName {
		greatHumpName := name
		smallHumpName := uncapitalize(greatHumpName)

		codeMap["Service层方法定义"] += fmt.Sprintf(entity.BackendTemplateServiceImpl, greatHumpName, greatHumpName,
			greatHumpName, greatHumpName, greatHumpName)
		codeMap["Service层接口声明"] += fmt.Sprintf(entity.BackendTemplateServiceInterface, greatHumpName, greatHumpName,
			greatHumpName, greatHumpName)
		codeMap["Endpoint层make函数定义"] += fmt.Sprintf(entity.BackendTemplateEndpointMake,
			greatHumpName, greatHumpName, greatHumpName)
		codeMap["Transport层handler"] += fmt.Sprintf(entity.BackendTemplateTransportHandler, smallHumpName,
			greatHumpName, greatHumpName, greatHumpName)
		codeMap["Transport层handler调用"] += fmt.Sprintf(entity.BackendTemplateTransportHandle, smallHumpName)
		codeMap["Transport层输入输出"] += fmt.Sprintf(entity.BackendTemplateTransportImpl, greatHumpName, greatHumpName,
			greatHumpName, greatHumpName)
		codeMap["UnitTest单测Service层"] += fmt.Sprintf(entity.TemplateBackendUnitTestServiceFuncImpl, greatHumpName,
			greatHumpName, greatHumpName, greatHumpName, greatHumpName)
	}
	return
}

func generateServiceTemplate(funcName []string, codeMap map[string]string) {
	for _, name := range funcName {
		greatHumpName := name
		smallHumpName := uncapitalize(greatHumpName)

		codeMap["Service层方法定义"] += fmt.Sprintf(entity.ServiceTemplateServiceImpl, greatHumpName, greatHumpName,
			greatHumpName, greatHumpName, greatHumpName)
		codeMap["Service层接口声明"] += fmt.Sprintf(entity.ServiceTemplateServiceInterface, greatHumpName, greatHumpName,
			greatHumpName, greatHumpName)
		codeMap["Endpoint层make函数定义"] += fmt.Sprintf(entity.ServiceTemplateEndpointMake, greatHumpName,
			greatHumpName, greatHumpName)
		codeMap["Transport层结构体成员"] += fmt.Sprintf(entity.ServiceTemplateTransportStructMember,
			smallHumpName)
		codeMap["Transport层handler"] += fmt.Sprintf(entity.ServiceTemplateTransportHandler, smallHumpName,
			greatHumpName, greatHumpName, greatHumpName)
		codeMap["Transport层handler调用"] += fmt.Sprintf(entity.ServiceTemplateTransportHandle, smallHumpName,
			smallHumpName)
		codeMap["Transport层输入输出"] += fmt.Sprintf(entity.ServiceTemplateTransportImpl, greatHumpName,
			greatHumpName, greatHumpName, smallHumpName, greatHumpName, greatHumpName, greatHumpName,
			greatHumpName, greatHumpName)
		codeMap["UnitTest单测Service层"] += fmt.Sprintf(entity.TemplateServiceUnitTestServiceFuncImpl, greatHumpName,
			greatHumpName, greatHumpName, greatHumpName, greatHumpName, greatHumpName)
	}
	return
}

func generateCommonTemplate(funcName []string, codeMap map[string]string) {
	for _, name := range funcName {
		greatHumpName := name

		codeMap["Endpoint层成员"] += fmt.Sprintf(entity.TemplateEndpointStructMember, greatHumpName)
		codeMap["Endpoint层make函数调用"] += fmt.Sprintf(entity.TemplateEndpointMakeCall, greatHumpName, greatHumpName)
		codeMap["Repo层Service调用函数"] += fmt.Sprintf(entity.TemplateRepoServiceCallImpl, greatHumpName, greatHumpName,
			greatHumpName, greatHumpName, greatHumpName, greatHumpName, greatHumpName)
		codeMap["Repo层base函数"] += fmt.Sprintf(entity.TemplateRepoSampleFuncImpl, greatHumpName, greatHumpName,
			greatHumpName)
	}
	return
}
