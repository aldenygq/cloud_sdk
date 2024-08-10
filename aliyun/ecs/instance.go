package ecs

import (
	ecs20140526  "github.com/alibabacloud-go/ecs-20140526/v4/client"
	util  "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

func (ecs *AliyunEcsClient) EcsListByRegion(region string) (*ecs20140526.DescribeInstancesResponse,error) {
	describeInstancesRequest := &ecs20140526.DescribeInstancesRequest{
		RegionId: tea.String(region),
	}
	runtime := &util.RuntimeOptions{}
	// 复制代码运行请自行打印 API 的返回值
	resp, err := ecs.Client.DescribeInstancesWithOptions(describeInstancesRequest, runtime)
	if err != nil {
	  return nil,err
	}
	return resp,nil 
}
  