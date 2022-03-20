package v2

type Config struct {
	DBHost string
	DBPort string
	DBName string
}

func (c Config) GetConnectionString() string {
	return c.DBHost + ":" + c.DBPort
}

func NewLocalConfig(dbName string) *Config {
	return &Config{
		DBHost: "",
		DBPort: "1729",
		DBName: dbName,
	}
}

func NewConfig(dbHost, dbPort, dbName string) *Config {
	return &Config{
		DBHost: dbHost,
		DBPort: dbPort,
		DBName: dbName,
	}
}
