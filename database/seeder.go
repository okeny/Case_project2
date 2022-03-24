package database

import (
	"log"
	"project-cases/entity"

	"github.com/jinzhu/gorm"
)

var cases = []entity.Case{
	entity.Case{
		ID:          1,
		Title:       "Steven victor",
		Description: "steven@gmail.com",
	},
	entity.Case{
		ID:          2,
		Title:       "Martin Luther",
		Description: "luther@gmail.com",
	},
}


func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&entity.Case{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&entity.Case{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}


	for i, _ := range cases {
		err = db.Debug().Model(&entity.Case{}).Create(&cases[i]).Error
		if err != nil {
			log.Fatalf("cannot seed Cases table: %v", err)
		}
		
	}
}
