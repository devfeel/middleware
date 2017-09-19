package etcd

import (
	"github.com/devfeel/dotweb"
	"io/ioutil"
	"os"
	"gopkg.in/yaml.v2"
	"fmt"
	"github.com/coreos/go-etcd/etcd"
	"strings"
)
const RootPath = "/dotweb"	//etcd的根路径

type Config struct {
	Machines []string `yaml:"machines,flow"`	//url of machines for etcd like http://127.0.0.1:6379
}

type  EtcdMiddleware struct{
	dotweb.BaseMiddlware
	Config *Config
	Client *etcd.Client
}


func (m *EtcdMiddleware) Handle(ctx dotweb.Context) error {
	fmt.Println("handle function of etcd middleware")
	return nil
}

/**
	return root node of etcd
 */
func (m *EtcdMiddleware) RootNode() (*etcd.Node, error) {
	response,err := m.Client.Get(RootPath,false,true)

	if err != nil {
		return nil, err
	}

	return response.Node, nil
}

/**
	add a new dir
 */
func (m *EtcdMiddleware) AddNewDirectory(path string) error {
	_,err := m.Client.CreateDir(RootPath + path,0)
	if err != nil{
		return err
	}
	return nil
}

/**
	add a new node
 */
func (m *EtcdMiddleware) AddNewNode(key string, value string, ttl uint64) error{
	_, err := m.Client.Create(RootPath + key, value, ttl)
	if err != nil{
		return err
	}
	return nil
}

func (m *EtcdMiddleware) RemoveNode(key string) error{
	_, err := m.Client.Delete(RootPath + key, false)
	if err != nil {
		return err
	}
	return nil
}

func (m *EtcdMiddleware) RemoveDir(key string) error{
	_, err := m.Client.Delete(RootPath + key, true)
	if err != nil {
		return err
	}
	return nil
}

func (m *EtcdMiddleware) UpdateDir(key string) error{
	_, err := m.Client.UpdateDir(RootPath + key, 0)
	if err != nil {
		return err
	}
	return nil

}

func (m *EtcdMiddleware) UpdateNode(key string) error{
	_, err := m.Client.UpdateDir(RootPath + key, 0)
	if err != nil {
		return err
	}
	return nil
}



// Middleware new create a Etcd Middleware
func Middleware(path string) *EtcdMiddleware {
	conf := InitConfig(path)		//default path is ./etcd_conf.yaml
	client := InitClientConnect(conf)
	InitDirtory(client)

	return &EtcdMiddleware{Config:conf,Client:client}
}

//初始化配置文件（yaml）
//默认
func InitConfig(configFilePath string) *Config {
	if configFilePath == "" {
		configFilePath = "./etcd_conf.yaml"
	}

	var config Config
	file, err := os.Open(configFilePath)
	if err != nil{
		panic("open config file of etcd middleware error\n" +
			"path: " + configFilePath + "\n" +
			err.Error())
		os.Exit(1)
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		panic("read config file of etcd middleware error\n" +
			"path: " + configFilePath + "\n" +
			err.Error())
	}
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		panic("parse config file of etcd middleware error\n" +
			"path: " + configFilePath + "\n" +
			err.Error())
	}
	return &config
}

//初始化客户端连接
func InitClientConnect(conf *Config)(client *etcd.Client){
	return etcd.NewClient(conf.Machines)
}

func InitDirtory(client *etcd.Client){
	if _,err := client.Get(RootPath,false,true);err != nil{
		if getEtcdErrorCode(err.Error()) == "100"{
			client.CreateDir(RootPath,0)
		}else {
			panic(err.Error())
		}
	}

}

/**
	返回etcd 客户端 错误码
 */
func getEtcdErrorCode(errmsg string) string{
	errs := strings.Split(errmsg,":")
	return errs[0]
}
