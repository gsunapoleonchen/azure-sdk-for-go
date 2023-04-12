//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armtimeseriesinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/timeseriesinsights/armtimeseriesinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/stable/2020-05-15/examples/ReferenceDataSetsCreate.json
func ExampleReferenceDataSetsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtimeseriesinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReferenceDataSetsClient().CreateOrUpdate(ctx, "rg1", "env1", "rds1", armtimeseriesinsights.ReferenceDataSetCreateOrUpdateParameters{
		Location: to.Ptr("West US"),
		Properties: &armtimeseriesinsights.ReferenceDataSetCreationProperties{
			KeyProperties: []*armtimeseriesinsights.ReferenceDataSetKeyProperty{
				{
					Name: to.Ptr("DeviceId1"),
					Type: to.Ptr(armtimeseriesinsights.ReferenceDataKeyPropertyTypeString),
				},
				{
					Name: to.Ptr("DeviceFloor"),
					Type: to.Ptr(armtimeseriesinsights.ReferenceDataKeyPropertyTypeDouble),
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ReferenceDataSetResource = armtimeseriesinsights.ReferenceDataSetResource{
	// 	Name: to.Ptr("rds1"),
	// 	Type: to.Ptr("Microsoft.TimeSeriesInsights/Environments/ReferenceDataSets"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.TimeSeriesInsights/Environments/env1/referenceDataSets/rds1"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armtimeseriesinsights.ReferenceDataSetResourceProperties{
	// 		KeyProperties: []*armtimeseriesinsights.ReferenceDataSetKeyProperty{
	// 			{
	// 				Name: to.Ptr("DeviceId1"),
	// 				Type: to.Ptr(armtimeseriesinsights.ReferenceDataKeyPropertyTypeString),
	// 			},
	// 			{
	// 				Name: to.Ptr("DeviceFloor"),
	// 				Type: to.Ptr(armtimeseriesinsights.ReferenceDataKeyPropertyTypeDouble),
	// 		}},
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-18T19:20:33.2288820Z"); return t}()),
	// 		ProvisioningState: to.Ptr(armtimeseriesinsights.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/stable/2020-05-15/examples/ReferenceDataSetsGet.json
func ExampleReferenceDataSetsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtimeseriesinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReferenceDataSetsClient().Get(ctx, "rg1", "env1", "rds1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ReferenceDataSetResource = armtimeseriesinsights.ReferenceDataSetResource{
	// 	Name: to.Ptr("rds1"),
	// 	Type: to.Ptr("Microsoft.TimeSeriesInsights/Environments/ReferenceDataSets"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.TimeSeriesInsights/Environments/env1/referenceDataSets/rds1"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armtimeseriesinsights.ReferenceDataSetResourceProperties{
	// 		KeyProperties: []*armtimeseriesinsights.ReferenceDataSetKeyProperty{
	// 			{
	// 				Name: to.Ptr("DeviceId1"),
	// 				Type: to.Ptr(armtimeseriesinsights.ReferenceDataKeyPropertyTypeString),
	// 			},
	// 			{
	// 				Name: to.Ptr("DeviceFloor"),
	// 				Type: to.Ptr(armtimeseriesinsights.ReferenceDataKeyPropertyTypeDouble),
	// 		}},
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-18T19:20:33.2288820Z"); return t}()),
	// 		ProvisioningState: to.Ptr(armtimeseriesinsights.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/stable/2020-05-15/examples/ReferenceDataSetsPatchTags.json
func ExampleReferenceDataSetsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtimeseriesinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReferenceDataSetsClient().Update(ctx, "rg1", "env1", "rds1", armtimeseriesinsights.ReferenceDataSetUpdateParameters{
		Tags: map[string]*string{
			"someKey": to.Ptr("someValue"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ReferenceDataSetResource = armtimeseriesinsights.ReferenceDataSetResource{
	// 	Name: to.Ptr("rds1"),
	// 	Type: to.Ptr("Microsoft.TimeSeriesInsights/Environments/ReferenceDataSets"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.TimeSeriesInsights/Environments/env1/referenceDataSets/rds1"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"someKey": to.Ptr("someValue"),
	// 	},
	// 	Properties: &armtimeseriesinsights.ReferenceDataSetResourceProperties{
	// 		KeyProperties: []*armtimeseriesinsights.ReferenceDataSetKeyProperty{
	// 			{
	// 				Name: to.Ptr("DeviceId1"),
	// 				Type: to.Ptr(armtimeseriesinsights.ReferenceDataKeyPropertyTypeString),
	// 			},
	// 			{
	// 				Name: to.Ptr("DeviceFloor"),
	// 				Type: to.Ptr(armtimeseriesinsights.ReferenceDataKeyPropertyTypeDouble),
	// 		}},
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-18T19:20:33.2288820Z"); return t}()),
	// 		ProvisioningState: to.Ptr(armtimeseriesinsights.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/stable/2020-05-15/examples/ReferenceDataSetsDelete.json
func ExampleReferenceDataSetsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtimeseriesinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewReferenceDataSetsClient().Delete(ctx, "rg1", "env1", "rds1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/stable/2020-05-15/examples/ReferenceDataSetsListByEnvironment.json
func ExampleReferenceDataSetsClient_ListByEnvironment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtimeseriesinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReferenceDataSetsClient().ListByEnvironment(ctx, "rg1", "env1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ReferenceDataSetListResponse = armtimeseriesinsights.ReferenceDataSetListResponse{
	// 	Value: []*armtimeseriesinsights.ReferenceDataSetResource{
	// 		{
	// 			Name: to.Ptr("rds1"),
	// 			Type: to.Ptr("Microsoft.TimeSeriesInsights/Environments/ReferenceDataSets"),
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.TimeSeriesInsights/Environments/env1/referenceDataSets/rds1"),
	// 			Location: to.Ptr("West US"),
	// 			Tags: map[string]*string{
	// 			},
	// 			Properties: &armtimeseriesinsights.ReferenceDataSetResourceProperties{
	// 				KeyProperties: []*armtimeseriesinsights.ReferenceDataSetKeyProperty{
	// 					{
	// 						Name: to.Ptr("DeviceId1"),
	// 						Type: to.Ptr(armtimeseriesinsights.ReferenceDataKeyPropertyTypeString),
	// 					},
	// 					{
	// 						Name: to.Ptr("DeviceFloor"),
	// 						Type: to.Ptr(armtimeseriesinsights.ReferenceDataKeyPropertyTypeDouble),
	// 				}},
	// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-18T19:20:33.2288820Z"); return t}()),
	// 				ProvisioningState: to.Ptr(armtimeseriesinsights.ProvisioningStateSucceeded),
	// 			},
	// 	}},
	// }
}