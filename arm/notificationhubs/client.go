// Package notificationhubs implements the Azure ARM Notificationhubs service
// API version 2014-09-01.
//
// .Net client wrapper for the REST API for Azure NotificationHub Service
package notificationhubs

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

const (
	// APIVersion is the version of the Notificationhubs
	APIVersion = "2014-09-01"

	// DefaultBaseURI is the default URI used for the service Notificationhubs
	DefaultBaseURI = "https://management.azure.com"
)

// ManagementClient is the base client for Notificationhubs.
type ManagementClient struct {
	autorest.Client
	BaseURI        string
	SubscriptionID string
}

// New creates an instance of the ManagementClient client.
func New(subscriptionID string) ManagementClient {
	return NewWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWithBaseURI creates an instance of the ManagementClient client.
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return ManagementClient{
		Client:         autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:        baseURI,
		SubscriptionID: subscriptionID,
	}
}

// CheckAvailability checks the availability of the given notificationHub in a
// namespace.
//
// resourceGroupName is the name of the resource group. namespaceName is the
// namespace name. parameters is the notificationHub name.
func (client ManagementClient) CheckAvailability(resourceGroupName string, namespaceName string, parameters CheckAvailabilityParameters) (result CheckAvailabilityResource, ae error) {
	req, err := client.CheckAvailabilityPreparer(resourceGroupName, namespaceName, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "notificationhubs/ManagementClient", "CheckAvailability", nil, "Failure preparing request")
	}

	resp, err := client.CheckAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "notificationhubs/ManagementClient", "CheckAvailability", resp, "Failure sending request")
	}

	result, err = client.CheckAvailabilityResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "notificationhubs/ManagementClient", "CheckAvailability", resp, "Failure responding to request")
	}

	return
}

