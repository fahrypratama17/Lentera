package database

import (
	featureBEntity "github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/entity"
	userEntity "github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/user/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&userEntity.User{},
		&featureBEntity.Material{},
		// TODO: daftarin entity baru di sini
	)
}
