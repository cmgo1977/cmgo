package comm

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Conf struct {
	Db struct {
		Mongodb struct {
			DbName      string `yaml:"mgo_Name"`
			Host      string `yaml:"mgo_Host"`
			Port      string `yaml:"mgo_Port"`
			User      string `yaml:"mgo_User"`
			Passwd    string `yaml:"mgo_Passwd"`
			TimeOut   int    `yaml:"mgo_TimeOut"`
			PoolLimit int    `yaml:"mgo_PoolLimit"`
			Direct    bool   `yaml:"mgo_Direct"`
		}
		PostgreSql struct {
			DbName    string `yaml:"pg_Name"`
			Host    string `yaml:"pg_Host"`
			Port    string `yaml:"pg_Port"`
			User    string `yaml:"pg_User"`
			Passwd  string `yaml:"pg_Passwd"`
			NetWork string `yaml:"pg_NetWork"`
			// TCP host:port or Unix socket depending on Network.
			Addr     string `yaml:"pg_Addr"`
			PoolSize int    `yaml:"pg_PoolSize"`
		}
		Redis struct {
			Host   string `yaml:"redis_Host"`
			Passwd string `yaml:"redis_Passwd"`
			Expire int `yaml:"redis_Expire"`
		}
	}
	Auth struct{
		SecretKey string `yaml:"jwt_SecretKey"`
		Expire int `yaml:"jwt_Expire"`
		ISS string `yaml:"jwt_ISS"`
	}
}

var C = &Conf{}

func InitConf() {
	var err error

	configFile, err := ioutil.ReadFile("config/conf")
	if err != nil {
		log.Fatalf("配置文件读取错误 %v ", err)
	}

	err = yaml.Unmarshal(configFile, C)
	if err != nil {
		log.Fatalf("反射配置文件内容错误 %v ", err)
	}
}

/*
func Getconf()string{
	//config, err := yaml.ReadFile("D:\\code\\cmgo\\src\\web\\config\\conf")
	config, err := yaml.ReadFile("config/conf")
	if err != nil {
		fmt.Println(err)
	}
	v,err := config.Get("mg_name")
	if err != nil {
		fmt.Println(err)
	}

	return v
	//config.Get("pg_host")
}
*/
