package api

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Aadil-Nabi/cmgarage/auth/jwtauth"
	"github.com/Aadil-Nabi/cmgarage/configs/envs"
	"github.com/Aadil-Nabi/cmgarage/internal/pkg/cmhttpclient"
)

// GetAuthDetails to get the Authentication details
var token_details = jwtauth.GetAuthDetails()

// token to store the token details
var token = jwtauth.JWTData{
	Jwt:              token_details.Jwt,
	Duration:         token_details.Duration,
	Token_type:       token_details.Token_type,
	Client_id:        token_details.Client_id,
	Refresh_token_id: token_details.Refresh_token_id,
	Refresh_token:    token_details.Refresh_token,
}

// Encrypting method to encrypt the data using the provided key
func Encrypting() {

	jwt_token := token.Jwt
	jwt_token_type := token.Token_type

	//Bearer
	bearer := jwt_token_type + " " + jwt_token

	url := "https://192.168.238.129/api/v1/crypto/encrypt"

	envs := envs.GetEnvs()

	// Encode the data to be encrypted in base64 string as CM only accepts a valid base64 string
	plaintext := "ksdnckjsbdhkcbsdkjncbkhsdbckjsfdvbfbvkjhjkbfvhbndfkjvbksjdncjksdvjkfbvh"
	plaintext = base64.StdEncoding.EncodeToString([]byte(plaintext))
	payload := map[string]string{
		"id":        envs["ENCRYPTION_KEY"],
		"plaintext": plaintext,
	}

	// Convert data into JSON encoded byte array
	encodedBody, _ := json.Marshal(payload)

	// convert the encoded JSON data to a type implemented by the io.Reader interface
	body := bytes.NewBuffer(encodedBody)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Fatalf("Something went wrong in the request  %v", err)
	}

	// Add the required headers to the request
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", bearer)

	//get client from a helper function
	client := cmhttpclient.GetClient()

	// Do method to send the http request to the CM to http response
	// this is used when we add headers to the request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Unable to Encrypt %v", err)
	}

	// close the response
	defer resp.Body.Close()

	// Read the response received from the CM
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))

}
