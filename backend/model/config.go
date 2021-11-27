package model

type Config struct {
	Mysql   Mysql
	Redis   Redis
	Http    Http
	Proxy   Proxy
	CronTab CronTab
}

type Mysql struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type Redis struct {
	Address  string
	Password string
}

type Http struct {
	Port string
}

type Proxy struct {
	Address string
}

type CronTab struct {
	Schedule string
}
