package main

import (
	"log"
	cart "whlxbd.github.com/gomall/kitex_gen/cart/cartservice"
)

func main() {
	svr := cart.NewServer(new(CartServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
