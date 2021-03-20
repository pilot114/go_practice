package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kylelemons/go-gypsy/yaml"
	"gopkg.in/gcfg.v1"
)

type configuration struct {
	Enabled bool
	Path    string
}

// json / yaml / ini - основные форматы конфигурации
func main() {
	// JSON
	file, _ := os.Open("config.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := configuration{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(conf.Path)

	// YAML
	config, err := yaml.ReadFile("config.yaml")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(config.Get("path"))
	fmt.Println(config.GetBool("enabled"))

	// INI
	configIni := struct {
		Section struct {
			Enabled bool
			Path    string
		}
	}{}
	err = gcfg.ReadFileInto(&configIni, "config.ini")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(configIni.Section.Enabled)
	fmt.Println(configIni.Section.Path)

	// ENV
	fmt.Println(os.Getenv("PATH"))
}
