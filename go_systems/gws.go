package main

import (
	// Native packages

	"flag"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/vsivarajah/Go-websockets/go_systems/procon_data"
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
			procon_data.SendMsg("^vAr^", "server-ws-connect-success-msg", "Hello...", c)
			break

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
