package main

import (
	"go-smaole/router"
)

func main() {
	r := router.NewRouter()
	r.Run(":80")
}
