package main

import (
	"fmt"

	"github.com/Aadil-Nabi/cmgarage/api/encryption"
	"github.com/Aadil-Nabi/cmgarage/internal/secrets"
)

func main() {
	fmt.Println("***********************************************************************************************")
	fmt.Println("**************************==========Welcome to the CMgarage==========**************************")
	fmt.Println("***********************************************************************************************")

	fmt.Println()

	// Call APIs of CM

	// cluster.ClusterStatus()
	encryption.Encrypting()
	// encryption.DiskEncryptionStatus()

	// backups.GetBackupStatus()

	secrets.Secrets()

	// time.Sleep(time.Millisecond * 500)

}
