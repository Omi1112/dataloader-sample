package model

type Todo struct {
	ID     int `gorm:primary_key;AUTO_INCREMENT`
	UserID int
	Body   string
}
