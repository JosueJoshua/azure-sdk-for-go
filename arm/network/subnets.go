package network

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 0.14.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
	"net/url"
)

// SubnetsClient is the the Microsoft Azure Network management API provides a
// RESTful set of web services that interact with Microsoft Azure Networks
// service to manage your network resrources. The API has entities that
// capture the relationship between an end user and the Microsoft Azure
// Networks service.
type SubnetsClient struct {
	ManagementClient
}

// NewSubnetsClient creates an instance of the SubnetsClient client.
func NewSubnetsClient(subscriptionID string) SubnetsClient {
	return NewSubnetsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSubnetsClientWithBaseURI creates an instance of the SubnetsClient client.
func NewSubnetsClientWithBaseURI(baseURI string, subscriptionID string) SubnetsClient {
	return SubnetsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate the Put Subnet operation creates/updates a subnet in
// thespecified virtual network
//
// resourceGroupName is the name of the resource group. virtualNetworkName is
// the name of the virtual network. subnetName is the name of the subnet.
// subnetParameters is parameters supplied to the create/update Subnet
// operation
func (client SubnetsClient) CreateOrUpdate(resourceGroupName string, virtualNetworkName string, subnetName string, subnetParameters Subnet) (result autorest.Response, ae error) {
	req, err := client.CreateOrUpdatePreparer(resourceGroupName, virtualNetworkName, subnetName, subnetParameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/SubnetsClient", "CreateOrUpdate", nil, "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "network/SubnetsClient", "CreateOrUpdate", resp, "Failure sending request")
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/SubnetsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SubnetsClient) CreateOrUpdatePreparer(resourceGroupName string, virtualNetworkName string, subnetName string, subnetParameters Subnet) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  url.QueryEscape(resourceGroupName),
		"subnetName":         url.QueryEscape(subnetName),
		"subscriptionId":     url.QueryEscape(client.SubscriptionID),
		"virtualNetworkName": url.QueryEscape(virtualNetworkName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualnetworks/{virtualNetworkName}/subnets/{subnetName}"),
		autorest.WithJSON(subnetParameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SubnetsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	resp, err := client.Send(req)
	if err == nil && azure.ResponseIsLongRunning(resp) {
		req, err := azure.NewAsyncPollingRequest(resp, client.Client)
		if err == nil {
			resp, err = autorest.SendWithSender(client, req,
				azure.WithAsyncPolling(autorest.DefaultPollingDelay))
		}
	}
	return resp, err
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SubnetsClient) CreateOrUpdateResponder(resp *http.Response) (result autorest.Response, ae error) {
	ae = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete the delete subnet operation deletes the specified subnet.
//
// resourceGroupName is the name of the resource group. virtualNetworkName is
// the name of the virtual network. subnetName is the name of the subnet.
func (client SubnetsClient) Delete(resourceGroupName string, virtualNetworkName string, subnetName string) (result autorest.Response, ae error) {
	req, err := client.DeletePreparer(resourceGroupName, virtualNetworkName, subnetName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/SubnetsClient", "Delete", nil, "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "network/SubnetsClient", "Delete", resp, "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/SubnetsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SubnetsClient) DeletePreparer(resourceGroupName string, virtualNetworkName string, subnetName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  url.QueryEscape(resourceGroupName),
		"subnetName":         url.QueryEscape(subnetName),
		"subscriptionId":     url.QueryEscape(client.SubscriptionID),
		"virtualNetworkName": url.QueryEscape(virtualNetworkName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualnetworks/{virtualNetworkName}/subnets/{subnetName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SubnetsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	resp, err := client.Send(req)
	if err == nil && azure.ResponseIsLongRunning(resp) {
		req, err := azure.NewAsyncPollingRequest(resp, client.Client)
		if err == nil {
			resp, err = autorest.SendWithSender(client, req,
				azure.WithAsyncPolling(autorest.DefaultPollingDelay))
		}
	}
	return resp, err
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SubnetsClient) DeleteResponder(resp *http.Response) (result autorest.Response, ae error) {
	ae = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get the Get subnet operation retreives information about the specified
// subnet.
//
// resourceGroupName is the name of the resource group. virtualNetworkName is
// the name of the virtual network. subnetName is the name of the subnet.
// expand is expand references resources.
func (client SubnetsClient) Get(resourceGroupName string, virtualNetworkName string, subnetName string, expand string) (result Subnet, ae error) {
	req, err := client.GetPreparer(resourceGroupName, virtualNetworkName, subnetName, expand)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/SubnetsClient", "Get", nil, "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/SubnetsClient", "Get", resp, "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/SubnetsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client SubnetsClient) GetPreparer(resourceGroupName string, virtualNetworkName string, subnetName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  url.QueryEscape(resourceGroupName),
		"subnetName":         url.QueryEscape(subnetName),
		"subscriptionId":     url.QueryEscape(client.SubscriptionID),
		"virtualNetworkName": url.QueryEscape(virtualNetworkName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = expand
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualnetworks/{virtualNetworkName}/subnets/{subnetName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SubnetsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SubnetsClient) GetResponder(resp *http.Response) (result Subnet, ae error) {
	ae = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List the List subnets opertion retrieves all the subnets in a virtual
// network.
//
// resourceGroupName is the name of the resource group. virtualNetworkName is
// the name of the virtual network.
func (client SubnetsClient) List(resourceGroupName string, virtualNetworkName string) (result SubnetListResult, ae error) {
	req, err := client.ListPreparer(resourceGroupName, virtualNetworkName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/SubnetsClient", "List", nil, "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/SubnetsClient", "List", resp, "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/SubnetsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client SubnetsClient) ListPreparer(resourceGroupName string, virtualNetworkName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  url.QueryEscape(resourceGroupName),
		"subscriptionId":     url.QueryEscape(client.SubscriptionID),
		"virtualNetworkName": url.QueryEscape(virtualNetworkName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualnetworks/{virtualNetworkName}/subnets"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client SubnetsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client SubnetsClient) ListResponder(resp *http.Response) (result SubnetListResult, ae error) {
	ae = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client SubnetsClient) ListNextResults(lastResults SubnetListResult) (result SubnetListResult, ae error) {
	req, err := lastResults.SubnetListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/SubnetsClient", "List", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/SubnetsClient", "List", resp, "Failure sending next results request request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/SubnetsClient", "List", resp, "Failure responding to next results request request")
	}

	return
}
