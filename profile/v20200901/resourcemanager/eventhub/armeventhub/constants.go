//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armeventhub

type AccessRights string

const (
	AccessRightsListen AccessRights = "Listen"
	AccessRightsManage AccessRights = "Manage"
	AccessRightsSend   AccessRights = "Send"
)

// PossibleAccessRightsValues returns the possible values for the AccessRights const type.
func PossibleAccessRightsValues() []AccessRights {
	return []AccessRights{
		AccessRightsListen,
		AccessRightsManage,
		AccessRightsSend,
	}
}

// ClusterSKUName - Name of this SKU.
type ClusterSKUName string

const (
	ClusterSKUNameDedicated ClusterSKUName = "Dedicated"
)

// PossibleClusterSKUNameValues returns the possible values for the ClusterSKUName const type.
func PossibleClusterSKUNameValues() []ClusterSKUName {
	return []ClusterSKUName{
		ClusterSKUNameDedicated,
	}
}

// EncodingCaptureDescription - Enumerates the possible values for the encoding format of capture description. Note: 'AvroDeflate'
// will be deprecated in New API Version
type EncodingCaptureDescription string

const (
	EncodingCaptureDescriptionAvro        EncodingCaptureDescription = "Avro"
	EncodingCaptureDescriptionAvroDeflate EncodingCaptureDescription = "AvroDeflate"
)

// PossibleEncodingCaptureDescriptionValues returns the possible values for the EncodingCaptureDescription const type.
func PossibleEncodingCaptureDescriptionValues() []EncodingCaptureDescription {
	return []EncodingCaptureDescription{
		EncodingCaptureDescriptionAvro,
		EncodingCaptureDescriptionAvroDeflate,
	}
}

// EndPointProvisioningState - Provisioning state of the Private Endpoint Connection.
type EndPointProvisioningState string

const (
	EndPointProvisioningStateCanceled  EndPointProvisioningState = "Canceled"
	EndPointProvisioningStateCreating  EndPointProvisioningState = "Creating"
	EndPointProvisioningStateDeleting  EndPointProvisioningState = "Deleting"
	EndPointProvisioningStateFailed    EndPointProvisioningState = "Failed"
	EndPointProvisioningStateSucceeded EndPointProvisioningState = "Succeeded"
	EndPointProvisioningStateUpdating  EndPointProvisioningState = "Updating"
)

// PossibleEndPointProvisioningStateValues returns the possible values for the EndPointProvisioningState const type.
func PossibleEndPointProvisioningStateValues() []EndPointProvisioningState {
	return []EndPointProvisioningState{
		EndPointProvisioningStateCanceled,
		EndPointProvisioningStateCreating,
		EndPointProvisioningStateDeleting,
		EndPointProvisioningStateFailed,
		EndPointProvisioningStateSucceeded,
		EndPointProvisioningStateUpdating,
	}
}

// EntityStatus - Enumerates the possible values for the status of the Event Hub.
type EntityStatus string

const (
	EntityStatusActive          EntityStatus = "Active"
	EntityStatusDisabled        EntityStatus = "Disabled"
	EntityStatusRestoring       EntityStatus = "Restoring"
	EntityStatusSendDisabled    EntityStatus = "SendDisabled"
	EntityStatusReceiveDisabled EntityStatus = "ReceiveDisabled"
	EntityStatusCreating        EntityStatus = "Creating"
	EntityStatusDeleting        EntityStatus = "Deleting"
	EntityStatusRenaming        EntityStatus = "Renaming"
	EntityStatusUnknown         EntityStatus = "Unknown"
)

// PossibleEntityStatusValues returns the possible values for the EntityStatus const type.
func PossibleEntityStatusValues() []EntityStatus {
	return []EntityStatus{
		EntityStatusActive,
		EntityStatusDisabled,
		EntityStatusRestoring,
		EntityStatusSendDisabled,
		EntityStatusReceiveDisabled,
		EntityStatusCreating,
		EntityStatusDeleting,
		EntityStatusRenaming,
		EntityStatusUnknown,
	}
}

// KeyType - The access key to regenerate.
type KeyType string

const (
	KeyTypePrimaryKey   KeyType = "PrimaryKey"
	KeyTypeSecondaryKey KeyType = "SecondaryKey"
)

// PossibleKeyTypeValues returns the possible values for the KeyType const type.
func PossibleKeyTypeValues() []KeyType {
	return []KeyType{
		KeyTypePrimaryKey,
		KeyTypeSecondaryKey,
	}
}

// PrivateLinkConnectionStatus - Status of the connection.
type PrivateLinkConnectionStatus string

const (
	PrivateLinkConnectionStatusApproved     PrivateLinkConnectionStatus = "Approved"
	PrivateLinkConnectionStatusDisconnected PrivateLinkConnectionStatus = "Disconnected"
	PrivateLinkConnectionStatusPending      PrivateLinkConnectionStatus = "Pending"
	PrivateLinkConnectionStatusRejected     PrivateLinkConnectionStatus = "Rejected"
)

// PossiblePrivateLinkConnectionStatusValues returns the possible values for the PrivateLinkConnectionStatus const type.
func PossiblePrivateLinkConnectionStatusValues() []PrivateLinkConnectionStatus {
	return []PrivateLinkConnectionStatus{
		PrivateLinkConnectionStatusApproved,
		PrivateLinkConnectionStatusDisconnected,
		PrivateLinkConnectionStatusPending,
		PrivateLinkConnectionStatusRejected,
	}
}

// SKUName - Name of this SKU.
type SKUName string

const (
	SKUNameBasic    SKUName = "Basic"
	SKUNameStandard SKUName = "Standard"
)

// PossibleSKUNameValues returns the possible values for the SKUName const type.
func PossibleSKUNameValues() []SKUName {
	return []SKUName{
		SKUNameBasic,
		SKUNameStandard,
	}
}

// SKUTier - The billing tier of this particular SKU.
type SKUTier string

const (
	SKUTierBasic    SKUTier = "Basic"
	SKUTierStandard SKUTier = "Standard"
)

// PossibleSKUTierValues returns the possible values for the SKUTier const type.
func PossibleSKUTierValues() []SKUTier {
	return []SKUTier{
		SKUTierBasic,
		SKUTierStandard,
	}
}

// UnavailableReason - Specifies the reason for the unavailability of the service.
type UnavailableReason string

const (
	UnavailableReasonNone                                  UnavailableReason = "None"
	UnavailableReasonInvalidName                           UnavailableReason = "InvalidName"
	UnavailableReasonSubscriptionIsDisabled                UnavailableReason = "SubscriptionIsDisabled"
	UnavailableReasonNameInUse                             UnavailableReason = "NameInUse"
	UnavailableReasonNameInLockdown                        UnavailableReason = "NameInLockdown"
	UnavailableReasonTooManyNamespaceInCurrentSubscription UnavailableReason = "TooManyNamespaceInCurrentSubscription"
)

// PossibleUnavailableReasonValues returns the possible values for the UnavailableReason const type.
func PossibleUnavailableReasonValues() []UnavailableReason {
	return []UnavailableReason{
		UnavailableReasonNone,
		UnavailableReasonInvalidName,
		UnavailableReasonSubscriptionIsDisabled,
		UnavailableReasonNameInUse,
		UnavailableReasonNameInLockdown,
		UnavailableReasonTooManyNamespaceInCurrentSubscription,
	}
}