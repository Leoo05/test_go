package main

import (
	"fmt"

	factorymethod "github.com/Leoo05/test_go/internal/infrastructure/designPatterns/factoryMethod"
)

func main() {
	// var wg sync.WaitGroup
	factorymethod.NewEchoAPIRestAdapter().RunServer()

	fmt.Println("test started")
}
