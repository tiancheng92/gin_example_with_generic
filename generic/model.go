package generic

import (
	"gorm.io/gorm"
	"time"
)

type SoftDeleteModel struct {
	Model     `yaml:",inline"`
	DeletedAt gorm.DeletedAt `json:"-" xml:"-" yaml:"-" gorm:"index"`
}

type Model struct {
	ID        int       `json:"id" xml:"id" yaml:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at" xml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `json:"updated_at" xml:"updated_at" yaml:"updated_at"`
}

func (Model) GetPrimaryKeyName() string {
	return "id"
}

func (m Model) GetPrimaryKey() any {
	return m.ID
}

type ModelInterface interface {
	GetPrimaryKeyName() string
	GetPrimaryKey() any
	GetFuzzySearchFieldList() []string
}
