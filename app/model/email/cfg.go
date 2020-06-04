package email

import (
	"os"
	"strconv"
)

type MailConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

func (c *MailConfig) Enabled() bool {
	return len(c.Username) > 0 && len(c.Host) > 0
}

func getCfg() MailConfig {
	port, _ := strconv.Atoi(env("port", "587"))
	if port == 0 {
		port = 587
	}
	un := env("username", "")
	return MailConfig{
		Host:     env("host", ""),
		Port:     port,
		Username: un,
		Password: env("password", ""),
		From:     env("from", un),
	}
}

func env(k string, def string) string {
	v := os.Getenv("rituals_mail_" + k)
	if len(v) == 0 {
		v = def
	}
	return v
}
