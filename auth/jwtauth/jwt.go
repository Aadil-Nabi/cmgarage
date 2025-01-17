package jwtauth

import (
	"bytes"
	"encoding/json"
	"io"
	"log"

	"github.com/Aadil-Nabi/cmgarage/internal/config"
	"github.com/Aadil-Nabi/cmgarage/internal/pkg/cmhttpclient"
	"github.com/Aadil-Nabi/cmgarage/internal/secrets"
)

// JWTData is a struct to store the key and values of JSON payload received after unmarshing.
type JWTData struct {
	Jwt              string
	Duration         int
	Token_type       string
	Client_id        string
	Refresh_token_id string
	Refresh_token    string
}

func GetAuthDetails() *JWTData {

	var dat1 JWTData

	jwt_details := getJwt()

	err := json.Unmarshal(jwt_details, &dat1)
	if err != nil {
		log.Fatalf("cannot unmarshal the data %v", err)
	}

	return &dat1
}

func getJwt() []byte {

	// Fetch the username and password from config.yaml file provide in the command line
	configs := config.MustLoad()

	// Get Secrets from the AKeyless account, here we fect the CM Password from the the AKeyless vault.
	// We can get the password from the config.yaml file, but we are fecthing the CM Password from the AKeyless Vault
	secrets := secrets.GetSecrets()
	cm_password := secrets["cm_pass"]
	cm_passwd := cm_password.(string) // this interface value is converted into string using "type assertion"

	// payload to be sent to get the JWT token
	payload := map[string]string{
		"grant_type": "password",
		"username":   configs.Cm_user,
		"password":   cm_passwd,
		"token_type": "Bearer",
	}

	// Encode the jason payload
	encodedBody, _ := json.Marshal(payload)

	// convert the encoded JSON data to a type implemented by the io.Reader interface
	body := bytes.NewBuffer(encodedBody)

	// get the base url and version from the local.yaml file.
	url := configs.Base_Url + configs.Version + "/auth/tokens"

	// get the client to perform the Post operation
	client := cmhttpclient.GetClient()

	// instead of http, use the client for the post request
	resp, err := client.Post(url, "application/json", body)
	if err != nil {
		log.Fatalf("Error:  %s", err)
	}
	// close the response once the function execution is done.
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// dat contains the jwt token and related details
	return dat

}
