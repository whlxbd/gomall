package main

import (
	"log"
	auth "whlxbd.github.com/gomall/kitex_gen/auth/authservice"
)

func main() {
	svr := auth.NewServer(new(AuthServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
