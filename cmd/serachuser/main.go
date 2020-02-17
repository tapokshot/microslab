package main

import (
    "flag"
    "github.com/BurntSushi/toml"
    "github.com/tapokshot/microslab/internal/app/searchuser"
    "log"
)

var (
    configPath string
)

func init() {
    flag.StringVar(&configPath, "config-path", "configs/application.toml", "Application config")
}

func main() {
    log.Println("Parse config file")
    flag.Parse()

    config := searchuser.CreateConfig()
    _, err := toml.DecodeFile(configPath, &config)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Server startup")
    server := searchuser.New(config)
    if err := server.Launch(); err != nil {
        log.Fatal(err)
    }
}
