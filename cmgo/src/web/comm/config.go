package comm

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

//测试


type Conf struct {
	Db struct {
		Mongodb struct {
			M_name      string `yaml:"mg_name"`
			M_host      string `yaml:"mg_host"`
			M_port      string `yaml:"mg_port"`
			M_user      string `yaml:"mg_user"`
			M_passwd    string `yaml:"mg_passwd"`
			M_timeout   int    `yaml:"mg_timeout"`
			M_poollimit int    `yaml:"mg_poollimit"`
			M_direct    bool   `yaml:"mg_direct"`
		}
		Postresql struct {
			P_name    string `yaml:"pg_name"`
			P_host    string `yaml:"pg_host"`
			P_port    string `yaml:"pg_port"`
			P_user    string `yaml:"pg_user"`
			P_passwd  string `yaml:"pg_passwd"`
			P_network string `yaml:"pg_network"`
			// TCP host:port or Unix socket depending on Network.
			P_addr     string `yaml:"pg_addr"`
			P_poolsize int    `yaml:"pg_poolsize"`
		}
		Redis struct {
			R_host   string `yaml:"rd_host"`
			R_passwd string `yaml:"rd_passwd"`
			R_expire int `yaml:"rd_expire"`
		}
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
