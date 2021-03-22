package acr

import (
	"context"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func AuthWithCLI(resourceURL string) (context.Context, autorest.Authorizer) {
	ctx := context.Background()
	authorizer, err := auth.NewAuthorizerFromCLIWithResource(resourceURL)
	if err != nil {
		panic(err)
	}
	return ctx, authorizer
}
