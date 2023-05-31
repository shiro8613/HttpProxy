package proxy

import (
	"fmt"
	"strings"

	"github.com/shiro8613/HttpProxy/config"
)

func Build(config config.Config) Proxys {
	handlers := make(Proxys)

	for k, v := range config.Location {
		var schem string
		var host string
		if strings.Contains(v.ProxyPass, "http://") {
			schem = "http"
			host = strings.Replace(v.ProxyPass, "http://", "", -1)
		} else if strings.Contains(v.ProxyPass, "https://") {
			schem = "https"
			host = strings.Replace(v.ProxyPass, "https://", "", -1)
		} else {
			continue
		}

		handlers[v.Path] = ProxyWare{
			host: host,
			schem: schem,
			path: v.Path,
		}

		fmt.Printf("[Proxy:%s] http://%s%s -> %s\n",k ,config.Listen, v.Path, v.ProxyPass)
	}

	return handlers
}
