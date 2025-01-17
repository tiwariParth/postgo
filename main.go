package main

import (
	"flag"
	"fmt"
	"os"
)

type Config struct{
	URL string
}
func parseFlags() (*Config, error){
	config:= &Config{}
	flag.StringVar(&config.URL, "url", "", "URL to send request to")
	flag.Parse()
	if config.URL == ""{
		return nil, fmt.Errorf("Error: URL is required")
	}
	return config, nil
}
func main() {
	config, err:= parseFlags()
	flag.Parse()
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("URL provided:%s\n", config.URL)
}
