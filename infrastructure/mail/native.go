package mail

import (
	"fmt"
	"net/smtp"
)

type Config struct {
	Username string `koanf:"username"`
	Password string `koanf:"password"`
	Host     string `koanf:"host"`
	Port     string `koanf:"port"`
}

// TODO - add this to config
const (
	User     = "hosseinhalaj1379@gmail.com"
	Password = "dqukorbjgndhamte"
	SmtpHost = "smtp.gmail.com"
	SmtpPort = ":587"
)

type Mail struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
	cfg     Config
}

func NewMail(cfg Config) IMail {
	return Mail{cfg: cfg}
}

func (m Mail) SendingMail(mail *Mail) error {
	auth := smtp.PlainAuth("", m.cfg.Username, m.cfg.Password, fmt.Sprintf(":%s", m.cfg.Port))

	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\nThis is auto message from Pech-Pech\n\n%s", mail.To, mail.Subject, mail.Body))

	// I don't know this is right or wrong and what is the best way
	// to make send email on goroutine
	var err error
	go func() {
		addr := fmt.Sprintf("%s:%s", m.cfg.Host, m.cfg.Port)
		err = smtp.SendMail(addr, auth, m.cfg.Username, []string{mail.To}, msg)
		if err != nil {
			fmt.Println(err)
		}
	}()

	return err
}
