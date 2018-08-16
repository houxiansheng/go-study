package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)
var global int=0;
func main() {
	fmt.Println("Hello, World!")
	var a = 1.5
	var b = 2
	fmt.Println(a,b)

	//db,err := sql.Open("mysql","root:root@192.168.186.129/sql-collector")
	//rows, err := db.Query("select * from sql_query")
	//if err != nil {
	//}
	//for rows.Next(){
	//
	//}
	getMysql();
	go coroutineTest();
	go coroutineTest();
	go coroutineTest();
	go coroutineTest();
	go coroutineTest();
	go coroutineTest();
	go coroutineTest();
	go coroutineTest();
	go coroutineTest();
	fmt.Println(a,a)
	time.Sleep(time.Duration(15)*time.Second)
}