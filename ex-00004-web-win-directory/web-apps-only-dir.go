package main
import "net/http"
func main(){
    http.ListenAndServe(":6767",http.FileServer(http.Dir("c:/Users/michael.cane/")));

}
