package seeds

import (
	"log"

	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/models"
	"gorm.io/gorm"
)

func RoleSeeder(db *gorm.DB) error {
	var roles []string = []string{"admin", "user"}

	var err error

	for _, v := range roles {
		var role models.Role = models.Role{
			Name: v,
		}

		err = db.Where(role).FirstOrCreate(&role).Error
	}

	log.Println("Role Seeder Success")

	return err
}
