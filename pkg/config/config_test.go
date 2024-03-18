package config

import (
	"testing"

	"github.com/spf13/viper"
)

func TestConfigInit(t *testing.T) {
	config := "./config.yaml"
	err := ConfigInit(&config)
	if err != nil {
		t.Fatal(err)
	}
	apikey := viper.GetString("wechat.apikey")
	t.Log(apikey)
}
