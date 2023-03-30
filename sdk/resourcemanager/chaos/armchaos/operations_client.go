//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armchaos

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// OperationsClient contains the methods for the Operations group.
// Don't use this type directly, use NewOperationsClient() instead.
type OperationsClient struct {
	internal *arm.Client
}

// NewOperationsClient creates a new instance of OperationsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewOperationsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*OperationsClient, error) {
	cl, err := arm.NewClient(moduleName+".OperationsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &OperationsClient{
		internal: cl,
	}
	return client, nil
}

// NewListAllPager - Get a list all available Operations.
//
// Generated from API version 2021-09-15-preview
//   - options - OperationsClientListAllOptions contains the optional parameters for the OperationsClient.NewListAllPager method.
func (client *OperationsClient) NewListAllPager(options *OperationsClientListAllOptions) *runtime.Pager[OperationsClientListAllResponse] {
	return runtime.NewPager(runtime.PagingHandler[OperationsClientListAllResponse]{
		More: func(page OperationsClientListAllResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OperationsClientListAllResponse) (OperationsClientListAllResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAllCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return OperationsClientListAllResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return OperationsClientListAllResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return OperationsClientListAllResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAllHandleResponse(resp)
		},
	})
}

// listAllCreateRequest creates the ListAll request.
func (client *OperationsClient) listAllCreateRequest(ctx context.Context, options *OperationsClientListAllOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Chaos/operations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *OperationsClient) listAllHandleResponse(resp *http.Response) (OperationsClientListAllResponse, error) {
	result := OperationsClientListAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationListResult); err != nil {
		return OperationsClientListAllResponse{}, err
	}
	return result, nil
}