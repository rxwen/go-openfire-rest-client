// Package main
package main

import (
	"fmt"

	openfirerest "github.com/rxwen/go-openfire-rest-client"
)

func main() {
	server := "http://127.0.0.1:9090"
	// calculate authorization echo -n "username:password" | base64
	authorization := "Basic Calculated code"
	items, _ := openfirerest.GetRoster(server, authorization, "user")
	fmt.Println(items)

	item := openfirerest.RosterItem{
		JID:              "contact",
		NickName:         "contact",
		SubscriptionType: "3",
		Groups: openfirerest.RosterGroup{
			Group: "group",
		},
	}
	fmt.Println(openfirerest.AddRoster(server, authorization, "user", item))
	item.SubscriptionType = "2"
	fmt.Println(openfirerest.UpdateRoster(server, authorization, "user", item))
}
