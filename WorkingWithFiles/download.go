package WorkingWithFiles

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)


func downloadFile(URL string) error {
	URL = "https://www.sevsu.ru" + URL
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("received non 200 response code (downloadFile)")
	}
	//Create a empty file
	file, err := os.Create("ИИТ-22-о.xlsx")
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the fiel
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func Parsing() {
	// Request the HTML page.
	res, err := http.Get("https://www.sevsu.ru/univers/shedule/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("a.document-link").Each(func(i int, s *goquery.Selection) {
		ss := strings.TrimSpace(s.Find("div.document-link__name").Text())
		if ss == "ИИТ-22-о" {
			// fmt.Println(ss)
			if link, ok := s.Attr("href"); ok{
				fmt.Println("downloading file...")
				if err := downloadFile(link); err != nil{
					fmt.Println(err)
					return
				}
			}
		}
	})
}
