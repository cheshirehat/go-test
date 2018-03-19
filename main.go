package main

import "github.com/cheshirehat/go-test/router"

func main() {
	r := router.GetRouter()
	r.Run(":8000")
}
