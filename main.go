package main

import (
	"fmt"

	"github.com/mrcampbell/stax-refund-service/config"
)

func main() {
	config.Load()
	fmt.Println(config.IsDev())
}
