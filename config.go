package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	ListenAddr     string
	TargetURL      string
	AppendPath     bool
	HttpStatusCode int
}

var ConfigInstance *Config = &Config{}

func GetConfig() *Config {
	return ConfigInstance
}

func (c *Config) ReadConfig() {
	c.ListenAddr = c.getEnv("LISTEN_ADDR", "0.0.0.0:8080")
	c.TargetURL = strings.TrimSuffix(c.getEnv("TARGET", "http://localhost/"), "/")
	c.AppendPath = c.getEnvBool("APPEND_PATH", true)
	c.HttpStatusCode = c.getEnvInt("STATUS_CODE", http.StatusMovedPermanently)
}

func (c *Config) Print() {
	s, _ := json.Marshal(c)
	log.Println("Using config: " + string(s))
}

func (c *Config) getEnvInt(key string, defaultValue int) int {
	val := c.getEnv(key, strconv.Itoa(defaultValue))
	if i, err := strconv.Atoi(val); err != nil {
		return defaultValue
	} else {
		return i
	}
}

func (c *Config) getEnvBool(key string, defaultValue bool) bool {
	def := ""
	if defaultValue {
		def = "1"
	}
	val := strings.ToLower(c.getEnv(key, def))
	return (val == "1") || (val == "true") || (val == "yes") || (val == "on")
}

func (c *Config) getEnv(key, defaultValue string) string {
	res := os.Getenv(key)
	if res == "" {
		return defaultValue
	}
	return res
}
