package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DatabaseConection() {
	dsn := "root@tcp(localhost:3306)/api-golang?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
}
