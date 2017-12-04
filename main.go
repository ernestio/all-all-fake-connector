/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"

	ecc "github.com/ernestio/ernest-config-client"
	"github.com/nats-io/nats"
)

// Transition is used to map a transitions arc, along with any extra information
type Transition struct {
	From  string
	To    string
	Extra string `json:"extra,omitempty"`
}

func getTransitions(path string) []Transition {
	var keys []Transition

	file, err := os.Open(path)
	log.Printf("Reading config from: %s\n", path)
	if err != nil {
		log.Panic("error:", err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&keys)
	if err != nil {
		log.Println("Config file is invalid")
		log.Panic("error:", err)
	}

	fmt.Printf("%#v\n", keys)
	return keys
}

func manage(n *nats.Conn, m *nats.Msg, ts []Transition) {
	var current Transition
	for _, t := range ts {
		if t.From == m.Subject {
			current = t
		}
	}
	fmt.Printf("%#v", current)
	if &current != nil && current.To != "" && current.From != "" {

		var x map[string]interface{}
		err := json.Unmarshal(m.Data, &x)
		if err != nil {
			return
		}
		x["type"] = current.To

		if current.Extra != "" {
			var y map[string]interface{}
			err = json.Unmarshal([]byte(current.Extra), &y)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%#v", y)
			for k, v := range y {
				x[k] = v
			}
		}

		b, err := json.Marshal(x)
		if err != nil {
			fmt.Println("error:", err)
		}
		log.Println("In")
		err = n.Publish(current.To, b)
		if err != nil {
			log.Println(err)
		}
	}
}

func main() {
	trnPath := flag.String("transitions", "transitions/default.json", "path to json transitions")
	flag.Parse()

	ts := getTransitions(*trnPath)
	n := ecc.NewConfig(os.Getenv("NATS_URI")).Nats()
	h := handler{Nats: n}

	subscriptions := []string{"*", "*.*", "*.*.*"}
	for _, s := range subscriptions {
		_, err := n.Subscribe(s, func(m *nats.Msg) {
			manage(n, m, ts)
		})
		if err != nil {
			log.Fatal(err)
		}
	}

	_, err := n.Subscribe("federation.auth-fake", h.federationAuth)
	if err != nil {
		log.Fatal(err)
	}

	runtime.Goexit()
}
