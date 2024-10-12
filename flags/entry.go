package flags

import (
	"flag"
)

type Option struct {
	DB     bool   // db
	User   string // -u admin  -u user
	Load   string // load
	Dump   bool   // dump
	Es     bool   // es
	ESDump bool   // esdump
	ESLoad string // esload
}

// コマンドラインを解析する
func Parse() (option *Option) {
	option = new(Option)
	flag.StringVar(&option.User, "u", "", "ユーザーを作成する")
	flag.BoolVar(&option.DB, "db", false, "データベースを初期化する")
	flag.BoolVar(&option.Es, "es", false, "インデックスを作成する")
	flag.BoolVar(&option.Dump, "dump", false, "SQLデータベースをエクスポートする")
	flag.StringVar(&option.Load, "load", "", "SQLデータベースをインポートする")
	flag.BoolVar(&option.ESDump, "esdump", false, "ESインデックスをエクスポートする")
	flag.StringVar(&option.ESLoad, "esload", "", "ESインデックスをインポートする")
	flag.Parse()
	return option
}

func (option Option) Run() bool {
	if option.DB {
		DB()
		return true
	}
	if option.Load != "" {
		Load(option.Load)
		return true
	}
	if option.Es {
		ESIndex()
		return true
	}
	if option.Dump {
		Dump()
		return true
	}
	if option.ESDump {
		ESDump()
		return true
	}
	if option.ESLoad != "" {
		ESLoad(option.ESLoad)
		return true
	}
	return false
}
