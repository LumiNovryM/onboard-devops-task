package connection

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"backend/app/utils"
	"backend/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func OpenConnection() (*sql.DB, error) {
	dsn := os.Getenv("POSTGRES_URL")
	if dsn == "" {
		dsn = buildDSNFromEnv()
	}
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitDB() {
	conf, err, location := utils.GET_CONFIG()
	if err != nil {
		utils.PRINT_LOG("ERROR", "GET CONFIG", location, fmt.Sprintf(err.Error()))
	}
	logger.Initialize(conf)

	dsn := os.Getenv("POSTGRES_URL")
	if dsn == "" {
		utils.PRINT_LOG("WARNING", "POSTGRES_URL not found, using individual env variables instead", utils.GetLocation(), "")
		dsn = buildDSNFromEnv()
	}

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		utils.PRINT_LOG("ERROR", "Failed To Connect Database With GORM...", utils.GetLocation(), err.Error())
		return
	}
}

// Fallback builder dari env kalau POSTGRES_URL tidak tersedia
func buildDSNFromEnv() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
}
