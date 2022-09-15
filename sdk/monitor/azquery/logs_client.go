//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azquery

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// LogsClient contains the methods for the LogsClient group.
// Don't use this type directly, use NewLogsClient() instead.
type LogsClient struct {
	pl runtime.Pipeline
}

// Batch - Executes a batch of Analytics queries for data. Here [https://dev.loganalytics.io/documentation/Using-the-API]
// is an example for using POST with an Analytics query.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-05-19_Preview
// body - The batch request body
// options - LogsClientBatchOptions contains the optional parameters for the LogsClient.Batch method.
func (client *LogsClient) Batch(ctx context.Context, body BatchRequest, options *LogsClientBatchOptions) (LogsClientBatchResponse, error) {
	req, err := client.batchCreateRequest(ctx, body, options)
	if err != nil {
		return LogsClientBatchResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LogsClientBatchResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LogsClientBatchResponse{}, runtime.NewResponseError(resp)
	}
	return client.batchHandleResponse(resp)
}

// batchCreateRequest creates the Batch request.
func (client *LogsClient) batchCreateRequest(ctx context.Context, body BatchRequest, options *LogsClientBatchOptions) (*policy.Request, error) {
	urlPath := "/$batch"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// batchHandleResponse handles the Batch response.
func (client *LogsClient) batchHandleResponse(resp *http.Response) (LogsClientBatchResponse, error) {
	result := LogsClientBatchResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BatchResponse); err != nil {
		return LogsClientBatchResponse{}, err
	}
	return result, nil
}

// QueryWorkspace - Executes an Analytics query for data. Here [https://dev.loganalytics.io/documentation/Using-the-API] is
// an example for using POST with an Analytics query.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-05-19_Preview
// workspaceID - ID of the workspace. This is Workspace ID from the Properties blade in the Azure portal.
// body - The Analytics query. Learn more about the Analytics query syntax [https://azure.microsoft.com/documentation/articles/app-insights-analytics-reference/]
// options - LogsClientQueryWorkspaceOptions contains the optional parameters for the LogsClient.QueryWorkspace method.
func (client *LogsClient) QueryWorkspace(ctx context.Context, workspaceID string, body Body, options *LogsClientQueryWorkspaceOptions) (LogsClientQueryWorkspaceResponse, error) {
	req, err := client.queryWorkspaceCreateRequest(ctx, workspaceID, body, options)
	if err != nil {
		return LogsClientQueryWorkspaceResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LogsClientQueryWorkspaceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LogsClientQueryWorkspaceResponse{}, runtime.NewResponseError(resp)
	}
	return client.queryWorkspaceHandleResponse(resp)
}

// queryWorkspaceCreateRequest creates the QueryWorkspace request.
func (client *LogsClient) queryWorkspaceCreateRequest(ctx context.Context, workspaceID string, body Body, options *LogsClientQueryWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/workspaces/{workspaceId}/query"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	if options != nil && options.Prefer != nil {
		req.Raw().Header["Prefer"] = []string{*options.Prefer}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// queryWorkspaceHandleResponse handles the QueryWorkspace response.
func (client *LogsClient) queryWorkspaceHandleResponse(resp *http.Response) (LogsClientQueryWorkspaceResponse, error) {
	result := LogsClientQueryWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Results); err != nil {
		return LogsClientQueryWorkspaceResponse{}, err
	}
	return result, nil
}