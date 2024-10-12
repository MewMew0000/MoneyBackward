package conf

type SiteInfo struct {
	CreatedAt   string `yaml:"createdAt" json:"createdAt"`
	Title       string `yaml:"title" json:"title"`
	Version     string `yaml:"version" json:"version"`
	Name        string `yaml:"name" json:"name"`
	Addr        string `yaml:"addr" json:"addr"`
	Slogan      string `yaml:"slogan" json:"slogan"`
	Email       string `yaml:"email" json:"email"`
	LineImage   string `yaml:"lineImage" json:"lineImage"`
	WechatImage string `yaml:"wechatImage" json:"wechatImage"`
	BilibiliUrl string `yaml:"bilibiliUrl" json:"bilibiliUrl"`
	GithubUrl   string `yaml:"githubUrl" json:"githubUrl"`
}
