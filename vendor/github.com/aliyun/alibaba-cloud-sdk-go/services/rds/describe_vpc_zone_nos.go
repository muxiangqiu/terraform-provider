package rds

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeVpcZoneNos invokes the rds.DescribeVpcZoneNos API synchronously
// api document: https://help.aliyun.com/api/rds/describevpczonenos.html
func (client *Client) DescribeVpcZoneNos(request *DescribeVpcZoneNosRequest) (response *DescribeVpcZoneNosResponse, err error) {
	response = CreateDescribeVpcZoneNosResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVpcZoneNosWithChan invokes the rds.DescribeVpcZoneNos API asynchronously
// api document: https://help.aliyun.com/api/rds/describevpczonenos.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVpcZoneNosWithChan(request *DescribeVpcZoneNosRequest) (<-chan *DescribeVpcZoneNosResponse, <-chan error) {
	responseChan := make(chan *DescribeVpcZoneNosResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVpcZoneNos(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeVpcZoneNosWithCallback invokes the rds.DescribeVpcZoneNos API asynchronously
// api document: https://help.aliyun.com/api/rds/describevpczonenos.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVpcZoneNosWithCallback(request *DescribeVpcZoneNosRequest, callback func(response *DescribeVpcZoneNosResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVpcZoneNosResponse
		var err error
		defer close(result)
		response, err = client.DescribeVpcZoneNos(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeVpcZoneNosRequest is the request struct for api DescribeVpcZoneNos
type DescribeVpcZoneNosRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	Region               string           `position:"Query" name:"Region"`
	ZoneId               string           `position:"Query" name:"ZoneId"`
}

// DescribeVpcZoneNosResponse is the response struct for api DescribeVpcZoneNos
type DescribeVpcZoneNosResponse struct {
	*responses.BaseResponse
	RequestId string                    `json:"RequestId" xml:"RequestId"`
	Items     ItemsInDescribeVpcZoneNos `json:"Items" xml:"Items"`
}

// CreateDescribeVpcZoneNosRequest creates a request to invoke DescribeVpcZoneNos API
func CreateDescribeVpcZoneNosRequest() (request *DescribeVpcZoneNosRequest) {
	request = &DescribeVpcZoneNosRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeVpcZoneNos", "rds", "openAPI")
	return
}

// CreateDescribeVpcZoneNosResponse creates a response to parse from DescribeVpcZoneNos response
func CreateDescribeVpcZoneNosResponse() (response *DescribeVpcZoneNosResponse) {
	response = &DescribeVpcZoneNosResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}