package config

import (
  "os"
  "strings"
)

type Config struct {
  Port        string
  Backends    []string
  Algorithm   string
}

func Load() *Config {
  return &Config{
    Port:       getEnv("PORT", "8000"),
    Backends:   strings.Split(getEnv("BACKENDS", "http:/localhost:8080,http:/localhost:8081"), ","),
    Algorithm:  getEnv("ALGORITHM", "round-robin"),
  }
}

func getEnv(key, fallback string) string {
  if value, ok := os.LookupEnv(key); ok {
    return value
  }
  return fallback
}
