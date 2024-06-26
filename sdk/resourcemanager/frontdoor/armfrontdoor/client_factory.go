//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armfrontdoor

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	internal       *arm.Client
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	internal, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID,
		internal:       internal,
	}, nil
}

// NewEndpointsClient creates a new instance of EndpointsClient.
func (c *ClientFactory) NewEndpointsClient() *EndpointsClient {
	return &EndpointsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewExperimentsClient creates a new instance of ExperimentsClient.
func (c *ClientFactory) NewExperimentsClient() *ExperimentsClient {
	return &ExperimentsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewFrontDoorsClient creates a new instance of FrontDoorsClient.
func (c *ClientFactory) NewFrontDoorsClient() *FrontDoorsClient {
	return &FrontDoorsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewFrontendEndpointsClient creates a new instance of FrontendEndpointsClient.
func (c *ClientFactory) NewFrontendEndpointsClient() *FrontendEndpointsClient {
	return &FrontendEndpointsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewManagedRuleSetsClient creates a new instance of ManagedRuleSetsClient.
func (c *ClientFactory) NewManagedRuleSetsClient() *ManagedRuleSetsClient {
	return &ManagedRuleSetsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewNameAvailabilityClient creates a new instance of NameAvailabilityClient.
func (c *ClientFactory) NewNameAvailabilityClient() *NameAvailabilityClient {
	return &NameAvailabilityClient{
		internal: c.internal,
	}
}

// NewNameAvailabilityWithSubscriptionClient creates a new instance of NameAvailabilityWithSubscriptionClient.
func (c *ClientFactory) NewNameAvailabilityWithSubscriptionClient() *NameAvailabilityWithSubscriptionClient {
	return &NameAvailabilityWithSubscriptionClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewNetworkExperimentProfilesClient creates a new instance of NetworkExperimentProfilesClient.
func (c *ClientFactory) NewNetworkExperimentProfilesClient() *NetworkExperimentProfilesClient {
	return &NetworkExperimentProfilesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewPoliciesClient creates a new instance of PoliciesClient.
func (c *ClientFactory) NewPoliciesClient() *PoliciesClient {
	return &PoliciesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewPreconfiguredEndpointsClient creates a new instance of PreconfiguredEndpointsClient.
func (c *ClientFactory) NewPreconfiguredEndpointsClient() *PreconfiguredEndpointsClient {
	return &PreconfiguredEndpointsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewReportsClient creates a new instance of ReportsClient.
func (c *ClientFactory) NewReportsClient() *ReportsClient {
	return &ReportsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewRulesEnginesClient creates a new instance of RulesEnginesClient.
func (c *ClientFactory) NewRulesEnginesClient() *RulesEnginesClient {
	return &RulesEnginesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}
