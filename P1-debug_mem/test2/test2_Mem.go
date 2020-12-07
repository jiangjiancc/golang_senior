package main

import (
	"log"
	"runtime"
	"time"
)

func test()  {
	//slice是一个动态扩容的，用slice来作堆内存的一个申请
	mySlice := make([]int,8)
	log.Println("--> loop begin...")
	for i := 0; i < 32*1000*1000; i++ {
		mySlice = append(mySlice,i)
	}
	log.Println("--> loop end...")
}
func main() {
	log.Println("Start ...")

	test()

	//强制调用GC回收
	log.Println("force GC...")
	runtime.GC()

	log.Println("Done..")

	time.Sleep(3600 * time.Second)
}
