package api

import (
	"WebWork/model"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AllImg(c *gin.Context) {
	dir := "static/img"
	// 读取目录下的所有文件和子目录信息
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		var res = model.Result{Code: -12, Count: 0, Msg: "没有该目录", Data: nil}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	var fileNames []string
	for _, file := range files {
		if !file.IsDir() {
			// 获取文件名（包含扩展名）并添加到切片中
			fileNames = append(fileNames, file.Name())
		}
	}
	var res = model.Result{Code: 0, Count: 0, Msg: "success", Data: fileNames}
	c.JSON(http.StatusOK, res)

}
