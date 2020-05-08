package modal

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

//GenValidateCode 生成随机数
func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

//Verification 获取验证码
func Verification(num string) string {
	shu := 5
	code := GenValidateCode(shu)

	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "LTAI4FhwZpHwHbyK8Ny8ZsDJ", "qrkQzeVX3JsJdvvbAc8vi3KzeEupBc")

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = num
	request.SignName = "number"
	request.TemplateCode = "SMS_185842473"
	request.TemplateParam = "{\"code\":\"" + code + "\"}"

	response, err := client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)

	return code
}
