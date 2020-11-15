package main

import (
	// Native packages

	"flag"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/vsivarajah/Go-websockets/go_systems/procon_config"
	"github.com/vsivarajah/Go-websockets/go_systems/procon_data"
	"github.com/vsivarajah/Go-websockets/go_systems/procon_jwt"
	//Our Packages
)

var addr = flag.String("addr", ":1200", "http service address")
var upgrader = websocket.Upgrader{} //use default options

func handleAPI(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Print("WTF @HandleAPI Ws Upgrade Error >", err)
		return
	}

	id, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
	}

	//Modified Mux websocket package conn strunct in conn.go
	c.Uuid = "ws-" + id.String()

Loop:
	for {
		in := procon_data.Msg{}
		err := c.ReadJSON(&in)
		if err != nil {
			c.Close()
			break Loop
		}
		switch in.Type {
		case "register-client-msg":
			procon_data.SendMsg("^vAr^", "server-ws-connect-success-msg", c.Uuid, c)
			jwt, err := procon_jwt.GenerateJWT(procon_config.PrivKeyFile, "fake-name", "fake-alias", "fake@gmail.com", "Admin")
			if err != nil {
				fmt.Println("JWT Generation Failed")
			} else {
				procon_data.SendMsg("^var^", "server-ws-connect-success-jwt", jwt, c)
			}
			break
		case "test-jwt-message":
			valid, err := procon_jwt.ValidateJWT(procon_config.PubKeyFile, in.Jwt)
			fmt.Println(in.Jwt)
			if err != nil {
				fmt.Println(err)
				procon_data.SendMsg("^var^", "jwt-token-invalid", err.Error(), c)
			} else if err == nil && valid {
				fmt.Println("Valid JWT")
			}
			//Redis Operations
		default:
			break
		}
	}
}

func main() {
	flag.Parse()

	//Look into subrouter stuffs

	r := mux.NewRouter()

	//Websocket API
	r.HandleFunc("/ws", handleAPI)

	// REST API
	fmt.Println("Server Running...")
	http.ListenAndServe(*addr, r)
}
