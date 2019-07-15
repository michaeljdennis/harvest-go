package main

import (
	"log"

	"github.com/michaeljdennis/harvest-go"
	"github.com/michaeljdennis/harvest-go/endpoint"
)

func main() {
	var err error

	config, err := harvest.NewConfig("../.env")
	if err != nil {
		log.Fatalln(err)
	}
	client := harvest.NewClient(config)

	// usersEndpoint := endpoint.NewUsers()
	// users := &endpoint.UsersModel{}
	// err = client.Get(usersEndpoint, users)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// endpoint.PrintModel(users)

	meEndpoint := endpoint.NewMe()
	me := &endpoint.UserModel{}
	err = client.Get(meEndpoint, me)
	if err != nil {
		log.Fatalln(err)
	}
	endpoint.PrintModel(me)
}
