// This file is auto-generated, don't edit it. Thanks.
package nodegroup

import (
  //"encoding/json"
  //"strings"
  //"fmt"
  //"os"
  cs20151215  "github.com/alibabacloud-go/cs-20151215/v5/client"
  //openapi  "github.com/alibabacloud-go/darabonba-openapi/v2/client"
  util  "github.com/alibabacloud-go/tea-utils/v2/service"
  "github.com/alibabacloud-go/tea/tea"
)

//node group list
func (ack *AliyunAckClient) NodeGroupList(clusterid string) (*cs20151215.DescribeClusterNodePoolsResponse,error) {

  describeClusterNodePoolsRequest := &cs20151215.DescribeClusterNodePoolsRequest{}
  runtime := &util.RuntimeOptions{}
  headers := make(map[string]*string)
    // 复制代码运行请自行打印 API 的返回值
  resp, err := ack.Client.DescribeClusterNodePoolsWithOptions(tea.String(clusterid), describeClusterNodePoolsRequest, headers, runtime)
  if err != nil {
    return nil,err
  }

  return resp,err
}
//node group info 
//func (ack *AliyunAckClient) NodeGroupInfo()
