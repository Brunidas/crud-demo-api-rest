package commons

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func GetConnections() *gorm.DB {
	db, error := gorm.Open("mysql", "root:@/test?charset=utf8")
	if error != nil {
		log.Fatal(error)
	}
	return db
}

func Migrate() {
	db := GetConnections()
	log.Println("Migrando...")

	db.AutoMigrate(&models.Persona{})
}
