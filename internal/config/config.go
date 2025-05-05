package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
)

type Config struct {
	DataPath string
	Language string
}

var DefaultConfig = Config{
	DataPath: getDefaultDataPath(),
	Language: "en",
}

func getDefaultDataPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return "./notes"
	}

	switch runtime.GOOS {
	case "windows":
		// Windows: %APPDATA%\gotodo\notes
		return filepath.Join(home, "AppData", "Roaming", "gotodo", "notes")
	case "darwin":
		// macOS: ~/.local/share/gotodo/notes
		return filepath.Join(home, ".local", "share", "gotodo", "notes")
	default:
		// Linux/Unix: ~/.local/share/gotodo/notes
		return filepath.Join(home, ".local", "share", "gotodo", "notes")
	}
}

func getConfigDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return "./config"
	}

	switch runtime.GOOS {
	case "windows":
		// Windows: %APPDATA%\gotodo
		return filepath.Join(home, "AppData", "Roaming", "gotodo")
	case "darwin":
		// macOS: ~/.config/gotodo
		return filepath.Join(home, ".config", "gotodo")
	default:
		// Linux/Unix: ~/.config/gotodo
		return filepath.Join(home, ".config", "gotodo")
	}
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

	configDir := getConfigDir()
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(configDir, "config.json"), data, 0644)
}

// loadConfig 从文件加载配置
func loadConfig() {
	configPath := filepath.Join(getConfigDir(), "config.json")
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
