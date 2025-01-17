package main

import (
	"fmt"

	"github.com/Aadil-Nabi/cmgarage/internal/secrets"
)

func main() {
	fmt.Println("***********************************************************************************************")
	fmt.Println("**************************==========Welcome to the CMgarage==========**************************")
	fmt.Println("***********************************************************************************************")

	fmt.Println()

	// CM APIS
	// validators.GetPasswordValidated()
	// encryption.DiskEncryptionStatus()
	// cluster.ClusterStatus()
	// backups.GetBackupStatus()

	password := secrets.GetSecrets()
	fmt.Println(password["cm_pass"])

}
