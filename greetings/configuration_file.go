package greetings

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"os"
)

var LoadedInstance Config

type Config struct {
	GithubToken   string   `json:"github_token"`
	RepoCount     int      `json:"repo_count"`
	OnlyPublic    bool     `json:"hide_private"`
	HideEmptyDesc bool     `json:"hide_empty_descriptions"`
	Socials       []Social `json:"socials"`
}

type Social struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func LoadConfiguration() Config {
	f, err := os.Open("config.json")
	if err != nil {
		logrus.Error(err)
	}
	defer f.Close()

	var cfg Config
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		logrus.Error(err)
	}

	LoadedInstance = cfg

	return cfg
}
