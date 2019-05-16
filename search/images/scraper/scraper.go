// Package scraper scrapes image urls from image search machines.
//
//     import "gitlab.innoq.com/daniel/go-workshop-utils/search/images/scraper"
//
//     ...
//
//     url := scraper.FlickrURL("apple")
//     imgs := []scraper.Image{}
//
//     resp, err := http.Get(url)
//     if err != nil {
//       return fmt.Errorf("error getting %s: %s\n", url, err)
//     }
//
//     for _, img := range scraper.ParseFlickrResult(resp.Body) {
//       imgs = append(imgs, img)
//     }
//
package scraper

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	"github.com/alecthomas/template"
)

// Image represents an image that gets returned by an image search.
//
//   URL:    url of the image
//   Source: name of the search engine that yielded this image
type Image struct {
	URL    string
	Source string
}

func (i Image) String() string {
	return fmt.Sprintf("%15s: %s", i.Source, i.URL)
}

const (
	bingURLFormatString         = "https://www.bing.com/images/search?q=%s&scope=images"
	pixabayURLFormatString	    = "https://pixabay.com/images/search/%s/"
	unsplashURLFormatString     = "https://unsplash.com/search/photos/%s/"
)

// BingURL returns the bing url that searches for a specific image
//
//   q: search term
func BingURL(q string) string {
	return fmt.Sprintf(bingURLFormatString, template.URLQueryEscaper(q))
}

// PixabayURL returns the pixabay url that searches for a specific image
//
//   q: search term
func PixabayURL(q string) string {
	return fmt.Sprintf(pixabayURLFormatString, template.URLQueryEscaper(q))
}

// UnsplashURL returns the unsplash url that searches for a specific image
//
//   q: search term
func UnsplashURL(q string) string {
	return fmt.Sprintf(unsplashURLFormatString, template.URLQueryEscaper(q))
}

func extractImageUrls(r io.Reader, urlPart string, source string) []Image {
	images := []Image{}
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	body := string(b)

	res := regexp.MustCompile(`src=['"](`+urlPart+`.*?)['"]`).FindAllStringSubmatch(body, -1)
	for _, url := range res {
		images = append(images, Image{URL: url[1], Source: source})
	}

	return images
}

func parseType1Results(r io.Reader, requiredSubString, source string) []Image {
	images := []Image{}
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	body := string(b)

	res := regexp.MustCompile(`<img`).ReplaceAllString(body, "\n<img")
	res = strings.Join(regexp.MustCompile("<img.*").FindAllString(res, -1), "\n")
	res = regexp.MustCompile(`.*src=["']`).ReplaceAllString(res, "")
	res = regexp.MustCompile(`['"].*`).ReplaceAllString(res, "")

	for _, url := range regexp.MustCompile(".*"+requiredSubString+".*").FindAllString(res, -1) {
		images = append(images, Image{URL: url, Source: source})
	}

	return images
}

// ParseBingResult parses a bing response and returns an array of the images the search returned
func ParseBingResult(r io.Reader) []Image {
	return extractImageUrls(r, "https://www.bing.com/th", "bing")
	//return parseType1Results(r, "http", "bing")
}

// ParseBingResult parses a bing response and returns an array of the images the search returned
func ParsePixabayResult(r io.Reader) []Image {
	return extractImageUrls(r, "https://cdn.pixabay.com/photo/", "pixabay")
}

// ParseBingResult parses a bing response and returns an array of the images the search returned
func ParseUnsplashResult(r io.Reader) []Image {
	return extractImageUrls(r, "https://images.unsplash.com/photo", "unsplash")
}