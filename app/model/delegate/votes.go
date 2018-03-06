package delegate

import "time"

type Vote struct {
	ID        string `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Address   string     `gorm:"type:varchar(100);unique_index`
	Balance   float64
	Status    bool
}
