package main

import (
	"fmt"
	"clean_architecture_sample/api/externalInteface"
)

func main() {
	var addr =  ":8080"
	err := externalInteface.Router().Run(addr)
	if err != nil {
		panic(fmt.Errorf("[FAILED] start sever. err: %v", err))
	}
}
