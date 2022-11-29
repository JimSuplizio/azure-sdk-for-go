//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/PoolList.json
func ExamplePoolClient_NewListByBatchAccountPager_listPool() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbatch.NewPoolClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByBatchAccountPager("default-azurebatch-japaneast", "sampleacct", &armbatch.PoolClientListByBatchAccountOptions{Maxresults: nil,
		Select: nil,
		Filter: nil,
	})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/PoolListWithFilter.json
func ExamplePoolClient_NewListByBatchAccountPager_listPoolWithFilter() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbatch.NewPoolClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByBatchAccountPager("default-azurebatch-japaneast", "sampleacct", &armbatch.PoolClientListByBatchAccountOptions{Maxresults: nil,
		Select: to.Ptr("properties/allocationState,properties/provisioningStateTransitionTime,properties/currentDedicatedNodes,properties/currentLowPriorityNodes"),
		Filter: to.Ptr("startswith(name, 'po') or (properties/allocationState eq 'Steady' and properties/provisioningStateTransitionTime lt datetime'2017-02-02')"),
	})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/PoolCreate_SharedImageGallery.json
func ExamplePoolClient_Create_createPoolCustomImage() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbatch.NewPoolClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx, "default-azurebatch-japaneast", "sampleacct", "testpool", armbatch.Pool{
		Properties: &armbatch.PoolProperties{
			DeploymentConfiguration: &armbatch.DeploymentConfiguration{
				VirtualMachineConfiguration: &armbatch.VirtualMachineConfiguration{
					ImageReference: &armbatch.ImageReference{
						ID: to.Ptr("/subscriptions/subid/resourceGroups/networking-group/providers/Microsoft.Compute/galleries/testgallery/images/testimagedef/versions/0.0.1"),
					},
					NodeAgentSKUID: to.Ptr("batch.node.ubuntu 18.04"),
				},
			},
			VMSize: to.Ptr("STANDARD_D4"),
		},
	}, &armbatch.PoolClientCreateOptions{IfMatch: nil,
		IfNoneMatch: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/PoolCreate_CloudServiceConfiguration.json
func ExamplePoolClient_Create_createPoolFullCloudServiceConfiguration() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbatch.NewPoolClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx, "default-azurebatch-japaneast", "sampleacct", "testpool", armbatch.Pool{
		Properties: &armbatch.PoolProperties{
			ApplicationLicenses: []*string{
				to.Ptr("app-license0"),
				to.Ptr("app-license1")},
			ApplicationPackages: []*armbatch.ApplicationPackageReference{
				{
					ID:      to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool/applications/app_1234"),
					Version: to.Ptr("asdf"),
				}},
			Certificates: []*armbatch.CertificateReference{
				{
					ID:            to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool/certificates/sha1-1234567"),
					StoreLocation: to.Ptr(armbatch.CertificateStoreLocationLocalMachine),
					StoreName:     to.Ptr("MY"),
					Visibility: []*armbatch.CertificateVisibility{
						to.Ptr(armbatch.CertificateVisibilityRemoteUser)},
				}},
			DeploymentConfiguration: &armbatch.DeploymentConfiguration{
				CloudServiceConfiguration: &armbatch.CloudServiceConfiguration{
					OSFamily:  to.Ptr("4"),
					OSVersion: to.Ptr("WA-GUEST-OS-4.45_201708-01"),
				},
			},
			DisplayName:            to.Ptr("my-pool-name"),
			InterNodeCommunication: to.Ptr(armbatch.InterNodeCommunicationStateEnabled),
			Metadata: []*armbatch.MetadataItem{
				{
					Name:  to.Ptr("metadata-1"),
					Value: to.Ptr("value-1"),
				},
				{
					Name:  to.Ptr("metadata-2"),
					Value: to.Ptr("value-2"),
				}},
			NetworkConfiguration: &armbatch.NetworkConfiguration{
				PublicIPAddressConfiguration: &armbatch.PublicIPAddressConfiguration{
					IPAddressIDs: []*string{
						to.Ptr("/subscriptions/subid1/resourceGroups/rg13/providers/Microsoft.Network/publicIPAddresses/ip135"),
						to.Ptr("/subscriptions/subid2/resourceGroups/rg24/providers/Microsoft.Network/publicIPAddresses/ip268")},
					Provision: to.Ptr(armbatch.IPAddressProvisioningTypeUserManaged),
				},
				SubnetID: to.Ptr("/subscriptions/subid/resourceGroups/rg1234/providers/Microsoft.Network/virtualNetworks/network1234/subnets/subnet123"),
			},
			ScaleSettings: &armbatch.ScaleSettings{
				FixedScale: &armbatch.FixedScaleSettings{
					NodeDeallocationOption: to.Ptr(armbatch.ComputeNodeDeallocationOptionTaskCompletion),
					ResizeTimeout:          to.Ptr("PT8M"),
					TargetDedicatedNodes:   to.Ptr[int32](6),
					TargetLowPriorityNodes: to.Ptr[int32](28),
				},
			},
			StartTask: &armbatch.StartTask{
				CommandLine: to.Ptr("cmd /c SET"),
				EnvironmentSettings: []*armbatch.EnvironmentSetting{
					{
						Name:  to.Ptr("MYSET"),
						Value: to.Ptr("1234"),
					}},
				MaxTaskRetryCount: to.Ptr[int32](6),
				ResourceFiles: []*armbatch.ResourceFile{
					{
						FileMode: to.Ptr("777"),
						FilePath: to.Ptr("c:\\temp\\gohere"),
						HTTPURL:  to.Ptr("https://testaccount.blob.core.windows.net/example-blob-file"),
					}},
				UserIdentity: &armbatch.UserIdentity{
					AutoUser: &armbatch.AutoUserSpecification{
						ElevationLevel: to.Ptr(armbatch.ElevationLevelAdmin),
						Scope:          to.Ptr(armbatch.AutoUserScopePool),
					},
				},
				WaitForSuccess: to.Ptr(true),
			},
			TaskSchedulingPolicy: &armbatch.TaskSchedulingPolicy{
				NodeFillType: to.Ptr(armbatch.ComputeNodeFillTypePack),
			},
			TaskSlotsPerNode: to.Ptr[int32](13),
			UserAccounts: []*armbatch.UserAccount{
				{
					Name:           to.Ptr("username1"),
					ElevationLevel: to.Ptr(armbatch.ElevationLevelAdmin),
					LinuxUserConfiguration: &armbatch.LinuxUserConfiguration{
						Gid:           to.Ptr[int32](4567),
						SSHPrivateKey: to.Ptr("sshprivatekeyvalue"),
						UID:           to.Ptr[int32](1234),
					},
					Password: to.Ptr("<ExamplePassword>"),
				}},
			VMSize: to.Ptr("STANDARD_D4"),
		},
	}, &armbatch.PoolClientCreateOptions{IfMatch: nil,
		IfNoneMatch: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/PoolCreate_VirtualMachineConfiguration.json
func ExamplePoolClient_Create_createPoolFullVirtualMachineConfiguration() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbatch.NewPoolClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx, "default-azurebatch-japaneast", "sampleacct", "testpool", armbatch.Pool{
		Properties: &armbatch.PoolProperties{
			DeploymentConfiguration: &armbatch.DeploymentConfiguration{
				VirtualMachineConfiguration: &armbatch.VirtualMachineConfiguration{
					DataDisks: []*armbatch.DataDisk{
						{
							Caching:            to.Ptr(armbatch.CachingTypeReadWrite),
							DiskSizeGB:         to.Ptr[int32](30),
							Lun:                to.Ptr[int32](0),
							StorageAccountType: to.Ptr(armbatch.StorageAccountTypePremiumLRS),
						},
						{
							Caching:            to.Ptr(armbatch.CachingTypeNone),
							DiskSizeGB:         to.Ptr[int32](200),
							Lun:                to.Ptr[int32](1),
							StorageAccountType: to.Ptr(armbatch.StorageAccountTypeStandardLRS),
						}},
					DiskEncryptionConfiguration: &armbatch.DiskEncryptionConfiguration{
						Targets: []*armbatch.DiskEncryptionTarget{
							to.Ptr(armbatch.DiskEncryptionTargetOsDisk),
							to.Ptr(armbatch.DiskEncryptionTargetTemporaryDisk)},
					},
					ImageReference: &armbatch.ImageReference{
						Offer:     to.Ptr("WindowsServer"),
						Publisher: to.Ptr("MicrosoftWindowsServer"),
						SKU:       to.Ptr("2016-Datacenter-SmallDisk"),
						Version:   to.Ptr("latest"),
					},
					LicenseType:    to.Ptr("Windows_Server"),
					NodeAgentSKUID: to.Ptr("batch.node.windows amd64"),
					NodePlacementConfiguration: &armbatch.NodePlacementConfiguration{
						Policy: to.Ptr(armbatch.NodePlacementPolicyTypeZonal),
					},
					OSDisk: &armbatch.OSDisk{
						EphemeralOSDiskSettings: &armbatch.DiffDiskSettings{
							Placement: to.Ptr("CacheDisk"),
						},
					},
					WindowsConfiguration: &armbatch.WindowsConfiguration{
						EnableAutomaticUpdates: to.Ptr(false),
					},
				},
			},
			NetworkConfiguration: &armbatch.NetworkConfiguration{
				EndpointConfiguration: &armbatch.PoolEndpointConfiguration{
					InboundNatPools: []*armbatch.InboundNatPool{
						{
							Name:                   to.Ptr("testnat"),
							BackendPort:            to.Ptr[int32](12001),
							FrontendPortRangeEnd:   to.Ptr[int32](15100),
							FrontendPortRangeStart: to.Ptr[int32](15000),
							NetworkSecurityGroupRules: []*armbatch.NetworkSecurityGroupRule{
								{
									Access:              to.Ptr(armbatch.NetworkSecurityGroupRuleAccessAllow),
									Priority:            to.Ptr[int32](150),
									SourceAddressPrefix: to.Ptr("192.100.12.45"),
									SourcePortRanges: []*string{
										to.Ptr("1"),
										to.Ptr("2")},
								},
								{
									Access:              to.Ptr(armbatch.NetworkSecurityGroupRuleAccessDeny),
									Priority:            to.Ptr[int32](3500),
									SourceAddressPrefix: to.Ptr("*"),
									SourcePortRanges: []*string{
										to.Ptr("*")},
								}},
							Protocol: to.Ptr(armbatch.InboundEndpointProtocolTCP),
						}},
				},
			},
			ScaleSettings: &armbatch.ScaleSettings{
				AutoScale: &armbatch.AutoScaleSettings{
					EvaluationInterval: to.Ptr("PT5M"),
					Formula:            to.Ptr("$TargetDedicatedNodes=1"),
				},
			},
			VMSize: to.Ptr("STANDARD_D4"),
		},
	}, &armbatch.PoolClientCreateOptions{IfMatch: nil,
		IfNoneMatch: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/PoolCreate_MinimalCloudServiceConfiguration.json
func ExamplePoolClient_Create_createPoolMinimalCloudServiceConfiguration() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbatch.NewPoolClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx, "default-azurebatch-japaneast", "sampleacct", "testpool", armbatch.Pool{
		Properties: &armbatch.PoolProperties{
			DeploymentConfiguration: &armbatch.DeploymentConfiguration{
				CloudServiceConfiguration: &armbatch.CloudServiceConfiguration{
					OSFamily: to.Ptr("5"),
				},
			},
			ScaleSettings: &armbatch.ScaleSettings{
				FixedScale: &armbatch.FixedScaleSettings{
					TargetDedicatedNodes: to.Ptr[int32](3),
				},
			},
			VMSize: to.Ptr("STANDARD_D4"),
		},
	}, &armbatch.PoolClientCreateOptions{IfMatch: nil,
		IfNoneMatch: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/PoolCreate_MinimalVirtualMachineConfiguration.json
func ExamplePoolClient_Create_createPoolMinimalVirtualMachineConfiguration() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbatch.NewPoolClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx, "default-azurebatch-japaneast", "sampleacct", "testpool", armbatch.Pool{
		Properties: &armbatch.PoolProperties{
			DeploymentConfiguration: &armbatch.DeploymentConfiguration{
				VirtualMachineConfiguration: &armbatch.VirtualMachineConfiguration{
					ImageReference: &armbatch.ImageReference{
						Offer:     to.Ptr("UbuntuServer"),
						Publisher: to.Ptr("Canonical"),
						SKU:       to.Ptr("18.04-LTS"),
						Version:   to.Ptr("latest"),
					},
					NodeAgentSKUID: to.Ptr("batch.node.ubuntu 18.04"),
				},
			},
			ScaleSettings: &armbatch.ScaleSettings{
				AutoScale: &armbatch.AutoScaleSettings{
					EvaluationInterval: to.Ptr("PT5M"),
					Formula:            to.Ptr("$TargetDedicatedNodes=1"),
				},
			},
			VMSize: to.Ptr("STANDARD_D4"),
		},
	}, &armbatch.PoolClientCreateOptions{IfMatch: nil,
		IfNoneMatch: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/PoolCreate_NoPublicIPAddresses.json
func ExamplePoolClient_Create_createPoolNoPublicIp() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbatch.NewPoolClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx, "default-azurebatch-japaneast", "sampleacct", "testpool", armbatch.Pool{
		Properties: &armbatch.PoolProperties{
			DeploymentConfiguration: &armbatch.DeploymentConfiguration{
				VirtualMachineConfiguration: &armbatch.VirtualMachineConfiguration{
					ImageReference: &armbatch.ImageReference{
						ID: to.Ptr("/subscriptions/subid/resourceGroups/networking-group/providers/Microsoft.Compute/galleries/testgallery/images/testimagedef/versions/0.0.1"),
					},
					NodeAgentSKUID: to.Ptr("batch.node.ubuntu 18.04"),
				},
			},
			NetworkConfiguration: &armbatch.NetworkConfiguration{
				PublicIPAddressConfiguration: &armbatch.PublicIPAddressConfiguration{
					Provision: to.Ptr(armbatch.IPAddressProvisioningTypeNoPublicIPAddresses),
				},
				SubnetID: to.Ptr("/subscriptions/subid/resourceGroups/rg1234/providers/Microsoft.Network/virtualNetworks/network1234/subnets/subnet123"),
			},
			VMSize: to.Ptr("STANDARD_D4"),
		},
	}, &armbatch.PoolClientCreateOptions{IfMatch: nil,
		IfNoneMatch: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/PoolCreate_PublicIPs.json
func ExamplePoolClient_Create_createPoolPublicIPs() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbatch.NewPoolClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx, "default-azurebatch-japaneast", "sampleacct", "testpool", armbatch.Pool{
		Properties: &armbatch.PoolProperties{
			DeploymentConfiguration: &armbatch.DeploymentConfiguration{
				VirtualMachineConfiguration: &armbatch.VirtualMachineConfiguration{
					ImageReference: &armbatch.ImageReference{
						ID: to.Ptr("/subscriptions/subid/resourceGroups/networking-group/providers/Microsoft.Compute/galleries/testgallery/images/testimagedef/versions/0.0.1"),
					},
					NodeAgentSKUID: to.Ptr("batch.node.ubuntu 18.04"),
				},
			},
			NetworkConfiguration: &armbatch.NetworkConfiguration{
				PublicIPAddressConfiguration: &armbatch.PublicIPAddressConfiguration{
					IPAddressIDs: []*string{
						to.Ptr("/subscriptions/subid1/resourceGroups/rg13/providers/Microsoft.Network/publicIPAddresses/ip135")},
					Provision: to.Ptr(armbatch.IPAddressProvisioningTypeUserManaged),
				},
				SubnetID: to.Ptr("/subscriptions/subid/resourceGroups/rg1234/providers/Microsoft.Network/virtualNetworks/network1234/subnets/subnet123"),
			},
			VMSize: to.Ptr("STANDARD_D4"),
		},
	}, &armbatch.PoolClientCreateOptions{IfMatch: nil,
		IfNoneMatch: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/PoolCreate_UserAssignedIdentities.json
func ExamplePoolClient_Create_createPoolUserAssignedIdentities() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbatch.NewPoolClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx, "default-azurebatch-japaneast", "sampleacct", "testpool", armbatch.Pool{
		Identity: &armbatch.PoolIdentity{
			Type: to.Ptr(armbatch.PoolIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armbatch.UserAssignedIdentities{
				"/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": {},
				"/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id2": {},
			},
		},
		Properties: &armbatch.PoolProperties{
			DeploymentConfiguration: &armbatch.DeploymentConfiguration{
				VirtualMachineConfiguration: &armbatch.VirtualMachineConfiguration{
					ImageReference: &armbatch.ImageReference{
						Offer:     to.Ptr("UbuntuServer"),
						Publisher: to.Ptr("Canonical"),
						SKU:       to.Ptr("18.04-LTS"),
						Version:   to.Ptr("latest"),
					},
					NodeAgentSKUID: to.Ptr("batch.node.ubuntu 18.04"),
				},
			},
			ScaleSettings: &armbatch.ScaleSettings{
				AutoScale: &armbatch.AutoScaleSettings{
					EvaluationInterval: to.Ptr("PT5M"),
					Formula:            to.Ptr("$TargetDedicatedNodes=1"),
				},
			},
			VMSize: to.Ptr("STANDARD_D4"),
		},
	}, &armbatch.PoolClientCreateOptions{IfMatch: nil,
		IfNoneMatch: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/PoolCreate_VirtualMachineConfiguration_Extensions.json
func ExamplePoolClient_Create_createPoolVirtualMachineConfigurationExtensions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbatch.NewPoolClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx, "default-azurebatch-japaneast", "sampleacct", "testpool", armbatch.Pool{
		Properties: &armbatch.PoolProperties{
			DeploymentConfiguration: &armbatch.DeploymentConfiguration{
				VirtualMachineConfiguration: &armbatch.VirtualMachineConfiguration{
					ImageReference: &armbatch.ImageReference{
						Offer:     to.Ptr("0001-com-ubuntu-server-focal"),
						Publisher: to.Ptr("Canonical"),
						SKU:       to.Ptr("20_04-lts"),
					},
					NodeAgentSKUID: to.Ptr("batch.node.ubuntu 20.04"),
					Extensions: []*armbatch.VMExtension{
						{
							Name:                    to.Ptr("batchextension1"),
							Type:                    to.Ptr("SecurityMonitoringForLinux"),
							AutoUpgradeMinorVersion: to.Ptr(true),
							ProtectedSettings: map[string]interface{}{
								"protectedSettingsKey": "protectedSettingsValue",
							},
							Publisher: to.Ptr("Microsoft.Azure.Security.Monitoring"),
							Settings: map[string]interface{}{
								"settingsKey": "settingsValue",
							},
							TypeHandlerVersion: to.Ptr("1.0"),
						}},
				},
			},
			ScaleSettings: &armbatch.ScaleSettings{
				AutoScale: &armbatch.AutoScaleSettings{
					EvaluationInterval: to.Ptr("PT5M"),
					Formula:            to.Ptr("$TargetDedicatedNodes=1"),
				},
			},
			TargetNodeCommunicationMode: to.Ptr(armbatch.NodeCommunicationModeDefault),
			VMSize:                      to.Ptr("STANDARD_D4"),
		},
	}, &armbatch.PoolClientCreateOptions{IfMatch: nil,
		IfNoneMatch: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/PoolUpdate_EnableAutoScale.json
func ExamplePoolClient_Update_updatePoolEnableAutoscale() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbatch.NewPoolClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx, "default-azurebatch-japaneast", "sampleacct", "testpool", armbatch.Pool{
		Properties: &armbatch.PoolProperties{
			ScaleSettings: &armbatch.ScaleSettings{
				AutoScale: &armbatch.AutoScaleSettings{
					Formula: to.Ptr("$TargetDedicatedNodes=34"),
				},
			},
		},
	}, &armbatch.PoolClientUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/PoolUpdate_OtherProperties.json
func ExamplePoolClient_Update_updatePoolOtherProperties() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbatch.NewPoolClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx, "default-azurebatch-japaneast", "sampleacct", "testpool", armbatch.Pool{
		Properties: &armbatch.PoolProperties{
			ApplicationPackages: []*armbatch.ApplicationPackageReference{
				{
					ID: to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool/applications/app_1234"),
				},
				{
					ID:      to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool/applications/app_5678"),
					Version: to.Ptr("1.0"),
				}},
			Certificates: []*armbatch.CertificateReference{
				{
					ID:            to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool/certificates/sha1-1234567"),
					StoreLocation: to.Ptr(armbatch.CertificateStoreLocationLocalMachine),
					StoreName:     to.Ptr("MY"),
				}},
			Metadata: []*armbatch.MetadataItem{
				{
					Name:  to.Ptr("key1"),
					Value: to.Ptr("value1"),
				}},
			TargetNodeCommunicationMode: to.Ptr(armbatch.NodeCommunicationModeSimplified),
		},
	}, &armbatch.PoolClientUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/PoolUpdate_RemoveStartTask.json
func ExamplePoolClient_Update_updatePoolRemoveStartTask() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbatch.NewPoolClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx, "default-azurebatch-japaneast", "sampleacct", "testpool", armbatch.Pool{
		Properties: &armbatch.PoolProperties{
			StartTask: &armbatch.StartTask{},
		},
	}, &armbatch.PoolClientUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/PoolUpdate_ResizePool.json
func ExamplePoolClient_Update_updatePoolResizePool() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbatch.NewPoolClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx, "default-azurebatch-japaneast", "sampleacct", "testpool", armbatch.Pool{
		Properties: &armbatch.PoolProperties{
			ScaleSettings: &armbatch.ScaleSettings{
				FixedScale: &armbatch.FixedScaleSettings{
					NodeDeallocationOption: to.Ptr(armbatch.ComputeNodeDeallocationOptionTaskCompletion),
					ResizeTimeout:          to.Ptr("PT8M"),
					TargetDedicatedNodes:   to.Ptr[int32](5),
					TargetLowPriorityNodes: to.Ptr[int32](0),
				},
			},
		},
	}, &armbatch.PoolClientUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/PoolDelete.json
func ExamplePoolClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbatch.NewPoolClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx, "default-azurebatch-japaneast", "sampleacct", "testpool", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/PoolGet.json
func ExamplePoolClient_Get_getPool() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbatch.NewPoolClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "default-azurebatch-japaneast", "sampleacct", "testpool", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/PoolGet_VirtualMachineConfiguration_Extensions.json
func ExamplePoolClient_Get_getPoolVirtualMachineConfigurationExtensions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbatch.NewPoolClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "default-azurebatch-japaneast", "sampleacct", "testpool", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/PoolDisableAutoScale.json
func ExamplePoolClient_DisableAutoScale() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbatch.NewPoolClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.DisableAutoScale(ctx, "default-azurebatch-japaneast", "sampleacct", "testpool", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/PoolStopResize.json
func ExamplePoolClient_StopResize() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbatch.NewPoolClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.StopResize(ctx, "default-azurebatch-japaneast", "sampleacct", "testpool", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}