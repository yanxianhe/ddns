// This file is auto-generated, don't edit it. Thanks.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v4/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	console "github.com/alibabacloud-go/tea-console/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

/**
 * 使用AK&SK初始化账号Client
 * @return Client
 * @throws Exception
 */
func CreateClient () (_result *alidns20150109.Client, _err error) {
  // 工程代码泄露可能会导致 AccessKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考。
  // 建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html。
  config := &openapi.Config{
    // 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_ID。
    AccessKeyId: tea.String(os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID")),
    // 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_SECRET。
    AccessKeySecret: tea.String(os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET")),
  }
  // Endpoint 请参考 https://api.aliyun.com/product/Alidns
  config.Endpoint = tea.String("alidns.cn-hangzhou.aliyuncs.com")
  _result = &alidns20150109.Client{}
  _result, _err = alidns20150109.NewClient(config)
  return _result, _err
}

func _main (RecordId, RR, Value string) (_err error) {
  client, _err := CreateClient()
  if _err != nil {
    return _err
  }

  updateDomainRecordRequest := &alidns20150109.UpdateDomainRecordRequest{
    RecordId: tea.String(RecordId),
    RR: tea.String(RR),
    Type: tea.String("A"),
    Value: tea.String(Value), // 
  }
  runtime := &util.RuntimeOptions{}
  tryErr := func()(_e error) {
    defer func() {
      if r := tea.Recover(recover()); r != nil {
        _e = r
      }
    }()
    resp, _err := client.UpdateDomainRecordWithOptions(updateDomainRecordRequest, runtime)
    if _err != nil {
      return _err
    }

    console.Log(util.ToJSONString(resp))

    return nil
  }()

  if tryErr != nil {
    var error = &tea.SDKError{}
    if _t, ok := tryErr.(*tea.SDKError); ok {
      error = _t
    } else {
      error.Message = tea.String(tryErr.Error())
    }
    // 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
    // 错误 message
    fmt.Println(tea.StringValue(error.Message))
    // 诊断地址
    var data interface{}
    d := json.NewDecoder(strings.NewReader(tea.StringValue(error.Data)))
    d.Decode(&data)
    if m, ok := data.(map[string]interface{}); ok {
      recommend, _ := m["Recommend"]
      fmt.Println(recommend)
    }
    _, _err = util.AssertAsString(error.Message)
    if _err != nil {
      return _err
    }
  }
  return _err
}


func main() {
  var recordId, domainName, ipValue string
  flag.StringVar(&recordId, "recordId", "", "RecordId for the domain record")
  flag.StringVar(&domainName, "domainName", "", "RR for the domain record")
	flag.StringVar(&ipValue, "ipValue", "", "Value for the domain record")
  flag.Parse()
  if recordId == "" || domainName == "" || ipValue == "" {
		fmt.Println("Usage: ddnsClient --recordId <RecordId> --domainName <RR> --ipValue <Value>")
		os.Exit(1)
	}
  err := _main(recordId,domainName,ipValue)
  if err != nil {
    panic(err)
  }
}
