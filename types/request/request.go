package request

import (
	"gin_example_with_generic/generic"
	"gin_example_with_generic/store/model"
)

type Country struct {
	Name      string `json:"name" xml:"name" yaml:"name" binding:"required,max=30" describe_zh:"国家名" describe_en:"country name" describe_ja:"国号"`
	NameCN    string `json:"name_cn" xml:"name_cn" yaml:"name_cn" binding:"required,max=30" describe_zh:"国家中文名" describe_en:"country chinese name" describe_ja:"国日文名"`
	ShortName string `json:"short_name" xml:"short_name" yaml:"short_name" binding:"required,max=30" describe_zh:"国家名简写" describe_en:"abbreviation of country name" describe_ja:"国名の略字"`
}

func (c Country) FormatToModel() generic.ModelInterface {
	return model.Country{
		Name:      c.Name,
		NameCN:    c.NameCN,
		ShortName: c.ShortName,
	}
}

type User struct {
	Name      string `json:"name" xml:"name" yaml:"name" binding:"required,max=30" describe_zh:"姓名" describe_en:"name" describe_ja:"名前"`
	CountryID int    `json:"country_id" xml:"country_id" yaml:"country_id" binding:"required,min=0" describe_zh:"国家ID" describe_en:"country id" describe_ja:"国ID"`
	Email     string `json:"email" xml:"email" yaml:"email" binding:"required,email" describe_zh:"邮箱" describe_en:"email" describe_ja:"メール"`
}

func (u User) FormatToModel() generic.ModelInterface {
	return model.User{
		Name:      u.Name,
		CountryID: u.CountryID,
		Email:     u.Email,
	}
}
