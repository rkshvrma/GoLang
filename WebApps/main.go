package main



import (
	"fmt"
	"net/http"	
)

func handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"Hello world")

}

func page1handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w ,"New page ")

}

func main(){
	http.HandleFunc("/", handler)
	http.HandleFunc("/page1", page1handler)
	http.ListenAndServe(":8081",nil)
}