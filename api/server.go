package main

import (
	"clean_architecture_sample/api/externalInteface"
	"fmt"
)

func main() {
	addr := ":8080"
	err := externalInteface.Router().Run(addr)

	if err != nil {
		panic(fmt.Errorf("[FAILED] start sever. err: %v", err))
	}
}
