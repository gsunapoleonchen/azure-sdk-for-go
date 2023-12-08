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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
	"net/http"
	"regexp"
)

// PropertyServer is a fake server for instances of the armbilling.PropertyClient type.
type PropertyServer struct {
	// Get is the fake for method PropertyClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, options *armbilling.PropertyClientGetOptions) (resp azfake.Responder[armbilling.PropertyClientGetResponse], errResp azfake.ErrorResponder)

	// Update is the fake for method PropertyClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, parameters armbilling.Property, options *armbilling.PropertyClientUpdateOptions) (resp azfake.Responder[armbilling.PropertyClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewPropertyServerTransport creates a new instance of PropertyServerTransport with the provided implementation.
// The returned PropertyServerTransport instance is connected to an instance of armbilling.PropertyClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPropertyServerTransport(srv *PropertyServer) *PropertyServerTransport {
	return &PropertyServerTransport{srv: srv}
}

// PropertyServerTransport connects instances of armbilling.PropertyClient to instances of PropertyServer.
// Don't use this type directly, use NewPropertyServerTransport instead.
type PropertyServerTransport struct {
	srv *PropertyServer
}

// Do implements the policy.Transporter interface for PropertyServerTransport.
func (p *PropertyServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PropertyClient.Get":
		resp, err = p.dispatchGet(req)
	case "PropertyClient.Update":
		resp, err = p.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PropertyServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Billing/billingProperty/default`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := p.srv.Get(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Property, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PropertyServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Billing/billingProperty/default`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armbilling.Property](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Update(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Property, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}