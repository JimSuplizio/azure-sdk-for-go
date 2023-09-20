//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SensitivityLabelsListByDatabaseWithSourceCurrent.json
func ExampleSensitivityLabelsClient_NewListCurrentByDatabasePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSensitivityLabelsClient().NewListCurrentByDatabasePager("myRG", "myServer", "myDatabase", &armsql.SensitivityLabelsClientListCurrentByDatabaseOptions{SkipToken: nil,
		Count:  nil,
		Filter: nil,
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.SensitivityLabelListResult = armsql.SensitivityLabelListResult{
		// 	Value: []*armsql.SensitivityLabel{
		// 		{
		// 			Name: to.Ptr("current"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases/schemas/tables/columns/sensitivityLabels"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myRG/providers/Microsoft.Sql/servers/myServer/databases/myDatabase/schemas/dbo/tables/myTable/columns/myColumn3/sensitivityLabels/current"),
		// 			Properties: &armsql.SensitivityLabelProperties{
		// 				ColumnName: to.Ptr("myColumn3"),
		// 				InformationType: to.Ptr("Financial"),
		// 				InformationTypeID: to.Ptr("d22fa6e9-5ee4-3bde-4c2b-a409604c4646"),
		// 				LabelID: to.Ptr("05e6eaa1-075a-4fb4-a732-a92215a2444a"),
		// 				LabelName: to.Ptr("Sensitive"),
		// 				Rank: to.Ptr(armsql.SensitivityLabelRankLow),
		// 				SchemaName: to.Ptr("dbo"),
		// 				TableName: to.Ptr("myTable"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("current"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases/schemas/tables/columns/sensitivityLabels"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myRG/providers/Microsoft.Sql/servers/myServer/databases/myDatabase/schemas/dbo/tables/myTable/columns/myColumn4/sensitivityLabels/current"),
		// 			Properties: &armsql.SensitivityLabelProperties{
		// 				ColumnName: to.Ptr("myColumn4"),
		// 				InformationType: to.Ptr("Email"),
		// 				InformationTypeID: to.Ptr("d22fa6e9-5ee4-3bde-4c2b-a409604c4646"),
		// 				LabelID: to.Ptr("bf91e08c-f4f0-478a-b016-25164b2a65ff"),
		// 				LabelName: to.Ptr("PII"),
		// 				Rank: to.Ptr(armsql.SensitivityLabelRankNone),
		// 				SchemaName: to.Ptr("dbo"),
		// 				TableName: to.Ptr("myTable"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SensitivityLabelsCurrentUpdate.json
func ExampleSensitivityLabelsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewSensitivityLabelsClient().Update(ctx, "myRG", "myServer", "myDatabase", armsql.SensitivityLabelUpdateList{
		Operations: []*armsql.SensitivityLabelUpdate{
			{
				Properties: &armsql.SensitivityLabelUpdateProperties{
					Schema: to.Ptr("dbo"),
					Column: to.Ptr("column1"),
					Op:     to.Ptr(armsql.SensitivityLabelUpdateKindSet),
					SensitivityLabel: &armsql.SensitivityLabel{
						Properties: &armsql.SensitivityLabelProperties{
							InformationType:   to.Ptr("Financial"),
							InformationTypeID: to.Ptr("1D3652D6-422C-4115-82F1-65DAEBC665C8"),
							LabelID:           to.Ptr("3A477B16-9423-432B-AA97-6069B481CEC3"),
							LabelName:         to.Ptr("Highly Confidential"),
							Rank:              to.Ptr(armsql.SensitivityLabelRankLow),
						},
					},
					Table: to.Ptr("table1"),
				},
			},
			{
				Properties: &armsql.SensitivityLabelUpdateProperties{
					Schema: to.Ptr("dbo"),
					Column: to.Ptr("column2"),
					Op:     to.Ptr(armsql.SensitivityLabelUpdateKindSet),
					SensitivityLabel: &armsql.SensitivityLabel{
						Properties: &armsql.SensitivityLabelProperties{
							InformationType:   to.Ptr("PhoneNumber"),
							InformationTypeID: to.Ptr("d22fa6e9-5ee4-3bde-4c2b-a409604c4646"),
							LabelID:           to.Ptr("bf91e08c-f4f0-478a-b016-25164b2a65ff"),
							LabelName:         to.Ptr("PII"),
							Rank:              to.Ptr(armsql.SensitivityLabelRankCritical),
						},
					},
					Table: to.Ptr("table2"),
				},
			},
			{
				Properties: &armsql.SensitivityLabelUpdateProperties{
					Schema: to.Ptr("dbo"),
					Column: to.Ptr("Column3"),
					Op:     to.Ptr(armsql.SensitivityLabelUpdateKindRemove),
					Table:  to.Ptr("Table1"),
				},
			}},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SensitivityLabelsListByDatabaseWithSourceRecommended.json
func ExampleSensitivityLabelsClient_NewListRecommendedByDatabasePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSensitivityLabelsClient().NewListRecommendedByDatabasePager("myRG", "myServer", "myDatabase", &armsql.SensitivityLabelsClientListRecommendedByDatabaseOptions{SkipToken: nil,
		IncludeDisabledRecommendations: nil,
		Filter:                         nil,
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.SensitivityLabelListResult = armsql.SensitivityLabelListResult{
		// 	Value: []*armsql.SensitivityLabel{
		// 		{
		// 			Name: to.Ptr("recommended"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases/schemas/tables/columns/sensitivityLabels"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myRG/providers/Microsoft.Sql/servers/myServer/databases/myDatabase/schemas/dbo/tables/myTable/columns/myColumn/sensitivityLabels/recommended"),
		// 			Properties: &armsql.SensitivityLabelProperties{
		// 				ColumnName: to.Ptr("myColumn"),
		// 				InformationType: to.Ptr("Financial"),
		// 				InformationTypeID: to.Ptr("d22fa6e9-5ee4-3bde-4c2b-a409604c4646"),
		// 				LabelID: to.Ptr("05e6eaa1-075a-4fb4-a732-a92215a2444a"),
		// 				LabelName: to.Ptr("Sensitive"),
		// 				SchemaName: to.Ptr("dbo"),
		// 				TableName: to.Ptr("myTable"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("recommended"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases/schemas/tables/columns/sensitivityLabels"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myRG/providers/Microsoft.Sql/servers/myServer/databases/myDatabase/schemas/dbo/tables/myTable/columns/myColumn2/sensitivityLabels/recommended"),
		// 			Properties: &armsql.SensitivityLabelProperties{
		// 				ColumnName: to.Ptr("myColumn2"),
		// 				InformationType: to.Ptr("Email"),
		// 				InformationTypeID: to.Ptr("d22fa6e9-5ee4-3bde-4c2b-a409604c4646"),
		// 				LabelID: to.Ptr("bf91e08c-f4f0-478a-b016-25164b2a65ff"),
		// 				LabelName: to.Ptr("PII"),
		// 				SchemaName: to.Ptr("dbo"),
		// 				TableName: to.Ptr("myTable"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/RecommendedColumnSensitivityLabelEnable.json
func ExampleSensitivityLabelsClient_EnableRecommendation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewSensitivityLabelsClient().EnableRecommendation(ctx, "myRG", "myServer", "myDatabase", "dbo", "myTable", "myColumn", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/RecommendedColumnSensitivityLabelDisable.json
func ExampleSensitivityLabelsClient_DisableRecommendation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewSensitivityLabelsClient().DisableRecommendation(ctx, "myRG", "myServer", "myDatabase", "dbo", "myTable", "myColumn", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ColumnSensitivityLabelGet.json
func ExampleSensitivityLabelsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSensitivityLabelsClient().Get(ctx, "myRG", "myServer", "myDatabase", "dbo", "myTable", "myColumn", armsql.SensitivityLabelSourceCurrent, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SensitivityLabel = armsql.SensitivityLabel{
	// 	Name: to.Ptr("current"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/schemas/tables/columns/sensitivityLabels"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myRG/providers/Microsoft.Sql/servers/myServer/databases/myDatabase/schemas/dbo/tables/myTable/columns/myColumn/sensitivityLabels/current"),
	// 	Properties: &armsql.SensitivityLabelProperties{
	// 		ColumnName: to.Ptr("myColumn"),
	// 		InformationType: to.Ptr("PhoneNumber"),
	// 		InformationTypeID: to.Ptr("d22fa6e9-5ee4-3bde-4c2b-a409604c4646"),
	// 		LabelID: to.Ptr("bf91e08c-f4f0-478a-b016-25164b2a65ff"),
	// 		LabelName: to.Ptr("PII"),
	// 		Rank: to.Ptr(armsql.SensitivityLabelRankHigh),
	// 		SchemaName: to.Ptr("dbo"),
	// 		TableName: to.Ptr("myTable"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ColumnSensitivityLabelCreateMax.json
func ExampleSensitivityLabelsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSensitivityLabelsClient().CreateOrUpdate(ctx, "myRG", "myServer", "myDatabase", "dbo", "myTable", "myColumn", armsql.SensitivityLabel{
		Properties: &armsql.SensitivityLabelProperties{
			InformationType:   to.Ptr("PhoneNumber"),
			InformationTypeID: to.Ptr("d22fa6e9-5ee4-3bde-4c2b-a409604c4646"),
			LabelID:           to.Ptr("bf91e08c-f4f0-478a-b016-25164b2a65ff"),
			LabelName:         to.Ptr("PII"),
			Rank:              to.Ptr(armsql.SensitivityLabelRankLow),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SensitivityLabel = armsql.SensitivityLabel{
	// 	Name: to.Ptr("current"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/schemas/tables/columns/sensitivityLabels"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myRG/providers/Microsoft.Sql/servers/myServer/databases/myDatabase/schemas/dbo/tables/myTable/columns/myColumn/sensitivityLabels/current"),
	// 	Properties: &armsql.SensitivityLabelProperties{
	// 		ColumnName: to.Ptr("myColumn"),
	// 		InformationType: to.Ptr("PhoneNumber"),
	// 		InformationTypeID: to.Ptr("d22fa6e9-5ee4-3bde-4c2b-a409604c4646"),
	// 		LabelID: to.Ptr("bf91e08c-f4f0-478a-b016-25164b2a65ff"),
	// 		LabelName: to.Ptr("PII"),
	// 		Rank: to.Ptr(armsql.SensitivityLabelRankMedium),
	// 		SchemaName: to.Ptr("dbo"),
	// 		TableName: to.Ptr("myTable"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ColumnSensitivityLabelDelete.json
func ExampleSensitivityLabelsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewSensitivityLabelsClient().Delete(ctx, "myRG", "myServer", "myDatabase", "dbo", "myTable", "myColumn", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
