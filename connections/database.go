package connections

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgresDB *gorm.DB

func Connect() {
	if PostgresDB == nil {
		var err error
		dsn := "host=" + "localhost" +
			" user=" + "postgres" +
			" password=" + "postgres" +
			" dbname=" + "pre_order_food" +
			" port=" + "5432" +
			" sslmode=disable"

		PostgresDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("[Connection], Error in opening db : ", err)
		}
		fmt.Println("DB connection established")
	}
}

func DB() *gorm.DB {
	if PostgresDB == nil {
		Connect()
	}
	return PostgresDB
}
