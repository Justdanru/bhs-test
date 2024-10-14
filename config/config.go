package config

type (
	Config struct {
		HTTP HTTPConfig
		PG   PostgreSQLConfig
	}

	HTTPConfig struct {
		Host string
		Port string
	}

	PostgreSQLConfig struct {
		User     string
		Password string
		Host     string
		Port     string
		DB       string
	}
)

func NewConfig() *Config {
	return &Config{
		HTTP: HTTPConfig{
			Host: "127.0.0.1",
			Port: "8085",
		},
		PG: PostgreSQLConfig{
			User:     "bhs",
			Password: "password",
			Host:     "localhost",
			Port:     "5432",
			DB:       "bhs",
		},
	}
}
