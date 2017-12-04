package main

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats"
)

type handler struct {
	Nats *nats.Conn
}

func (h *handler) federationAuth(msg *nats.Msg) {
	var u user

	err := json.Unmarshal(msg.Data, &u)
	if err != nil {
		log.Println(err)
	}

	if u.Username == "john" && u.Password == "secret123" {
		err = h.Nats.Publish(msg.Reply, []byte(`{"ok": true}`))
		if err != nil {
			log.Println(err)
		}
		return
	}

	err = h.Nats.Publish(msg.Reply, []byte(`{"ok": false}`))
	if err != nil {
		log.Println(err)
	}
}

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
