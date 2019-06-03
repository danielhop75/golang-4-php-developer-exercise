package main

 
import (
	"log"	
	"fmt"
	"github.com/gin-gonic/gin"
	
	
 )
 
 type Book struct {
	 Title string `json:"title"`
	 Author string `json:"autor"`
	 
 }
 

func main() {
	r := gin.Default()

	r.POST("/jsonconsume", func(c *gin.Context) {
		var stBook1 Book

		err := c.BindJSON(&stBook1)
		if err != nil {
			log.Fatal(err)
        }        

        
		fmt.Println(stBook1)
		fmt.Println("=================")
		fmt.Println(stBook1.Author)
		fmt.Println(stBook1.Title)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
