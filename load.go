package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type configFile struct {
	Version              string   `yaml:"version"`
	Local_subscriptions  []string `yaml:"local_subscriptions"`
	Remote_subscriptions []string `yaml:"remote_subscriptions"`
}

func loadSubscriptions(path string) ([]string, error) {
	conf, err := loadConfig(path)
	if err != nil {
		return nil, err
	}
	fmt.Println("版本：", conf.Version)
	fmt.Println("本地配置文件：", conf.Local_subscriptions)
	fmt.Println("远程配置文件：", conf.Remote_subscriptions)
	var subscriptions []string
	for _, s := range conf.Local_subscriptions {
		sub, err := readFile(s)
		if err != nil {
			return nil, err
		}
		subscriptions = append(subscriptions, sub...)
	}
	for _, s := range conf.Remote_subscriptions {
		sub, err := readNetworkFile(s)
		if err != nil {
			return nil, err
		}
		subscriptions = append(subscriptions, sub...)
	}
	return subscriptions, nil
}

func loadConfig(path string) (configFile, error) {
	file, err := os.ReadFile(path)
	var conf configFile
	if err != nil {
		return conf, err
	}
	err = yaml.Unmarshal(file, &conf)
	if err != nil {
		return conf, err
	}
	return conf, nil
}
