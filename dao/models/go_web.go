package models

import (
	"github.com/lascape/gopkg/gormx"
)

type GoWeb struct {
	gormx.Model
}

// TableName sets the name of the table for GORM to use
func (*GoWeb) TableName() string {
	return "go_web"
}
