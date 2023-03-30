//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armattestation

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationList
}

// PrivateEndpointConnectionsClientCreateResponse contains the response from method PrivateEndpointConnectionsClient.Create.
type PrivateEndpointConnectionsClientCreateResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.Delete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListResponse contains the response from method PrivateEndpointConnectionsClient.NewListPager.
type PrivateEndpointConnectionsClientListResponse struct {
	PrivateEndpointConnectionListResult
}

// ProvidersClientCreateResponse contains the response from method ProvidersClient.Create.
type ProvidersClientCreateResponse struct {
	Provider
}

// ProvidersClientDeleteResponse contains the response from method ProvidersClient.Delete.
type ProvidersClientDeleteResponse struct {
	// placeholder for future response values
}

// ProvidersClientGetDefaultByLocationResponse contains the response from method ProvidersClient.GetDefaultByLocation.
type ProvidersClientGetDefaultByLocationResponse struct {
	Provider
}

// ProvidersClientGetResponse contains the response from method ProvidersClient.Get.
type ProvidersClientGetResponse struct {
	Provider
}

// ProvidersClientListByResourceGroupResponse contains the response from method ProvidersClient.ListByResourceGroup.
type ProvidersClientListByResourceGroupResponse struct {
	ProviderListResult
}

// ProvidersClientListDefaultResponse contains the response from method ProvidersClient.ListDefault.
type ProvidersClientListDefaultResponse struct {
	ProviderListResult
}

// ProvidersClientListResponse contains the response from method ProvidersClient.List.
type ProvidersClientListResponse struct {
	ProviderListResult
}

// ProvidersClientUpdateResponse contains the response from method ProvidersClient.Update.
type ProvidersClientUpdateResponse struct {
	Provider
}