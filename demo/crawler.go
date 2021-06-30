package demo

import (
	"WebWork/tool"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/simplifiedchinese"
)

//秀人网获取图片和百度网盘
var url = ""

func Xiuren(newUrl string) {
	url = newUrl
	imgCovers := getImageCover(url)
	getBaidu(imgCovers)
}

func getImageCover(url string) (imgCovers []string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("request fail", err)
		return
	}
	reader := simplifiedchinese.GB18030.NewDecoder().Reader(res.Body) //需要将bgk转utf8
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		fmt.Println("解析HTML失败", err)
		return
	}
	s := doc.Find("div.bus_vtem")
	s.Each(func(i int, s *goquery.Selection) { //获取节点集合并遍历
		href, _ := s.Children().First().Attr("href")
		if !strings.HasPrefix(href, "http") {
			href = "https://www.heisiyun.com/" + href
			imgCovers = append(imgCovers, href)
		}

	})
	return
}

func downloadImage(imgCovers []string) {

}

func getBaidu(imgCovers []string) {
	for _, v := range imgCovers {
		res, err := http.Get(v)
		if err != nil {
			fmt.Println("request fail", err)
			return
		}
		reader := simplifiedchinese.GB18030.NewDecoder().Reader(res.Body) //需要将bgk转utf8
		doc, err := goquery.NewDocumentFromReader(reader)
		if err != nil {
			fmt.Println("解析HTML失败", err)
			return
		}
		s := doc.Find("td.t_f")
		txt := s.First().Text()
		txt = strings.Replace(txt, " ", "", -1)
		txt = strings.Replace(txt, "\n", "", -1)
		if strings.Contains(txt, "百度网盘") {

			index := strings.Index(txt, "链接：")
			index2 := strings.Index(txt, "提取码：")
			link := txt[index:index2]
			code := txt[index2+4:]
			fmt.Println(link)
			writeBaiduLink(link, code)
		}

	}
}

func writeBaiduLink(link string, code string) {
	var fileName = "./秀人" + tool.Substr(url, 31, 5) + ".txt"
	var f *os.File
	var err1 error
	if tool.CheckFileIsExist(fileName) {
		f, err1 = os.OpenFile(fileName, os.O_APPEND, 0666)
		fmt.Println("文件存在")
	} else {
		f, err1 = os.Create(fileName)
		fmt.Println("文件不存在")
	}
	if err1 != nil {
		panic(err1)

	}
	n, err1 := io.WriteString(f, link+"           "+code+"\n")
	if err1 != nil {
		panic(err1)

	}
	fmt.Println(n)
}
