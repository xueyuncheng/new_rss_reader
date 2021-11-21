package model

type Config struct {
	Mysql Mysql
	Redis Redis
	Http  Http
	Proxy Proxy
}

type Mysql struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type Redis struct {
	Host     string
	Port     string
	Password string
}

type Http struct {
	Port string
}

type Proxy struct {
	Host string
	Port string
}
