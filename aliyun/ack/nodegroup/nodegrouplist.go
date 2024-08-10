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
func (ack *AliyunAckClient) NodeGroupInfo(clusterid,nodegroupid string) (*cs20151215.DescribeClusterNodePoolDetailResponse,error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	// 复制代码运行请自行打印 API 的返回值
	resp, err := ack.Client.DescribeClusterNodePoolDetailWithOptions(tea.String(clusterid), tea.String(nodegroupid), headers, runtime)
	if err != nil {
	  return nil,err
	}

	return resp,err 
}
func (ack *AliyunAckClient) NodeListByNodeGroup(clusterid,nodegroupid string) (*cs20151215.DescribeClusterNodesResponse,error) {
	describeClusterNodesRequest := &cs20151215.DescribeClusterNodesRequest{
		NodepoolId: tea.String(nodegroupid),
	}
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	resp, err := ack.Client.DescribeClusterNodesWithOptions(tea.String(clusterid), describeClusterNodesRequest, headers, runtime)
    if err != nil {
      return nil,err
    }

	return resp,nil 
}
