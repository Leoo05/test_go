package main

import (
	factorymethod "github.com/Leoo05/test_go/internal/infrastructure/designPatterns/factoryMethod"
)

func main() {
	// var wg sync.WaitGroup
	// Here we create the echo api rest server
	factorymethod.NewEchoAPIRestAdapter().RunServer()
}
