package main

import (
    "flag"
    "github.com/BurntSushi/toml"
    "github.com/tapokshot/microslab/internal/app/microslab"
    "log"
)

var (
    configPath string
)

func init() {
    flag.StringVar(&configPath, "config-path", "configs/application.toml", "Application config")
}

func main() {
    flag.Parse()

    config := microslab.CreateConfig()
    _, err := toml.DecodeFile(configPath, &config)
    if err != nil {
        log.Fatal(err)
    }
    launcher := microslab.New(config)

    if err := launcher.Launch(); err != nil {
        log.Fatal(err)
    }

}
