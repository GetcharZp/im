package main

import "im/router"

func main() {
	e := router.Router()
	e.Run(":8080")
}
