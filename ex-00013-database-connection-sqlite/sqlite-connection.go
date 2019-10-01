package main

import (
	"database/sql"
	"fmt"
	"strconv"
	_ "github.com/mattn/go-sqlite3"
)



func main() {

database,_:=sql.Open("sqlite3","./bazka.db")
statement,_:=database.Prepare("CREATE TABLE IF NOT EXISTS MYDATA (id integer primary key, value1 text, value2 text)")
statement.Exec()

statement,_=database.Prepare("Insert into MYDATA (value1,value2) values (?,?)")
statement.Exec("Pole1 wartosc","Pole2 wartosc")

rows,_:=database.Query("select id, value1, value2 from MYDATA")

var id int
var field1 string
var field2 string


for rows.Next(){

	 rows.Next()
	 
	 rows.Scan(&id,&field1,&field2)
	 fmt.Println(strconv.Itoa(id) + ":" + field1 + ":" + field2)
	
	}


}

