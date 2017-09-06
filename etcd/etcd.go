package etcd

import (
	//etcd "github.com/coreos/go-etcd/etcd"
	"github.com/devfeel/dotweb"
	"io/ioutil"
	"os"
	"gopkg.in/yaml.v2"
	"path/filepath"
	"fmt"
	"github.com/coreos/go-etcd/etcd"
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
func init(){

}

func (m *EtcdMiddleware) Handle(ctx dotweb.Context) error {
	fmt.Println("handle function of etcd middleware")
	return nil
}

/**
	list root path dirs of etcd
 */
func (m *EtcdMiddleware) List() (etcd.Nodes, error) {
	response,err := m.Client.Get(RootPath,false,true)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%+v \n",*response.Node)
	return response.Node.Nodes, nil
}

func (m *EtcdMiddleware) AddNewDirectory(path string) error {
	_,err := m.Client.CreateDir(RootPath + path,0)
	if err != nil{
		return err
	}
	return nil
}

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

// Middleware new create a AccessLog Middleware
func Middleware(path string) *EtcdMiddleware {
	conf := InitConfig(path)		//default path is ./etcd_conf.yaml
	client := InitClientConnect(conf)
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
		os.Exit(1)
	}
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		panic("parse config file of etcd middleware error\n" +
			"path: " + configFilePath + "\n" +
			err.Error())
		os.Exit(1)
	}
	return &config
}

//初始化客户端连接
func InitClientConnect(conf *Config)(client *etcd.Client){
	return etcd.NewClient(conf.Machines)
}

func InitDirtory(client *etcd.Client){
	if _,err := client.Get(RootPath,false,true);err != nil{

	}

}


func lookupFile(configFile string) (realFile string, exists bool) {
	//add default file lookup
	//1、按绝对路径检查
	//2、尝试在当前进程根目录下寻找
	//3、尝试在当前进程根目录/config/ 下寻找
	//fixed for (#3 当使用json配置的时候，运行会抛出panic)
	realFile = configFile
	exists = true
	if !fileExists(realFile) {
		realFile = getCurrentDirectory() + "/" + configFile
		exists = false
	}
	if !exists && !fileExists(realFile) {
		realFile = getCurrentDirectory() + "/config/" + configFile
	} else {
		exists = true
	}
	if !exists && fileExists(realFile) {
		exists = true
	}
	return realFile, exists
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func getCurrentDirectory() string {
	return filepath.Clean(filepath.Dir(os.Args[0])) + "/"
}