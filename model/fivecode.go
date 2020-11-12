package model

import (
	"fmt"
	"strings"

	"WebWork/global"

	"github.com/jinzhu/gorm"
)

type FiveCode struct {
	gorm.Model  `json:"-"`
	Word        string
	PinYin      string
	FiveCode    string
	ImgCode     string
	ImgKeyboard string
}

func GetCodes(letters []string) (codes []FiveCode, err error) {
	var sql = fmt.Sprintf(`select id,word,pin_yin,five_code,img_code,img_keyboard from five_code where word in ('%s')`, strings.Join(letters, ","))
	fmt.Println(sql)
	global.G_DB.Raw(sql).Scan(&codes)
	return
}

func SaveFive(code FiveCode) {
	global.G_DB.Create(&code)
}
