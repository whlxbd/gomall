package main

import (
	"log"
	order "whlxbd.github.com/gomall/kitex_gen/order/orderservice"
)

func main() {
	svr := order.NewServer(new(OrderServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
