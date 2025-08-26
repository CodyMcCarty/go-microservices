package database

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DatabaseClient interface {
	Ready() bool
}

type Client struct { // 1:44 set up db
	DB *gorm.DB
}

func NewDatabaseClient() (DatabaseClient, error) {
	// todo pass in config. Frank says he typically passes in the config, but not for this course. 1.SetUpTheDatabaseClient 1:40
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		"localhost", // or the remote instance
		"postgres",
		"postgres",
		"postgres",
		5432,
		"disable",
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "wisdom.", // bc he defined everything in a schema, we need to add this or reference it everywhere and that's too painful 1.SetUpDBClient 4:02
		},
		NowFunc: func() time.Time { // why do I need this function?
			return time.Now().UTC()
		},
		QueryFields: true,
	})
	if err != nil {
		return nil, err
	}
	client := Client{
		DB: db,
	}
	return client, nil
}

func (c Client) Ready() bool {
	var ready string
	// tx = transaction
	tx := c.DB.Raw("SELECT 1 as ready").Scan(&ready)
	if tx.Error != nil {
		return false
	}
	if ready == "1" {
		return true
	}
	return false
}
