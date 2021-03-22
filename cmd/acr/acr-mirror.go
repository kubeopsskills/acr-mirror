package main

import (
	"github.com/Azure/azure-sdk-for-go/profiles/latest/containerregistry/mgmt/containerregistry"
	"github.com/kubeopsskills/acr-mirror/internal/acr"
	"github.com/kubeopsskills/acr-mirror/internal/utils"
)

func main() {

	configFileName := utils.GetEnv("CONFIG_FILE_NAME", "config.yaml")
	resourceURL := utils.GetEnv("RESOURCE_URL", "https://management.azure.com")
	ctx, authorizer := acr.AuthWithCLI(resourceURL)
	var config acr.Config
	config.GetConfig(configFileName)

	sourceRegistryClient := containerregistry.NewRegistriesClient(config.SourceSubscriptionID)
	sourceRegistryClient.Authorizer = authorizer

	targetRegistryClient := containerregistry.NewRegistriesClient(config.TargetSubscriptionID)
	targetRegistryClient.Authorizer = authorizer

	acr.FromSourceRegistryToTargetRegistry(sourceRegistryClient, targetRegistryClient, config, ctx)

}
