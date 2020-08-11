package conf

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"saas-demo/common/conf"
	"saas-demo/pkg/util"
	"strings"
)

func init() {
	if !cfg.parseConfig() {
		showHelp()
		os.Exit(-1)
	}
}

const (
	RunEnv = "SAAS_ENV"
)

var (
	cfg         = &config{}
	Nacos       = &cfg.Nacos
	HTTPServer  = &cfg.HTTPServer
	Discovery   = &cfg.Discovery
	Application = &cfg.Application
	IP          = util.LocalIP()
)

type config struct {
	Nacos       conf.NacosConfig       `yaml:"nacos"`
	HTTPServer  conf.HTTPServerConfig  `yaml:"http"`
	Discovery   conf.DiscoveryConfig   `yaml:"discovery"`
	Application conf.ApplicationConfig `yaml:"application"`
	CfgFile     string
}

func showHelp() {
	fmt.Printf("Usage:%s {params}\n", os.Args[0])
	fmt.Println("      -c {config file}")
	fmt.Println("      -h (show help info)")
}

func (c *config) parseConfig() bool {
	flag.StringVar(&c.CfgFile, "c", "conf/conf.yaml", "config file")
	env := os.Getenv(RunEnv)
	if env != "" {
		c.CfgFile = strings.Replace(c.CfgFile, ".yaml", env+".yaml", 1)
	}
	flag.Parse()
	if !c.load() {
		return false
	}
	return true
}

func (c *config) load() bool {
	_, err := os.Stat(c.CfgFile)
	if err != nil {
		return false
	}
	bytes, err := ioutil.ReadFile(c.CfgFile)
	if err != nil {
		_ = fmt.Errorf("load configfile[%s] err:%v\n", c.CfgFile, err)
		return false
	}
	fmt.Printf("confs values[\n%s\n]\n", string(bytes))
	if err = yaml.Unmarshal(bytes, c); err != nil {
		_ = fmt.Errorf("parse configfile[%s] err:%v\n", c.CfgFile, err)
		return false
	}
	return true
}
