package main

import (
	"fmt"

	"github.com/Aadil-Nabi/cmgarage/api/backups"
	"github.com/Aadil-Nabi/cmgarage/api/cluster"
	"github.com/Aadil-Nabi/cmgarage/api/encryption"
)

func main() {
	fmt.Println("***********************************************************************************************")
	fmt.Println("**************************==========Welcome to the CMgarage==========**************************")
	fmt.Println("***********************************************************************************************")

	fmt.Println()

	// var wg sync.WaitGroup

	// for i := 1; i <= 2; i++ {
	// 	wg.Add(1)

	// 	go func() {
	// 		defer wg.Done()
	// 		diskEncryptionStatus()
	// 	}()

	// 	go func() {
	// 		defer wg.Done()
	// 		cmBackupStatus()
	// 	}()
	// }

	// wg.Wait()

	// CM APIS
	encryption.DiskEncryptionStatus()
	// encryption.Encrypting()
	cluster.ClusterStatus()
	backups.GetBackupStatus()

}

// func diskEncryptionStatus() {
// 	encryption.DiskEncryptionStatus()

// }

// func cmClusterStatus() {
// 	cluster.ClusterStatus()

// }

// func cmBackupStatus() {
// 	backups.GetBackupStatus()

// }
