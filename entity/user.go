package entity

// User represents users table in database
type User struct {
	ID       uint64 `gorm:"primaryKey:autoIncrement" json:"id"`
	Name     string `gorm:"type:varchar(255)" json:"name"`
	Email    string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password string `gorm:"->;<-;not null" json:"-"` // allow read & create access, not shown in returned data
	Token    string `gorm:"-" json:"token,omitempty"`
}
