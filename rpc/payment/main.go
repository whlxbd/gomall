package main

import (
	"log"
	payment "whlxbd.github.com/gomall/kitex_gen/payment/paymentservice"
)

func main() {
	svr := payment.NewServer(new(PaymentServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
