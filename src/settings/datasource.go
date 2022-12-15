package settings

import (
	"abix360/src/infraestructure/shared"
	"log"
	"os"
	"sync"

	"gopkg.in/yaml.v2"
)

var lock = &sync.Mutex{}

type configFile struct {
	Database struct {
		Driver string `yaml:"driver"`
		Port   int    `yaml:"port"`
		User   string `yaml:"user"`
		Pass   string `yaml:"pass"`
		Name   string `yaml:"name"`
	}
}

type settings struct {
	config configFile
}

func getConfig() *configFile {
	content, err := os.ReadFile(shared.GetRootPath() + "/config.yml")
	if err != nil {
		log.Fatal("dao / datasource / Config / os.ReadFile: ", err)
	}

	var configFile configFile
	err = yaml.Unmarshal(content, &configFile)
	if err != nil {
		log.Fatal("dao / datasource / Config / yaml.Unmarshal: ", err)
	}
	return &configFile
}

var instanceSetting *settings

func DataSource() *settings {
	if instanceSetting == nil {
		lock.Lock()
		defer lock.Unlock()
		if instanceSetting == nil {
			instanceSetting = &settings{
				config: *getConfig(),
			}
		}
	}
	return instanceSetting
}

func (s *settings) DriverDB() string {
	return s.config.Database.Driver
}

func (s *settings) PortDB() int {
	return s.config.Database.Port
}

func (s *settings) UserDB() string {
	return s.config.Database.User
}

func (s *settings) PassDB() string {
	return s.config.Database.Pass
}

func (s *settings) NameDB() string {
	return s.config.Database.Name
}
