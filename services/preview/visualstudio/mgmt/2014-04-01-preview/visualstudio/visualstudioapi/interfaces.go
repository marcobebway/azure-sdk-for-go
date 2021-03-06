package visualstudioapi

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
	"github.com/Azure/azure-sdk-for-go/services/preview/visualstudio/mgmt/2014-04-01-preview/visualstudio"
	"github.com/Azure/go-autorest/autorest"
	"github.com/gofrs/uuid"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result visualstudio.OperationListResult, err error)
}

var _ OperationsClientAPI = (*visualstudio.OperationsClient)(nil)

// AccountsClientAPI contains the set of methods on the AccountsClient type.
type AccountsClientAPI interface {
	CheckNameAvailability(ctx context.Context, body visualstudio.CheckNameAvailabilityParameter) (result visualstudio.CheckNameAvailabilityResult, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, body visualstudio.AccountResourceRequest, resourceName string) (result visualstudio.AccountResource, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string) (result visualstudio.AccountResource, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result visualstudio.AccountResourceListResult, err error)
	Update(ctx context.Context, resourceGroupName string, body visualstudio.AccountTagRequest, resourceName string) (result visualstudio.AccountResource, err error)
}

var _ AccountsClientAPI = (*visualstudio.AccountsClient)(nil)

// ExtensionsClientAPI contains the set of methods on the ExtensionsClient type.
type ExtensionsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, body visualstudio.ExtensionResourceRequest, accountResourceName string, extensionResourceName string) (result visualstudio.ExtensionResource, err error)
	Delete(ctx context.Context, resourceGroupName string, accountResourceName string, extensionResourceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, accountResourceName string, extensionResourceName string) (result visualstudio.ExtensionResource, err error)
	ListByAccount(ctx context.Context, resourceGroupName string, accountResourceName string) (result visualstudio.ExtensionResourceListResult, err error)
	Update(ctx context.Context, resourceGroupName string, body visualstudio.ExtensionResourceRequest, accountResourceName string, extensionResourceName string) (result visualstudio.ExtensionResource, err error)
}

var _ ExtensionsClientAPI = (*visualstudio.ExtensionsClient)(nil)

// ProjectsClientAPI contains the set of methods on the ProjectsClient type.
type ProjectsClientAPI interface {
	Create(ctx context.Context, body visualstudio.ProjectResource, resourceGroupName string, rootResourceName string, resourceName string, validating string) (result visualstudio.ProjectsCreateFuture, err error)
	Get(ctx context.Context, resourceGroupName string, rootResourceName string, resourceName string) (result visualstudio.ProjectResource, err error)
	GetJobStatus(ctx context.Context, resourceGroupName string, rootResourceName string, resourceName string, subContainerName string, operation string, jobID *uuid.UUID) (result visualstudio.ProjectResource, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, rootResourceName string) (result visualstudio.ProjectResourceListResult, err error)
	Update(ctx context.Context, resourceGroupName string, body visualstudio.ProjectResource, rootResourceName string, resourceName string) (result visualstudio.ProjectResource, err error)
}

var _ ProjectsClientAPI = (*visualstudio.ProjectsClient)(nil)
