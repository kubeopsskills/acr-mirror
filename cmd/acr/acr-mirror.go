package main

import (
	"github.com/Azure/azure-sdk-for-go/profiles/latest/containerregistry/mgmt/containerregistry"
	"github.com/kubeopsskills/acr-mirror/internal/acr"
	"github.com/kubeopsskills/acr-mirror/internal/utils"
	log "github.com/sirupsen/logrus"
)

func main() {

	configFileName := utils.GetEnv("CONFIG_FILE_NAME", "config.yaml")
	resourceURL := utils.GetEnv("RESOURCE_URL", "https://management.azure.com")

	log.Info("Authenticating to Azure ...")
	ctx, authorizer := acr.AuthWithCLI(resourceURL)
	log.Info("Authenticated to Azure")
	var config acr.Config
	config.GetConfig(configFileName)

	sourceRegistryClient := containerregistry.NewRegistriesClient(config.SourceSubscriptionID)
	sourceRegistryClient.Authorizer = authorizer

	targetRegistryClient := containerregistry.NewRegistriesClient(config.TargetSubscriptionID)
	targetRegistryClient.Authorizer = authorizer

	log.Info("Starting Mirroring Images ...")
	acr.FromSourceRegistryToTargetRegistry(sourceRegistryClient, targetRegistryClient, config, ctx)
	log.Info("Completed Mirroring Images")

}
