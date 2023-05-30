package config

type Config struct {
	Listen		string				`yaml:"listen"`
	Location	map[string]Location	`yaml:"location"`
}

type Location struct {
	Path		string	`yaml:"path"`
	ProxyPass	string	`yaml:"proxy_pass"`
}