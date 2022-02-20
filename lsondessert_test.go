package main

import "testing"

func TestGetImageURLs(t *testing.T) {
	domainUrl := "https://www.lawson.co.jp"
	srcDirUrl := "/recommend/original/dessert/"
	imgDirUrl := "/recommend/original/detail/img/"
	imgUrls, err := GetImageURLs(domainUrl, srcDirUrl, imgDirUrl)
	if err != nil {
		t.Error(err)
	}
	if len(imgUrls) == 0 {
		t.Error("imgUrls is empty")
	}
}

func TestDownloadImages(t *testing.T) {
	domainUrl := "https://www.lawson.co.jp"
	srcDirUrl := "/recommend/original/dessert/"
	imgDirUrl := "/recommend/original/detail/img/"
	imgUrls, err := GetImageURLs(domainUrl, srcDirUrl, imgDirUrl)
	if err != nil {
		t.Error(err)
	}
	if len(imgUrls) == 0 {
		t.Error("imgUrls is empty")
	}
	imgDir := "./test_images"
	DownloadImages(imgUrls, imgDir)
}
