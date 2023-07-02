package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port     int    `yaml:"port"`
		Hostname string `yaml:"hostname"`
	} `yaml:"server"`
}

func main() {
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		fmt.Println("err in taking the config file", err)
	}
	fmt.Println("config file data :", data)

	var config Config

	if err := yaml.Unmarshal(data, &config); err != nil {
		fmt.Println(err)
	}
	fmt.Println("config data:", config)
	fmt.Println("Server port:", config.Server.Port)
	fmt.Println("server host:", config.Server.Hostname)
	//start server....
	addr := fmt.Sprintf("%s:%d", config.Server.Hostname, config.Server.Port)

	fmt.Println(data)
	log.Printf("server listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
