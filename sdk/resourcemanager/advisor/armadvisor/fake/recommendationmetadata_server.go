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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor"
	"net/http"
	"net/url"
	"regexp"
)

// RecommendationMetadataServer is a fake server for instances of the armadvisor.RecommendationMetadataClient type.
type RecommendationMetadataServer struct {
	// Get is the fake for method RecommendationMetadataClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, name string, options *armadvisor.RecommendationMetadataClientGetOptions) (resp azfake.Responder[armadvisor.RecommendationMetadataClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method RecommendationMetadataClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armadvisor.RecommendationMetadataClientListOptions) (resp azfake.PagerResponder[armadvisor.RecommendationMetadataClientListResponse])
}

// NewRecommendationMetadataServerTransport creates a new instance of RecommendationMetadataServerTransport with the provided implementation.
// The returned RecommendationMetadataServerTransport instance is connected to an instance of armadvisor.RecommendationMetadataClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRecommendationMetadataServerTransport(srv *RecommendationMetadataServer) *RecommendationMetadataServerTransport {
	return &RecommendationMetadataServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armadvisor.RecommendationMetadataClientListResponse]](),
	}
}

// RecommendationMetadataServerTransport connects instances of armadvisor.RecommendationMetadataClient to instances of RecommendationMetadataServer.
// Don't use this type directly, use NewRecommendationMetadataServerTransport instead.
type RecommendationMetadataServerTransport struct {
	srv          *RecommendationMetadataServer
	newListPager *tracker[azfake.PagerResponder[armadvisor.RecommendationMetadataClientListResponse]]
}

// Do implements the policy.Transporter interface for RecommendationMetadataServerTransport.
func (r *RecommendationMetadataServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "RecommendationMetadataClient.Get":
		resp, err = r.dispatchGet(req)
	case "RecommendationMetadataClient.NewListPager":
		resp, err = r.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *RecommendationMetadataServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Advisor/metadata/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), nameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).MetadataEntity, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RecommendationMetadataServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := r.newListPager.get(req)
	if newListPager == nil {
		resp := r.srv.NewListPager(nil)
		newListPager = &resp
		r.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armadvisor.RecommendationMetadataClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		r.newListPager.remove(req)
	}
	return resp, nil
}