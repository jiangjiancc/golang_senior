package main

import (
	"log"
	"net/http"
	"runtime"
	"time"
	_ "net/http/pprof"
)

//读取当前内存信息的方法
func printMemStats()  {
	//定义一个runtime.MemStats对象
	var ms runtime.MemStats
	//通过对象的三个属性，查询内存的信息
	//1.将内存中的数据加载到ms对象中
	runtime.ReadMemStats(&ms)

	//2.将ms对象信息打印出来
	log.Printf("===> Alloc:%d(bytes),HeapIdle:%d(bytes),HeapReleased:%d(bytes)",ms.Alloc,ms.HeapIdle,ms.HeapReleased)
}
func test()  {

	//slice是一个动态扩容的，用slice来做堆内存的一个申请
	mySlice := make([]int,8)

	log.Println("----> loop begin...")
	for i := 0; i < 32*1000*1000; i++ {
		mySlice = append(mySlice,i)

		if i == 16 *1000*1000 {
			printMemStats()
		}
	}
	log.Println("---> loop end...")
}
func main()  {
	//启动pprof调试监听
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:10000",nil))
	}()

	log.Println("Start...")
	printMemStats()

	test()

	//强制调用GC回收
	log.Println("force GC...")
	runtime.GC()

	log.Println("Done...")
	printMemStats()

	//开辟一个协程，定期的打印当前的内存信息
	go func() {
		for  {
			printMemStats()
			time.Sleep(10 * time.Second)
		}
	}()

	//主线程 睡眠等待
	select {

	}
}