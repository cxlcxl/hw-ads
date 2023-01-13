package main

import (
	_ "bs.mobgi.cc/bootstrap"
	"bs.mobgi.cc/router"
	"log"
)

func main() {
	if err := router.Router(); err != nil {
		log.Fatal(err)
	}
}
