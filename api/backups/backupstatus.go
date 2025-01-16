package backups

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

type Resources struct {
	Scope          string
	Account        string
	Version        string
	Id             string
	CreatedAt      string
	Status         string
	TiedToHSM      bool
	Description    string
	BackupKey      string
	ProductVersion string
}

type Backups struct {
	Limit     int
	Total     int
	Resources []Resources
}

func GetBackupStatus() {

	jwt_details := jwtauth.GetAuthDetails()
	bearer := jwt_details.Token_type + " " + jwt_details.Jwt

	configs := config.MustLoad()

	url := configs.Base_Url + configs.Version + "/backups"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Unable to create a new GET request", err)
	}

	req.Header.Add("Authorization", bearer)

	client := cmhttpclient.GetClient()
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Unable to send request to the server", err)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Unable to read data from the provided response body", err)
	}

	defer resp.Body.Close()

	var backups Backups
	err = json.Unmarshal(data, &backups)
	if err != nil {
		log.Fatal("Unmarshal the response from CM for listing the Backups", err)
	}

	total_backups := backups.Total
	fmt.Println("=> Total backups on system are: ", total_backups)

	for k, v := range backups.Resources {
		fmt.Println("✔", k, v.Status)
	}

}
