package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Command struct {
	Path        string   `json:"path"`
	Command     string   `json:"command"`
	Arguments   []string `json:"arguments"`
	Description string   `json:"description"`
}

type ProxyRequest struct {
	Path        string `json:"path"`
	Url         string `json:"url"`
	Description string `json:"description"`
}

type Settings struct {
	Address       string         `json:"address"`
	ApiToken      string         `json:apiToken`
	Commands      []Command      `json:commands`
	ProxyRequests []ProxyRequest `json:proxyRequests`
}

type ApiEndpoint struct {
	Path        string `json:"path"`
	Description string `json:"description"`
}

func (c Command) GetEndPoint() ApiEndpoint {
	return ApiEndpoint{Path: c.Path, Description: c.Description}
}

func (c ProxyRequest) GetEndPoint() ApiEndpoint {
	return ApiEndpoint{Path: c.Path, Description: c.Description}
}

func getSettings() Settings {

	var settingsFilePath string
	args := os.Args[1:]
	if len(args) == 2 && args[0] == "--settings" {
		settingsFilePath = args[1]
	} else if len(args) == 0 {
		settingsFilePath = "settings.json"
	} else {
		fmt.Println("invalid arguments")
		fmt.Println("use --settings path-to-settings/settings.json")
		fmt.Println("default config path is ./settings.json")
		os.Exit(2)
	}

	settingsFile, err := ioutil.ReadFile(settingsFilePath)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	var settings Settings
	err = json.Unmarshal(settingsFile, &settings)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(13)
	}
	return settings
}
