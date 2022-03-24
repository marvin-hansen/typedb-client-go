// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package v2

type Config struct {
	DBHost    string
	DBPort    string
	DBName    string
	Infer     bool
	Explain   bool
	BatchSize int
}

func (c Config) GetConnectionString() string {
	return c.DBHost + ":" + c.DBPort
}

func NewLocalConfig(dbName string) *Config {
	return &Config{
		DBHost:    "",
		DBPort:    "1729",
		DBName:    dbName,
		Infer:     true,
		Explain:   true,
		BatchSize: 0,
	}
}

func NewConfig(dbHost, dbPort, dbName string, infer, explain bool, batchSize int) *Config {
	return &Config{
		DBHost:    dbHost,
		DBPort:    dbPort,
		DBName:    dbName,
		Infer:     infer,
		Explain:   explain,
		BatchSize: batchSize,
	}
}
