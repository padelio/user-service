package seeder

import (
	"user-service/domain/models"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RunUserSeeder(db *gorm.DB) {
	password, _ := bcrypt.GenerateFromPassword([]byte("adminpassword123"), bcrypt.DefaultCost)
	user := models.User{
		UUID:        uuid.New(),
		Name:        "Administrator",
		Username:    "admin",
		Password:    string(password),
		PhoneNumber: "081234567890",
		Email:       "admin@gmail.com",
		RoleID:      1,
	}

	err := db.FirstOrCreate(&user, models.User{
		Username: user.Username,
	}).Error
	if err != nil {
		logrus.Errorf("failed to seed user: %v", err)
		panic(err)
	}
	logrus.Infof("user %s successfully seeded", user.Username)
}
