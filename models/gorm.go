package models

import (
	"log"
	"os"
	"time"

	"github.com/PrinceNorin/monga/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var ORM *gorm.DB

func InitGorm() error {
	c := config.Get()
	gorm.NowFunc = func() time.Time {
		return time.Now().In(c.GetLocation())
	}

	db, err := gorm.Open(c.DB.Type, c.DB.Params)
	if err != nil {
		return err
	}

	if err := db.DB().Ping(); err != nil {
		return err
	}

	db.DB().SetMaxIdleConns(c.DB.MaxIdle)
	db.DB().SetMaxOpenConns(c.DB.MaxOpen)

	if c.Logger.Enabled {
		db.LogMode(true)
		if c.Logger.Filename != "" {
			f, err := os.Create(c.Logger.Filename)
			if err != nil {
				return err
			}
			db.SetLogger(log.New(f, "\r\n", 0))
		}
	}

	if err := runMigrate(db); err != nil {
		return err
	}

	ORM = db
	return nil
}
