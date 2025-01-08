package jwtauth

import (
	"bytes"
	"encoding/json"
	"io"
	"log"

	"github.com/Aadil-Nabi/CMgarage/configs/envs"
	"github.com/Aadil-Nabi/CMgarage/internal/pkg/cmhttpclient"
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

	// Store the JWT details in a JWTData struct
	dat1 = JWTData{
		Jwt:              dat1.Jwt,
		Duration:         dat1.Duration,
		Token_type:       dat1.Token_type,
		Client_id:        dat1.Client_id,
		Refresh_token_id: dat1.Refresh_token_id,
		Refresh_token:    dat1.Refresh_token,
	}

	// fmt.Println(dat1)

	return &dat1
}

func getJwt() []byte {

	// Fetch the username and password from .env file
	envs := envs.GetEnvs()

	// payload to be sent to get the JWT token
	payload := map[string]string{
		"grant_type": "password",
		"username":   envs["CM_USER"],
		"password":   envs["CM_PASSWORD"],
		"token_type": "Bearer",
	}
	// Encode the jason payload
	encodedBody, _ := json.Marshal(payload)

	// convert the encoded JSON data to a type implemented by the io.Reader interface
	body := bytes.NewBuffer(encodedBody)

	url := "https://192.168.238.129/api/v1/auth/tokens"

	client := cmhttpclient.GetClient()

	// instead of http, use the client for the post request
	resp, err := client.Post(url, "application/json", body)
	if err != nil {
		log.Fatalf("Can't post %s", err)
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
