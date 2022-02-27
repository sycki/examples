package email

type Config struct {
	From     string `yaml:"from"`
	Password string `yaml:"password"`
	SmtpHost string `yaml:"smtpHost"`
	SmtpPort int    `yaml:"smtpPort"`
}
