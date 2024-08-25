package migrations

import "os"

func BaseMigrations() string {
	dbname := os.Getenv("DB_DATABASE")
	timezone := os.Getenv("DB_TIMEZONE")

	schema := `
		CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
		SELECT * FROM pg_timezone_names;
		ALTER DATABASE ` + dbname + ` SET timezone TO '` + timezone + `';
	`

	return schema
}
