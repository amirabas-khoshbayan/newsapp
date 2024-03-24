package main

import (
	"fmt"
	"newsapp/config"
)

func main() {
	//get config
	cfg := config.GetConfig()
	fmt.Println(cfg)

}
