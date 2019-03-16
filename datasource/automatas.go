package datasource

import (
	"yorha-api/datamodels"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func database() []datamodels.Automata {
	db, err := gorm.Open(
		"mysql",
		"yorha_user:yorha_pass@/yorha_db?charset=utf8&parseTime=True&loc=Local",
	)
	if err != nil {
		panic("Connection failed to open!")
	}
	defer db.Close()

	var character []datamodels.Automata
	db.Find(&character)

	return character
}

// Automatas retrieve everything in the database and use it as source data
var Automatas = database()
