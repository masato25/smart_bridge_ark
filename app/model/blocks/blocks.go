package blocks

import "time"

type Block struct {
	ID        string `gorm:"type:varchar(100);primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Height    int
	Reward    int64
}

func (self *Block) TableName() string {
	return "blocks"
}
