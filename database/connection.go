package database

import (
	"fmt"
	"log"

	"github.com/IsraelTeo/api-registry-court-files-MVP/configuration"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var GDB *gorm.DB

func Connection(cfg *configuration.Config) error {
	DSN := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Lima",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
	)

	var err error
	GDB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		return err
	}

	log.Println("✅ Conexión a PostgreSQL exitosa")
	return nil
}
