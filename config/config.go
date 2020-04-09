package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	App struct {
		Name string `yaml:"name"`
		Port string `yaml:"port"`
	} `yaml:"app"`
	Mongo struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"mongo"`
	DatabaseNames   map[string]string `yaml:"databaseNames"`
	CollectionNames map[string]string `yaml:"collectionNames"`
}

var AppConfig *Config

func Init() {

	AppConfig = &Config{}

	yamlFile, err := ioutil.ReadFile("./app.yaml")
	if err != nil {
		log.Printf("yamlFile. Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, AppConfig)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	log.Println("Config")
	log.Println(AppConfig)
}
