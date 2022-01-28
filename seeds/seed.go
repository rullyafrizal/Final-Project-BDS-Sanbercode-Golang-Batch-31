package seeds

import (
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	err := RoleSeeder(db)

	if err != nil {
		return err
	}

	err = UserSeeder(db)

	if err != nil {
		return err
	}

	return err
}
