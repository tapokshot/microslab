package searchuser

type Searchuser struct {
    Port     string `toml:"port"`
    Network  string `toml:"network"`
    LogLevel string `toml:"log_level"`
}

type Config struct {
    Searchuser `toml:"searchuser"`
}

func CreateConfig() *Config {
    result := &Config{}
    result.LogLevel = "debug"
    result.Network = "tcp"
    result.Port = "5300"
    return result
}
