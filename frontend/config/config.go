package config

type Config struct {
	Address    string
	AddService string
	SubService string
	MulService string
}

var CurrentConfig *Config = &Config{
	Address:    ":8090",
	AddService: "add",
	SubService: "sub",
	MulService: "mul",
}
