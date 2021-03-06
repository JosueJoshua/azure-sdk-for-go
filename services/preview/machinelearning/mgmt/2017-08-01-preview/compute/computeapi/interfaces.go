package computeapi

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
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2017-08-01-preview/compute"
)

// OperationalizationClustersClientAPI contains the set of methods on the OperationalizationClustersClient type.
type OperationalizationClustersClientAPI interface {
	CheckSystemServicesUpdatesAvailable(ctx context.Context, resourceGroupName string, clusterName string) (result compute.CheckSystemServicesUpdatesAvailableResponse, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, parameters compute.OperationalizationCluster) (result compute.OperationalizationClustersCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, clusterName string, deleteAll *bool) (result compute.OperationalizationClustersDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, clusterName string) (result compute.OperationalizationCluster, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, skiptoken string) (result compute.PaginatedOperationalizationClustersListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, skiptoken string) (result compute.PaginatedOperationalizationClustersListIterator, err error)
	ListBySubscriptionID(ctx context.Context, skiptoken string) (result compute.PaginatedOperationalizationClustersListPage, err error)
	ListBySubscriptionIDComplete(ctx context.Context, skiptoken string) (result compute.PaginatedOperationalizationClustersListIterator, err error)
	ListKeys(ctx context.Context, resourceGroupName string, clusterName string) (result compute.OperationalizationClusterCredentials, err error)
	Update(ctx context.Context, resourceGroupName string, clusterName string, parameters compute.OperationalizationClusterUpdateParameters) (result compute.OperationalizationCluster, err error)
	UpdateSystemServices(ctx context.Context, resourceGroupName string, clusterName string) (result compute.OperationalizationClustersUpdateSystemServicesFuture, err error)
}

var _ OperationalizationClustersClientAPI = (*compute.OperationalizationClustersClient)(nil)

// MachineLearningComputeClientAPI contains the set of methods on the MachineLearningComputeClient type.
type MachineLearningComputeClientAPI interface {
	ListAvailableOperations(ctx context.Context) (result compute.AvailableOperations, err error)
}

var _ MachineLearningComputeClientAPI = (*compute.MachineLearningComputeClient)(nil)
