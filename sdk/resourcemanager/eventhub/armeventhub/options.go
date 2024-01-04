//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventhub

// ApplicationGroupClientCreateOrUpdateApplicationGroupOptions contains the optional parameters for the ApplicationGroupClient.CreateOrUpdateApplicationGroup
// method.
type ApplicationGroupClientCreateOrUpdateApplicationGroupOptions struct {
	// placeholder for future optional parameters
}

// ApplicationGroupClientDeleteOptions contains the optional parameters for the ApplicationGroupClient.Delete method.
type ApplicationGroupClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ApplicationGroupClientGetOptions contains the optional parameters for the ApplicationGroupClient.Get method.
type ApplicationGroupClientGetOptions struct {
	// placeholder for future optional parameters
}

// ApplicationGroupClientListByNamespaceOptions contains the optional parameters for the ApplicationGroupClient.NewListByNamespacePager
// method.
type ApplicationGroupClientListByNamespaceOptions struct {
	// placeholder for future optional parameters
}

// ClustersClientBeginCreateOrUpdateOptions contains the optional parameters for the ClustersClient.BeginCreateOrUpdate method.
type ClustersClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ClustersClientBeginDeleteOptions contains the optional parameters for the ClustersClient.BeginDelete method.
type ClustersClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ClustersClientBeginUpdateOptions contains the optional parameters for the ClustersClient.BeginUpdate method.
type ClustersClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ClustersClientGetOptions contains the optional parameters for the ClustersClient.Get method.
type ClustersClientGetOptions struct {
	// placeholder for future optional parameters
}

// ClustersClientListAvailableClusterRegionOptions contains the optional parameters for the ClustersClient.ListAvailableClusterRegion
// method.
type ClustersClientListAvailableClusterRegionOptions struct {
	// placeholder for future optional parameters
}

// ClustersClientListByResourceGroupOptions contains the optional parameters for the ClustersClient.NewListByResourceGroupPager
// method.
type ClustersClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// ClustersClientListBySubscriptionOptions contains the optional parameters for the ClustersClient.NewListBySubscriptionPager
// method.
type ClustersClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// ClustersClientListNamespacesOptions contains the optional parameters for the ClustersClient.ListNamespaces method.
type ClustersClientListNamespacesOptions struct {
	// placeholder for future optional parameters
}

// ConfigurationClientGetOptions contains the optional parameters for the ConfigurationClient.Get method.
type ConfigurationClientGetOptions struct {
	// placeholder for future optional parameters
}

// ConfigurationClientPatchOptions contains the optional parameters for the ConfigurationClient.Patch method.
type ConfigurationClientPatchOptions struct {
	// placeholder for future optional parameters
}

// ConsumerGroupsClientCreateOrUpdateOptions contains the optional parameters for the ConsumerGroupsClient.CreateOrUpdate
// method.
type ConsumerGroupsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ConsumerGroupsClientDeleteOptions contains the optional parameters for the ConsumerGroupsClient.Delete method.
type ConsumerGroupsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ConsumerGroupsClientGetOptions contains the optional parameters for the ConsumerGroupsClient.Get method.
type ConsumerGroupsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ConsumerGroupsClientListByEventHubOptions contains the optional parameters for the ConsumerGroupsClient.NewListByEventHubPager
// method.
type ConsumerGroupsClientListByEventHubOptions struct {
	// Skip is only used if a previous operation returned a partial result. If a previous response contains a nextLink element,
	// the value of the nextLink element will include a skip parameter that specifies
	// a starting point to use for subsequent calls.
	Skip *int32

	// May be used to limit the number of results to the most recent N usageDetails.
	Top *int32
}

// DisasterRecoveryConfigsClientBreakPairingOptions contains the optional parameters for the DisasterRecoveryConfigsClient.BreakPairing
// method.
type DisasterRecoveryConfigsClientBreakPairingOptions struct {
	// placeholder for future optional parameters
}

// DisasterRecoveryConfigsClientCheckNameAvailabilityOptions contains the optional parameters for the DisasterRecoveryConfigsClient.CheckNameAvailability
// method.
type DisasterRecoveryConfigsClientCheckNameAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// DisasterRecoveryConfigsClientCreateOrUpdateOptions contains the optional parameters for the DisasterRecoveryConfigsClient.CreateOrUpdate
// method.
type DisasterRecoveryConfigsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// DisasterRecoveryConfigsClientDeleteOptions contains the optional parameters for the DisasterRecoveryConfigsClient.Delete
// method.
type DisasterRecoveryConfigsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// DisasterRecoveryConfigsClientFailOverOptions contains the optional parameters for the DisasterRecoveryConfigsClient.FailOver
// method.
type DisasterRecoveryConfigsClientFailOverOptions struct {
	// placeholder for future optional parameters
}

