package api

import (
	"WebWork/model"
	"fmt"
	"math/rand"
	"net/http"
	"path"
	"time"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println(err)
		var res = model.Result{Code: -11, Count: 0, Msg: "request not contail file", Data: nil}
		c.JSON(http.StatusOK, res)
		return
	}
	files := form.File["pictures[]"]

	for _, file := range files {
		// 可以在这里对每个文件进行操作，例如保存到本地
		ext := path.Ext(file.Filename)
		randomNum := rand.Intn(10000)
		now := time.Now().UnixNano()
		dst := fmt.Sprintf("img_%d_%d%s", now, randomNum, ext)
		c.SaveUploadedFile(file, "static/img/"+dst)
	}
	// gin 简单做了封装,拷贝了文件流
	var res = model.Result{Code: 0, Count: 0, Msg: "success", Data: nil}
	c.JSON(http.StatusOK, res)

}
