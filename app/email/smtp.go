package email

import (
	"fmt"
	"net/smtp"
	"strings"
)

func (s *Service) sendSMTP(to []string, subject string, htmlBody string) error {
	if !s.Enabled {
		s.logger.Warn("email send requested when sending is disabled")
		return nil
	}
	cfg := getCfg()

	var ret []string
	ret = append(ret, "From: "+cfg.From)
	ret = append(ret, "Subject: "+subject)
	ret = append(ret, "MIME-version: 1.0;")
	ret = append(ret, "Content-Type: text/html; charset=\"UTF-8\";")
	ret = append(ret, "")
	ret = append(ret, htmlBody)
	msg := strings.Join(ret, "\n")

	auth := smtp.PlainAuth("", cfg.Username, cfg.Password, cfg.Host)
	err := smtp.SendMail(fmt.Sprintf("%v:%v", cfg.Host, cfg.Port), auth, cfg.Username, to, []byte(msg))
	if err != nil {
		return err
	}
	return nil
}
