package seeds

import (
	"log"

	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/models"
	"gorm.io/gorm"
)

func UserSeeder(db *gorm.DB) error {
	var role models.Role

	db.Where("name = ?", "admin").First(&role)

	var user models.User = models.User{
		Name:     "Admin",
		Email:    "admin@admin.com",
		Password: "Admin123",
		RoleId:   role.ID,
	}

	err := db.Where(models.User{Email: user.Email}).FirstOrCreate(&user).Error

	log.Println("User Seeder Success")

	return err
}
