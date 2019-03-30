package main
import (
   "encoding/json"
   "net/http"
)

type Book struct {
    Title string `json:"title"`
    Author string `json:"autor"`
    
}


func main (){
	http.HandleFunc("/",ShowBooks)
	http.ListenAndServe(":8080",nil)
}



func ShowBooks(w http.ResponseWriter, r *http.Request){
	
	book:= Book{"Mickiewicz","Iliada"}

	js, err :=json.Marshal(book)
	

	if err!=nil {
		http.Error(w, err.Error(),http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type","application/json")
    w.Write(js)

}


