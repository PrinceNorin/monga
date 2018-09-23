package models

import "github.com/jinzhu/gorm"

func runMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&Manga{}, &Theme{}).Error; err != nil {
		return err
	}
	return nil
}