// CheckAvailabilityPreparer prepares the CheckAvailability request.
func (client ManagementClient) CheckAvailabilityPreparer(resourceGroupName string, namespaceName string, parameters CheckAvailabilityParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"namespaceName":     url.QueryEscape(namespaceName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/checkNotificationHubAvailability"),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CheckAvailabilitySender sends the CheckAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) CheckAvailabilitySender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// CheckAvailabilityResponder handles the response to the CheckAvailability request. The method always
// closes the http.Response Body.
func (client ManagementClient) CheckAvailabilityResponder(resp *http.Response) (result CheckAvailabilityResource, ae error) {
	ae = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Create creates a new NotificationHub in a namespace.
//
// resourceGroupName is the name of the resource group. namespaceName is the
// namespace name. notificationHubName is the notification hub name.
// parameters is parameters supplied to the create a Namespace Resource.
func (client ManagementClient) Create(resourceGroupName string, namespaceName string, notificationHubName string, parameters NotificationHubCreateOrUpdateParameters) (result NotificationHubResource, ae error) {
	req, err := client.CreatePreparer(resourceGroupName, namespaceName, notificationHubName, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "notificationhubs/ManagementClient", "Create", nil, "Failure preparing request")
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "notificationhubs/ManagementClient", "Create", resp, "Failure sending request")
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "notificationhubs/ManagementClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ManagementClient) CreatePreparer(resourceGroupName string, namespaceName string, notificationHubName string, parameters NotificationHubCreateOrUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"namespaceName":       url.QueryEscape(namespaceName),
		"notificationHubName": url.QueryEscape(notificationHubName),
		"resourceGroupName":   url.QueryEscape(resourceGroupName),
		"subscriptionId":      url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/notificationHubs/{notificationHubName}"),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ManagementClient) CreateResponder(resp *http.Response) (result NotificationHubResource, ae error) {
	ae = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdateAuthorizationRule the create NotificationHub authorization
// rule operation creates an authorization rule for a NotificationHub
//
// resourceGroupName is the name of the resource group. namespaceName is the
// namespace name. notificationHubName is the notification hub name.
// authorizationRuleName is the namespace authorizationRuleName name.
// parameters is the shared access authorization rule.
func (client ManagementClient) CreateOrUpdateAuthorizationRule(resourceGroupName string, namespaceName string, notificationHubName string, authorizationRuleName string, parameters SharedAccessAuthorizationRuleCreateOrUpdateParameters) (result SharedAccessAuthorizationRuleResource, ae error) {
	req, err := client.CreateOrUpdateAuthorizationRulePreparer(resourceGroupName, namespaceName, notificationHubName, authorizationRuleName, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "notificationhubs/ManagementClient", "CreateOrUpdateAuthorizationRule", nil, "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateAuthorizationRuleSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "notificationhubs/ManagementClient", "CreateOrUpdateAuthorizationRule", resp, "Failure sending request")
	}

	result, err = client.CreateOrUpdateAuthorizationRuleResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "notificationhubs/ManagementClient", "CreateOrUpdateAuthorizationRule", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdateAuthorizationRulePreparer prepares the CreateOrUpdateAuthorizationRule request.
func (client ManagementClient) CreateOrUpdateAuthorizationRulePreparer(resourceGroupName string, namespaceName string, notificationHubName string, authorizationRuleName string, parameters SharedAccessAuthorizationRuleCreateOrUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"authorizationRuleName": url.QueryEscape(authorizationRuleName),
		"namespaceName":         url.QueryEscape(namespaceName),
		"notificationHubName":   url.QueryEscape(notificationHubName),
		"resourceGroupName":     url.QueryEscape(resourceGroupName),
		"subscriptionId":        url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/notificationHubs/{notificationHubName}/AuthorizationRules/{authorizationRuleName}"),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CreateOrUpdateAuthorizationRuleSender sends the CreateOrUpdateAuthorizationRule request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) CreateOrUpdateAuthorizationRuleSender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// CreateOrUpdateAuthorizationRuleResponder handles the response to the CreateOrUpdateAuthorizationRule request. The method always
// closes the http.Response Body.
func (client ManagementClient) CreateOrUpdateAuthorizationRuleResponder(resp *http.Response) (result SharedAccessAuthorizationRuleResource, ae error) {
	ae = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a notification hub associated with a namespace.
//
// resourceGroupName is the name of the resource group. namespaceName is the
// namespace name. notificationHubName is the notification hub name.
func (client ManagementClient) Delete(resourceGroupName string, namespaceName string, notificationHubName string) (result autorest.Response, ae error) {
	req, err := client.DeletePreparer(resourceGroupName, namespaceName, notificationHubName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "notificationhubs/ManagementClient", "Delete", nil, "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "notificationhubs/ManagementClient", "Delete", resp, "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "notificationhubs/ManagementClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ManagementClient) DeletePreparer(resourceGroupName string, namespaceName string, notificationHubName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"namespaceName":       url.QueryEscape(namespaceName),
		"notificationHubName": url.QueryEscape(notificationHubName),
		"resourceGroupName":   url.QueryEscape(resourceGroupName),
		"subscriptionId":      url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/notificationHubs/{notificationHubName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ManagementClient) DeleteResponder(resp *http.Response) (result autorest.Response, ae error) {
	ae = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteAuthorizationRule the delete a notificationHub authorization rule
// operation
//
// resourceGroupName is the name of the resource group. namespaceName is the
// namespace name. notificationHubName is the notification hub name.
// authorizationRuleName is the namespace authorizationRuleName name.
func (client ManagementClient) DeleteAuthorizationRule(resourceGroupName string, namespaceName string, notificationHubName string, authorizationRuleName string) (result autorest.Response, ae error) {
	req, err := client.DeleteAuthorizationRulePreparer(resourceGroupName, namespaceName, notificationHubName, authorizationRuleName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "notificationhubs/ManagementClient", "DeleteAuthorizationRule", nil, "Failure preparing request")
	}

	resp, err := client.DeleteAuthorizationRuleSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "notificationhubs/ManagementClient", "DeleteAuthorizationRule", resp, "Failure sending request")
	}

	result, err = client.DeleteAuthorizationRuleResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "notificationhubs/ManagementClient", "DeleteAuthorizationRule", resp, "Failure responding to request")
	}

	return
}

// DeleteAuthorizationRulePreparer prepares the DeleteAuthorizationRule request.
func (client ManagementClient) DeleteAuthorizationRulePreparer(resourceGroupName string, namespaceName string, notificationHubName string, authorizationRuleName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"authorizationRuleName": url.QueryEscape(authorizationRuleName),
		"namespaceName":         url.QueryEscape(namespaceName),
		"notificationHubName":   url.QueryEscape(notificationHubName),
		"resourceGroupName":     url.QueryEscape(resourceGroupName),
		"subscriptionId":        url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/notificationHubs/{notificationHubName}/AuthorizationRules/{authorizationRuleName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteAuthorizationRuleSender sends the DeleteAuthorizationRule request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) DeleteAuthorizationRuleSender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// DeleteAuthorizationRuleResponder handles the response to the DeleteAuthorizationRule request. The method always
// closes the http.Response Body.
func (client ManagementClient) DeleteAuthorizationRuleResponder(resp *http.Response) (result autorest.Response, ae error) {
	ae = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get lists the notification hubs associated with a namespace.
//
// resourceGroupName is the name of the resource group. namespaceName is the
// namespace name. notificationHubName is the notification hub name.
func (client ManagementClient) Get(resourceGroupName string, namespaceName string, notificationHubName string) (result NotificationHubResource, ae error) {
	req, err := client.GetPreparer(resourceGroupName, namespaceName, notificationHubName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "notificationhubs/ManagementClient", "Get", nil, "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "notificationhubs/ManagementClient", "Get", resp, "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "notificationhubs/ManagementClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ManagementClient) GetPreparer(resourceGroupName string, namespaceName string, notificationHubName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"namespaceName":       url.QueryEscape(namespaceName),
		"notificationHubName": url.QueryEscape(notificationHubName),
		"resourceGroupName":   url.QueryEscape(resourceGroupName),
		"subscriptionId":      url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/notificationHubs/{notificationHubName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ManagementClient) GetResponder(resp *http.Response) (result NotificationHubResource, ae error) {
	ae = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetAuthorizationRule the get authorization rule operation gets an
// authorization rule for a NotificationHub by name.
//
// resourceGroupName is the name of the resource group. namespaceName is the
// namespace to get the authorization rule for. notificationHubName is the
// notification hub name. authorizationRuleName is the entity name to get the
// authorization rule for.
func (client ManagementClient) GetAuthorizationRule(resourceGroupName string, namespaceName string, notificationHubName string, authorizationRuleName string) (result SharedAccessAuthorizationRuleResource, ae error) {
	req, err := client.GetAuthorizationRulePreparer(resourceGroupName, namespaceName, notificationHubName, authorizationRuleName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "notificationhubs/ManagementClient", "GetAuthorizationRule", nil, "Failure preparing request")
	}

	resp, err := client.GetAuthorizationRuleSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "notificationhubs/ManagementClient", "GetAuthorizationRule", resp, "Failure sending request")
	}

	result, err = client.GetAuthorizationRuleResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "notificationhubs/ManagementClient", "GetAuthorizationRule", resp, "Failure responding to request")
	}

	return
}

// GetAuthorizationRulePreparer prepares the GetAuthorizationRule request.
func (client ManagementClient) GetAuthorizationRulePreparer(resourceGroupName string, namespaceName string, notificationHubName string, authorizationRuleName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"authorizationRuleName": url.QueryEscape(authorizationRuleName),
		"namespaceName":         url.QueryEscape(namespaceName),
		"notificationHubName":   url.QueryEscape(notificationHubName),
		"resourceGroupName":     url.QueryEscape(resourceGroupName),
		"subscriptionId":        url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/notificationHubs/{notificationHubName}/AuthorizationRules/{authorizationRuleName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetAuthorizationRuleSender sends the GetAuthorizationRule request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) GetAuthorizationRuleSender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// GetAuthorizationRuleResponder handles the response to the GetAuthorizationRule request. The method always
// closes the http.Response Body.
func (client ManagementClient) GetAuthorizationRuleResponder(resp *http.Response) (result SharedAccessAuthorizationRuleResource, ae error) {
	ae = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetPnsCredentials lists the PNS Credentials associated with a notification
// hub .
//
// resourceGroupName is the name of the resource group. namespaceName is the
// namespace name. notificationHubName is the notification hub name.
func (client ManagementClient) GetPnsCredentials(resourceGroupName string, namespaceName string, notificationHubName string) (result NotificationHubResource, ae error) {
	req, err := client.GetPnsCredentialsPreparer(resourceGroupName, namespaceName, notificationHubName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "notificationhubs/ManagementClient", "GetPnsCredentials", nil, "Failure preparing request")
	}

	resp, err := client.GetPnsCredentialsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "notificationhubs/ManagementClient", "GetPnsCredentials", resp, "Failure sending request")
	}

	result, err = client.GetPnsCredentialsResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "notificationhubs/ManagementClient", "GetPnsCredentials", resp, "Failure responding to request")
	}

	return
}

// GetPnsCredentialsPreparer prepares the GetPnsCredentials request.
func (client ManagementClient) GetPnsCredentialsPreparer(resourceGroupName string, namespaceName string, notificationHubName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"namespaceName":       url.QueryEscape(namespaceName),
		"notificationHubName": url.QueryEscape(notificationHubName),
		"resourceGroupName":   url.QueryEscape(resourceGroupName),
		"subscriptionId":      url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/notificationHubs/{notificationHubName}/pnsCredentials"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetPnsCredentialsSender sends the GetPnsCredentials request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) GetPnsCredentialsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// GetPnsCredentialsResponder handles the response to the GetPnsCredentials request. The method always
// closes the http.Response Body.
func (client ManagementClient) GetPnsCredentialsResponder(resp *http.Response) (result NotificationHubResource, ae error) {
	ae = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists the notification hubs associated with a namespace.
//
// resourceGroupName is the name of the resource group. namespaceName is the
// namespace name.
func (client ManagementClient) List(resourceGroupName string, namespaceName string) (result NotificationHubListResult, ae error) {
	req, err := client.ListPreparer(resourceGroupName, namespaceName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "notificationhubs/ManagementClient", "List", nil, "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "notificationhubs/ManagementClient", "List", resp, "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "notificationhubs/ManagementClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ManagementClient) ListPreparer(resourceGroupName string, namespaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"namespaceName":     url.QueryEscape(namespaceName),
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
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/notificationHubs"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ManagementClient) ListResponder(resp *http.Response) (result NotificationHubListResult, ae error) {
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
func (client ManagementClient) ListNextResults(lastResults NotificationHubListResult) (result NotificationHubListResult, ae error) {
	req, err := lastResults.NotificationHubListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "notificationhubs/ManagementClient", "List", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "notificationhubs/ManagementClient", "List", resp, "Failure sending next results request request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "notificationhubs/ManagementClient", "List", resp, "Failure responding to next results request request")
	}

	return
}

// ListAuthorizationRules the get authorization rules operation gets the
// authorization rules for a NotificationHub.
//
// resourceGroupName is the name of the resource group. namespaceName is the
// NotificationHub to get the authorization rule for. notificationHubName is
// the notification hub name.
func (client ManagementClient) ListAuthorizationRules(resourceGroupName string, namespaceName string, notificationHubName string) (result SharedAccessAuthorizationRuleListResult, ae error) {
	req, err := client.ListAuthorizationRulesPreparer(resourceGroupName, namespaceName, notificationHubName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "notificationhubs/ManagementClient", "ListAuthorizationRules", nil, "Failure preparing request")
	}

	resp, err := client.ListAuthorizationRulesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "notificationhubs/ManagementClient", "ListAuthorizationRules", resp, "Failure sending request")
	}

	result, err = client.ListAuthorizationRulesResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "notificationhubs/ManagementClient", "ListAuthorizationRules", resp, "Failure responding to request")
	}

	return
}

