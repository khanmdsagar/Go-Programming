package main

import "fmt"
import "net/http"
//import "log"

func main(){
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)

	fmt.Println("Server running at - http://localhost:8080")
	http.ListenAndServe(":8080", nil)
	
}

func home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Home Page")
}

func about(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "About Page")
}