package controllers

import (
	"fmt"
	"net/http"
	"strings"
)

//HandleHome handles the home function
func HandleHome(w http.ResponseWriter, r *http.Request) {

	message := r.URL.Path
	fmt.Println(r.URL.Query())
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	fmt.Println(message)
	w.Write([]byte(message))

}