// ListAuthorizationRulesPreparer prepares the ListAuthorizationRules request.
func (client ManagementClient) ListAuthorizationRulesPreparer(resourceGroupName string, namespaceName string, notificationHubName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"namespaceName":       url.QueryEscape(namespaceName),
		"notificationHubName": url.QueryEscape(notificationHubName),
		"resourceGroupName":   url.QueryEscape(resourceGroupName),
		"subscriptionId":      url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/notificationHubs/{notificationHubName}/AuthorizationRules"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListAuthorizationRulesSender sends the ListAuthorizationRules request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) ListAuthorizationRulesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// ListAuthorizationRulesResponder handles the response to the ListAuthorizationRules request. The method always
// closes the http.Response Body.
func (client ManagementClient) ListAuthorizationRulesResponder(resp *http.Response) (result SharedAccessAuthorizationRuleListResult, ae error) {
	ae = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListAuthorizationRulesNextResults retrieves the next set of results, if any.
func (client ManagementClient) ListAuthorizationRulesNextResults(lastResults SharedAccessAuthorizationRuleListResult) (result SharedAccessAuthorizationRuleListResult, ae error) {
	req, err := lastResults.SharedAccessAuthorizationRuleListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "notificationhubs/ManagementClient", "ListAuthorizationRules", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListAuthorizationRulesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "notificationhubs/ManagementClient", "ListAuthorizationRules", resp, "Failure sending next results request request")
	}

	result, err = client.ListAuthorizationRulesResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "notificationhubs/ManagementClient", "ListAuthorizationRules", resp, "Failure responding to next results request request")
	}

	return
}

