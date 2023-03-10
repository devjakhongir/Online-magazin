package config

type Config struct {
	HTTPPort string

	PostgresHost           string
	PostgresUser           string
	PostgresDatabase       string
	PostgresPassword       string
	PostgresPort           string
	PostgresMaxConnections int32

	RedisAddr     string
	RedisPassword string
	RedisDB       int

	AuthSecretKey string
	SuperAdmin    string
	Client        string
}

func Load() Config {

	var cfg Config

	cfg.HTTPPort = ":4000"

	cfg.PostgresHost = "localhost"
	cfg.PostgresUser = "jahongir2"
	cfg.PostgresDatabase = "texnomart"
	cfg.PostgresPassword = "1005"
	cfg.PostgresPort = "5432"
	cfg.PostgresMaxConnections = 20

	return cfg
}
