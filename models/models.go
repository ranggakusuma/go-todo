package models

import (
	"github.com/jinzhu/gorm"
	"github.com/ranggakusuma/go-todo/utils"
	uuid "github.com/satori/go.uuid"
)

// Todo is for todo type
type Todo struct {
	ID       string `json:"id" gorm:"type:uuid;primary_key;"`
	Title    string `json:"title"`
	Complete bool   `json:"complete"`
}

// BeforeCreate function
func (t *Todo) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()

	return scope.SetColumn("ID", uuid.String())
}

// Insert for insert todo anjing
func (t *Todo) Insert(db *gorm.DB) error {
	err := db.Save(&t).Error
	return err
}

// Get for get one
func (t *Todo) Get(db *gorm.DB, ID string) ([]Todo, error) {
	if ID != "" {
		err := db.First(&t).Where("ID = $1", ID).Error
		return nil, err
	}

	var dataTodo []Todo
	err := db.Table("todos").Scan(&dataTodo).Error
	return dataTodo, err
}

func init() {
	db := utils.DB()
	db.AutoMigrate(&Todo{})
}
