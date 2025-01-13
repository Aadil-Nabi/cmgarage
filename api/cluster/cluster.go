package cluster

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Aadil-Nabi/cmgarage/auth/jwtauth"
	"github.com/Aadil-Nabi/cmgarage/internal/config"
	"github.com/Aadil-Nabi/cmgarage/internal/pkg/cmhttpclient"
)

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
	jwt_details := jwtauth.GetAuthDetails()
	Bearer := jwt_details.Token_type + " " + jwt_details.Jwt

	configs := config.MustLoad()
	clusterInfoUrl := configs.Base_Url + configs.Version + "/cluster"

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
		fmt.Println("=>Primary cluster node is UP and Running")
		// slog.Info("Primary cluster node is UP and Running")
	}

	errMessage := ""

	// GetclusterErrors function to check if cluster nodes are up or down
	nodes := GetclusterErrors()
	for _, node := range *nodes {
		// log.Println(node.IsThisNode)
		fmt.Println("=>HA Status: ", node.IsThisNode)
		for _, err := range node.ClusterErrors {
			errMessage = err.ErrorMessage
		}
		if !node.IsThisNode {
			fmt.Println("=>HA is Broken, Seems other nodes in cluster are DOWN")
			// slog.Error("HA is Broken, Seems other nodes in cluster are DOWN")
			fmt.Println(errMessage)
		}
	}

}
