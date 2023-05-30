package main

import (
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
	config, err := config.LoadAndCreate("./config.yml")
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
        fmt.Printf("Shutdown gracefully....\n")
        os.Exit(0)
    }()
	
	err = http.ListenAndServe(config.Listen, logger.Logger(proxy.Proxy(proxy.Build(config))))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
