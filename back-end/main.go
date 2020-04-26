package main

import (
	"github.com/chaihanij/hyperledger-fabric-application/back-end/server"
)

func main() {
	s := server.NewServer()

	if err := s.Init(3000); err != nil {
		panic(err)
	}

	s.Start()
}
