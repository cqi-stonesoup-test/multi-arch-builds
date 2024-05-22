package main

/**
TODO: do not set argument idx in the querystring
TODO: set argument n in the querystring
TODO: iterate the images rather than just handling the first one
*/

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path"
	"strings"
	"time"
)

type DailyPhotoInfo struct {
	Url     string `json:"url"`
	Urlbase string `json:"urlbase"`
}

type BingDailyPhotoInfo struct {
	Images []DailyPhotoInfo `json:"images"`
}

const BingImageEndpoint = "https://www.bing.com/HPImageArchive.aspx"

func httpGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if content, err := io.ReadAll(resp.Body); err == nil {
		return content, nil
	} else {
		return nil, err
	}
}

func getBingDailyPhotoInfo(url string) (*BingDailyPhotoInfo, error) {
	var bingDailyPhotoInfo BingDailyPhotoInfo
	if data, err := httpGet(url); err == nil {
		if err := json.Unmarshal(data, &bingDailyPhotoInfo); err == nil {
			return &bingDailyPhotoInfo, nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

// Download a photo specified by a given URL.
func downloadPhoto(url string) ([]byte, error) {
	return httpGet(url)
}

func main() {
	var photoDir string
	var market string
	var resolution string

	thisUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	defaultPhotoDir := path.Join(thisUser.HomeDir, "Pictures/BingWallpaper")

	flag.StringVar(&photoDir, "download-dir", defaultPhotoDir, "Path to the download directory.")
	flag.StringVar(&market, "market", "zh-CN", "Limit the photo to be downloaded from this market. For what market is, please refer to bing.com")
	flag.StringVar(&resolution, "resoluation", "1920x1080", "Photo's resolution.")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(7) + 1

	downloadUrl := fmt.Sprintf("%s?format=js&idx=%d&n=1&mkt=%s", BingImageEndpoint, idx, market)
	bingDailyPhotoInfo, err := getBingDailyPhotoInfo(downloadUrl)
	if err != nil {
		log.Fatalln(err)
	}

	urlInfo, err := url.Parse(bingDailyPhotoInfo.Images[0].Url)
	if err != nil {
		log.Fatalln(err)
	}

	/*** So far, we have the daily photo information. Next step is to download it. ***/

	// Remove prefix OHR. from filename. Example:
	// OHR.WhaleHug_EN-CN8835644169_1920x1080.jpg -> WhaleHug_EN-CN8835644169_1920x1080.jpg
	localFilename := path.Join(photoDir, strings.ReplaceAll(urlInfo.Query()["id"][0], "OHR.", ""))

	_, err = os.Stat(localFilename)
	if !os.IsNotExist(err) {
		log.Println("Photo exists:", localFilename)
		return
	}

	// Example URL: https://www.bing.com//th?id=OHR.WhaleHug_EN-CN8835644169_1920x1080.jpg
	photoDownloadUrl := fmt.Sprintf("https://www.bing.com%s_%s.jpg", bingDailyPhotoInfo.Images[0].Urlbase, resolution)
	log.Println("Downloading new photo:", photoDownloadUrl)
	photoBytes, err := downloadPhoto(photoDownloadUrl)
	if err != nil {
		log.Fatalln(err)
	}
	if err = os.WriteFile(localFilename, photoBytes, 0644); err != nil {
		log.Fatalln(err)
	}
	log.Println("New photo:", localFilename)
}
