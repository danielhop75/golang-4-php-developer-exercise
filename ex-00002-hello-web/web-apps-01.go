package main
import "net/http"

func main () {
    http.ListenAndServe(":6868",http.FileServer(http.Dir(".")));
    
}