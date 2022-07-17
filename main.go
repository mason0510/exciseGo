package main

import (
	"exciseGo/router"
)

//define gin function
func main() {
	r:=router.Router()
	r.Run()
}