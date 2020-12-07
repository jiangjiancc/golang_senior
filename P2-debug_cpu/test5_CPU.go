package main

import (
	"bytes"
	"log"
	"math/rand"
	"net/http"
	"time"
	_ "net/http/pprof"
)

//循环打印字符串

func test()  {
	log.Println(" ----> loop begin...")

	for i := 0; i < 1000; i++ {
		log.Println(genSomeBytes())
	}
	log.Println("----> loop end...")
}
//随机生成一个字符串
func genSomeBytes()*bytes.Buffer  {
	var buff bytes.Buffer

	for i := 0; i < 20000; i++ {
		buff.Write([]byte{'0'+ byte(rand.Intn(10))})
	}
	return &buff
}
func main() {
	//调用test()打印随机字符串
	go func() {
		for  {
			test()
			time.Sleep(time.Second * 1)
		}
	}()

	//main goroutine中开启pprof的监听的服务
	http.ListenAndServe("0.0.0.0:10000",nil)
}
