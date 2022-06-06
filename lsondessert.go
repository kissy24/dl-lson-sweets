package main

import (
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Image struct {
	URL      string
	FileName string
}

/*
Get image URLs
Args:
	domainUrl: domain URL
	srcDirUrl: source directory URL
	imgDirUrl: image directory URL
Returns:
	[]Image: image URLs
*/
func GetImageURLs(domainUrl string, srcDirUrl string, imgDirUrl string) ([]Image, error) {
	doc, err := goquery.NewDocument(domainUrl + srcDirUrl)
	if err != nil {
		return nil, err
	}
	var images []Image
	doc.Find("img").Each(func(_ int, s *goquery.Selection) {
		imgFileUrl, _ := s.Attr("src")
		imgFileName, _ := s.Attr("alt")
		hasImgPattern := strings.Contains(imgFileUrl, imgDirUrl)
		if hasImgPattern {
			imgUrl := domainUrl + imgFileUrl
			img := Image{imgUrl, imgFileName}
			images = append(images, img)
		}
	})
	return images, nil
}

/*
Download images
Args:
	imgUrls: image URLs
	imgDir: image directory
*/
func DownloadImages(imgs []Image, imgDir string) {
	for _, img := range imgs {
		response, err := http.Get(img.URL)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()
		if _, err := os.Stat(imgDir); os.IsNotExist(err) {
			os.MkdirAll(imgDir, 0777)
		}
		if strings.Contains(img.FileName, "/") {
			img.FileName = strings.Replace(img.FileName, "/", "-", 1)
		}
		file, err := os.Create(imgDir + "/" + img.FileName + ".jpg")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		io.Copy(file, response.Body)
	}
}

func main() {
	domainUrl := "https://www.lawson.co.jp"
	srcDirUrl := "/recommend/original/dessert/"
	imgDirUrl := "/recommend/original/detail/img/"
	imgUrls, err := GetImageURLs(domainUrl, srcDirUrl, imgDirUrl)
	if err != nil {
		panic(err)
	}
	dateDir := time.Now().Format("2006-01-02")
	DownloadImages(imgUrls, "./images/"+dateDir)
}
