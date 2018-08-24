package main

import (
	_ "github.com/go-sql-driver/mysql"
	"cache"
	"fmt"
	"time"
)
var global int=0;
func main() {

	timestamp := time.Now().Unix()

	fmt.Println(timestamp)

	cache.Instance();
	var str=cache.Set("a","xcvbnm,.cvbnm,./",1);
	fmt.Println(str)
	time.Sleep(time.Duration(3)*time.Second)
	fmt.Println(cache.Get("a"));
	go cache.Server();
	time.Sleep(time.Duration(3)*time.Second)
	go cache.Client();
	go cache.Client();
	time.Sleep(time.Duration(15)*time.Second)
	//var a = 1.5
	//var b = 2
	//fmt.Println(a,b)
	//test.TestArr();

	//db,err := sql.Open("mysql","root:root@192.168.186.129/sql-collector")
	//rows, err := db.Query("select * from sql_query")
	//if err != nil {
	//}
	//for rows.Next(){
	//
	//}
	//getMysql();
	//go coroutineLock();
	//go coroutineLock();
	//go coroutineLock();
	//go coroutineLock();
	//go coroutineLock();
	//go coroutineLock();
	//go coroutineTest();
	//go coroutineTest();
	//go coroutineTest();
	//go coroutineTest();
	//go coroutineTest();
	//go coroutineTest();
	////fmt.Println(a,a)
	//time.Sleep(time.Duration(15)*time.Second)
}