package infrastructure

import (
	"log"
	"ms_auth/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connect to msql database
func ConnectDatabases() {
	var err error
	db, err = gorm.Open(mysql.Open(msql_dns), &gorm.Config{})
	if err != nil {
		log.Println("failed to connect database")
	}
	log.Println("Connect to database successfully")
}

func MigrateDatabases(isMigrate bool) {
	if !isMigrate {
		return
	}

	err := db.AutoMigrate(
		&model.User{},
		&model.CacheMem{},
	)
	if err != nil {
		log.Println("failed MigrateDatabases:", err)
	}
}
