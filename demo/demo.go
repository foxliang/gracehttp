package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/foxliang/gracehttp"
)

func main() {
	//curl "http://127.0.0.1:8000?duration=2s"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		duration, err := time.ParseDuration(r.FormValue("duration"))
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		time.Sleep(duration)
		fmt.Fprintf( //打印在请求端信息
			w,
			"started at %s slept for %d nanoseconds from pid %d.\n",
			time.Now(),
			duration.Nanoseconds(),
			os.Getpid(),
		)
	})

	pid := os.Getpid() //获取进程id
	address := ":8080"

	log.Printf("process with pid %d serving %s.\n", pid, address)
	err := gracehttp.ListenAndServe(address, nil)
	log.Printf("process with pid %d stoped, error: %s.\n", pid, err)
}
