package storage

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	retryTimes = 3
)

var db *gorm.DB

func init() {
	db = connect()
}

func GetInstance() *gorm.DB {
	return db
}

func connect() *gorm.DB {
	pass, err := ioutil.ReadFile("/run/secrets/db-password")
	if err != nil {
		log.Fatal("fail to load db password")
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		string(pass),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DB"),
	)

	for i := 0; i < retryTimes; i++ {
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Printf("connect to DB failed, retry: %d\n err: %v\n", i+1, err)
			time.Sleep(time.Second * (1 << i))
			continue
		}
		return db
	}

	log.Fatal("fail to connect DB")
}
