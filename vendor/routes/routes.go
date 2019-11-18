package routes

import(
	"net/http"
)
//Handle Home SETUPS the routes 
function HandleHome(){
	http.HandleFunc("/",controllers.HandleHome);
}