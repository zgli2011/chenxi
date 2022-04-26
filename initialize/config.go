package initialize

import (
	"io/ioutil"
	"log"
	"net"
	"os"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	System struct {
		HTTPRequestTimeout int      `yaml:"http_request_timeout" json:"http_request_timeout"`
		CommandTimeout     int      `yaml:"command_timeout" json:"command_timeout"`
		PageSize           int      `yaml:"page_size" json:"page_size"`
		Version            string   `yaml:"version" json:"version"`
		SecurityKey        string   `yaml:"security_key" json:"security_key"`
		AccessIPWhiteList  []string `yaml:"access_ip_white_list" json:"access_ip_white_list"`
		APIRegister        string   `yaml:"api_register" json:"api_register"`
		CmdbGatewayAddress string   `yaml:"cmdb_gateway_address" json:"cmdb_gateway_address"`
		SystemTokenName    string   `yaml:"system_token_name" json:"system_token_name"`
		SystemToken        string   `yaml:"system_token" json:"system_token"`
		ProxyToken         string   `yaml:"proxy_token" json:"proxy_token"`
		LocalIP            string   `yaml:"local_ip" json:"local_ip"`
		LocalHostName      string   `yaml:"local_hostname" json:"local_hostname"`
		ServerPort         string   `yaml:"server_port" json:"server_port"`
	} `yaml:"system" json:"system"`
	Log           LogOptions   `yaml:"log" json:"log"`
	Database      MysqlOptions `yaml:"database" json:"database"`
	Redis         RedisOptoins `yaml:"redis" json:"redis"`
	Elasticsearch []string     `yaml:"elasticsearch" json:"elasticsearch"`
}

var Config *Configuration

func LoadConfiguration(config_path string) error {
	log.Println("开始加载配置文件......")
	// current_path, err := os.Getwd()
	// if err != nil {
	// 	log.Panicf("load config failed: %v", err)
	// }
	// config_path := path.Join("/Users/10013341/go_code/transaction_orchestration", "config.yaml")
	data, err := ioutil.ReadFile(config_path)
	if err != nil {
		log.Panicf("read config failed: %v", err)
		return err
	}

	// var config Configuration
	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		log.Panicf("config content is illegal: %v", err)
		return err
	}
	Config.System.LocalIP = CurrentIP()
	log.Println("加载配置文件成功......")
	return nil
}

func CurrentIP() string {
	if addrs, err := net.InterfaceAddrs(); err != nil {
		return "127.0.0.1"
	} else {
		for _, address := range addrs {
			if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() { // 检查ip地址判断是否回环地址
				if ipnet.IP.To4() != nil {
					return ipnet.IP.String()
				}
			}
		}
	}
	return "127.0.0.1"
}

func CurrentHostName() string {
	if name, err := os.Hostname(); err != nil {
		return name
	}
	return "localhost"
}
