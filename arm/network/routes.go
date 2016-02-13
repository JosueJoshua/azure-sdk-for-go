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

// RoutesClient is the the Microsoft Azure Network management API provides a
// RESTful set of web services that interact with Microsoft Azure Networks
// service to manage your network resrources. The API has entities that
// capture the relationship between an end user and the Microsoft Azure
// Networks service.
type RoutesClient struct {
	ManagementClient
}

// NewRoutesClient creates an instance of the RoutesClient client.
func NewRoutesClient(subscriptionID string) RoutesClient {
	return NewRoutesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRoutesClientWithBaseURI creates an instance of the RoutesClient client.
func NewRoutesClientWithBaseURI(baseURI string, subscriptionID string) RoutesClient {
	return RoutesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate the Put route operation creates/updates a route in the
// specified route table
//
// resourceGroupName is the name of the resource group. routeTableName is the
// name of the route table. routeName is the name of the route.
// routeParameters is parameters supplied to the create/update routeoperation
func (client RoutesClient) CreateOrUpdate(resourceGroupName string, routeTableName string, routeName string, routeParameters Route) (result autorest.Response, ae error) {
	req, err := client.CreateOrUpdatePreparer(resourceGroupName, routeTableName, routeName, routeParameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/RoutesClient", "CreateOrUpdate", nil, "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "network/RoutesClient", "CreateOrUpdate", resp, "Failure sending request")
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/RoutesClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client RoutesClient) CreateOrUpdatePreparer(resourceGroupName string, routeTableName string, routeName string, routeParameters Route) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"routeName":         url.QueryEscape(routeName),
		"routeTableName":    url.QueryEscape(routeTableName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}/routes/{routeName}"),
		autorest.WithJSON(routeParameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client RoutesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
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
func (client RoutesClient) CreateOrUpdateResponder(resp *http.Response) (result autorest.Response, ae error) {
	ae = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete the delete route operation deletes the specified route from a route
// table.
//
// resourceGroupName is the name of the resource group. routeTableName is the
// name of the route table. routeName is the name of the route.
func (client RoutesClient) Delete(resourceGroupName string, routeTableName string, routeName string) (result autorest.Response, ae error) {
	req, err := client.DeletePreparer(resourceGroupName, routeTableName, routeName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/RoutesClient", "Delete", nil, "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "network/RoutesClient", "Delete", resp, "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/RoutesClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client RoutesClient) DeletePreparer(resourceGroupName string, routeTableName string, routeName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"routeName":         url.QueryEscape(routeName),
		"routeTableName":    url.QueryEscape(routeTableName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}/routes/{routeName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client RoutesClient) DeleteSender(req *http.Request) (*http.Response, error) {
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
func (client RoutesClient) DeleteResponder(resp *http.Response) (result autorest.Response, ae error) {
	ae = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusAccepted, http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get the Get route operation retreives information about the specified route
// from the route table.
//
// resourceGroupName is the name of the resource group. routeTableName is the
// name of the route table. routeName is the name of the route.
func (client RoutesClient) Get(resourceGroupName string, routeTableName string, routeName string) (result Route, ae error) {
	req, err := client.GetPreparer(resourceGroupName, routeTableName, routeName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/RoutesClient", "Get", nil, "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/RoutesClient", "Get", resp, "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/RoutesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client RoutesClient) GetPreparer(resourceGroupName string, routeTableName string, routeName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"routeName":         url.QueryEscape(routeName),
		"routeTableName":    url.QueryEscape(routeTableName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}/routes/{routeName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RoutesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RoutesClient) GetResponder(resp *http.Response) (result Route, ae error) {
	ae = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List the List network security rule opertion retrieves all the routes in a
// route table.
//
// resourceGroupName is the name of the resource group. routeTableName is the
// name of the route table.
func (client RoutesClient) List(resourceGroupName string, routeTableName string) (result RouteListResult, ae error) {
	req, err := client.ListPreparer(resourceGroupName, routeTableName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/RoutesClient", "List", nil, "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/RoutesClient", "List", resp, "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/RoutesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client RoutesClient) ListPreparer(resourceGroupName string, routeTableName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"routeTableName":    url.QueryEscape(routeTableName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}/routes"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client RoutesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client RoutesClient) ListResponder(resp *http.Response) (result RouteListResult, ae error) {
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
func (client RoutesClient) ListNextResults(lastResults RouteListResult) (result RouteListResult, ae error) {
	req, err := lastResults.RouteListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/RoutesClient", "List", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/RoutesClient", "List", resp, "Failure sending next results request request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/RoutesClient", "List", resp, "Failure responding to next results request request")
	}

	return
}
