package hybriddataapi

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
	"github.com/Azure/azure-sdk-for-go/services/hybriddatamanager/mgmt/2019-06-01/hybriddata"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result hybriddata.AvailableProviderOperationsPage, err error)
	ListComplete(ctx context.Context) (result hybriddata.AvailableProviderOperationsIterator, err error)
}

var _ OperationsClientAPI = (*hybriddata.OperationsClient)(nil)

// DataManagersClientAPI contains the set of methods on the DataManagersClient type.
type DataManagersClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, dataManagerName string, dataManager hybriddata.DataManager) (result hybriddata.DataManagersCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, dataManagerName string) (result hybriddata.DataManagersDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, dataManagerName string) (result hybriddata.DataManager, err error)
	List(ctx context.Context) (result hybriddata.DataManagerList, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result hybriddata.DataManagerList, err error)
	Update(ctx context.Context, resourceGroupName string, dataManagerName string, dataManagerUpdateParameter hybriddata.DataManagerUpdateParameter, ifMatch string) (result hybriddata.DataManagersUpdateFuture, err error)
}

var _ DataManagersClientAPI = (*hybriddata.DataManagersClient)(nil)

// DataServicesClientAPI contains the set of methods on the DataServicesClient type.
type DataServicesClientAPI interface {
	Get(ctx context.Context, dataServiceName string, resourceGroupName string, dataManagerName string) (result hybriddata.DataService, err error)
	ListByDataManager(ctx context.Context, resourceGroupName string, dataManagerName string) (result hybriddata.DataServiceListPage, err error)
	ListByDataManagerComplete(ctx context.Context, resourceGroupName string, dataManagerName string) (result hybriddata.DataServiceListIterator, err error)
}

var _ DataServicesClientAPI = (*hybriddata.DataServicesClient)(nil)

// JobDefinitionsClientAPI contains the set of methods on the JobDefinitionsClient type.
type JobDefinitionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, dataServiceName string, jobDefinitionName string, jobDefinition hybriddata.JobDefinition, resourceGroupName string, dataManagerName string) (result hybriddata.JobDefinitionsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, dataServiceName string, jobDefinitionName string, resourceGroupName string, dataManagerName string) (result hybriddata.JobDefinitionsDeleteFuture, err error)
	Get(ctx context.Context, dataServiceName string, jobDefinitionName string, resourceGroupName string, dataManagerName string) (result hybriddata.JobDefinition, err error)
	ListByDataManager(ctx context.Context, resourceGroupName string, dataManagerName string, filter string) (result hybriddata.JobDefinitionListPage, err error)
	ListByDataManagerComplete(ctx context.Context, resourceGroupName string, dataManagerName string, filter string) (result hybriddata.JobDefinitionListIterator, err error)
	ListByDataService(ctx context.Context, dataServiceName string, resourceGroupName string, dataManagerName string, filter string) (result hybriddata.JobDefinitionListPage, err error)
	ListByDataServiceComplete(ctx context.Context, dataServiceName string, resourceGroupName string, dataManagerName string, filter string) (result hybriddata.JobDefinitionListIterator, err error)
	Run(ctx context.Context, dataServiceName string, jobDefinitionName string, runParameters hybriddata.RunParameters, resourceGroupName string, dataManagerName string) (result hybriddata.JobDefinitionsRunFuture, err error)
}

var _ JobDefinitionsClientAPI = (*hybriddata.JobDefinitionsClient)(nil)

// JobsClientAPI contains the set of methods on the JobsClient type.
type JobsClientAPI interface {
	Cancel(ctx context.Context, dataServiceName string, jobDefinitionName string, jobID string, resourceGroupName string, dataManagerName string) (result hybriddata.JobsCancelFuture, err error)
	Get(ctx context.Context, dataServiceName string, jobDefinitionName string, jobID string, resourceGroupName string, dataManagerName string, expand string) (result hybriddata.Job, err error)
	ListByDataManager(ctx context.Context, resourceGroupName string, dataManagerName string, filter string) (result hybriddata.JobListPage, err error)
	ListByDataManagerComplete(ctx context.Context, resourceGroupName string, dataManagerName string, filter string) (result hybriddata.JobListIterator, err error)
	ListByDataService(ctx context.Context, dataServiceName string, resourceGroupName string, dataManagerName string, filter string) (result hybriddata.JobListPage, err error)
	ListByDataServiceComplete(ctx context.Context, dataServiceName string, resourceGroupName string, dataManagerName string, filter string) (result hybriddata.JobListIterator, err error)
	ListByJobDefinition(ctx context.Context, dataServiceName string, jobDefinitionName string, resourceGroupName string, dataManagerName string, filter string) (result hybriddata.JobListPage, err error)
	ListByJobDefinitionComplete(ctx context.Context, dataServiceName string, jobDefinitionName string, resourceGroupName string, dataManagerName string, filter string) (result hybriddata.JobListIterator, err error)
	Resume(ctx context.Context, dataServiceName string, jobDefinitionName string, jobID string, resourceGroupName string, dataManagerName string) (result hybriddata.JobsResumeFuture, err error)
}

var _ JobsClientAPI = (*hybriddata.JobsClient)(nil)

// DataStoresClientAPI contains the set of methods on the DataStoresClient type.
type DataStoresClientAPI interface {
	CreateOrUpdate(ctx context.Context, dataStoreName string, dataStore hybriddata.DataStore, resourceGroupName string, dataManagerName string) (result hybriddata.DataStoresCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, dataStoreName string, resourceGroupName string, dataManagerName string) (result hybriddata.DataStoresDeleteFuture, err error)
	Get(ctx context.Context, dataStoreName string, resourceGroupName string, dataManagerName string) (result hybriddata.DataStore, err error)
	ListByDataManager(ctx context.Context, resourceGroupName string, dataManagerName string, filter string) (result hybriddata.DataStoreListPage, err error)
	ListByDataManagerComplete(ctx context.Context, resourceGroupName string, dataManagerName string, filter string) (result hybriddata.DataStoreListIterator, err error)
}

var _ DataStoresClientAPI = (*hybriddata.DataStoresClient)(nil)

// DataStoreTypesClientAPI contains the set of methods on the DataStoreTypesClient type.
type DataStoreTypesClientAPI interface {
	Get(ctx context.Context, dataStoreTypeName string, resourceGroupName string, dataManagerName string) (result hybriddata.DataStoreType, err error)
	ListByDataManager(ctx context.Context, resourceGroupName string, dataManagerName string) (result hybriddata.DataStoreTypeListPage, err error)
	ListByDataManagerComplete(ctx context.Context, resourceGroupName string, dataManagerName string) (result hybriddata.DataStoreTypeListIterator, err error)
}

var _ DataStoreTypesClientAPI = (*hybriddata.DataStoreTypesClient)(nil)

// PublicKeysClientAPI contains the set of methods on the PublicKeysClient type.
type PublicKeysClientAPI interface {
	Get(ctx context.Context, publicKeyName string, resourceGroupName string, dataManagerName string) (result hybriddata.PublicKey, err error)
	ListByDataManager(ctx context.Context, resourceGroupName string, dataManagerName string) (result hybriddata.PublicKeyListPage, err error)
	ListByDataManagerComplete(ctx context.Context, resourceGroupName string, dataManagerName string) (result hybriddata.PublicKeyListIterator, err error)
}

var _ PublicKeysClientAPI = (*hybriddata.PublicKeysClient)(nil)