//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdatafactory_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6051d2b126f5b1e4b623cde8edfa3e25cf730685/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/PipelineRuns_QueryByFactory.json
func ExamplePipelineRunsClient_QueryByFactory() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPipelineRunsClient().QueryByFactory(ctx, "exampleResourceGroup", "exampleFactoryName", armdatafactory.RunFilterParameters{
		Filters: []*armdatafactory.RunQueryFilter{
			{
				Operand:  to.Ptr(armdatafactory.RunQueryFilterOperandPipelineName),
				Operator: to.Ptr(armdatafactory.RunQueryFilterOperatorEquals),
				Values: []*string{
					to.Ptr("examplePipeline")},
			}},
		LastUpdatedAfter:  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:36:44.334Z"); return t }()),
		LastUpdatedBefore: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:49:48.368Z"); return t }()),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PipelineRunsQueryResponse = armdatafactory.PipelineRunsQueryResponse{
	// 	Value: []*armdatafactory.PipelineRun{
	// 		{
	// 			AdditionalProperties: map[string]any{
	// 				"annotations": []any{
	// 				},
	// 				"runDimension": map[string]any{
	// 					"JobId": "79c1cc52-265f-41a5-9553-be65e736fbd3",
	// 				},
	// 			},
	// 			DurationInMs: to.Ptr[int32](28105),
	// 			InvokedBy: &armdatafactory.PipelineRunInvokedBy{
	// 				Name: to.Ptr("Manual"),
	// 				ID: to.Ptr("80a01654a9d34ad18b3fcac5d5d76b67"),
	// 			},
	// 			LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:38:12.731Z"); return t}()),
	// 			Message: to.Ptr(""),
	// 			Parameters: map[string]*string{
	// 				"OutputBlobNameList": to.Ptr("[\"exampleoutput.csv\"]"),
	// 			},
	// 			PipelineName: to.Ptr("examplePipeline"),
	// 			RunEnd: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:38:12.731Z"); return t}()),
	// 			RunID: to.Ptr("2f7fdb90-5df1-4b8e-ac2f-064cfa58202b"),
	// 			RunStart: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:37:44.625Z"); return t}()),
	// 			Status: to.Ptr("Succeeded"),
	// 		},
	// 		{
	// 			AdditionalProperties: map[string]any{
	// 				"annotations": []any{
	// 				},
	// 				"runDimension": map[string]any{
	// 					"JobId": "84a3c493-0628-4b44-852f-ef5b3a11bdab",
	// 				},
	// 			},
	// 			InvokedBy: &armdatafactory.PipelineRunInvokedBy{
	// 				Name: to.Ptr("Manual"),
	// 				ID: to.Ptr("7c5fd7ef7e8a464b98b931cf15fcac66"),
	// 			},
	// 			LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:39:51.216Z"); return t}()),
	// 			Message: to.Ptr(""),
	// 			Parameters: map[string]*string{
	// 				"OutputBlobNameList": to.Ptr("[\"exampleoutput.csv\"]"),
	// 			},
	// 			PipelineName: to.Ptr("examplePipeline"),
	// 			RunID: to.Ptr("16ac5348-ff82-4f95-a80d-638c1d47b721"),
	// 			RunStart: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:39:49.274Z"); return t}()),
	// 			Status: to.Ptr("Cancelled"),
	// 	}},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6051d2b126f5b1e4b623cde8edfa3e25cf730685/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/PipelineRuns_Get.json
func ExamplePipelineRunsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPipelineRunsClient().Get(ctx, "exampleResourceGroup", "exampleFactoryName", "2f7fdb90-5df1-4b8e-ac2f-064cfa58202b", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PipelineRun = armdatafactory.PipelineRun{
	// 	AdditionalProperties: map[string]any{
	// 		"annotations": []any{
	// 		},
	// 	},
	// 	DurationInMs: to.Ptr[int32](28105),
	// 	InvokedBy: &armdatafactory.PipelineRunInvokedBy{
	// 		Name: to.Ptr("Manual"),
	// 		ID: to.Ptr("80a01654a9d34ad18b3fcac5d5d76b67"),
	// 	},
	// 	LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:38:12.731Z"); return t}()),
	// 	Message: to.Ptr(""),
	// 	Parameters: map[string]*string{
	// 		"OutputBlobNameList": to.Ptr("[\"exampleoutput.csv\"]"),
	// 	},
	// 	PipelineName: to.Ptr("examplePipeline"),
	// 	RunEnd: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:38:12.731Z"); return t}()),
	// 	RunID: to.Ptr("2f7fdb90-5df1-4b8e-ac2f-064cfa58202b"),
	// 	RunStart: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:37:44.625Z"); return t}()),
	// 	Status: to.Ptr("Succeeded"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6051d2b126f5b1e4b623cde8edfa3e25cf730685/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/PipelineRuns_Cancel.json
func ExamplePipelineRunsClient_Cancel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewPipelineRunsClient().Cancel(ctx, "exampleResourceGroup", "exampleFactoryName", "16ac5348-ff82-4f95-a80d-638c1d47b721", &armdatafactory.PipelineRunsClientCancelOptions{IsRecursive: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
