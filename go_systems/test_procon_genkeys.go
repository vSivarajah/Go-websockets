package main

import (
	"fmt"

	"github.com/vsivarajah/Go-websockets/go_systems/procon_config"
	"github.com/vsivarajah/Go-websockets/go_systems/procon_fs"
	"github.com/vsivarajah/Go-websockets/go_systems/procon_genkeys"
)

func main() {

	// openssl rsa -in mykey.pem -pubout > mykey.pub
	pk, err := procon_genkeys.PrivateKeyToEncrpytedPEM(1028, "SOMEHARDPASSWORD")
	if err != nil {
		fmt.Println(err)
	}

	f, ok, err := procon_fs.CreateFile(procon_config.KeyCertPath, "mykey.pem")
	if !ok {
		fmt.Println(err)
	} else {
		procon_fs.WriteFile(f, pk)
	}
}