// DisasterRecoveryConfigsClientGetAuthorizationRuleOptions contains the optional parameters for the DisasterRecoveryConfigsClient.GetAuthorizationRule
// method.
type DisasterRecoveryConfigsClientGetAuthorizationRuleOptions struct {
	// placeholder for future optional parameters
}

// DisasterRecoveryConfigsClientGetOptions contains the optional parameters for the DisasterRecoveryConfigsClient.Get method.
type DisasterRecoveryConfigsClientGetOptions struct {
	// placeholder for future optional parameters
}

// DisasterRecoveryConfigsClientListAuthorizationRulesOptions contains the optional parameters for the DisasterRecoveryConfigsClient.NewListAuthorizationRulesPager
// method.
type DisasterRecoveryConfigsClientListAuthorizationRulesOptions struct {
	// placeholder for future optional parameters
}

// DisasterRecoveryConfigsClientListKeysOptions contains the optional parameters for the DisasterRecoveryConfigsClient.ListKeys
// method.
type DisasterRecoveryConfigsClientListKeysOptions struct {
	// placeholder for future optional parameters
}

// DisasterRecoveryConfigsClientListOptions contains the optional parameters for the DisasterRecoveryConfigsClient.NewListPager
// method.
type DisasterRecoveryConfigsClientListOptions struct {
	// placeholder for future optional parameters
}

// EventHubsClientCreateOrUpdateAuthorizationRuleOptions contains the optional parameters for the EventHubsClient.CreateOrUpdateAuthorizationRule
// method.
type EventHubsClientCreateOrUpdateAuthorizationRuleOptions struct {
	// placeholder for future optional parameters
}

// EventHubsClientCreateOrUpdateOptions contains the optional parameters for the EventHubsClient.CreateOrUpdate method.
type EventHubsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// EventHubsClientDeleteAuthorizationRuleOptions contains the optional parameters for the EventHubsClient.DeleteAuthorizationRule
// method.
type EventHubsClientDeleteAuthorizationRuleOptions struct {
	// placeholder for future optional parameters
}

// EventHubsClientDeleteOptions contains the optional parameters for the EventHubsClient.Delete method.
type EventHubsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// EventHubsClientGetAuthorizationRuleOptions contains the optional parameters for the EventHubsClient.GetAuthorizationRule
// method.
type EventHubsClientGetAuthorizationRuleOptions struct {
	// placeholder for future optional parameters
}

// EventHubsClientGetOptions contains the optional parameters for the EventHubsClient.Get method.
type EventHubsClientGetOptions struct {
	// placeholder for future optional parameters
}

// EventHubsClientListAuthorizationRulesOptions contains the optional parameters for the EventHubsClient.NewListAuthorizationRulesPager
// method.
type EventHubsClientListAuthorizationRulesOptions struct {
	// placeholder for future optional parameters
}

// EventHubsClientListByNamespaceOptions contains the optional parameters for the EventHubsClient.NewListByNamespacePager
// method.
type EventHubsClientListByNamespaceOptions struct {
	// Skip is only used if a previous operation returned a partial result. If a previous response contains a nextLink element,
	// the value of the nextLink element will include a skip parameter that specifies
	// a starting point to use for subsequent calls.
	Skip *int32

	// May be used to limit the number of results to the most recent N usageDetails.
	Top *int32
}

// EventHubsClientListKeysOptions contains the optional parameters for the EventHubsClient.ListKeys method.
type EventHubsClientListKeysOptions struct {
	// placeholder for future optional parameters
}

// EventHubsClientRegenerateKeysOptions contains the optional parameters for the EventHubsClient.RegenerateKeys method.
type EventHubsClientRegenerateKeysOptions struct {
	// placeholder for future optional parameters
}

// NamespacesClientBeginCreateOrUpdateOptions contains the optional parameters for the NamespacesClient.BeginCreateOrUpdate
// method.
type NamespacesClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// NamespacesClientBeginDeleteOptions contains the optional parameters for the NamespacesClient.BeginDelete method.
type NamespacesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// NamespacesClientCheckNameAvailabilityOptions contains the optional parameters for the NamespacesClient.CheckNameAvailability
// method.
type NamespacesClientCheckNameAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// NamespacesClientCreateOrUpdateAuthorizationRuleOptions contains the optional parameters for the NamespacesClient.CreateOrUpdateAuthorizationRule
// method.
type NamespacesClientCreateOrUpdateAuthorizationRuleOptions struct {
	// placeholder for future optional parameters
}

// NamespacesClientCreateOrUpdateNetworkRuleSetOptions contains the optional parameters for the NamespacesClient.CreateOrUpdateNetworkRuleSet
// method.
type NamespacesClientCreateOrUpdateNetworkRuleSetOptions struct {
	// placeholder for future optional parameters
}

