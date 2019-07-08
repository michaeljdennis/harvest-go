package main

import (
	"github.com/michaeljdennis/harvest-go/client"
	"github.com/michaeljdennis/harvest-go/config"
	"github.com/michaeljdennis/harvest-go/endpoint"
)

func main() {
	config := config.New("./.env")
	client := client.New(config)

	// usersEndpoint := endpoint.NewUsers()
	// users := &endpoint.UsersModel{}
	// client.Get(usersEndpoint, users)
	// endpoint.PrintModel(users)

	meEndpoint := endpoint.NewMe()
	me := &endpoint.UserModel{}
	client.Get(meEndpoint, me)
	endpoint.PrintModel(me)
}
