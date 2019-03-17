package datasource

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func database() *gorm.DB {
	db, err := gorm.Open(
		"mysql",
		"yorha_user:yorha_pass@/yorha_db?charset=utf8&parseTime=True&loc=Local",
	)
	if err != nil {
		panic("Connection failed to open!")
	}
	//defer db.Close()

	return db
}

// Database retrieve the instance for the database
var Database = database()
