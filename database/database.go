package database

import (
	"log"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
}

func NewDatabase(host string, port int, user string, password string, dbname string) *gorm.DB {
	Database := Database{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		dbname:   dbname}

	return Database.GetConnection()
}

func (d *Database) GetConnection() *gorm.DB {

	dsn := "host="+d.host+" user="+d.user+" password="+d.password+" dbname="+d.dbname+" port="+
	strconv.Itoa(d.port)+ " sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to the database", err)
	} else {

		log.Println("Connected to Database")
	}
	return db

}
