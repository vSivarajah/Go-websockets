package procon_config

import (
	"crypto/rsa"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/vsivarajah/Go-websockets/go_systems/procon_fs"
)

var (
	PubKeyFile  *rsa.PublicKey
	PrivKeyFile *rsa.PrivateKey
)

const (
	PKPWD             = "SOMEHARDPASSWORD"
	File_Storage_Path = "/Users/vigneshsivarajah/Go/src/github.com/vsivarajah/Go-websockets/uploads/"
	KeyCertPath       = "/Users/vigneshsivarajah/Go/src/github.com/vsivarajah/Go-websockets/keycertz/"
	PrivKeyPath       = "/Users/vigneshsivarajah/Go/src/github.com/vsivarajah/Go-websockets/keycertz/mykey.pem"
	PubKeyPath        = "/Users/vigneshsivarajah/Go/src/github.com/vsivarajah/Go-websockets/keycertz/mykey.pub"
)

func init() {
	f, ok, err := procon_fs.ReadFile(PubKeyPath)
	if !ok || err != nil {
		fmt.Println(err)
	} else {
		PubKeyFile, err = jwt.ParseRSAPublicKeyFromPEM(f)
		if err != nil {
			fmt.Println(err)
		}
	}
	f, ok, err = procon_fs.ReadFile(PrivKeyPath)
	if !ok || err != nil {
		fmt.Println(err)
	} else {
		PrivKeyFile, err = jwt.ParseRSAPrivateKeyFromPEMWithPassword(f, PKPWD)
		if err != nil {
			fmt.Println(err)
		}
	}
}
