package config

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func SetUpConfig() {
	var setting Setting
	conf := setting.getConf()
	data, err := json.Marshal(conf)
	if err != nil {
		fmt.Println("===error=={}==", err.Error())
		return
	}
	ConfSetting = setting
	fmt.Println("yaml ini success:==={}", string(data))
}

/*
*
读取配置文件赋值
*/
func (setting *Setting) getConf() *Setting {

	yamlFile, err := ioutil.ReadFile("config/app.yaml")
	if err != nil {
		fmt.Println("===error===={}===", err.Error())
	}
	err = yaml.Unmarshal(yamlFile, setting)

	if err != nil {
		fmt.Println("==file error=={}====", err.Error())
	}
	return setting
}

var ConfSetting Setting

type Setting struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
	Redis    Redis    `yaml:"redis"`
}

/*
*
服务器配置
*/
type Server struct {
	RunMode      string `yaml:"runMode"`
	HttpPort     string `yaml:"httpPort"`
	ReadTimeout  string `yaml:"readTimeout"`
	WriteTimeout string `yaml:"writeTimeout"`
}

/*
*
数据库配置
*/
type Database struct {
	Type        string `yaml:"type"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Host        string `yaml:"host"`
	Name        string `yaml:"name"`
	TablePrefix string `yaml:"tablePrefix"`
}

/*
*
redis配置
*/
type Redis struct {
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}
