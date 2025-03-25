package gospschema

import (
	"github.com/AlexandrShapkin/gosp-schema/models"
	"gorm.io/gorm"
)

func ConnectDB(dialector *gorm.Dialector) (*gorm.DB, error) {
	db, err := gorm.Open(*dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Chat{},
		&models.ChatParticipants{},
		&models.Message{},
		&models.Token{},
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}
