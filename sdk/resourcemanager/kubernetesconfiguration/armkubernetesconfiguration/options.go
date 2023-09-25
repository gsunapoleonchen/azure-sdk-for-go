//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkubernetesconfiguration

// ExtensionsClientBeginCreateOptions contains the optional parameters for the ExtensionsClient.BeginCreate method.
type ExtensionsClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ExtensionsClientBeginDeleteOptions contains the optional parameters for the ExtensionsClient.BeginDelete method.
type ExtensionsClientBeginDeleteOptions struct {
	// Delete the extension resource in Azure - not the normal asynchronous delete.
	ForceDelete *bool

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ExtensionsClientBeginUpdateOptions contains the optional parameters for the ExtensionsClient.BeginUpdate method.
type ExtensionsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ExtensionsClientGetOptions contains the optional parameters for the ExtensionsClient.Get method.
type ExtensionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ExtensionsClientListOptions contains the optional parameters for the ExtensionsClient.NewListPager method.
type ExtensionsClientListOptions struct {
	// placeholder for future optional parameters
}

// FluxConfigOperationStatusClientGetOptions contains the optional parameters for the FluxConfigOperationStatusClient.Get
// method.
type FluxConfigOperationStatusClientGetOptions struct {
	// placeholder for future optional parameters
}

// FluxConfigurationsClientBeginCreateOrUpdateOptions contains the optional parameters for the FluxConfigurationsClient.BeginCreateOrUpdate
// method.
type FluxConfigurationsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// FluxConfigurationsClientBeginDeleteOptions contains the optional parameters for the FluxConfigurationsClient.BeginDelete
// method.
type FluxConfigurationsClientBeginDeleteOptions struct {
	// Delete the extension resource in Azure - not the normal asynchronous delete.
	ForceDelete *bool

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// FluxConfigurationsClientBeginUpdateOptions contains the optional parameters for the FluxConfigurationsClient.BeginUpdate
// method.
type FluxConfigurationsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// FluxConfigurationsClientGetOptions contains the optional parameters for the FluxConfigurationsClient.Get method.
type FluxConfigurationsClientGetOptions struct {
	// placeholder for future optional parameters
}

// FluxConfigurationsClientListOptions contains the optional parameters for the FluxConfigurationsClient.NewListPager method.
type FluxConfigurationsClientListOptions struct {
	// placeholder for future optional parameters
}

// OperationStatusClientGetOptions contains the optional parameters for the OperationStatusClient.Get method.
type OperationStatusClientGetOptions struct {
	// placeholder for future optional parameters
}

// OperationStatusClientListOptions contains the optional parameters for the OperationStatusClient.NewListPager method.
type OperationStatusClientListOptions struct {
	// placeholder for future optional parameters
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// SourceControlConfigurationsClientBeginDeleteOptions contains the optional parameters for the SourceControlConfigurationsClient.BeginDelete
// method.
type SourceControlConfigurationsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// SourceControlConfigurationsClientCreateOrUpdateOptions contains the optional parameters for the SourceControlConfigurationsClient.CreateOrUpdate
// method.
type SourceControlConfigurationsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// SourceControlConfigurationsClientGetOptions contains the optional parameters for the SourceControlConfigurationsClient.Get
// method.
type SourceControlConfigurationsClientGetOptions struct {
	// placeholder for future optional parameters
}

// SourceControlConfigurationsClientListOptions contains the optional parameters for the SourceControlConfigurationsClient.NewListPager
// method.
type SourceControlConfigurationsClientListOptions struct {
	// placeholder for future optional parameters
}