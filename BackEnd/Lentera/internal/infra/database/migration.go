package database

import (
	featureAEntity "github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/feature_a/entity"
	featureBEntity "github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/feature_b/entity"
	userEntity "github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/user/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&userEntity.User{},
		&featureAEntity.FeatureA{},
		&featureBEntity.FeatureB{},
		// TODO: daftarin entity baru di sini
	)
}
