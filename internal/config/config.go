package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env        string     `yaml:"env" env-default:"local"` // we can add env-required:"true" to be sure that env isn't default
	DB         Database   `yaml:"database"`
	Server     HTTPServer `yaml:"http_server"`
	CacheStore Cache      `yaml:"cache"`
}

type Database struct {
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
	User    string `yaml:"user"`
	DBName  string `yaml:"dbname"`
	SSLMode string `yaml:"sslmode"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:":8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
	User        string        `yaml:"user" env-required:"true"`
	Password    string        `yaml:"password" env-required:"true" env:"HTTP_SERVER_PASSWORD"`
}

type Cache struct {
	Addr        string        `yaml:"address"`
	Timeout     time.Duration `yaml:"timeout"`
	DialTimeout time.Duration `yaml:"dial_timeout"`
}

// Upload config path from .env
// Use this to have an option of choosing our
// environment(dev,prod,local etc.)
func MustLoad() string {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("no .env file found")
	}
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH isn't set")
	}

	//check the presence of the file
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s\n", configPath)
	}

	return configPath
}

func LoadAllConfig() *Config {
	configPath := MustLoad()

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s\n", err.Error())
	}
	return &cfg
}

func (c *Config) PostgreDSN() string {
	password := os.Getenv("DB_PASSWORD")

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.DB.User,
		password,
		c.DB.Host,
		c.DB.Port,
		c.DB.DBName,
		c.DB.SSLMode,
	)
}
