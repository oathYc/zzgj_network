package global

import (
	"github.com/jinzhu/gorm"
)

var (
	DbWrapper map[string]*gorm.DB
)
