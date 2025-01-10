package api

import (
	"encoding/json"
	"io"
	"log"
	"log/slog"
	"net/http"

	"github.com/Aadil-Nabi/cmgarage/auth/jwtauth"
	"github.com/Aadil-Nabi/cmgarage/internal/pkg/cmhttpclient"
)

var jwt_details = jwtauth.GetAuthDetails()

type CmClusterStatus struct {
	Code        string
	Description string
}

type CMclusterInfo struct {
	NodeID    string
	Status    CmClusterStatus
	NodeCount int
}

func ClusterStatus() {

	// Get the Bearer Token
	Bearer := jwt_details.Token_type + " " + jwt_details.Jwt

	clusterInfoUrl := "https://192.168.238.129/api/v1/cluster"

	// Create a new request for cluster info API
	reqClusterInfo, err := http.NewRequest("GET", clusterInfoUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Add Headers to the request for cluster info
	reqClusterInfo.Header.Add("Authorization", Bearer)

	// Get http client and call a Do function on it to send the http request for cluster info
	clientClusterInfo := cmhttpclient.GetClient()
	respClusterInfo, err := clientClusterInfo.Do(reqClusterInfo)
	if err != nil {
		log.Fatalf("error making request %v", err)
	}

	// Close the response body
	defer respClusterInfo.Body.Close()

	// Read data from response received from Do function called above for cluster info
	data_info, err := io.ReadAll(respClusterInfo.Body)
	if err != nil {
		log.Fatal(err)
	}

	var cmClusterInfo CMclusterInfo
	err = json.Unmarshal(data_info, &cmClusterInfo)
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	statusDescription := cmClusterInfo.Status.Description
	if statusDescription == "ready" {
		slog.Info("Primary cluster node is UP and Running")
	}

	GetclusterErrors()

}
