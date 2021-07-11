package app

import (
	"github/tqcenglish/word-english/app/http_server"
	"github/tqcenglish/word-english/app/model"
	"github/tqcenglish/word-english/configs"
)

func RunApp() {
	configs.LoadValue()
	model.CreateDBInstance()
	go http_server.CreateServer(configs.ConfigGlobal.WebPort)

}
