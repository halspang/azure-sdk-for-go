//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armguestconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/guestconfiguration/armguestconfiguration"
)

// x-ms-original-file: specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2020-06-25/examples/createOrUpdateGuestConfigurationHCRPAssignment.json
func ExampleGuestConfigurationHCRPAssignmentsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armguestconfiguration.NewGuestConfigurationHCRPAssignmentsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<guest-configuration-assignment-name>",
		"<resource-group-name>",
		"<machine-name>",
		armguestconfiguration.GuestConfigurationAssignment{
			ProxyResource: armguestconfiguration.ProxyResource{
				Resource: armguestconfiguration.Resource{
					Name:     to.StringPtr("<name>"),
					Location: to.StringPtr("<location>"),
				},
			},
			Properties: &armguestconfiguration.GuestConfigurationAssignmentProperties{
				Context: to.StringPtr("<context>"),
				GuestConfiguration: &armguestconfiguration.GuestConfigurationNavigation{
					Name:           to.StringPtr("<name>"),
					AssignmentType: armguestconfiguration.AssignmentTypeApplyAndAutoCorrect.ToPtr(),
					ConfigurationParameter: []*armguestconfiguration.ConfigurationParameter{
						{
							Name:  to.StringPtr("<name>"),
							Value: to.StringPtr("<value>"),
						}},
					ContentHash: to.StringPtr("<content-hash>"),
					ContentURI:  to.StringPtr("<content-uri>"),
					Version:     to.StringPtr("<version>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("GuestConfigurationAssignment.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2020-06-25/examples/getGuestConfigurationHCRPAssignment.json
func ExampleGuestConfigurationHCRPAssignmentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armguestconfiguration.NewGuestConfigurationHCRPAssignmentsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<guest-configuration-assignment-name>",
		"<machine-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("GuestConfigurationAssignment.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2020-06-25/examples/deleteGuestConfigurationHCRPAssignment.json
func ExampleGuestConfigurationHCRPAssignmentsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armguestconfiguration.NewGuestConfigurationHCRPAssignmentsClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<guest-configuration-assignment-name>",
		"<machine-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2020-06-25/examples/listGuestConfigurationHCRPAssignments.json
func ExampleGuestConfigurationHCRPAssignmentsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armguestconfiguration.NewGuestConfigurationHCRPAssignmentsClient("<subscription-id>", cred, nil)
	_, err = client.List(ctx,
		"<resource-group-name>",
		"<machine-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
