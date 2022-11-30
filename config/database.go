package config

import (
    "fmt"
    "log"
    "os"

    "github.com/Drack112/Crud-Golang-API/entity"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

func SetupDatabasePostgres() *gorm.DB {

    dbUser := os.Getenv("DB_USER")
    dbPass := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbName := os.Getenv("DB_NAME")
    dbPort := os.Getenv("DB_PORT")

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPass, dbName, dbPort)

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default,
    })
    if err != nil {
        log.Panicf("Error in opening the database Postgres\n %s", err.Error())
    }

    db.AutoMigrate(&entity.User{}, &entity.Book{})
    return db
}

func CloseDatabasePostgres(db *gorm.DB) {
    dbSQL, err := db.DB()
    if err != nil {
        log.Panicf("Error in opening the database Postgres\n %s", err.Error())
    }

    dbSQL.Close()
}
