package config

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type Config struct {
	MensualPerc float32 `json: "mensual_perc"`
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

func IsFirstDayOfMonth() bool {
	day := time.Now().Day()
	if day == 1 || day == 01 {
		return true
	}
	return false
}
