package model

import "gorm.io/gorm"

var DbConnect *gorm.DB

type Tabler interface {
    TableName() string
}
