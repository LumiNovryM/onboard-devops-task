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
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
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
		utils.PRINT_LOG("ERROR", "Missing POSTGRES_URL in .env", utils.GetLocation(), "")
		return
	}

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		utils.PRINT_LOG("ERROR", "Failed To Connect Database With GORM...", utils.GetLocation(), err.Error())
		return
	}
}
