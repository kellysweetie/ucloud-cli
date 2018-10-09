//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UHost DescribeUHostInstance

package uhost

import (
	"github.com/ucloud/ucloud-sdk-go/sdk"
	"github.com/ucloud/ucloud-sdk-go/sdk/request"
	"github.com/ucloud/ucloud-sdk-go/sdk/response"
)

// DescribeUHostInstanceRequest is request schema for DescribeUHostInstance action
type DescribeUHostInstanceRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	Zone *string `required:"false"`

	// 【数组】UHost主机的资源ID，例如UHostIds.0代表希望获取信息 的主机1，UHostIds.1代表主机2。 如果不传入，则返回当前Region 所有符合条件的UHost实例。
	UHostIds []string `required:"false"`

	// 要查询的业务组名称
	Tag *string `required:"false"`

	// 1：普通云主机；2：抢占型云主机；如不传此参数，默认全部获取
	LifeCycle *int `required:"false"`

	// 列表起始位置偏移量，默认为0
	Offset *int `required:"false"`

	// 返回数据长度，默认为20，最大100
	Limit *int `required:"false"`
}

// DescribeUHostInstanceResponse is response schema for DescribeUHostInstance action
type DescribeUHostInstanceResponse struct {
	response.CommonBase

	// UHostInstance总数
	TotalCount int

	// 云主机实例列表，每项参数可见下面 UHostInstanceSet
	UHostSet []UHostInstanceSet
}

// NewDescribeUHostInstanceRequest will create request of DescribeUHostInstance action.
func (c *UHostClient) NewDescribeUHostInstanceRequest() *DescribeUHostInstanceRequest {
	cfg := c.client.GetConfig()

	return &DescribeUHostInstanceRequest{
		CommonBase: request.CommonBase{
			Region:    sdk.String(cfg.Region),
			ProjectId: sdk.String(cfg.ProjectId),
		},
	}
}

// DescribeUHostInstance - 获取主机或主机列表信息，并可根据数据中心，主机ID等参数进行过滤。
func (c *UHostClient) DescribeUHostInstance(req *DescribeUHostInstanceRequest) (*DescribeUHostInstanceResponse, error) {
	var err error
	var res DescribeUHostInstanceResponse

	err = c.client.InvokeAction("DescribeUHostInstance", req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
