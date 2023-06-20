package model

import (
	"gin_example_with_generic/generic"
)

type Country struct {
	generic.Model `yaml:",inline"`
	Name          string `json:"name" xml:"name" yaml:"name" gorm:"size:30;uniqueIndex"`
	NameCN        string `json:"name_cn" xml:"name_cn" yaml:"name_cn" gorm:"size:30;uniqueIndex"`
	ShortName     string `json:"short_name" xml:"short_name" yaml:"short_name" gorm:"size:5;uniqueIndex"`
}

func (Country) GetFuzzySearchFieldList() []string {
	return []string{"name", "name_cn", "short_name"}
}
