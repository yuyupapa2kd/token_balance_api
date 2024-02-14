package main

import (
	"fmt"

	"github.com/the-medium/token-balance-api/internal/api"
	"github.com/the-medium/token-balance-api/internal/config"
)

func main() {
	// profile, err := ioutil.ReadFile("./PROFILE")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("read file PROFILE")
	profile := "prd"
	config.SetRuntimeConfig(profile)
	fmt.Println("======================================================")
	fmt.Println("connect to network type of ", profile)
	fmt.Println("======================================================")

	r := api.SetRouter()
	go r.Run(":" + config.RuntimeConf.Server.Port)
	fmt.Println("server start")

	select {}
}
