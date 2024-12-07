package global

import (
	"context"

	"gorm.io/gorm"
)

var _db = &gorm.DB{}

func SetDB(db *gorm.DB) {
	_db = db
}

func DB(ctx context.Context) *gorm.DB {
	return _db.WithContext(ctx)
}

var _dbBot = &gorm.DB{}

func SetBotDB(db *gorm.DB) {
	_dbBot = db
}
