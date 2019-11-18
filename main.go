package main

import (
	"fmt"
	"net/http"
	"routes"
)

func main() {
	fmt.Println("Launching main")
	routes.HandleHome()
	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	}
}
