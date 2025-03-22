package models

type Chat struct {
	ID   string `gorm:"primaryKey"`
	Name string
	Type string `gorm:"type:enum('private','public');not null"`

	Messages    []Message         `gorm:"foreignKey:ChatID"`
	Paticipants []ChatPaticipants `gorm:"foreignKey:ChatID"`
}
