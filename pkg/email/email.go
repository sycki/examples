package email

import (
	"fmt"
	"net/smtp"
	"strings"

	"github.com/pkg/errors"
)

type Email struct {
	config *Config
}

func NewEmail(config *Config) *Email {
	return &Email{
		config: config,
	}
}

func (p *Email) SendTo(to []string, title, message string) error {
	bodyTemplate := `Subject: %s
To: %s
MIME-version: 1.0;
Content-Type: text/html; charset="UTF-8";



%s
`

	body := fmt.Sprintf(bodyTemplate, title, strings.Join(to, ";"), message)
	auth := NewPlainAuth("", p.config.From, p.config.Password, p.config.SmtpHost)
	err := smtp.SendMail(fmt.Sprintf("%s:%d", p.config.SmtpHost, p.config.SmtpPort), auth, p.config.From, to, []byte(body))
	if err != nil {
		return errors.Wrap(err, "smtp.SendMail")
	}

	return nil
}

type plainAuth struct {
	identity, username, password string
	host                         string
}

func NewPlainAuth(identity, username, password, host string) smtp.Auth {
	return &plainAuth{identity, username, password, host}
}

func isLocalhost(name string) bool {
	return name == "localhost" || name == "127.0.0.1" || name == "::1"
}

func (a *plainAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	if server.Name != a.host {
		return "", nil, errors.New("wrong host name")
	}
	resp := []byte(a.identity + "\x00" + a.username + "\x00" + a.password)
	return "PLAIN", resp, nil
}

func (a *plainAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		// We've already sent everything.
		return nil, errors.New("unexpected server challenge")
	}
	return nil, nil
}
