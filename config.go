package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var (
	configPath string
	homeDir    string
	config     Configuration
)

type Configuration struct {
	Root string `yaml:"root"`
}

func saveRootPath(path string) error {
	config.Root = path
	buf, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, buf, os.ModePerm)
}

func init() {
	var err error
	homeDir, err = os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	configDirPath := filepath.Join(homeDir, ".config/motes")
	_, err = os.Lstat(configDirPath)
	if err != nil {
		err = os.MkdirAll(configDirPath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	configPath = filepath.Join(configDirPath, "config.yaml")
	f, err := os.OpenFile(configPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}

	s, _ := f.Stat()
	buf := make([]byte, s.Size())
	_, err = f.Read(buf)
	fmt.Printf("Path: %s, Content: %q, Size: %d\n", configPath, string(buf), s.Size())
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		panic(err)
	}
}
