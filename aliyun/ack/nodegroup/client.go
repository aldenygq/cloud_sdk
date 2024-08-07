package nodegroup 

// This file is auto-generated, don't edit it. Thanks.
import (
  //"encoding/json"
  //"strings"
  //"fmt"
  //"os"
  cs20151215  "github.com/alibabacloud-go/cs-20151215/v5/client"
  openapi  "github.com/alibabacloud-go/darabonba-openapi/v2/client"
  //util  "github.com/alibabacloud-go/tea-utils/v2/service"
  "github.com/alibabacloud-go/tea/tea"
)


type AliyunAckClient struct {
	Client *cs20151215.Client
}
// Description:
// 
// 使用AK&SK初始化账号Client
// 
// @return Client
// 
// @throws Exception
func NewAliyunAckClient(ak,sk,endpoint string) (*AliyunAckClient, error) {
  // 工程代码泄露可能会导致 AccessKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考。
  // 建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html。
  config := &openapi.Config{
    // 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_ID。
    AccessKeyId: tea.String(ak),
    // 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_SECRET。
    AccessKeySecret: tea.String(sk),
  }
  // Endpoint 请参考 https://api.aliyun.com/product/CS
  config.Endpoint = tea.String(endpoint)
  result := &cs20151215.Client{}
  result, err := cs20151215.NewClient(config)
  return &AliyunAckClient{
	Client: result,
  },err 
}

