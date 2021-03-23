package acr

import (
	"context"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	log "github.com/sirupsen/logrus"
)

func AuthWithCLI(resourceURL string) (context.Context, autorest.Authorizer) {
	ctx := context.Background()
	authorizer, err := auth.NewAuthorizerFromCLIWithResource(resourceURL)
	if err != nil {
		log.Error("Error from authenticating to Azure: %v", err)
	}
	return ctx, authorizer
}
