//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armaad

// PrivateEndpointConnectionsClientCreateResponse contains the response from method PrivateEndpointConnectionsClient.BeginCreate.
type PrivateEndpointConnectionsClientCreateResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.BeginDelete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListByPolicyNameResponse contains the response from method PrivateEndpointConnectionsClient.NewListByPolicyNamePager.
type PrivateEndpointConnectionsClientListByPolicyNameResponse struct {
	PrivateEndpointConnectionListResult
}

// PrivateLinkForAzureAdClientCreateResponse contains the response from method PrivateLinkForAzureAdClient.BeginCreate.
type PrivateLinkForAzureAdClientCreateResponse struct {
	PrivateLinkPolicy
}

// PrivateLinkForAzureAdClientDeleteResponse contains the response from method PrivateLinkForAzureAdClient.Delete.
type PrivateLinkForAzureAdClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateLinkForAzureAdClientGetResponse contains the response from method PrivateLinkForAzureAdClient.Get.
type PrivateLinkForAzureAdClientGetResponse struct {
	PrivateLinkPolicy
}

// PrivateLinkForAzureAdClientListBySubscriptionResponse contains the response from method PrivateLinkForAzureAdClient.NewListBySubscriptionPager.
type PrivateLinkForAzureAdClientListBySubscriptionResponse struct {
	PrivateLinkPolicyListResult
}

// PrivateLinkForAzureAdClientListResponse contains the response from method PrivateLinkForAzureAdClient.NewListPager.
type PrivateLinkForAzureAdClientListResponse struct {
	PrivateLinkPolicyListResult
}

// PrivateLinkForAzureAdClientUpdateResponse contains the response from method PrivateLinkForAzureAdClient.Update.
type PrivateLinkForAzureAdClientUpdateResponse struct {
	PrivateLinkPolicy
}

// PrivateLinkResourcesClientGetResponse contains the response from method PrivateLinkResourcesClient.Get.
type PrivateLinkResourcesClientGetResponse struct {
	PrivateLinkResource
}

// PrivateLinkResourcesClientListByPrivateLinkPolicyResponse contains the response from method PrivateLinkResourcesClient.NewListByPrivateLinkPolicyPager.
type PrivateLinkResourcesClientListByPrivateLinkPolicyResponse struct {
	PrivateLinkResourceListResult
}