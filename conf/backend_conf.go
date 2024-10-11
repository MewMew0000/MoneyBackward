package conf

type List struct {
	Mysql  Mysql  `yaml:"mysql"`
	Logger Logger `yaml:"logger"`
	Server Server `yaml:"server"`
}
