package persistence

import "gorm.io/gorm"

type Db interface {
	Create(v interface{}) *gorm.DB
}
