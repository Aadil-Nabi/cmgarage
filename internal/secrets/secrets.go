package secrets

import (
	"context"
	"errors"
	"log"

	"github.com/Aadil-Nabi/cmgarage/internal/config"
	"github.com/akeylesslabs/akeyless-go/v4"
)

// GetSecrets contains the code to get the secret stored in Akeyless account
func GetSecrets() map[string]interface{} {

	ctx := context.Background()
	client := akeyless.NewAPIClient(&akeyless.Configuration{
		Servers: []akeyless.ServerConfiguration{
			{
				// default: public API Gateway
				URL: "https://api.akeyless.io",

				// use port 8081 exposed by the deployment:
				// URL: "https://gateway.company.com:8081",

				// use port 8080 exposed by the deployment with /v2 prefix:
				// URL: "https://gateway.company.com:8080/v2",
			},
		},
	}).V2Api

	// Load configuartions from the config.yaml file
	configs := config.MustLoad()
	email := configs.AkeylessUsername
	password := configs.AkeylessPassword

	authBody := akeyless.NewAuthWithDefaults()
	authBody.AdminEmail = akeyless.PtrString(email)
	authBody.AdminPassword = akeyless.PtrString(password)

	var apiErr akeyless.GenericOpenAPIError

	authOut, _, err := client.Auth(ctx).Body(*authBody).Execute()
	if err != nil {
		if errors.As(err, &apiErr) {
			log.Fatalln("authentication failed:", string(apiErr.Body()))
		}
		log.Fatalln("authentication failed:", err)
	}

	token := authOut.GetToken()

	gsvBody := akeyless.GetSecretValue{
		Names: []string{"cm_pass"},
		Token: &token,
	}
	gsvOut, _, err := client.GetSecretValue(ctx).Body(gsvBody).Execute()
	if err != nil {
		if errors.As(err, &apiErr) {
			log.Fatalln("can't get secret value:", string(apiErr.Body()))
		}
		log.Fatalln("can't get secret value:", err)
	}

	return gsvOut
}
