package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/shiro8613/HttpProxy/config"
	"github.com/shiro8613/HttpProxy/logger"
	"github.com/shiro8613/HttpProxy/proxy"
)

func main() {

	var c = flag.String("c" , "./config.yml", "-c /a/b/config.yml")
	flag.Parse()

	config, err := config.LoadAndCreate(*c)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Proxy server started!")
	fmt.Printf("Listen on %s\n", config.Listen)

	go func() {
        trap := make(chan os.Signal, 1)
        signal.Notify(trap, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT)
        s := <-trap
        fmt.Printf("Received shutdown signal %s\n", s)
		exit()
	}()

	go func () {
		for {
			var str string
			fmt.Scan(&str)
			if str == "stop" {
				exit()
			} else {
				fmt.Println("Commands not found")
			}
		}
	}()
	
	err = http.ListenAndServe(config.Listen, logger.Logger(proxy.Proxy(proxy.Build(config))))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func exit() {
	fmt.Printf("Shutdown gracefully....\n")
	os.Exit(0)
}