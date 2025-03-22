package models

type ChatPaticipants struct {
	ChatID string `gorm:"primaryKey"`
	UserID string `gorm:"primaryKey"`

	Chat Chat `gorm:"foreignKey:ChatID;constraint:OnDelete:CASCADE"`
	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
