package infrastructure

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	Host       string
	Username   string
	Password   string
	DBName     string
	Connection *gorm.DB
}

func NewDB(c *Config) *DB {
	fmt.Println(c.Language)
	return newDB(&DB{
		Host:     c.DB.Local.Host,
		Username: c.DB.Local.Username,
		Password: c.DB.Local.Password,
		DBName:   c.DB.Local.DBName,
	})
}

func NewTestDB(c *Config) *DB {
	fmt.Println(c.Language)
	return newDB(&DB{
		Host:     c.DB.Test.Host,
		Username: c.DB.Test.Username,
		Password: c.DB.Test.Password,
		DBName:   c.DB.Test.DBName,
	})
}

func newDB(d *DB) *DB {
	fmt.Println(d.Username + ":" + d.Password + "@tcp(" + d.Host + ")/" + d.DBName + "?charset=utf8&parseTime=True&loc=Local")
	// db, err := gorm.Open(mysql.Open(d.Username + ":" + d.Password + "@tcp(" + d.Host + ")/" + d.DBName + "?charset=utf8&parseTime=True&loc=Local"))
	db, err := gorm.Open(
		postgres.Open(
			"host=" + d.Host + " user=" + d.Username + " password=" + d.Password + " dbname=" + d.DBName + " port=5432",
		),
	)
	if err != nil {
		panic(err.Error())
	}
	d.Connection = db
	return d
}

func (db *DB) Begin() *gorm.DB {
	return db.Connection.Begin()
}

func (db *DB) Connect() *gorm.DB {
	return db.Connection
}
