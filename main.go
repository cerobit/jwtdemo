package  main

import (
	"fmt"
	"net/http"
)

func  srvlogin(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "This is a seeerver response")
}

func main(){

    http.HandleFunc("/login", srvlogin)
    http.ListenAndServe(":8080", nil)
}