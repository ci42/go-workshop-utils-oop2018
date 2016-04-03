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
	bingURLFormatString         = "http://www.bing.com/images/search?q=$%s&scope=images"
	shutterstockURLFormatString = "http://www.shutterstock.com/cat.mhtml?autocomplete_id=&language=de&lang=de&search_source=&safesearch=1&version=llv1&searchterm=%s&media_type=images"
	flickrURLFormatString       = "https://www.flickr.com/search/?text=%s"
)

// BingURL returns the bing url that searches for a specific image
//
//   q: search term
func BingURL(q string) string {
	return fmt.Sprintf(bingURLFormatString, template.URLQueryEscaper(q))
}

// FlickURL returns the flickr url that searches for a specific image
//
//   q: search term
func FlickrURL(q string) string {
	return fmt.Sprintf(flickrURLFormatString, template.URLQueryEscaper(q))
}

// ShutterstokURL returns the shutterstock url that searches for a specific image
//
//   q: search term
func ShutterstockURL(q string) string {
	return fmt.Sprintf(shutterstockURLFormatString, template.URLQueryEscaper(q))
}

// ParseFlickrResult parses a flickr response and returns an array of the images the search returned
func ParseFlickrResult(r io.Reader) []Image {
	images := []Image{}
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	body := string(b)

	res := strings.Join(regexp.MustCompile(`img\.src=.*`).FindAllString(body, -1), "\n")
	res = regexp.MustCompile(`img\.src='`).ReplaceAllString(res, "http:")
	res = regexp.MustCompile(`'.*`).ReplaceAllString(res, "")

	for _, url := range regexp.MustCompile("http.*").FindAllString(res, -1) {
		images = append(images, Image{URL: url, Source: "flickr"})
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
	return parseType1Results(r, "http", "bing")
}

// ParseShutterstockResult parses a shutterstock response and returns an array of the images the search returned
func ParseShutterstockResult(r io.Reader) []Image {
	return parseType1Results(r, "thumb", "shutterstock")
}
