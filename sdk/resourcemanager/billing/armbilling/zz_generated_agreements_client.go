//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbilling

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
	"strings"
)

// AgreementsClient contains the methods for the Agreements group.
// Don't use this type directly, use NewAgreementsClient() instead.
type AgreementsClient struct {
	ep string
	pl runtime.Pipeline
}

// NewAgreementsClient creates a new instance of AgreementsClient with the specified values.
func NewAgreementsClient(credential azcore.TokenCredential, options *arm.ClientOptions) *AgreementsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &AgreementsClient{ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Gets an agreement by ID.
// If the operation fails it returns the *ErrorResponse error type.
func (client *AgreementsClient) Get(ctx context.Context, billingAccountName string, agreementName string, options *AgreementsGetOptions) (AgreementsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, billingAccountName, agreementName, options)
	if err != nil {
		return AgreementsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AgreementsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AgreementsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AgreementsClient) getCreateRequest(ctx context.Context, billingAccountName string, agreementName string, options *AgreementsGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/agreements/{agreementName}"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if agreementName == "" {
		return nil, errors.New("parameter agreementName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agreementName}", url.PathEscape(agreementName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AgreementsClient) getHandleResponse(resp *http.Response) (AgreementsGetResponse, error) {
	result := AgreementsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Agreement); err != nil {
		return AgreementsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *AgreementsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByBillingAccount - Lists the agreements for a billing account.
// If the operation fails it returns the *ErrorResponse error type.
func (client *AgreementsClient) ListByBillingAccount(billingAccountName string, options *AgreementsListByBillingAccountOptions) *AgreementsListByBillingAccountPager {
	return &AgreementsListByBillingAccountPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByBillingAccountCreateRequest(ctx, billingAccountName, options)
		},
		advancer: func(ctx context.Context, resp AgreementsListByBillingAccountResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AgreementListResult.NextLink)
		},
	}
}

// listByBillingAccountCreateRequest creates the ListByBillingAccount request.
func (client *AgreementsClient) listByBillingAccountCreateRequest(ctx context.Context, billingAccountName string, options *AgreementsListByBillingAccountOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/agreements"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByBillingAccountHandleResponse handles the ListByBillingAccount response.
func (client *AgreementsClient) listByBillingAccountHandleResponse(resp *http.Response) (AgreementsListByBillingAccountResponse, error) {
	result := AgreementsListByBillingAccountResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AgreementListResult); err != nil {
		return AgreementsListByBillingAccountResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByBillingAccountHandleError handles the ListByBillingAccount error response.
func (client *AgreementsClient) listByBillingAccountHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