// ListKeys gets the Primary and Secondary ConnectionStrings to the
// NotificationHub
//
// resourceGroupName is the name of the resource group. namespaceName is the
// namespace name. notificationHubName is the notification hub name.
// authorizationRuleName is the connection string of the NotificationHub for
// the specified authorizationRule.
func (client ManagementClient) ListKeys(resourceGroupName string, namespaceName string, notificationHubName string, authorizationRuleName string) (result ResourceListKeys, ae error) {
	req, err := client.ListKeysPreparer(resourceGroupName, namespaceName, notificationHubName, authorizationRuleName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "notificationhubs/ManagementClient", "ListKeys", nil, "Failure preparing request")
	}

	resp, err := client.ListKeysSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "notificationhubs/ManagementClient", "ListKeys", resp, "Failure sending request")
	}

	result, err = client.ListKeysResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "notificationhubs/ManagementClient", "ListKeys", resp, "Failure responding to request")
	}

	return
}

// ListKeysPreparer prepares the ListKeys request.
func (client ManagementClient) ListKeysPreparer(resourceGroupName string, namespaceName string, notificationHubName string, authorizationRuleName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"authorizationRuleName": url.QueryEscape(authorizationRuleName),
		"namespaceName":         url.QueryEscape(namespaceName),
		"notificationHubName":   url.QueryEscape(notificationHubName),
		"resourceGroupName":     url.QueryEscape(resourceGroupName),
		"subscriptionId":        url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/notificationHubs/{notificationHubName}/AuthorizationRules/{authorizationRuleName}/listKeys"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListKeysSender sends the ListKeys request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) ListKeysSender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// ListKeysResponder handles the response to the ListKeys request. The method always
// closes the http.Response Body.
func (client ManagementClient) ListKeysResponder(resp *http.Response) (result ResourceListKeys, ae error) {
	ae = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
