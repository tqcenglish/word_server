package configs

import (
	"io/ioutil"
	"log"
	"strings"
)

var BasicPath = "/Users/tqcenglish/SrcCode/Flutter/word_english/wordPic-master"

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
	byteValue, err := ioutil.ReadFile("/Users/tqcenglish/SrcCode/Flutter/word_english/wordPic-master/google-10000-english.txt")
	if err != nil {
		log.Printf("have error %s", err)
		return nil, err
	}

	data := strings.Split(string(byteValue), "\n")
	return data, nil
}
