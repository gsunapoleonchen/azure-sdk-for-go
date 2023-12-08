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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicelinker/armservicelinker"
	"net/http"
	"net/url"
	"regexp"
)

// LinkerServer is a fake server for instances of the armservicelinker.LinkerClient type.
type LinkerServer struct {
	// BeginCreateOrUpdate is the fake for method LinkerClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceURI string, linkerName string, parameters armservicelinker.LinkerResource, options *armservicelinker.LinkerClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armservicelinker.LinkerClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method LinkerClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceURI string, linkerName string, options *armservicelinker.LinkerClientBeginDeleteOptions) (resp azfake.PollerResponder[armservicelinker.LinkerClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method LinkerClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceURI string, linkerName string, options *armservicelinker.LinkerClientGetOptions) (resp azfake.Responder[armservicelinker.LinkerClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method LinkerClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceURI string, options *armservicelinker.LinkerClientListOptions) (resp azfake.PagerResponder[armservicelinker.LinkerClientListResponse])

	// ListConfigurations is the fake for method LinkerClient.ListConfigurations
	// HTTP status codes to indicate success: http.StatusOK
	ListConfigurations func(ctx context.Context, resourceURI string, linkerName string, options *armservicelinker.LinkerClientListConfigurationsOptions) (resp azfake.Responder[armservicelinker.LinkerClientListConfigurationsResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method LinkerClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginUpdate func(ctx context.Context, resourceURI string, linkerName string, parameters armservicelinker.LinkerPatch, options *armservicelinker.LinkerClientBeginUpdateOptions) (resp azfake.PollerResponder[armservicelinker.LinkerClientUpdateResponse], errResp azfake.ErrorResponder)

	// BeginValidate is the fake for method LinkerClient.BeginValidate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginValidate func(ctx context.Context, resourceURI string, linkerName string, options *armservicelinker.LinkerClientBeginValidateOptions) (resp azfake.PollerResponder[armservicelinker.LinkerClientValidateResponse], errResp azfake.ErrorResponder)
}

// NewLinkerServerTransport creates a new instance of LinkerServerTransport with the provided implementation.
// The returned LinkerServerTransport instance is connected to an instance of armservicelinker.LinkerClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewLinkerServerTransport(srv *LinkerServer) *LinkerServerTransport {
	return &LinkerServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armservicelinker.LinkerClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armservicelinker.LinkerClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armservicelinker.LinkerClientListResponse]](),
		beginUpdate:         newTracker[azfake.PollerResponder[armservicelinker.LinkerClientUpdateResponse]](),
		beginValidate:       newTracker[azfake.PollerResponder[armservicelinker.LinkerClientValidateResponse]](),
	}
}

// LinkerServerTransport connects instances of armservicelinker.LinkerClient to instances of LinkerServer.
// Don't use this type directly, use NewLinkerServerTransport instead.
type LinkerServerTransport struct {
	srv                 *LinkerServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armservicelinker.LinkerClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armservicelinker.LinkerClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armservicelinker.LinkerClientListResponse]]
	beginUpdate         *tracker[azfake.PollerResponder[armservicelinker.LinkerClientUpdateResponse]]
	beginValidate       *tracker[azfake.PollerResponder[armservicelinker.LinkerClientValidateResponse]]
}

// Do implements the policy.Transporter interface for LinkerServerTransport.
func (l *LinkerServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "LinkerClient.BeginCreateOrUpdate":
		resp, err = l.dispatchBeginCreateOrUpdate(req)
	case "LinkerClient.BeginDelete":
		resp, err = l.dispatchBeginDelete(req)
	case "LinkerClient.Get":
		resp, err = l.dispatchGet(req)
	case "LinkerClient.NewListPager":
		resp, err = l.dispatchNewListPager(req)
	case "LinkerClient.ListConfigurations":
		resp, err = l.dispatchListConfigurations(req)
	case "LinkerClient.BeginUpdate":
		resp, err = l.dispatchBeginUpdate(req)
	case "LinkerClient.BeginValidate":
		resp, err = l.dispatchBeginValidate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (l *LinkerServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if l.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := l.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceLinker/linkers/(?P<linkerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armservicelinker.LinkerResource](req)
		if err != nil {
			return nil, err
		}
		resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
		if err != nil {
			return nil, err
		}
		linkerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("linkerName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := l.srv.BeginCreateOrUpdate(req.Context(), resourceURIParam, linkerNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		l.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		l.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		l.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (l *LinkerServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if l.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := l.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceLinker/linkers/(?P<linkerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
		if err != nil {
			return nil, err
		}
		linkerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("linkerName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := l.srv.BeginDelete(req.Context(), resourceURIParam, linkerNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		l.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		l.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		l.beginDelete.remove(req)
	}

	return resp, nil
}

func (l *LinkerServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if l.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceLinker/linkers/(?P<linkerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
	if err != nil {
		return nil, err
	}
	linkerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("linkerName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := l.srv.Get(req.Context(), resourceURIParam, linkerNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).LinkerResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *LinkerServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if l.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := l.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceLinker/linkers`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
		if err != nil {
			return nil, err
		}
		resp := l.srv.NewListPager(resourceURIParam, nil)
		newListPager = &resp
		l.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armservicelinker.LinkerClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		l.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		l.newListPager.remove(req)
	}
	return resp, nil
}

func (l *LinkerServerTransport) dispatchListConfigurations(req *http.Request) (*http.Response, error) {
	if l.srv.ListConfigurations == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListConfigurations not implemented")}
	}
	const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceLinker/linkers/(?P<linkerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listConfigurations`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
	if err != nil {
		return nil, err
	}
	linkerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("linkerName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := l.srv.ListConfigurations(req.Context(), resourceURIParam, linkerNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SourceConfigurationResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *LinkerServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if l.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := l.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceLinker/linkers/(?P<linkerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armservicelinker.LinkerPatch](req)
		if err != nil {
			return nil, err
		}
		resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
		if err != nil {
			return nil, err
		}
		linkerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("linkerName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := l.srv.BeginUpdate(req.Context(), resourceURIParam, linkerNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		l.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		l.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		l.beginUpdate.remove(req)
	}

	return resp, nil
}

func (l *LinkerServerTransport) dispatchBeginValidate(req *http.Request) (*http.Response, error) {
	if l.srv.BeginValidate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginValidate not implemented")}
	}
	beginValidate := l.beginValidate.get(req)
	if beginValidate == nil {
		const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceLinker/linkers/(?P<linkerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/validateLinker`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
		if err != nil {
			return nil, err
		}
		linkerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("linkerName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := l.srv.BeginValidate(req.Context(), resourceURIParam, linkerNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginValidate = &respr
		l.beginValidate.add(req, beginValidate)
	}

	resp, err := server.PollerResponderNext(beginValidate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		l.beginValidate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginValidate) {
		l.beginValidate.remove(req)
	}

	return resp, nil
}