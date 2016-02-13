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

// ExpressRouteCircuitAuthorizationsClient is the the Microsoft Azure Network
// management API provides a RESTful set of web services that interact with
// Microsoft Azure Networks service to manage your network resrources. The
// API has entities that capture the relationship between an end user and the
// Microsoft Azure Networks service.
type ExpressRouteCircuitAuthorizationsClient struct {
	ManagementClient
}

// NewExpressRouteCircuitAuthorizationsClient creates an instance of the
// ExpressRouteCircuitAuthorizationsClient client.
func NewExpressRouteCircuitAuthorizationsClient(subscriptionID string) ExpressRouteCircuitAuthorizationsClient {
	return NewExpressRouteCircuitAuthorizationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewExpressRouteCircuitAuthorizationsClientWithBaseURI creates an instance
// of the ExpressRouteCircuitAuthorizationsClient client.
func NewExpressRouteCircuitAuthorizationsClientWithBaseURI(baseURI string, subscriptionID string) ExpressRouteCircuitAuthorizationsClient {
	return ExpressRouteCircuitAuthorizationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate the Put Authorization operation creates/updates an
// authorization in thespecified ExpressRouteCircuits
//
// resourceGroupName is the name of the resource group. circuitName is the
// name of the express route circuit. authorizationName is the name of the
// authorization. authorizationParameters is parameters supplied to the
// create/update ExpressRouteCircuitAuthorization operation
func (client ExpressRouteCircuitAuthorizationsClient) CreateOrUpdate(resourceGroupName string, circuitName string, authorizationName string, authorizationParameters ExpressRouteCircuitAuthorization) (result autorest.Response, ae error) {
	req, err := client.CreateOrUpdatePreparer(resourceGroupName, circuitName, authorizationName, authorizationParameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitAuthorizationsClient", "CreateOrUpdate", nil, "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitAuthorizationsClient", "CreateOrUpdate", resp, "Failure sending request")
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/ExpressRouteCircuitAuthorizationsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ExpressRouteCircuitAuthorizationsClient) CreateOrUpdatePreparer(resourceGroupName string, circuitName string, authorizationName string, authorizationParameters ExpressRouteCircuitAuthorization) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"authorizationName": url.QueryEscape(authorizationName),
		"circuitName":       url.QueryEscape(circuitName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/authorizations/{authorizationName}"),
		autorest.WithJSON(authorizationParameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitAuthorizationsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
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
func (client ExpressRouteCircuitAuthorizationsClient) CreateOrUpdateResponder(resp *http.Response) (result autorest.Response, ae error) {
	ae = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete the delete authorization operation deletes the specified
// authorization from the specified ExpressRouteCircuit.
//
// resourceGroupName is the name of the resource group. circuitName is the
// name of the express route circuit. authorizationName is the name of the
// authorization.
func (client ExpressRouteCircuitAuthorizationsClient) Delete(resourceGroupName string, circuitName string, authorizationName string) (result autorest.Response, ae error) {
	req, err := client.DeletePreparer(resourceGroupName, circuitName, authorizationName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitAuthorizationsClient", "Delete", nil, "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitAuthorizationsClient", "Delete", resp, "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/ExpressRouteCircuitAuthorizationsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ExpressRouteCircuitAuthorizationsClient) DeletePreparer(resourceGroupName string, circuitName string, authorizationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"authorizationName": url.QueryEscape(authorizationName),
		"circuitName":       url.QueryEscape(circuitName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/authorizations/{authorizationName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitAuthorizationsClient) DeleteSender(req *http.Request) (*http.Response, error) {
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
func (client ExpressRouteCircuitAuthorizationsClient) DeleteResponder(resp *http.Response) (result autorest.Response, ae error) {
	ae = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusAccepted, http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get the GET authorization operation retrieves the specified authorization
// from the specified ExpressRouteCircuit.
//
// resourceGroupName is the name of the resource group. circuitName is the
// name of the express route circuit. authorizationName is the name of the
// authorization.
func (client ExpressRouteCircuitAuthorizationsClient) Get(resourceGroupName string, circuitName string, authorizationName string) (result ExpressRouteCircuitAuthorization, ae error) {
	req, err := client.GetPreparer(resourceGroupName, circuitName, authorizationName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitAuthorizationsClient", "Get", nil, "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitAuthorizationsClient", "Get", resp, "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/ExpressRouteCircuitAuthorizationsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ExpressRouteCircuitAuthorizationsClient) GetPreparer(resourceGroupName string, circuitName string, authorizationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"authorizationName": url.QueryEscape(authorizationName),
		"circuitName":       url.QueryEscape(circuitName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/authorizations/{authorizationName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitAuthorizationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitAuthorizationsClient) GetResponder(resp *http.Response) (result ExpressRouteCircuitAuthorization, ae error) {
	ae = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List the List authorization operation retrieves all the authorizations in
// an ExpressRouteCircuit.
//
// resourceGroupName is the name of the resource group. circuitName is the
// name of the curcuit.
func (client ExpressRouteCircuitAuthorizationsClient) List(resourceGroupName string, circuitName string) (result AuthorizationListResult, ae error) {
	req, err := client.ListPreparer(resourceGroupName, circuitName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitAuthorizationsClient", "List", nil, "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitAuthorizationsClient", "List", resp, "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/ExpressRouteCircuitAuthorizationsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ExpressRouteCircuitAuthorizationsClient) ListPreparer(resourceGroupName string, circuitName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"circuitName":       url.QueryEscape(circuitName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/authorizations"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitAuthorizationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitAuthorizationsClient) ListResponder(resp *http.Response) (result AuthorizationListResult, ae error) {
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
func (client ExpressRouteCircuitAuthorizationsClient) ListNextResults(lastResults AuthorizationListResult) (result AuthorizationListResult, ae error) {
	req, err := lastResults.AuthorizationListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitAuthorizationsClient", "List", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitAuthorizationsClient", "List", resp, "Failure sending next results request request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/ExpressRouteCircuitAuthorizationsClient", "List", resp, "Failure responding to next results request request")
	}

	return
}
