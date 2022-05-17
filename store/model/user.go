package model

import "gin_example_with_generic/generic"

type User struct {
	generic.SoftDeleteModel `yaml:",inline"`
	Name                    string   `json:"name" xml:"name" yaml:"name" gorm:"size:30;uniqueIndex"`
	CountryID               int      `json:"-" xml:"-" yaml:"-"`
	Country                 *Country `json:"country" xml:"country" yaml:"country" gorm:"foreignkey:CountryID;constraint:OnDelete:SET NULL"`
	Email                   string   `json:"email" xml:"email" yaml:"email" gorm:"size:100"`
}

func (User) GetFuzzySearchFieldList() []string {
	return []string{"name", "email"}
}
