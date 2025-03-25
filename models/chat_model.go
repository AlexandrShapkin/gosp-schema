package models

type Chat struct {
	ID   string `gorm:"primaryKey"`
	Name string
	Type string `gorm:"not null;check:type IN ('private', 'public')"`

	Messages    []Message         `gorm:"foreignKey:ChatID"`
	Paticipants []ChatPaticipants `gorm:"foreignKey:ChatID"`
}
