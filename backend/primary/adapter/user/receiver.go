package user_primary_adapter

import "gorm.io/gorm"

// FIXME: おそらくヘキサゴナルアーキテクチャらしくないので後から修正する
type myDB struct {
	*gorm.DB
}

func NewMyDB(db *gorm.DB) *myDB {
	return &myDB{DB: db}
}