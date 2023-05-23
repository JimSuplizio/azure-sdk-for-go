//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armresourcehealth

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// EmergingIssuesClient contains the methods for the EmergingIssues group.
// Don't use this type directly, use NewEmergingIssuesClient() instead.
type EmergingIssuesClient struct {
	internal *arm.Client
}

// NewEmergingIssuesClient creates a new instance of EmergingIssuesClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewEmergingIssuesClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*EmergingIssuesClient, error) {
	cl, err := arm.NewClient(moduleName+".EmergingIssuesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &EmergingIssuesClient{
		internal: cl,
	}
	return client, nil
}

// Get - Gets Azure services' emerging issues.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01
//   - issueName - The name of the emerging issue.
//   - options - EmergingIssuesClientGetOptions contains the optional parameters for the EmergingIssuesClient.Get method.
func (client *EmergingIssuesClient) Get(ctx context.Context, issueName IssueNameParameter, options *EmergingIssuesClientGetOptions) (EmergingIssuesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, issueName, options)
	if err != nil {
		return EmergingIssuesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EmergingIssuesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EmergingIssuesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *EmergingIssuesClient) getCreateRequest(ctx context.Context, issueName IssueNameParameter, options *EmergingIssuesClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.ResourceHealth/emergingIssues/{issueName}"
	if issueName == "" {
		return nil, errors.New("parameter issueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{issueName}", url.PathEscape(string(issueName)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *EmergingIssuesClient) getHandleResponse(resp *http.Response) (EmergingIssuesClientGetResponse, error) {
	result := EmergingIssuesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EmergingIssuesGetResult); err != nil {
		return EmergingIssuesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists Azure services' emerging issues.
//
// Generated from API version 2022-10-01
//   - options - EmergingIssuesClientListOptions contains the optional parameters for the EmergingIssuesClient.NewListPager method.
func (client *EmergingIssuesClient) NewListPager(options *EmergingIssuesClientListOptions) *runtime.Pager[EmergingIssuesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[EmergingIssuesClientListResponse]{
		More: func(page EmergingIssuesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *EmergingIssuesClientListResponse) (EmergingIssuesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return EmergingIssuesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return EmergingIssuesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return EmergingIssuesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *EmergingIssuesClient) listCreateRequest(ctx context.Context, options *EmergingIssuesClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.ResourceHealth/emergingIssues"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *EmergingIssuesClient) listHandleResponse(resp *http.Response) (EmergingIssuesClientListResponse, error) {
	result := EmergingIssuesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EmergingIssueListResult); err != nil {
		return EmergingIssuesClientListResponse{}, err
	}
	return result, nil
}