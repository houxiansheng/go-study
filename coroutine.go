package main

import (
	"fmt"
)

func coroutineTest(){
	var a int;
	for a=0; a<10;a++  {
		global++;
		fmt.Println(global);
		//time.Sleep(time.Duration(1)*time.Second)
	}
}
