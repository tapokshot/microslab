package microslab

// Server config
type Config struct {
    BindAddr    string `toml:"bind_addr"`
    LogLevel    string `toml:"log_level"`
    DatabaseUrl string `toml:"database_url"`
}

func CreateConfig() *Config {
    return &Config{
        BindAddr: "8080",
        LogLevel: "debug",
    }
}
