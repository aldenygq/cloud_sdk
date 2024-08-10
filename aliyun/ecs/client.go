package ecs

import (
  //"encoding/json"
  //"strings"
  //"fmt"
  //"os"
  ecs20140526  "github.com/alibabacloud-go/ecs-20140526/v4/client"
  openapi  "github.com/alibabacloud-go/darabonba-openapi/client"
  //util  "github.com/alibabacloud-go/tea-utils/v2/service"
  "github.com/alibabacloud-go/tea/tea"
)
type AliyunEcsClient struct {
	Client *ecs20140526.Client
}

func NewAliyunEcsClient(ak,sk,endpoint string) (*AliyunEcsClient, error) {
	config := &openapi.Config{
		// 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_ID。
		AccessKeyId: tea.String(ak),
		// 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_SECRET。
		AccessKeySecret: tea.String(sk),
	}
	  // Endpoint 请参考 https://api.aliyun.com/product/Ecs
	config.Endpoint = tea.String(endpoint)
	result := &ecs20140526.Client{}
	result, err := ecs20140526.NewClient(config)
	if err != nil {
		return nil,err 
	  }
	  return &AliyunEcsClient{
		Client: result,
	  },err 
}
