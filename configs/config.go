package configs

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var BasicPath = "./data"
var WordData []string

type configStruct struct {
	WebPort int
	WebHost string

	ConfigPath  string
	Key         string
	IdentityKey string

	SlaveId     int
	Retry       int
	ShowVersion bool
}

var ConfigGlobal = &configStruct{}

func LoadValue() ([]string, error) {
	if WordData != nil {
		return WordData, nil
	}

	byteValue, err := ioutil.ReadFile(fmt.Sprintf("%s/google-10000-english.txt", BasicPath))
	if err != nil {
		log.Printf("have error %s", err)
		return nil, err
	}

	WordData = strings.Split(string(byteValue), "\n")
	return WordData, nil
}
