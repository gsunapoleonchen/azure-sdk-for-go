//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhealthcareapis

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DicomServicesClient contains the methods for the DicomServices group.
// Don't use this type directly, use NewDicomServicesClient() instead.
type DicomServicesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDicomServicesClient creates a new instance of DicomServicesClient with the specified values.
// subscriptionID - The subscription identifier.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDicomServicesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DicomServicesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &DicomServicesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a DICOM Service resource with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the service instance.
// workspaceName - The name of workspace resource.
// dicomServiceName - The name of DICOM Service resource.
// dicomservice - The parameters for creating or updating a Dicom Service resource.
// options - DicomServicesClientBeginCreateOrUpdateOptions contains the optional parameters for the DicomServicesClient.BeginCreateOrUpdate
// method.
func (client *DicomServicesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, dicomServiceName string, dicomservice DicomService, options *DicomServicesClientBeginCreateOrUpdateOptions) (*armruntime.Poller[DicomServicesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, dicomServiceName, dicomservice, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[DicomServicesClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[DicomServicesClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a DICOM Service resource with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DicomServicesClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, dicomServiceName string, dicomservice DicomService, options *DicomServicesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, dicomServiceName, dicomservice, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DicomServicesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, dicomServiceName string, dicomservice DicomService, options *DicomServicesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/dicomservices/{dicomServiceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if dicomServiceName == "" {
		return nil, errors.New("parameter dicomServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dicomServiceName}", url.PathEscape(dicomServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, dicomservice)
}

// BeginDelete - Deletes a DICOM Service.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the service instance.
// dicomServiceName - The name of DICOM Service resource.
// workspaceName - The name of workspace resource.
// options - DicomServicesClientBeginDeleteOptions contains the optional parameters for the DicomServicesClient.BeginDelete
// method.
func (client *DicomServicesClient) BeginDelete(ctx context.Context, resourceGroupName string, dicomServiceName string, workspaceName string, options *DicomServicesClientBeginDeleteOptions) (*armruntime.Poller[DicomServicesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, dicomServiceName, workspaceName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[DicomServicesClientDeleteResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[DicomServicesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a DICOM Service.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DicomServicesClient) deleteOperation(ctx context.Context, resourceGroupName string, dicomServiceName string, workspaceName string, options *DicomServicesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, dicomServiceName, workspaceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DicomServicesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, dicomServiceName string, workspaceName string, options *DicomServicesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/dicomservices/{dicomServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dicomServiceName == "" {
		return nil, errors.New("parameter dicomServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dicomServiceName}", url.PathEscape(dicomServiceName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets the properties of the specified DICOM Service.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the service instance.
// workspaceName - The name of workspace resource.
// dicomServiceName - The name of DICOM Service resource.
// options - DicomServicesClientGetOptions contains the optional parameters for the DicomServicesClient.Get method.
func (client *DicomServicesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, dicomServiceName string, options *DicomServicesClientGetOptions) (DicomServicesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, dicomServiceName, options)
	if err != nil {
		return DicomServicesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DicomServicesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DicomServicesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DicomServicesClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, dicomServiceName string, options *DicomServicesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/dicomservices/{dicomServiceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if dicomServiceName == "" {
		return nil, errors.New("parameter dicomServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dicomServiceName}", url.PathEscape(dicomServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DicomServicesClient) getHandleResponse(resp *http.Response) (DicomServicesClientGetResponse, error) {
	result := DicomServicesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DicomService); err != nil {
		return DicomServicesClientGetResponse{}, err
	}
	return result, nil
}

// ListByWorkspace - Lists all DICOM Services for the given workspace
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the service instance.
// workspaceName - The name of workspace resource.
// options - DicomServicesClientListByWorkspaceOptions contains the optional parameters for the DicomServicesClient.ListByWorkspace
// method.
func (client *DicomServicesClient) ListByWorkspace(resourceGroupName string, workspaceName string, options *DicomServicesClientListByWorkspaceOptions) *runtime.Pager[DicomServicesClientListByWorkspaceResponse] {
	return runtime.NewPager(runtime.PageProcessor[DicomServicesClientListByWorkspaceResponse]{
		More: func(page DicomServicesClientListByWorkspaceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DicomServicesClientListByWorkspaceResponse) (DicomServicesClientListByWorkspaceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByWorkspaceCreateRequest(ctx, resourceGroupName, workspaceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DicomServicesClientListByWorkspaceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DicomServicesClientListByWorkspaceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DicomServicesClientListByWorkspaceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByWorkspaceHandleResponse(resp)
		},
	})
}

// listByWorkspaceCreateRequest creates the ListByWorkspace request.
func (client *DicomServicesClient) listByWorkspaceCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *DicomServicesClientListByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/dicomservices"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByWorkspaceHandleResponse handles the ListByWorkspace response.
func (client *DicomServicesClient) listByWorkspaceHandleResponse(resp *http.Response) (DicomServicesClientListByWorkspaceResponse, error) {
	result := DicomServicesClientListByWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DicomServiceCollection); err != nil {
		return DicomServicesClientListByWorkspaceResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Patch DICOM Service details.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the service instance.
// dicomServiceName - The name of DICOM Service resource.
// workspaceName - The name of workspace resource.
// dicomservicePatchResource - The parameters for updating a Dicom Service.
// options - DicomServicesClientBeginUpdateOptions contains the optional parameters for the DicomServicesClient.BeginUpdate
// method.
func (client *DicomServicesClient) BeginUpdate(ctx context.Context, resourceGroupName string, dicomServiceName string, workspaceName string, dicomservicePatchResource DicomServicePatchResource, options *DicomServicesClientBeginUpdateOptions) (*armruntime.Poller[DicomServicesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, dicomServiceName, workspaceName, dicomservicePatchResource, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[DicomServicesClientUpdateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[DicomServicesClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Patch DICOM Service details.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DicomServicesClient) update(ctx context.Context, resourceGroupName string, dicomServiceName string, workspaceName string, dicomservicePatchResource DicomServicePatchResource, options *DicomServicesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, dicomServiceName, workspaceName, dicomservicePatchResource, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *DicomServicesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, dicomServiceName string, workspaceName string, dicomservicePatchResource DicomServicePatchResource, options *DicomServicesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/dicomservices/{dicomServiceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if dicomServiceName == "" {
		return nil, errors.New("parameter dicomServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dicomServiceName}", url.PathEscape(dicomServiceName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, dicomservicePatchResource)
}
