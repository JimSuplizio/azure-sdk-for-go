//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhealthcareapis

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
//   - subscriptionID - The ID of the target subscription.
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

// NewDicomServicesClient creates a new instance of DicomServicesClient.
func (c *ClientFactory) NewDicomServicesClient() *DicomServicesClient {
	return &DicomServicesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewFhirDestinationsClient creates a new instance of FhirDestinationsClient.
func (c *ClientFactory) NewFhirDestinationsClient() *FhirDestinationsClient {
	return &FhirDestinationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewFhirServicesClient creates a new instance of FhirServicesClient.
func (c *ClientFactory) NewFhirServicesClient() *FhirServicesClient {
	return &FhirServicesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewIotConnectorFhirDestinationClient creates a new instance of IotConnectorFhirDestinationClient.
func (c *ClientFactory) NewIotConnectorFhirDestinationClient() *IotConnectorFhirDestinationClient {
	return &IotConnectorFhirDestinationClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewIotConnectorsClient creates a new instance of IotConnectorsClient.
func (c *ClientFactory) NewIotConnectorsClient() *IotConnectorsClient {
	return &IotConnectorsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewOperationResultsClient creates a new instance of OperationResultsClient.
func (c *ClientFactory) NewOperationResultsClient() *OperationResultsClient {
	return &OperationResultsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	return &OperationsClient{
		internal: c.internal,
	}
}

// NewPrivateEndpointConnectionsClient creates a new instance of PrivateEndpointConnectionsClient.
func (c *ClientFactory) NewPrivateEndpointConnectionsClient() *PrivateEndpointConnectionsClient {
	return &PrivateEndpointConnectionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewPrivateLinkResourcesClient creates a new instance of PrivateLinkResourcesClient.
func (c *ClientFactory) NewPrivateLinkResourcesClient() *PrivateLinkResourcesClient {
	return &PrivateLinkResourcesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewServicesClient creates a new instance of ServicesClient.
func (c *ClientFactory) NewServicesClient() *ServicesClient {
	return &ServicesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewWorkspacePrivateEndpointConnectionsClient creates a new instance of WorkspacePrivateEndpointConnectionsClient.
func (c *ClientFactory) NewWorkspacePrivateEndpointConnectionsClient() *WorkspacePrivateEndpointConnectionsClient {
	return &WorkspacePrivateEndpointConnectionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewWorkspacePrivateLinkResourcesClient creates a new instance of WorkspacePrivateLinkResourcesClient.
func (c *ClientFactory) NewWorkspacePrivateLinkResourcesClient() *WorkspacePrivateLinkResourcesClient {
	return &WorkspacePrivateLinkResourcesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewWorkspacesClient creates a new instance of WorkspacesClient.
func (c *ClientFactory) NewWorkspacesClient() *WorkspacesClient {
	return &WorkspacesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}
