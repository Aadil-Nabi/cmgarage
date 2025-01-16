package main

import (
	"fmt"

	"github.com/Aadil-Nabi/cmgarage/api/backups"
	"github.com/Aadil-Nabi/cmgarage/api/cluster"
	"github.com/Aadil-Nabi/cmgarage/api/encryption"
	"github.com/Aadil-Nabi/cmgarage/internal/validators"
)

func main() {
	fmt.Println("***********************************************************************************************")
	fmt.Println("**************************==========Welcome to the CMgarage==========**************************")
	fmt.Println("***********************************************************************************************")

	fmt.Println()

	// CM APIS
	validators.GetPasswordValidated()
	encryption.DiskEncryptionStatus()
	cluster.ClusterStatus()
	backups.GetBackupStatus()

}
