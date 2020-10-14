package main

import (
	"github.com/open-function-computers-llc/payment-taker/server"
)

func main() {
	s := server.Create()
	s.Serve()
}
