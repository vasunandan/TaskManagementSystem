package repositories

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "TaskManagementSystem/pkg/logging"
    "TaskManagementSystem/internal/models/postgres1"
    "fmt"
)

type PostgresValue struct {
    ServerName string
    Port       string
    User       string
    Password   string
    DBName     string
}

func (pgv *PostgresValue)ConnectToProstgresDB() *gorm.DB {
    var err error
    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",pgv.ServerName,pgv.Port,pgv.User,pgv.Password,pgv.DBName)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        logging.Log("Error connecting to database. Error: " + err.Error())
    }
    db.AutoMigrate(&postgres1.Task{})
    logging.Log("Connected to db sucessfully")
    return db
}
