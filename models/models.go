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

	return scope.SetColumn("id", uuid.String())
}

// Insert for insert todo anjing
func (t *Todo) Insert(db *gorm.DB) error {
	err := db.Save(&t).Error
	return err
}

// Get for get one
func (t *Todo) Get(db *gorm.DB, ID string) ([]Todo, error) {
	if ID != "" {
		err := db.Where("id = $1", ID).First(t).Error
		return nil, err
	}

	var dataTodo []Todo
	err := db.Table("todos").Scan(&dataTodo).Error
	return dataTodo, err
}

// Delete function
func (t *Todo) Delete(db *gorm.DB, ID string) error {
	_, err := t.Get(db, ID)
	if err != nil {
		return err
	}

	err = db.Where("id = $1", ID).Delete(t).Error
	return err
}

func init() {
	db := utils.DB()
	db.AutoMigrate(&Todo{})
}
