package config

type AppConfig struct {
    Port string `envconfig:"PORT" default:"3000"`
}
