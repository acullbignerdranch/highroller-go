package main

import "highroller-go/src/controller"

func handleRequests() {
	controller.HandleRouters()
}

func main() {
	handleRequests()
}
