package procon_data

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Msg struct {
	Jwt  string `json:"jwt"`
	Type string `json:"type"`
	Data string `json:"data"`
}

func SendMsg(jwt string, t string, data string, c *websocket.Conn) {
	m := Msg{jwt, t, data}
	if err := c.WriteJSON(m); err != nil {
		fmt.Println(err)
	}
}
