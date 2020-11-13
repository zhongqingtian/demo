package yaml

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"
)

type Yaml struct {
	Mysql struct {
		User     string `yaml:"user"`
		Host     string `yaml:"host"`
		Password string `yaml:"password"`
		Port     string `yaml:"port"`
		Name     string `yaml:"name"`
	}
	Cache struct {
		Enable bool     `yaml:"enable"`
		List   []string `yaml:"list"`
	}
}

type Yaml1 struct {
	SQLConf   Mysql `yaml:"mysql"`
	CacheConf Cache `yaml:"cache"`
}

type Mysql struct {
	User     string `yaml:"user"`
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
}

type Cache struct {
	Enable bool     `yaml:"enable"`
	List   []string `yaml:"list"`
}

func MkYaml() {
	conf := new(Yaml)
	yamlFile, err := ioutil.ReadFile("./yaml/test.yaml")

	log.Println("yamlFile:", yamlFile)
	if err != nil {
		log.Printf("yamlFile.Get err #%v", err)
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	log.Println("conf", conf)
	time.Sleep(1 * time.Second)
}
