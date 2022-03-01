package config

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type Config struct {
	MensualPerc      float32 `json: "MensualPerc"`
	UpdatedThisMonth bool    `json:	"UpdatedThisMonth"`
}

func WriteConfig(c Config) {
	file, _ := json.MarshalIndent(c, "", " ")
	_ = ioutil.WriteFile("config.json", file, 0644)
}

func ReadConfig() (Config, error) {
	conf := Config{}
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		return conf, err
	}
	_ = json.Unmarshal([]byte(file), &conf)
	return conf, nil
}

func IsDayOfMonth(d int) bool {
	day := time.Now().Day()
	return day == d
}

func IsOneOfThisDays(days []int) (bool, int) {
	day := time.Now().Day()
	for i, d := range days {
		return d == day, days[i]
	}
	return false, -1
}
