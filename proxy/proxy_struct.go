package proxy

type Proxys map[string]ProxyWare

type ProxyWare struct {
	host string
	schem string
	path string
}