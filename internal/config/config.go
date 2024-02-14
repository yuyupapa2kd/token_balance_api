package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var RuntimeConf = RuntimeConfig{}

type RuntimeConfig struct {
	RpcEndpoint string  `yaml:"rpcEndpoint"`
	TokenCA     TokenCA `yaml:"tokenCA"`
	Server      Server  `yaml:"server"`
}

type TokenCA struct {
	SOP    string `yaml:"SOP"`
	LOUI   string `yaml:"LOUI"`
	KsETH  string `yaml:"ksETH"`
	KsUSDT string `yaml:"ksUSDT"`
	KsXRP  string `yaml:"ksXRP"`
	KsBNB  string `yaml:"ksBNB"`
	KsKLAY string `yaml:"ksKLAY"`
	InKSTA string `yaml:"inKSTA"`
	DLT    string `yaml:"DLT"`
	XABT   string `yaml:"XABT"`
	BOM    string `yaml:"BOM"`
}

type Server struct {
	Ip   string `yaml:"ip"`
	Port string `yaml:"port"`
}

func SetRuntimeConfig(profile string) {
	viper.AddConfigPath("./internal/config")
	viper.SetConfigName(profile)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&RuntimeConf)
	if err != nil {
		panic(err)
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed: ", e.Name)
		var err error
		err = viper.ReadInConfig()
		if err != nil {
			fmt.Println(err)
			return
		}
		err = viper.Unmarshal(&RuntimeConf)
		if err != nil {
			fmt.Println(err)
			return
		}
	})

	viper.WatchConfig()
}
