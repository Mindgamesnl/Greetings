package greetings

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"os"
)

type Language struct {
	Color string `json:"color"`
	Url string `json:"url"`
}

type Languages struct {
	Languages map[string]Language `json:"colors"`
}

var KnownLanguages Languages

func LoadLanguages()  {
	KnownLanguages = getAll()
}

func getAll() Languages {
	f, err := os.Open("resources/lang_colors.json")
	if err != nil {
		logrus.Error(err)
	}
	defer f.Close()

	var cfg Languages
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		logrus.Error(err)
	}

	return cfg
}