package acr

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/containerregistry/mgmt/2019-05-01/containerregistry"
	"github.com/kubeopsskills/acr-mirror/internal/docker"
	log "github.com/sirupsen/logrus"
)

func FromSourceRegistryToTargetRegistry(sourceRegistryClient containerregistry.RegistriesClient, targetRegistryClient containerregistry.RegistriesClient, config Config, ctx context.Context) {

	sourceRegistries := config.Registries
	for _, sourceRegistry := range sourceRegistries {
		registryCredentials, err := sourceRegistryClient.ListCredentials(ctx, sourceRegistry.SourceResourceGroupName, sourceRegistry.Name)
		if err != nil {
			log.Fatalf("Error getting credentials: %v", err)
		}
		username := *registryCredentials.Username
		password := *(*registryCredentials.Passwords)[0].Value

		for _, repository := range sourceRegistry.Repositories {

			tags, err := docker.GetRemoteTags(sourceRegistry.Name, repository, username, password)
			if err != nil {
				log.Fatalf("Error getting remote tags for %s from registry %s: %v", repository, sourceRegistry.Name, err)
			}
			tags = docker.FilterTags(tags, sourceRegistry.Tags)

			for _, targetRegistry := range sourceRegistry.Target {
				for _, tag := range tags {
					repositoryURL := sourceRegistry.Name + ".azurecr.io"
					repositoryTag := repository + ":" + tag
					targetRepositoryTag := ""
					if targetRegistry.RePrefix != nil {
						repository = strings.Replace(repository, strings.Split(repository, "/")[0], *targetRegistry.RePrefix, -1)
					}
					if targetRegistry.Prefix != nil {
						targetRepositoryTag = *targetRegistry.Prefix + "/" + repository + ":" + tag
					} else {
						targetRepositoryTag = repository + ":" + tag
					}
					targetRepositoryTags := make([]string, 0)

					log.Info("Importing image %s to registry %s", repositoryURL+"/"+repositoryTag, targetRegistry.Name)
					targetRepositoryTags = append(targetRepositoryTags, targetRepositoryTag)
					_, err := targetRegistryClient.ImportImage(ctx, targetRegistry.TargetResourceGroupName, targetRegistry.Name, containerregistry.ImportImageParameters{
						Source: &containerregistry.ImportSource{
							RegistryURI: &repositoryURL,
							Credentials: &containerregistry.ImportSourceCredentials{
								Username: &username,
								Password: &password,
							},
							SourceImage: &repositoryTag,
						},
						TargetTags: &targetRepositoryTags,
					})
					if err != nil {
						log.Fatalf("Error importing image %s to registry %s: %v", repositoryURL+"/"+repositoryTag, targetRegistry.Name, err)
					}
					log.Info("Imported image %s to registry %s", repositoryURL+"/"+repositoryTag, targetRegistry.Name)
				}
			}

		}
	}
}
