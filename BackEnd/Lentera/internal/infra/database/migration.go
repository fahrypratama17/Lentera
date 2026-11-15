package database

import (
	cohortEntity "github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/cohort/entity"
	materiEntity "github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/entity"
	userEntity "github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/user/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		// user (Mone)
		&userEntity.User{},
		&userEntity.RefreshToken{},

		// cohort (Mone)
		&cohortEntity.Cohort{},
		&cohortEntity.Membership{},
		&cohortEntity.Invitation{},

		// materi (Neval) — Material dulu, baru Progress (ada FK ke materials)
		&materiEntity.Material{},
		&materiEntity.Progress{},

		// TODO: daftarin entity baru di sini
	)
}