// NamespacesClientDeleteAuthorizationRuleOptions contains the optional parameters for the NamespacesClient.DeleteAuthorizationRule
// method.
type NamespacesClientDeleteAuthorizationRuleOptions struct {
	// placeholder for future optional parameters
}

// NamespacesClientGetAuthorizationRuleOptions contains the optional parameters for the NamespacesClient.GetAuthorizationRule
// method.
type NamespacesClientGetAuthorizationRuleOptions struct {
	// placeholder for future optional parameters
}

// NamespacesClientGetNetworkRuleSetOptions contains the optional parameters for the NamespacesClient.GetNetworkRuleSet method.
type NamespacesClientGetNetworkRuleSetOptions struct {
	// placeholder for future optional parameters
}

// NamespacesClientGetOptions contains the optional parameters for the NamespacesClient.Get method.
type NamespacesClientGetOptions struct {
	// placeholder for future optional parameters
}

// NamespacesClientListAuthorizationRulesOptions contains the optional parameters for the NamespacesClient.NewListAuthorizationRulesPager
// method.
type NamespacesClientListAuthorizationRulesOptions struct {
	// placeholder for future optional parameters
}

// NamespacesClientListByResourceGroupOptions contains the optional parameters for the NamespacesClient.NewListByResourceGroupPager
// method.
type NamespacesClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// NamespacesClientListKeysOptions contains the optional parameters for the NamespacesClient.ListKeys method.
type NamespacesClientListKeysOptions struct {
	// placeholder for future optional parameters
}

// NamespacesClientListNetworkRuleSetOptions contains the optional parameters for the NamespacesClient.ListNetworkRuleSet
// method.
type NamespacesClientListNetworkRuleSetOptions struct {
	// placeholder for future optional parameters
}

// NamespacesClientListOptions contains the optional parameters for the NamespacesClient.NewListPager method.
type NamespacesClientListOptions struct {
	// placeholder for future optional parameters
}

// NamespacesClientRegenerateKeysOptions contains the optional parameters for the NamespacesClient.RegenerateKeys method.
type NamespacesClientRegenerateKeysOptions struct {
	// placeholder for future optional parameters
}

// NamespacesClientUpdateOptions contains the optional parameters for the NamespacesClient.Update method.
type NamespacesClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// NetworkSecurityPerimeterConfigurationClientListOptions contains the optional parameters for the NetworkSecurityPerimeterConfigurationClient.List
// method.
type NetworkSecurityPerimeterConfigurationClientListOptions struct {
	// placeholder for future optional parameters
}

// NetworkSecurityPerimeterConfigurationsClientBeginCreateOrUpdateOptions contains the optional parameters for the NetworkSecurityPerimeterConfigurationsClient.BeginCreateOrUpdate
// method.
type NetworkSecurityPerimeterConfigurationsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientBeginDeleteOptions contains the optional parameters for the PrivateEndpointConnectionsClient.BeginDelete
// method.
type PrivateEndpointConnectionsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PrivateEndpointConnectionsClientCreateOrUpdateOptions contains the optional parameters for the PrivateEndpointConnectionsClient.CreateOrUpdate
// method.
type PrivateEndpointConnectionsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientGetOptions contains the optional parameters for the PrivateEndpointConnectionsClient.Get
// method.
type PrivateEndpointConnectionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientListOptions contains the optional parameters for the PrivateEndpointConnectionsClient.NewListPager
// method.
type PrivateEndpointConnectionsClientListOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkResourcesClientGetOptions contains the optional parameters for the PrivateLinkResourcesClient.Get method.
type PrivateLinkResourcesClientGetOptions struct {
	// placeholder for future optional parameters
}

// SchemaRegistryClientCreateOrUpdateOptions contains the optional parameters for the SchemaRegistryClient.CreateOrUpdate
// method.
type SchemaRegistryClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// SchemaRegistryClientDeleteOptions contains the optional parameters for the SchemaRegistryClient.Delete method.
type SchemaRegistryClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// SchemaRegistryClientGetOptions contains the optional parameters for the SchemaRegistryClient.Get method.
type SchemaRegistryClientGetOptions struct {
	// placeholder for future optional parameters
}

// SchemaRegistryClientListByNamespaceOptions contains the optional parameters for the SchemaRegistryClient.NewListByNamespacePager
// method.
type SchemaRegistryClientListByNamespaceOptions struct {
	// Skip is only used if a previous operation returned a partial result. If a previous response contains a nextLink element,
	// the value of the nextLink element will include a skip parameter that specifies
	// a starting point to use for subsequent calls.
	Skip *int32

	// May be used to limit the number of results to the most recent N usageDetails.
	Top *int32
}