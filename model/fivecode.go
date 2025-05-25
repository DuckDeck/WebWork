package model

import (



	"WebWork/global"

	"github.com/jinzhu/gorm"
)

type FiveCode struct {
	gorm.Model  `json:"-"`
	Zi       string `json:"word"`
	Py      string `json:"spell"`
	Wubi    string	`json:"code"`
}
func (FiveCode) TableName() string {
    return "xhzd_surnfu" // 替换为实际表名
}

func GetCodes(letters []string) (codes []FiveCode, err error) {
	global.G_DB.Where(`zi IN (?)`,letters).Find(&codes)
	return
}

func SaveFive(code FiveCode) {
	global.G_DB.Create(&code)
}
