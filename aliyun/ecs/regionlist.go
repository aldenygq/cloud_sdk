package ecs

import (
	//"encoding/json"
	//"strings"
	//"fmt"
	//"os"
	ecs20140526  "github.com/alibabacloud-go/ecs-20140526/v4/client"
	//openapi  "github.com/alibabacloud-go/darabonba-openapi/client"
	util  "github.com/alibabacloud-go/tea-utils/service"
	//"github.com/alibabacloud-go/tea/tea"
)

func (ecs *AliyunEcsClient) RegionList() (*ecs20140526.DescribeRegionsResponse,error) {
	describeRegionsRequest := &ecs20140526.DescribeRegionsRequest{}
	runtime := &util.RuntimeOptions{}
	// 复制代码运行请自行打印 API 的返回值
	resp, err := ecs.Client.DescribeRegionsWithOptions(describeRegionsRequest, runtime)
	if err != nil {
	  return nil,err
	}
	return resp,nil 
} 