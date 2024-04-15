package common

import (
	"context"
	"os"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	armNetwork "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v5"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
)

func TestNsgRule(t *testing.T, ctx types.TestContext) {

	subscriptionID := os.Getenv("ARM_SUBSCRIPTION_ID")
	if len(subscriptionID) == 0 {
		t.Fatal("ARM_SUBSCRIPTION_ID is not set in the environment variables ")
	}

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		t.Fatalf("Unable to get credentials: %e\n", err)
	}

	options := arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Cloud: cloud.AzurePublic,
		},
	}
	// Create network security group client
	nsgClient, err := armNetwork.NewSecurityGroupsClient(subscriptionID, credential, &options)
	if err != nil {
		t.Fatalf("Error getting NSG client: %v", err)
	}

	t.Run("doesNsgRuleExist", func(t *testing.T) {
		checkNsgRulesExistence(t, nsgClient, ctx)
	})
}

func checkNsgRulesExistence(t *testing.T, nsgClient *armNetwork.SecurityGroupsClient, ctx types.TestContext) {
	resourceGroupName := terraform.Output(t, ctx.TerratestTerraformOptions(), "resource_group_name")
	nsgName := terraform.Output(t, ctx.TerratestTerraformOptions(), "nsg_name")

	nsg, err := nsgClient.Get(context.Background(), resourceGroupName, nsgName, nil)
	if err != nil {
		t.Fatalf("Error getting NSG: %v", err)
	}
	if nsg.Name == nil {
		t.Fatalf("NSG does not exist")
	}

	assert.NotEmpty(t, nsg.Properties.SecurityRules)
}
