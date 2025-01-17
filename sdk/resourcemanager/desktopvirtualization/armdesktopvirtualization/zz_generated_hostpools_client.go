//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdesktopvirtualization

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// HostPoolsClient contains the methods for the HostPools group.
// Don't use this type directly, use NewHostPoolsClient() instead.
type HostPoolsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewHostPoolsClient creates a new instance of HostPoolsClient with the specified values.
func NewHostPoolsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *HostPoolsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &HostPoolsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CreateOrUpdate - Create or update a host pool.
// If the operation fails it returns the *CloudError error type.
func (client *HostPoolsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, hostPoolName string, hostPool HostPool, options *HostPoolsCreateOrUpdateOptions) (HostPoolsCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, hostPoolName, hostPool, options)
	if err != nil {
		return HostPoolsCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HostPoolsCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return HostPoolsCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *HostPoolsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, hostPoolName string, hostPool HostPool, options *HostPoolsCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostPoolName == "" {
		return nil, errors.New("parameter hostPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostPoolName}", url.PathEscape(hostPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-03-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, hostPool)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *HostPoolsClient) createOrUpdateHandleResponse(resp *http.Response) (HostPoolsCreateOrUpdateResponse, error) {
	result := HostPoolsCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.HostPool); err != nil {
		return HostPoolsCreateOrUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *HostPoolsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Remove a host pool.
// If the operation fails it returns the *CloudError error type.
func (client *HostPoolsClient) Delete(ctx context.Context, resourceGroupName string, hostPoolName string, options *HostPoolsDeleteOptions) (HostPoolsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, hostPoolName, options)
	if err != nil {
		return HostPoolsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HostPoolsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return HostPoolsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return HostPoolsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *HostPoolsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, hostPoolName string, options *HostPoolsDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostPoolName == "" {
		return nil, errors.New("parameter hostPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostPoolName}", url.PathEscape(hostPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-03-preview")
	if options != nil && options.Force != nil {
		reqQP.Set("force", strconv.FormatBool(*options.Force))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *HostPoolsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Get a host pool.
// If the operation fails it returns the *CloudError error type.
func (client *HostPoolsClient) Get(ctx context.Context, resourceGroupName string, hostPoolName string, options *HostPoolsGetOptions) (HostPoolsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, hostPoolName, options)
	if err != nil {
		return HostPoolsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HostPoolsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HostPoolsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *HostPoolsClient) getCreateRequest(ctx context.Context, resourceGroupName string, hostPoolName string, options *HostPoolsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostPoolName == "" {
		return nil, errors.New("parameter hostPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostPoolName}", url.PathEscape(hostPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-03-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *HostPoolsClient) getHandleResponse(resp *http.Response) (HostPoolsGetResponse, error) {
	result := HostPoolsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.HostPool); err != nil {
		return HostPoolsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *HostPoolsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - List hostPools in subscription.
// If the operation fails it returns the *CloudError error type.
func (client *HostPoolsClient) List(options *HostPoolsListOptions) *HostPoolsListPager {
	return &HostPoolsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp HostPoolsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.HostPoolList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *HostPoolsClient) listCreateRequest(ctx context.Context, options *HostPoolsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DesktopVirtualization/hostPools"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-03-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *HostPoolsClient) listHandleResponse(resp *http.Response) (HostPoolsListResponse, error) {
	result := HostPoolsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.HostPoolList); err != nil {
		return HostPoolsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *HostPoolsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByResourceGroup - List hostPools.
// If the operation fails it returns the *CloudError error type.
func (client *HostPoolsClient) ListByResourceGroup(resourceGroupName string, options *HostPoolsListByResourceGroupOptions) *HostPoolsListByResourceGroupPager {
	return &HostPoolsListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp HostPoolsListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.HostPoolList.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *HostPoolsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *HostPoolsListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-03-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *HostPoolsClient) listByResourceGroupHandleResponse(resp *http.Response) (HostPoolsListByResourceGroupResponse, error) {
	result := HostPoolsListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.HostPoolList); err != nil {
		return HostPoolsListByResourceGroupResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *HostPoolsClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// RetrieveRegistrationToken - Registration token of the host pool.
// If the operation fails it returns the *CloudError error type.
func (client *HostPoolsClient) RetrieveRegistrationToken(ctx context.Context, resourceGroupName string, hostPoolName string, options *HostPoolsRetrieveRegistrationTokenOptions) (HostPoolsRetrieveRegistrationTokenResponse, error) {
	req, err := client.retrieveRegistrationTokenCreateRequest(ctx, resourceGroupName, hostPoolName, options)
	if err != nil {
		return HostPoolsRetrieveRegistrationTokenResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HostPoolsRetrieveRegistrationTokenResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HostPoolsRetrieveRegistrationTokenResponse{}, client.retrieveRegistrationTokenHandleError(resp)
	}
	return client.retrieveRegistrationTokenHandleResponse(resp)
}

// retrieveRegistrationTokenCreateRequest creates the RetrieveRegistrationToken request.
func (client *HostPoolsClient) retrieveRegistrationTokenCreateRequest(ctx context.Context, resourceGroupName string, hostPoolName string, options *HostPoolsRetrieveRegistrationTokenOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/retrieveRegistrationToken"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostPoolName == "" {
		return nil, errors.New("parameter hostPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostPoolName}", url.PathEscape(hostPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-03-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// retrieveRegistrationTokenHandleResponse handles the RetrieveRegistrationToken response.
func (client *HostPoolsClient) retrieveRegistrationTokenHandleResponse(resp *http.Response) (HostPoolsRetrieveRegistrationTokenResponse, error) {
	result := HostPoolsRetrieveRegistrationTokenResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RegistrationInfo); err != nil {
		return HostPoolsRetrieveRegistrationTokenResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// retrieveRegistrationTokenHandleError handles the RetrieveRegistrationToken error response.
func (client *HostPoolsClient) retrieveRegistrationTokenHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Update - Update a host pool.
// If the operation fails it returns the *CloudError error type.
func (client *HostPoolsClient) Update(ctx context.Context, resourceGroupName string, hostPoolName string, options *HostPoolsUpdateOptions) (HostPoolsUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, hostPoolName, options)
	if err != nil {
		return HostPoolsUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HostPoolsUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HostPoolsUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *HostPoolsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, hostPoolName string, options *HostPoolsUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostPoolName == "" {
		return nil, errors.New("parameter hostPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostPoolName}", url.PathEscape(hostPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-03-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.HostPool != nil {
		return req, runtime.MarshalAsJSON(req, *options.HostPool)
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *HostPoolsClient) updateHandleResponse(resp *http.Response) (HostPoolsUpdateResponse, error) {
	result := HostPoolsUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.HostPool); err != nil {
		return HostPoolsUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *HostPoolsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
