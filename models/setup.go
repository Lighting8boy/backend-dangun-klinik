package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/finpro_web"))
	if err != nil {
		panic(err)

	}

	database.AutoMigrate(&Dokters{})
	database.AutoMigrate(&Jenis_Konsultasi{})
	database.AutoMigrate(&User{})
	database.AutoMigrate(&Riwayat{})
	database.AutoMigrate(&Konsultasi{})

	DB = database
}
