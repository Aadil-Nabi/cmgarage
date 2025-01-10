package api

import (
	"encoding/json"
	"io"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/Aadil-Nabi/cmgarage/internal/pkg/cmhttpclient"
)

type ClusterErrors struct {
	ErrorMessage string
	ErrorTime    time.Time
}

type Node struct {
	IsThisNode    bool
	NodeID        string
	ClusterErrors []ClusterErrors
}

var nodes []Node

func GetclusterErrors() {
	// Get the Bearer Token
	Bearer := jwt_details.Token_type + " " + jwt_details.Jwt

	clusterErrorUrl := "https://192.168.238.129/api/v1/cluster/errors"

	// Create a new request for cluster error API
	reqErrorInfo, err := http.NewRequest("GET", clusterErrorUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Add Headers to the request for cluster errors
	reqErrorInfo.Header.Add("Authorization", Bearer)

	// Get http client and call a Do function on it to send the http request for cluster info
	clientClusterError := cmhttpclient.GetClient()
	respClusterError, err := clientClusterError.Do(reqErrorInfo)
	if err != nil {
		log.Fatalf("error making request %v", err)
	}

	defer respClusterError.Body.Close()

	// Read data from response received from Do function called above for cluster errors
	data_err, err := io.ReadAll(respClusterError.Body)
	if err != nil {
		log.Fatal(err)
	}

	// var cmClusterStatus CMclusterErrStatus
	err = json.Unmarshal(data_err, &nodes)
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	errMessage := ""
	for _, node := range nodes {
		for _, err := range node.ClusterErrors {
			errMessage = err.ErrorMessage
		}

		if !node.IsThisNode {
			slog.Error("HA is Broken, Seems other nodes in cluster are DOWN")
			slog.Error(errMessage)
		}
	}

}
