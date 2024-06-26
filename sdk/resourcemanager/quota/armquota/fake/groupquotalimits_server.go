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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota"
	"net/http"
	"net/url"
	"regexp"
)

// GroupQuotaLimitsServer is a fake server for instances of the armquota.GroupQuotaLimitsClient type.
type GroupQuotaLimitsServer struct {
	// Get is the fake for method GroupQuotaLimitsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, managementGroupID string, groupQuotaName string, resourceProviderName string, resourceName string, filter string, options *armquota.GroupQuotaLimitsClientGetOptions) (resp azfake.Responder[armquota.GroupQuotaLimitsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method GroupQuotaLimitsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(managementGroupID string, groupQuotaName string, resourceProviderName string, filter string, options *armquota.GroupQuotaLimitsClientListOptions) (resp azfake.PagerResponder[armquota.GroupQuotaLimitsClientListResponse])
}

// NewGroupQuotaLimitsServerTransport creates a new instance of GroupQuotaLimitsServerTransport with the provided implementation.
// The returned GroupQuotaLimitsServerTransport instance is connected to an instance of armquota.GroupQuotaLimitsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewGroupQuotaLimitsServerTransport(srv *GroupQuotaLimitsServer) *GroupQuotaLimitsServerTransport {
	return &GroupQuotaLimitsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armquota.GroupQuotaLimitsClientListResponse]](),
	}
}

// GroupQuotaLimitsServerTransport connects instances of armquota.GroupQuotaLimitsClient to instances of GroupQuotaLimitsServer.
// Don't use this type directly, use NewGroupQuotaLimitsServerTransport instead.
type GroupQuotaLimitsServerTransport struct {
	srv          *GroupQuotaLimitsServer
	newListPager *tracker[azfake.PagerResponder[armquota.GroupQuotaLimitsClientListResponse]]
}

// Do implements the policy.Transporter interface for GroupQuotaLimitsServerTransport.
func (g *GroupQuotaLimitsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "GroupQuotaLimitsClient.Get":
		resp, err = g.dispatchGet(req)
	case "GroupQuotaLimitsClient.NewListPager":
		resp, err = g.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GroupQuotaLimitsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if g.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Management/managementGroups/(?P<managementGroupId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Quota/groupQuotas/(?P<groupQuotaName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceProviders/(?P<resourceProviderName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/groupQuotaLimits/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	managementGroupIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("managementGroupId")])
	if err != nil {
		return nil, err
	}
	groupQuotaNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("groupQuotaName")])
	if err != nil {
		return nil, err
	}
	resourceProviderNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceProviderName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	filterParam, err := url.QueryUnescape(qp.Get("$filter"))
	if err != nil {
		return nil, err
	}
	respr, errRespr := g.srv.Get(req.Context(), managementGroupIDParam, groupQuotaNameParam, resourceProviderNameParam, resourceNameParam, filterParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).GroupQuotaLimit, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *GroupQuotaLimitsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if g.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := g.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/providers/Microsoft\.Management/managementGroups/(?P<managementGroupId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Quota/groupQuotas/(?P<groupQuotaName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceProviders/(?P<resourceProviderName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/groupQuotaLimits`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		managementGroupIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("managementGroupId")])
		if err != nil {
			return nil, err
		}
		groupQuotaNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("groupQuotaName")])
		if err != nil {
			return nil, err
		}
		resourceProviderNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceProviderName")])
		if err != nil {
			return nil, err
		}
		filterParam, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		resp := g.srv.NewListPager(managementGroupIDParam, groupQuotaNameParam, resourceProviderNameParam, filterParam, nil)
		newListPager = &resp
		g.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armquota.GroupQuotaLimitsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		g.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		g.newListPager.remove(req)
	}
	return resp, nil
}
