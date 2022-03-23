package utility

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Database struct{
	host string
	port int
	user string
	password string
	dbname string
}


func NewDatabase(host string,port int, user string,password string, dbname string )*Database{
   return &Database{
	   host,
	   port,
	   user,
	   password,
	   dbname }
}

func (d *Database)GetConnection() *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		d.host, d.port, d.user, d.password, d.dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic("failed to connect to the database")
	}
	return db
}
