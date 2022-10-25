package config

type AppConfig struct {
	Title  string `yaml:"title"`
	Server Server `yaml:"server"`
	Logger Logger `yaml:"logger"`
	DB     DB     `yaml:"postgres"`
}

type Server struct {
	Bind string `yaml:"bind"`
}

type Logger struct {
	Level string `yaml:"level"`
}

type DB struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SslMode  string `yaml:"sslmode"`
}
