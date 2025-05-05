package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DataPath string
	Language string
}

var DefaultConfig = Config{
	DataPath: "./notes",
	Language: "en",
}

func GetConfig() Config {
	loadConfig() // 尝试加载配置
	return DefaultConfig
}

// SaveConfig 保存配置到文件
func SaveConfig() error {
	data, err := json.MarshalIndent(DefaultConfig, "", "  ")
	if err != nil {
		return err
	}

	configDir := "./config"
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(configDir, "config.json"), data, 0644)
}

// loadConfig 从文件加载配置
func loadConfig() {
	configPath := filepath.Join("./config", "config.json")
	data, err := os.ReadFile(configPath)
	if err != nil {
		return // 如果配置文件不存在，使用默认配置
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return // 如果解析失败，使用默认配置
	}

	DefaultConfig = cfg
}
