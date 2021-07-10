package main

import (
	"flag"
	"fmt"
	"github/tqcenglish/word-english/app"
	"github/tqcenglish/word-english/configs"
	"github/tqcenglish/word-english/pkg/utils"
)

// release 编译信息
var (
	gitCommitCode string
	buildDateTime string
	goVersion     string
)

func init() {
	flag.StringVar(&configs.ConfigGlobal.IdentityKey, "jwtkey", "", "jwt 认证")

	flag.IntVar(&configs.ConfigGlobal.WebPort, "port", 8006, "web 端口")
	flag.StringVar(&configs.ConfigGlobal.WebHost, "host", "0.0.0.0", "web host")
	flag.BoolVar(&configs.ConfigGlobal.ShowVersion, "showVersion", false, "showVersion")
}

func main() {
	flag.Parse()
	if configs.ConfigGlobal.ShowVersion {
		fmt.Printf("gitCommitCode: %s, buildDateTime: %s %s", gitCommitCode, buildDateTime, goVersion)
		return
	}
	app.RunApp()
	utils.Exit()
}
