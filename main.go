package main

import (
	"log"

	"github.com/michaeljdennis/harvest-go/client"
	"github.com/michaeljdennis/harvest-go/config"
	"github.com/michaeljdennis/harvest-go/endpoint"
)

func main() {
	config, err := config.New("./.env")
	if err != nil {
		log.Fatalln(err)
	}
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
