package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Application struct {
		Get     string	`yaml:"get"`
		Apps    string	`yaml:"apps"`
		Create	string	`yaml:"create"`
		Update  string	`yaml:"update"`
		Delete  string	`yaml:"delete"`
	} `yaml:"application"`
}

func (c *Config) InitConfig() *Config {
	yamlFile, err := ioutil.ReadFile("./cfg/config.yaml")
	if err != nil {
		log.Fatal(err)
		return nil
	}

	if err := yaml.Unmarshal(yamlFile, c); err != nil {
		log.Fatal(err)
		return nil
	}

	return c
}

type Conf struct {
	Commond struct {
		Application struct {
			Get struct {
				Url    string `yaml:"url"`
				Method string `yaml:"method"`
				State  string `yaml:"state"`
			} `yaml:get`
			Apps struct {
				Url    string `yaml:"url"`
				Method string `yaml:"method"`
				State  string `yaml:"state"`
			} `yaml:"apps"`
			Create struct {
				Url    string `yaml:"url"`
				Method string `yaml:"method"`
				State  string `yaml:"state"`
			} `yaml:"create"`
			Update struct {
				Url    string `yaml:"url"`
				Method string `yaml:"method"`
				State  string `yaml:"state"`
			} `yaml:"update"`
			Delete struct {
				Url    string `yaml:"url"`
				Method string `yaml:"method"`
				State  string `yaml:"state"`
			} `yaml:"delete"`
		} `yaml:"application"`
		Config struct {
			
		} `yaml:"config"`
		User struct {

		} `yaml:"user"`
		SDK struct {
			
		} `yaml:"sdk"`
		Permission struct {
			
		} `yaml:"permission"`
	} `yaml:"commond"`
}

func (c *Conf) InitConf() *Conf {
	yamlFile, err := ioutil.ReadFile("./cfg/conf.yaml")
	if err != nil {
		log.Fatal(err)
		return nil
	}

	if err := yaml.Unmarshal(yamlFile, c); err != nil {
		log.Fatal(err)
		return nil
	}

	return c
}

func main() {
	// var conf Config
	// conf.InitConfig()
	// fmt.Println(conf)

	var conf Conf
	conf.InitConf()
	fmt.Println(conf)
}
