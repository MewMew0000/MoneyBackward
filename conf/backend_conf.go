package conf

type ListOfConf struct {
	Mysql    Mysql    `yaml:"mysql"`
	Logger   Logger   `yaml:"logger"`
	Server   Server   `yaml:"server"`
	SiteInfo SiteInfo `yaml:"siteInfo"`
}
