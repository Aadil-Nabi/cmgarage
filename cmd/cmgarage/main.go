package main

import (
	"fmt"

	"github.com/Aadil-Nabi/cmgarage/api/cluster"
	"github.com/Aadil-Nabi/cmgarage/api/encryption"
)

func main() {
	fmt.Println("***********************************************************************************************")
	fmt.Println("**************************==========Welcome to the CMgarage==========**************************")
	fmt.Println("***********************************************************************************************")

	fmt.Println()

	// Call APIs of CM

	cluster.ClusterStatus()
	encryption.Encrypting()
	encryption.DiskEncryptionStatus()

	// time.Sleep(time.Millisecond * 500)

}
