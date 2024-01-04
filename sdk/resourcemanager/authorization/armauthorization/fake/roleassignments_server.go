//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
	"net/http"
	"net/url"
	"regexp"
)

// RoleAssignmentsServer is a fake server for instances of the armauthorization.RoleAssignmentsClient type.
type RoleAssignmentsServer struct {
	// Create is the fake for method RoleAssignmentsClient.Create
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	Create func(ctx context.Context, scope string, roleAssignmentName string, parameters armauthorization.RoleAssignmentCreateParameters, options *armauthorization.RoleAssignmentsClientCreateOptions) (resp azfake.Responder[armauthorization.RoleAssignmentsClientCreateResponse], errResp azfake.ErrorResponder)

	// CreateByID is the fake for method RoleAssignmentsClient.CreateByID
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateByID func(ctx context.Context, roleAssignmentID string, parameters armauthorization.RoleAssignmentCreateParameters, options *armauthorization.RoleAssignmentsClientCreateByIDOptions) (resp azfake.Responder[armauthorization.RoleAssignmentsClientCreateByIDResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method RoleAssignmentsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, scope string, roleAssignmentName string, options *armauthorization.RoleAssignmentsClientDeleteOptions) (resp azfake.Responder[armauthorization.RoleAssignmentsClientDeleteResponse], errResp azfake.ErrorResponder)

	// DeleteByID is the fake for method RoleAssignmentsClient.DeleteByID
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	DeleteByID func(ctx context.Context, roleAssignmentID string, options *armauthorization.RoleAssignmentsClientDeleteByIDOptions) (resp azfake.Responder[armauthorization.RoleAssignmentsClientDeleteByIDResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method RoleAssignmentsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, scope string, roleAssignmentName string, options *armauthorization.RoleAssignmentsClientGetOptions) (resp azfake.Responder[armauthorization.RoleAssignmentsClientGetResponse], errResp azfake.ErrorResponder)

	// GetByID is the fake for method RoleAssignmentsClient.GetByID
	// HTTP status codes to indicate success: http.StatusOK
	GetByID func(ctx context.Context, roleAssignmentID string, options *armauthorization.RoleAssignmentsClientGetByIDOptions) (resp azfake.Responder[armauthorization.RoleAssignmentsClientGetByIDResponse], errResp azfake.ErrorResponder)

	// NewListForResourcePager is the fake for method RoleAssignmentsClient.NewListForResourcePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListForResourcePager func(resourceGroupName string, resourceProviderNamespace string, resourceType string, resourceName string, options *armauthorization.RoleAssignmentsClientListForResourceOptions) (resp azfake.PagerResponder[armauthorization.RoleAssignmentsClientListForResourceResponse])

	// NewListForResourceGroupPager is the fake for method RoleAssignmentsClient.NewListForResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListForResourceGroupPager func(resourceGroupName string, options *armauthorization.RoleAssignmentsClientListForResourceGroupOptions) (resp azfake.PagerResponder[armauthorization.RoleAssignmentsClientListForResourceGroupResponse])

	// NewListForScopePager is the fake for method RoleAssignmentsClient.NewListForScopePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListForScopePager func(scope string, options *armauthorization.RoleAssignmentsClientListForScopeOptions) (resp azfake.PagerResponder[armauthorization.RoleAssignmentsClientListForScopeResponse])

	// NewListForSubscriptionPager is the fake for method RoleAssignmentsClient.NewListForSubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListForSubscriptionPager func(options *armauthorization.RoleAssignmentsClientListForSubscriptionOptions) (resp azfake.PagerResponder[armauthorization.RoleAssignmentsClientListForSubscriptionResponse])
}

// NewRoleAssignmentsServerTransport creates a new instance of RoleAssignmentsServerTransport with the provided implementation.
// The returned RoleAssignmentsServerTransport instance is connected to an instance of armauthorization.RoleAssignmentsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRoleAssignmentsServerTransport(srv *RoleAssignmentsServer) *RoleAssignmentsServerTransport {
	return &RoleAssignmentsServerTransport{
		srv:                          srv,
		newListForResourcePager:      newTracker[azfake.PagerResponder[armauthorization.RoleAssignmentsClientListForResourceResponse]](),
		newListForResourceGroupPager: newTracker[azfake.PagerResponder[armauthorization.RoleAssignmentsClientListForResourceGroupResponse]](),
		newListForScopePager:         newTracker[azfake.PagerResponder[armauthorization.RoleAssignmentsClientListForScopeResponse]](),
		newListForSubscriptionPager:  newTracker[azfake.PagerResponder[armauthorization.RoleAssignmentsClientListForSubscriptionResponse]](),
	}
}

// RoleAssignmentsServerTransport connects instances of armauthorization.RoleAssignmentsClient to instances of RoleAssignmentsServer.
// Don't use this type directly, use NewRoleAssignmentsServerTransport instead.
type RoleAssignmentsServerTransport struct {
	srv                          *RoleAssignmentsServer
	newListForResourcePager      *tracker[azfake.PagerResponder[armauthorization.RoleAssignmentsClientListForResourceResponse]]
	newListForResourceGroupPager *tracker[azfake.PagerResponder[armauthorization.RoleAssignmentsClientListForResourceGroupResponse]]
	newListForScopePager         *tracker[azfake.PagerResponder[armauthorization.RoleAssignmentsClientListForScopeResponse]]
	newListForSubscriptionPager  *tracker[azfake.PagerResponder[armauthorization.RoleAssignmentsClientListForSubscriptionResponse]]
}

// Do implements the policy.Transporter interface for RoleAssignmentsServerTransport.
func (r *RoleAssignmentsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "RoleAssignmentsClient.Create":
		resp, err = r.dispatchCreate(req)
	case "RoleAssignmentsClient.CreateByID":
		resp, err = r.dispatchCreateByID(req)
	case "RoleAssignmentsClient.Delete":
		resp, err = r.dispatchDelete(req)
	case "RoleAssignmentsClient.DeleteByID":
		resp, err = r.dispatchDeleteByID(req)
	case "RoleAssignmentsClient.Get":
		resp, err = r.dispatchGet(req)
	case "RoleAssignmentsClient.GetByID":
		resp, err = r.dispatchGetByID(req)
	case "RoleAssignmentsClient.NewListForResourcePager":
		resp, err = r.dispatchNewListForResourcePager(req)
	case "RoleAssignmentsClient.NewListForResourceGroupPager":
		resp, err = r.dispatchNewListForResourceGroupPager(req)
	case "RoleAssignmentsClient.NewListForScopePager":
		resp, err = r.dispatchNewListForScopePager(req)
	case "RoleAssignmentsClient.NewListForSubscriptionPager":
		resp, err = r.dispatchNewListForSubscriptionPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *RoleAssignmentsServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if r.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/roleAssignments/(?P<roleAssignmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armauthorization.RoleAssignmentCreateParameters](req)
	if err != nil {
		return nil, err
	}
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	roleAssignmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("roleAssignmentName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Create(req.Context(), scopeParam, roleAssignmentNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RoleAssignment, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RoleAssignmentsServerTransport) dispatchCreateByID(req *http.Request) (*http.Response, error) {
	if r.srv.CreateByID == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateByID not implemented")}
	}
	const regexStr = `/(?P<roleAssignmentId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armauthorization.RoleAssignmentCreateParameters](req)
	if err != nil {
		return nil, err
	}
	roleAssignmentIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("roleAssignmentId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.CreateByID(req.Context(), roleAssignmentIDParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RoleAssignment, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RoleAssignmentsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if r.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/roleAssignments/(?P<roleAssignmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	roleAssignmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("roleAssignmentName")])
	if err != nil {
		return nil, err
	}
	tenantIDUnescaped, err := url.QueryUnescape(qp.Get("tenantId"))
	if err != nil {
		return nil, err
	}
	tenantIDParam := getOptional(tenantIDUnescaped)
	var options *armauthorization.RoleAssignmentsClientDeleteOptions
	if tenantIDParam != nil {
		options = &armauthorization.RoleAssignmentsClientDeleteOptions{
			TenantID: tenantIDParam,
		}
	}
	respr, errRespr := r.srv.Delete(req.Context(), scopeParam, roleAssignmentNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RoleAssignment, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RoleAssignmentsServerTransport) dispatchDeleteByID(req *http.Request) (*http.Response, error) {
	if r.srv.DeleteByID == nil {
		return nil, &nonRetriableError{errors.New("fake for method DeleteByID not implemented")}
	}
	const regexStr = `/(?P<roleAssignmentId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	roleAssignmentIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("roleAssignmentId")])
	if err != nil {
		return nil, err
	}
	tenantIDUnescaped, err := url.QueryUnescape(qp.Get("tenantId"))
	if err != nil {
		return nil, err
	}
	tenantIDParam := getOptional(tenantIDUnescaped)
	var options *armauthorization.RoleAssignmentsClientDeleteByIDOptions
	if tenantIDParam != nil {
		options = &armauthorization.RoleAssignmentsClientDeleteByIDOptions{
			TenantID: tenantIDParam,
		}
	}
	respr, errRespr := r.srv.DeleteByID(req.Context(), roleAssignmentIDParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RoleAssignment, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RoleAssignmentsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/roleAssignments/(?P<roleAssignmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	roleAssignmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("roleAssignmentName")])
	if err != nil {
		return nil, err
	}
	tenantIDUnescaped, err := url.QueryUnescape(qp.Get("tenantId"))
	if err != nil {
		return nil, err
	}
	tenantIDParam := getOptional(tenantIDUnescaped)
	var options *armauthorization.RoleAssignmentsClientGetOptions
	if tenantIDParam != nil {
		options = &armauthorization.RoleAssignmentsClientGetOptions{
			TenantID: tenantIDParam,
		}
	}
	respr, errRespr := r.srv.Get(req.Context(), scopeParam, roleAssignmentNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RoleAssignment, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RoleAssignmentsServerTransport) dispatchGetByID(req *http.Request) (*http.Response, error) {
	if r.srv.GetByID == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetByID not implemented")}
	}
	const regexStr = `/(?P<roleAssignmentId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	roleAssignmentIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("roleAssignmentId")])
	if err != nil {
		return nil, err
	}
	tenantIDUnescaped, err := url.QueryUnescape(qp.Get("tenantId"))
	if err != nil {
		return nil, err
	}
	tenantIDParam := getOptional(tenantIDUnescaped)
	var options *armauthorization.RoleAssignmentsClientGetByIDOptions
	if tenantIDParam != nil {
		options = &armauthorization.RoleAssignmentsClientGetByIDOptions{
			TenantID: tenantIDParam,
		}
	}
	respr, errRespr := r.srv.GetByID(req.Context(), roleAssignmentIDParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RoleAssignment, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RoleAssignmentsServerTransport) dispatchNewListForResourcePager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListForResourcePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListForResourcePager not implemented")}
	}
	newListForResourcePager := r.newListForResourcePager.get(req)
	if newListForResourcePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/(?P<resourceProviderNamespace>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<resourceType>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/roleAssignments`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resourceProviderNamespaceParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceProviderNamespace")])
		if err != nil {
			return nil, err
		}
		resourceTypeParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceType")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		tenantIDUnescaped, err := url.QueryUnescape(qp.Get("tenantId"))
		if err != nil {
			return nil, err
		}
		tenantIDParam := getOptional(tenantIDUnescaped)
		var options *armauthorization.RoleAssignmentsClientListForResourceOptions
		if filterParam != nil || tenantIDParam != nil {
			options = &armauthorization.RoleAssignmentsClientListForResourceOptions{
				Filter:   filterParam,
				TenantID: tenantIDParam,
			}
		}
		resp := r.srv.NewListForResourcePager(resourceGroupNameParam, resourceProviderNamespaceParam, resourceTypeParam, resourceNameParam, options)
		newListForResourcePager = &resp
		r.newListForResourcePager.add(req, newListForResourcePager)
		server.PagerResponderInjectNextLinks(newListForResourcePager, req, func(page *armauthorization.RoleAssignmentsClientListForResourceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListForResourcePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListForResourcePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListForResourcePager) {
		r.newListForResourcePager.remove(req)
	}
	return resp, nil
}

func (r *RoleAssignmentsServerTransport) dispatchNewListForResourceGroupPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListForResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListForResourceGroupPager not implemented")}
	}
	newListForResourceGroupPager := r.newListForResourceGroupPager.get(req)
	if newListForResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/roleAssignments`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		tenantIDUnescaped, err := url.QueryUnescape(qp.Get("tenantId"))
		if err != nil {
			return nil, err
		}
		tenantIDParam := getOptional(tenantIDUnescaped)
		var options *armauthorization.RoleAssignmentsClientListForResourceGroupOptions
		if filterParam != nil || tenantIDParam != nil {
			options = &armauthorization.RoleAssignmentsClientListForResourceGroupOptions{
				Filter:   filterParam,
				TenantID: tenantIDParam,
			}
		}
		resp := r.srv.NewListForResourceGroupPager(resourceGroupNameParam, options)
		newListForResourceGroupPager = &resp
		r.newListForResourceGroupPager.add(req, newListForResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListForResourceGroupPager, req, func(page *armauthorization.RoleAssignmentsClientListForResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListForResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListForResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListForResourceGroupPager) {
		r.newListForResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (r *RoleAssignmentsServerTransport) dispatchNewListForScopePager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListForScopePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListForScopePager not implemented")}
	}
	newListForScopePager := r.newListForScopePager.get(req)
	if newListForScopePager == nil {
		const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/roleAssignments`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		tenantIDUnescaped, err := url.QueryUnescape(qp.Get("tenantId"))
		if err != nil {
			return nil, err
		}
		tenantIDParam := getOptional(tenantIDUnescaped)
		skipTokenUnescaped, err := url.QueryUnescape(qp.Get("$skipToken"))
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(skipTokenUnescaped)
		var options *armauthorization.RoleAssignmentsClientListForScopeOptions
		if filterParam != nil || tenantIDParam != nil || skipTokenParam != nil {
			options = &armauthorization.RoleAssignmentsClientListForScopeOptions{
				Filter:    filterParam,
				TenantID:  tenantIDParam,
				SkipToken: skipTokenParam,
			}
		}
		resp := r.srv.NewListForScopePager(scopeParam, options)
		newListForScopePager = &resp
		r.newListForScopePager.add(req, newListForScopePager)
		server.PagerResponderInjectNextLinks(newListForScopePager, req, func(page *armauthorization.RoleAssignmentsClientListForScopeResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListForScopePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListForScopePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListForScopePager) {
		r.newListForScopePager.remove(req)
	}
	return resp, nil
}

func (r *RoleAssignmentsServerTransport) dispatchNewListForSubscriptionPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListForSubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListForSubscriptionPager not implemented")}
	}
	newListForSubscriptionPager := r.newListForSubscriptionPager.get(req)
	if newListForSubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/roleAssignments`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		tenantIDUnescaped, err := url.QueryUnescape(qp.Get("tenantId"))
		if err != nil {
			return nil, err
		}
		tenantIDParam := getOptional(tenantIDUnescaped)
		var options *armauthorization.RoleAssignmentsClientListForSubscriptionOptions
		if filterParam != nil || tenantIDParam != nil {
			options = &armauthorization.RoleAssignmentsClientListForSubscriptionOptions{
				Filter:   filterParam,
				TenantID: tenantIDParam,
			}
		}
		resp := r.srv.NewListForSubscriptionPager(options)
		newListForSubscriptionPager = &resp
		r.newListForSubscriptionPager.add(req, newListForSubscriptionPager)
		server.PagerResponderInjectNextLinks(newListForSubscriptionPager, req, func(page *armauthorization.RoleAssignmentsClientListForSubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListForSubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListForSubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListForSubscriptionPager) {
		r.newListForSubscriptionPager.remove(req)
	}
	return resp, nil
}