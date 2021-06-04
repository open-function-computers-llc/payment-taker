package main

import (
	"github.com/open-function-computers-llc/payment-taker/server"
)

func main() {
	s, err := server.Create()
	if err != nil {
		panic(err.Error())
	}
	s.Serve()
}
