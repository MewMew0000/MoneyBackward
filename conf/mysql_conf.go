package conf

type Mysql struct {
	Host          string `yaml:"host"`
	Port          string `yaml:"port"`
	Config        string `yaml:"config"`
	DefaultSchema string `yaml:"defaultSchema"`
	Username      string `yaml:"username"`
	Password      string `yaml:"password"`
	LogLevel      string `yaml:"logLevel"`
}

func (m Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.DefaultSchema + "?" + m.Config
}
