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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations/v3"
	"net/http"
	"net/url"
	"regexp"
)

// CalculateRefundServer is a fake server for instances of the armreservations.CalculateRefundClient type.
type CalculateRefundServer struct {
	// Post is the fake for method CalculateRefundClient.Post
	// HTTP status codes to indicate success: http.StatusOK
	Post func(ctx context.Context, reservationOrderID string, body armreservations.CalculateRefundRequest, options *armreservations.CalculateRefundClientPostOptions) (resp azfake.Responder[armreservations.CalculateRefundClientPostResponse], errResp azfake.ErrorResponder)
}

// NewCalculateRefundServerTransport creates a new instance of CalculateRefundServerTransport with the provided implementation.
// The returned CalculateRefundServerTransport instance is connected to an instance of armreservations.CalculateRefundClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCalculateRefundServerTransport(srv *CalculateRefundServer) *CalculateRefundServerTransport {
	return &CalculateRefundServerTransport{srv: srv}
}

// CalculateRefundServerTransport connects instances of armreservations.CalculateRefundClient to instances of CalculateRefundServer.
// Don't use this type directly, use NewCalculateRefundServerTransport instead.
type CalculateRefundServerTransport struct {
	srv *CalculateRefundServer
}

// Do implements the policy.Transporter interface for CalculateRefundServerTransport.
func (c *CalculateRefundServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "CalculateRefundClient.Post":
		resp, err = c.dispatchPost(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CalculateRefundServerTransport) dispatchPost(req *http.Request) (*http.Response, error) {
	if c.srv.Post == nil {
		return nil, &nonRetriableError{errors.New("fake for method Post not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Capacity/reservationOrders/(?P<reservationOrderId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/calculateRefund`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armreservations.CalculateRefundRequest](req)
	if err != nil {
		return nil, err
	}
	reservationOrderIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("reservationOrderId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Post(req.Context(), reservationOrderIDParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CalculateRefundResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}