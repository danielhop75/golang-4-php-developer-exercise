package main

 
import (
	"log"	
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
 )
 
 type Cidb struct {
	 Id string `json:"id"`
	 Fname string `json:"fname"`
	 Sname string `json:"sname"`

 }
 

 func AppendtoCSV(str Cidb,sepstr string){

	//f,err:=os.OpenFile("test.csv",os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)

	f,err:=os.OpenFile("test.csv",os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	
	fmt.Println("=================")

	if err !=nil {
		log.Fatalf("Failed creating file: %s",err)
	}

	f.Write([]byte(str.Id));
	f.Write([]byte(sepstr));

	f.Write([]byte(str.Fname));
	f.Write([]byte(sepstr));
	
	f.Write([]byte(str.Sname));
	f.Write([]byte(sepstr));
	
	f.Write([]byte("\n"));
	
	f.Close()

 }



func main() {
	r := gin.Default()

	r.POST("/jsonconsume", func(c *gin.Context) {
		var ciElement Cidb

		err := c.BindJSON(&ciElement)
		if err != nil {
			log.Fatal(err)
        }        

        
		fmt.Println(ciElement)
		fmt.Println("=================")
		fmt.Println(ciElement.Id)
		fmt.Println(ciElement.Fname)
		fmt.Println(ciElement.Sname)
		AppendtoCSV(ciElement,"|")

	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
