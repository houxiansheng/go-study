package main

import (
	"fmt"
	"sync"
	"time"
)

func coroutineTest(){
	var mutex sync.Mutex;
	var a int;
	for a=0; a<10;a++  {
		mutex.Lock();
		global++;
		fmt.Println(a);
		mutex.Unlock();
		//time.Sleep(time.Duration(1)*time.Second)
	}
}
func coroutineLock(){
	var mutex sync.Mutex;
	var a int;
	for a=0; a<10;a++  {
		mutex.Lock();
		global++;
		fmt.Println("lock--",a);
		time.Sleep(time.Duration(1)*time.Second)
		mutex.Unlock();
	}
}
