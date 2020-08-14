package models

import (
  "fmt"
  "log"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
  USER = "root"
  PASS = "1605"
  HOST = "127.0.0.1"
  PORT = 3306
  DBNAME = "books-pagination"
)

func Connect() *gorm.DB {
  URL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
  db, err := gorm.Open("mysql", URL)
  if err != nil {
    log.Fatal(err)
    return nil
  }
  return db
}