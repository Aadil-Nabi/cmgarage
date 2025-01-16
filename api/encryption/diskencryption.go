package encryption

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

type DiskEncryption struct {
	EncryptionStatus string
	HasDEK           bool
	AttendedBoot     bool
}

func DiskEncryptionStatus() {
	authDetails := jwtauth.GetAuthDetails()

	Bearer := authDetails.Token_type + " " + authDetails.Jwt

	configs := config.MustLoad()
	url := configs.Base_Url + configs.Version + "/locker/diskenc/status"

	// Create a new GET request.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Unable to create a New Request: ", err)
	}

	// Add Headers
	req.Header.Add("Authorization", Bearer)

	// Create a client
	client := cmhttpclient.GetClient()
	resp, err := client.Do(req)
	if err != nil {
		log.Println("unable to send a request to the CipherTrust Manager", err)
	}

	// Close
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("", err)
	}

	var diskEncryption DiskEncryption
	err = json.Unmarshal(data, &diskEncryption)
	if err != nil {
		log.Println(err)
	}

	encStatus := diskEncryption.EncryptionStatus
	if encStatus == "not encrypted" {
		// slog.Info("Disk is NOT Encrypted on the targeted node")
		fmt.Println("=> 🛑 Disk is NOT Encrypted on the targeted node.")

	}
}
