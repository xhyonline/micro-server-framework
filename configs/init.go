package configs

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/xhyonline/xutil/helper"
)

// Env 当前环境
var Env = "dev"

type Config struct {
	MySQL *MySQL `toml:"mysql"`
	Redis *Redis `toml:"redis"`
}

type Redis struct {
	dbCommon
	DB int `toml:"db"`
}

type MySQL struct {
	dbCommon
	DB string `toml:"db"`
}

type dbCommon struct {
	Host        string `toml:"host"`
	User        string `toml:"user"`
	Password    string `toml:"password"`
	Port        int    `toml:"port"`
	MaxConnNum  int    `toml:"max_conn_num"`
	IdleConnNum int    `toml:"idle_conn_num"`
}

var Instance = &Config{
	Redis: new(Redis),
	MySQL: new(MySQL),
}

type Option func() string

// filePath 默认配置文件地址
var filePath = "./configs/common/"

const (
	// 生产环境读取配置文件的地址
	productConfigPath = "/usr/local/go-micro/common/"
)

// Init 初始化配置文件信息
func Init(options ...Option) {
	// 判断生产环境的配置文件是否存在,如果存在优先读取
	exists, _ := helper.PathExists(productConfigPath)
	if exists {
		Env = "product"
		filePath = productConfigPath
	}
	for _, v := range options {
		load(v)
	}
}

// load 载入配置文件
func load(option Option) {
	if exists, _ := helper.PathExists(option()); exists {
		body, _ := ioutil.ReadFile(option())
		if _, err := toml.Decode(string(body), Instance); err != nil {
			fmt.Println("配置文件加载失败")
			os.Exit(0)
		}
	}
}

func WithMySQL() Option {
	return func() string {
		return filePath + "mysql.toml"
	}
}

func WithRedis() Option {
	return func() string {
		return filePath + "redis.toml"
	}
}
