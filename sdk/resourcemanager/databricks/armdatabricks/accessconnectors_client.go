//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatabricks

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

// AccessConnectorsClient contains the methods for the AccessConnectors group.
// Don't use this type directly, use NewAccessConnectorsClient() instead.
type AccessConnectorsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAccessConnectorsClient creates a new instance of AccessConnectorsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAccessConnectorsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AccessConnectorsClient, error) {
	cl, err := arm.NewClient(moduleName+".AccessConnectorsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AccessConnectorsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates azure databricks accessConnector.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - connectorName - The name of the azure databricks accessConnector.
//   - parameters - Parameters supplied to the create or update an azure databricks accessConnector.
//   - options - AccessConnectorsClientBeginCreateOrUpdateOptions contains the optional parameters for the AccessConnectorsClient.BeginCreateOrUpdate
//     method.
func (client *AccessConnectorsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, connectorName string, parameters AccessConnector, options *AccessConnectorsClientBeginCreateOrUpdateOptions) (*runtime.Poller[AccessConnectorsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, connectorName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[AccessConnectorsClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[AccessConnectorsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates azure databricks accessConnector.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
func (client *AccessConnectorsClient) createOrUpdate(ctx context.Context, resourceGroupName string, connectorName string, parameters AccessConnector, options *AccessConnectorsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, connectorName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AccessConnectorsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, connectorName string, parameters AccessConnector, options *AccessConnectorsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Databricks/accessConnectors/{connectorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if connectorName == "" {
		return nil, errors.New("parameter connectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectorName}", url.PathEscape(connectorName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the azure databricks accessConnector.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - connectorName - The name of the azure databricks accessConnector.
//   - options - AccessConnectorsClientBeginDeleteOptions contains the optional parameters for the AccessConnectorsClient.BeginDelete
//     method.
func (client *AccessConnectorsClient) BeginDelete(ctx context.Context, resourceGroupName string, connectorName string, options *AccessConnectorsClientBeginDeleteOptions) (*runtime.Poller[AccessConnectorsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, connectorName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[AccessConnectorsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[AccessConnectorsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the azure databricks accessConnector.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
func (client *AccessConnectorsClient) deleteOperation(ctx context.Context, resourceGroupName string, connectorName string, options *AccessConnectorsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, connectorName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AccessConnectorsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, connectorName string, options *AccessConnectorsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Databricks/accessConnectors/{connectorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if connectorName == "" {
		return nil, errors.New("parameter connectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectorName}", url.PathEscape(connectorName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets an azure databricks accessConnector.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - connectorName - The name of the azure databricks accessConnector.
//   - options - AccessConnectorsClientGetOptions contains the optional parameters for the AccessConnectorsClient.Get method.
func (client *AccessConnectorsClient) Get(ctx context.Context, resourceGroupName string, connectorName string, options *AccessConnectorsClientGetOptions) (AccessConnectorsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, connectorName, options)
	if err != nil {
		return AccessConnectorsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AccessConnectorsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AccessConnectorsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AccessConnectorsClient) getCreateRequest(ctx context.Context, resourceGroupName string, connectorName string, options *AccessConnectorsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Databricks/accessConnectors/{connectorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if connectorName == "" {
		return nil, errors.New("parameter connectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectorName}", url.PathEscape(connectorName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AccessConnectorsClient) getHandleResponse(resp *http.Response) (AccessConnectorsClientGetResponse, error) {
	result := AccessConnectorsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessConnector); err != nil {
		return AccessConnectorsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets all the azure databricks accessConnectors within a resource group.
//
// Generated from API version 2023-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - AccessConnectorsClientListByResourceGroupOptions contains the optional parameters for the AccessConnectorsClient.NewListByResourceGroupPager
//     method.
func (client *AccessConnectorsClient) NewListByResourceGroupPager(resourceGroupName string, options *AccessConnectorsClientListByResourceGroupOptions) *runtime.Pager[AccessConnectorsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[AccessConnectorsClientListByResourceGroupResponse]{
		More: func(page AccessConnectorsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AccessConnectorsClientListByResourceGroupResponse) (AccessConnectorsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AccessConnectorsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return AccessConnectorsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AccessConnectorsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *AccessConnectorsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *AccessConnectorsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Databricks/accessConnectors"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *AccessConnectorsClient) listByResourceGroupHandleResponse(resp *http.Response) (AccessConnectorsClientListByResourceGroupResponse, error) {
	result := AccessConnectorsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessConnectorListResult); err != nil {
		return AccessConnectorsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Gets all the azure databricks accessConnectors within a subscription.
//
// Generated from API version 2023-05-01
//   - options - AccessConnectorsClientListBySubscriptionOptions contains the optional parameters for the AccessConnectorsClient.NewListBySubscriptionPager
//     method.
func (client *AccessConnectorsClient) NewListBySubscriptionPager(options *AccessConnectorsClientListBySubscriptionOptions) *runtime.Pager[AccessConnectorsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[AccessConnectorsClientListBySubscriptionResponse]{
		More: func(page AccessConnectorsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AccessConnectorsClientListBySubscriptionResponse) (AccessConnectorsClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AccessConnectorsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return AccessConnectorsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AccessConnectorsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *AccessConnectorsClient) listBySubscriptionCreateRequest(ctx context.Context, options *AccessConnectorsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Databricks/accessConnectors"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *AccessConnectorsClient) listBySubscriptionHandleResponse(resp *http.Response) (AccessConnectorsClientListBySubscriptionResponse, error) {
	result := AccessConnectorsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessConnectorListResult); err != nil {
		return AccessConnectorsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates an azure databricks accessConnector.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - connectorName - The name of the azure databricks accessConnector.
//   - parameters - The update to the azure databricks accessConnector.
//   - options - AccessConnectorsClientBeginUpdateOptions contains the optional parameters for the AccessConnectorsClient.BeginUpdate
//     method.
func (client *AccessConnectorsClient) BeginUpdate(ctx context.Context, resourceGroupName string, connectorName string, parameters AccessConnectorUpdate, options *AccessConnectorsClientBeginUpdateOptions) (*runtime.Poller[AccessConnectorsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, connectorName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[AccessConnectorsClientUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[AccessConnectorsClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Updates an azure databricks accessConnector.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
func (client *AccessConnectorsClient) update(ctx context.Context, resourceGroupName string, connectorName string, parameters AccessConnectorUpdate, options *AccessConnectorsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, connectorName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *AccessConnectorsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, connectorName string, parameters AccessConnectorUpdate, options *AccessConnectorsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Databricks/accessConnectors/{connectorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if connectorName == "" {
		return nil, errors.New("parameter connectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectorName}", url.PathEscape(connectorName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}