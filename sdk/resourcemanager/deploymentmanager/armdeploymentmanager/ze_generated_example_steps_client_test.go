//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdeploymentmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deploymentmanager/armdeploymentmanager"
)

// x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/step_health_check_createorupdate.json
func ExampleStepsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdeploymentmanager.NewStepsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<step-name>",
		&armdeploymentmanager.StepsCreateOrUpdateOptions{StepInfo: &armdeploymentmanager.StepResource{
			TrackedResource: armdeploymentmanager.TrackedResource{
				Location: to.StringPtr("<location>"),
				Tags:     map[string]*string{},
			},
			Properties: &armdeploymentmanager.HealthCheckStepProperties{
				StepProperties: armdeploymentmanager.StepProperties{
					StepType: armdeploymentmanager.StepTypeHealthCheck.ToPtr(),
				},
				Attributes: &armdeploymentmanager.RestHealthCheckStepAttributes{
					HealthCheckStepAttributes: armdeploymentmanager.HealthCheckStepAttributes{
						Type:                 to.StringPtr("<type>"),
						HealthyStateDuration: to.StringPtr("<healthy-state-duration>"),
						MaxElasticDuration:   to.StringPtr("<max-elastic-duration>"),
						WaitDuration:         to.StringPtr("<wait-duration>"),
					},
					Properties: &armdeploymentmanager.RestParameters{
						HealthChecks: []*armdeploymentmanager.RestHealthCheck{
							{
								Name: to.StringPtr("<name>"),
								Response: &armdeploymentmanager.RestResponse{
									Regex: &armdeploymentmanager.RestResponseRegex{
										MatchQuantifier: armdeploymentmanager.RestMatchQuantifierAll.ToPtr(),
										Matches: []*string{
											to.StringPtr("(?i)Contoso-App"),
											to.StringPtr("(?i)\"health_status\":((.|\n)*)\"(green|yellow)\""),
											to.StringPtr("(?mi)^(\"application_host\": 94781052)$")},
									},
									SuccessStatusCodes: []*string{
										to.StringPtr("OK")},
								},
								Request: &armdeploymentmanager.RestRequest{
									Method: armdeploymentmanager.RestRequestMethodGET.ToPtr(),
									Authentication: &armdeploymentmanager.APIKeyAuthentication{
										RestRequestAuthentication: armdeploymentmanager.RestRequestAuthentication{
											Type: armdeploymentmanager.RestAuthTypeAPIKey.ToPtr(),
										},
										Name:  to.StringPtr("<name>"),
										In:    armdeploymentmanager.RestAuthLocationQuery.ToPtr(),
										Value: to.StringPtr("<value>"),
									},
									URI: to.StringPtr("<uri>"),
								},
							},
							{
								Name: to.StringPtr("<name>"),
								Response: &armdeploymentmanager.RestResponse{
									Regex: &armdeploymentmanager.RestResponseRegex{
										MatchQuantifier: armdeploymentmanager.RestMatchQuantifierAll.ToPtr(),
										Matches: []*string{
											to.StringPtr("(?i)Contoso-Service-EndToEnd"),
											to.StringPtr("(?i)\"health_status\":((.|\n)*)\"(green)\"")},
									},
									SuccessStatusCodes: []*string{
										to.StringPtr("OK")},
								},
								Request: &armdeploymentmanager.RestRequest{
									Method: armdeploymentmanager.RestRequestMethodGET.ToPtr(),
									Authentication: &armdeploymentmanager.APIKeyAuthentication{
										RestRequestAuthentication: armdeploymentmanager.RestRequestAuthentication{
											Type: armdeploymentmanager.RestAuthTypeAPIKey.ToPtr(),
										},
										Name:  to.StringPtr("<name>"),
										In:    armdeploymentmanager.RestAuthLocationHeader.ToPtr(),
										Value: to.StringPtr("<value>"),
									},
									URI: to.StringPtr("<uri>"),
								},
							}},
					},
				},
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("StepResource.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/step_get.json
func ExampleStepsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdeploymentmanager.NewStepsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<step-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("StepResource.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/step_delete.json
func ExampleStepsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdeploymentmanager.NewStepsClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<step-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/steps_list.json
func ExampleStepsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdeploymentmanager.NewStepsClient("<subscription-id>", cred, nil)
	_, err = client.List(ctx,
		"<resource-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
