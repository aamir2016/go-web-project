//Server to handle multiple requests at the same time

package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Team we are working on sample Go web server") //Fprintf used in case of web request

}

func main() {
	http.HandleFunc("/", handler) //Handles root url
	fmt.Println("Server starting in port 8090..........")
	http.ListenAndServe(":8090", nil)
}
