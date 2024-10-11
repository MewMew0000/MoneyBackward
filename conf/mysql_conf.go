package conf

type Mysql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Conn     string `yaml:"conn"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"logLevel"`
}
