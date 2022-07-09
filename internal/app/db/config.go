package db

type Config struct {
	driverName  string
	databaseUrl string
}

func NewConfig() *Config {
	return &Config{
		driverName: "postgres",
		databaseUrl: "host=localhost port=5432 user=postgres " +
			"password= dbname= sslmode=disable",
	}
}
