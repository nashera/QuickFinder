package finderconfig

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/larspensjo/config"
)

// Config 设置选项
type Config struct {
	SampleSheet  string
	ReportFolder string
	OutputFolder string
	DBPath       string
}

// GetConfig Config Constructor
func GetConfig() *Config {
	path, err := os.Executable()
	if err != nil {
		log.Printf(err.Error())
	}
	configFile := filepath.Dir(path) + fmt.Sprintf("%c", os.PathSeparator) + "config.ini"
	// configFile := "D:/Project/QuickFinder/config.ini"
	//set config file std
	cfg, err := config.ReadDefault(configFile)
	if err != nil {
		// log.Fatalf("Fail to find", configFile, err.Error())
		log.Printf(err.Error())
	}
	//set config file std End

	//topic list
	var TOPIC = make(map[string]string)

	//Initialized topic from the configuration
	if cfg.HasSection("DBPath") {
		section, err := cfg.SectionOptions("DBPath")
		if err == nil {
			for _, v := range section {
				options, err := cfg.String("DBPath", v)
				if err == nil {
					TOPIC[v] = options
				}
			}
		}
	}
	//Initialized topic from the configuration END
	if TOPIC["SampleSheet"] == "Default" {
		TOPIC["SampleSheet"] = "Z:/Project/项目进度表/阅尔基因项目进度表2-20180509.xlsx"
	}

	if TOPIC["OutputFolder"] == "Default" {
		TOPIC["OutputFolder"] = filepath.Dir(path) + fmt.Sprintf("%c", os.PathSeparator) + "output"
	}

	if TOPIC["DBPath"] == "Default" {
		TOPIC["DBPath"] = filepath.Dir(path) + fmt.Sprintf("%c", os.PathSeparator) + "local_db.sqlite"
	}
	// fmt.Println(TOPIC)
	// fmt.Println(TOPIC["SampleSheet"])
	return &Config{
		SampleSheet:  TOPIC["SampleSheet"],
		ReportFolder: TOPIC["ReportFolder"],
		OutputFolder: TOPIC["OutputFolder"],
		DBPath:       TOPIC["DBPath"],
	}
}
