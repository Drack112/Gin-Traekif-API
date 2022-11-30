package config

import (
    "fmt"
    "os"
    "testing"

    "github.com/Drack112/Crud-Golang-API/entity"
    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

func GetEnvs(env string, t *testing.T) string {
    err := godotenv.Load("../.env")
    if err != nil {
        t.Error("Error in loading the .env file\n", err.Error())
    }

    return os.Getenv(env)
}

func TestSetupDatabasePostgres(t *testing.T) {

    dbUser := GetEnvs("DB_USER", t)
    dbPass := GetEnvs("DB_PASSWORD", t)
    dbHost := GetEnvs("DB_HOST", t)
    dbName := GetEnvs("DB_NAME", t)
    dbPort := GetEnvs("DB_PORT", t)

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPass, dbName, dbPort)

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default,
    })
    if err != nil {
        t.Error("Error in opening the database Postgres\n", err.Error())
    }

    db.AutoMigrate(&entity.User{}, &entity.Book{})
}
