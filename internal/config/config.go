package config

type Config struct {
	Theme    string `json:"theme"`
	DataPath string `json:"dataPath"`
}

var DefaultConfig = Config{
	Theme:    "matrix",
	DataPath: "./notes",
}

func GetConfig() Config {
	return DefaultConfig
}
