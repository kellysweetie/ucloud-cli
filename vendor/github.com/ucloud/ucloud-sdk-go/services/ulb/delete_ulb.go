//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api ULB DeleteULB

package ulb

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DeleteULBRequest is request schema for DeleteULB action
type DeleteULBRequest struct {
	request.CommonBase

	// 负载均衡实例的ID
	ULBId *string `required:"true"`
}

// DeleteULBResponse is response schema for DeleteULB action
type DeleteULBResponse struct {
	response.CommonBase
}

// NewDeleteULBRequest will create request of DeleteULB action.
func (c *ULBClient) NewDeleteULBRequest() *DeleteULBRequest {
	req := &DeleteULBRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DeleteULB - 删除负载均衡实例
func (c *ULBClient) DeleteULB(req *DeleteULBRequest) (*DeleteULBResponse, error) {
	var err error
	var res DeleteULBResponse

	err = c.client.InvokeAction("DeleteULB", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
