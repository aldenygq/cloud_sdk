package ecs

import (
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v4/client"
	//openapi  "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

func (ecs *AliyunEcsClient) ZoneList(region string) (*ecs20140526.DescribeZonesResponse, error) {
	describeZonesRequest := &ecs20140526.DescribeZonesRequest{
		RegionId: tea.String(region),
		Verbose:  tea.Bool(true),
	}
	runtime := &util.RuntimeOptions{}
	// 复制代码运行请自行打印 API 的返回值
	resp, err := ecs.Client.DescribeZonesWithOptions(describeZonesRequest, runtime)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
