package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "root:123456@tcp(127.0.0.1:3002)/personalidades?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(stringDeConexao), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
		log.Fatal(err.Error())
	}

	//erroPing := sqlDB.Ping()
	//
	//if erroPing != nil {
	//	log.Fatal(erroPing.Error())
	//
	//}
}
