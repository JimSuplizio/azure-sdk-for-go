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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/guestconfiguration/armguestconfiguration"
	"net/http"
	"net/url"
	"regexp"
)

// AssignmentReportsVMSSServer is a fake server for instances of the armguestconfiguration.AssignmentReportsVMSSClient type.
type AssignmentReportsVMSSServer struct {
	// Get is the fake for method AssignmentReportsVMSSClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, vmssName string, name string, id string, options *armguestconfiguration.AssignmentReportsVMSSClientGetOptions) (resp azfake.Responder[armguestconfiguration.AssignmentReportsVMSSClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method AssignmentReportsVMSSClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, vmssName string, name string, options *armguestconfiguration.AssignmentReportsVMSSClientListOptions) (resp azfake.PagerResponder[armguestconfiguration.AssignmentReportsVMSSClientListResponse])
}

// NewAssignmentReportsVMSSServerTransport creates a new instance of AssignmentReportsVMSSServerTransport with the provided implementation.
// The returned AssignmentReportsVMSSServerTransport instance is connected to an instance of armguestconfiguration.AssignmentReportsVMSSClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAssignmentReportsVMSSServerTransport(srv *AssignmentReportsVMSSServer) *AssignmentReportsVMSSServerTransport {
	return &AssignmentReportsVMSSServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armguestconfiguration.AssignmentReportsVMSSClientListResponse]](),
	}
}

// AssignmentReportsVMSSServerTransport connects instances of armguestconfiguration.AssignmentReportsVMSSClient to instances of AssignmentReportsVMSSServer.
// Don't use this type directly, use NewAssignmentReportsVMSSServerTransport instead.
type AssignmentReportsVMSSServerTransport struct {
	srv          *AssignmentReportsVMSSServer
	newListPager *tracker[azfake.PagerResponder[armguestconfiguration.AssignmentReportsVMSSClientListResponse]]
}

// Do implements the policy.Transporter interface for AssignmentReportsVMSSServerTransport.
func (a *AssignmentReportsVMSSServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AssignmentReportsVMSSClient.Get":
		resp, err = a.dispatchGet(req)
	case "AssignmentReportsVMSSClient.NewListPager":
		resp, err = a.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AssignmentReportsVMSSServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/virtualMachineScaleSets/(?P<vmssName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.GuestConfiguration/guestConfigurationAssignments/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reports/(?P<id>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	vmssNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vmssName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	idParam, err := url.PathUnescape(matches[regex.SubexpIndex("id")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, vmssNameParam, nameParam, idParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AssignmentReport, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AssignmentReportsVMSSServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := a.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/virtualMachineScaleSets/(?P<vmssName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.GuestConfiguration/guestConfigurationAssignments/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reports`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		vmssNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vmssName")])
		if err != nil {
			return nil, err
		}
		nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListPager(resourceGroupNameParam, vmssNameParam, nameParam, nil)
		newListPager = &resp
		a.newListPager.add(req, newListPager)
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		a.newListPager.remove(req)
	}
	return resp, nil
}