package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	// DSN (Data Source Name) de MySQL
	dsn := "root:root@tcp(127.0.0.1:3306)/arqui2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// var tables []string
	// err = db.Raw("SHOW TABLES").Scan(&tables).Error
	// if err != nil {
	// 	log.Printf("Error querying tables: %v", err)
	// } else {
	// 	log.Println("Database tables:")
	// 	for _, table := range tables {
	// 		log.Printf("- %s", table)
	// 	}
	// }
	return db
}
