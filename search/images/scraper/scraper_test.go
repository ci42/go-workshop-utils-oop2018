package scraper

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func ExampleBingURL() {
	fmt.Printf("%s", BingURL("innoq"))
	// Output:
	// http://www.bing.com/images/search?q=$innoq&scope=images
}

func TestBingURL(t *testing.T) {
	expected := "http://www.bing.com/images/search?q=$innoq&scope=images"
	if got := BingURL("innoq"); got != expected {
		t.Errorf("\nexpected: %s\ngot      :%s\n")
	}
}

func TestBingParser(t *testing.T) {
	expected := []Image{
		{Source: "bing", URL: "http://tse3.mm.bing.net/th?id=OIP.M01828aeb8c6437d79d6f7122bff6e8ffH0&amp;w=230&amp;h=170&amp;rs=1&amp;pcl=dddddd&amp;pid=1.1"},
		{Source: "bing", URL: "http://tse4.mm.bing.net/th?id=OIP.M15860f7f017559efe72aebc23f87d426H0&amp;w=230&amp;h=170&amp;rs=1&amp;pcl=dddddd&amp;pid=1.1"},
		{Source: "bing", URL: "http://tse1.mm.bing.net/th?id=OIP.Mf82e65d5df2ccd76d0a9b08359b09e54o0&amp;w=230&amp;h=170&amp;rs=1&amp;pcl=dddddd&amp;pid=1.1"},
		{Source: "bing", URL: "http://tse1.mm.bing.net/th?id=OIP.M0f6e59faacfeac0cca37a9a6e9431dcbo0&amp;w=230&amp;h=170&amp;rs=1&amp;pcl=dddddd&amp;pid=1.1"},
		{Source: "bing", URL: "http://tse2.mm.bing.net/th?id=OIP.Mcbe7cc53aeec29b3ac2966968cc536b0o0&amp;w=230&amp;h=170&amp;rs=1&amp;pcl=dddddd&amp;pid=1.1"},
		{Source: "bing", URL: "http://tse2.mm.bing.net/th?id=OIP.M11603aba9fc50998da3d08edf04a0e53o0&amp;w=230&amp;h=170&amp;rs=1&amp;pcl=dddddd&amp;pid=1.1"},
	}

	actual := ParseBingResult(strings.NewReader(bingResponse))
	if len := len(actual); len != 28 {
		t.Errorf("error: number of images expected\n%#v\ngot:\n%#v\n", 28, len)
	}

	if !reflect.DeepEqual(expected, actual[:6]) {
		t.Errorf("error: expected\n%#v\ngot:\n%#v\n", expected, actual[:6])
	}

}

func TestShutterstockParser(t *testing.T) {
	expected := []Image{
		{Source: "shutterstock", URL: "http://thumb1.shutterstock.com/display_pic_with_logo/2061440/248662924/stock-photo-close-up-of-a-gopher-248662924.jpg"},
		{Source: "shutterstock", URL: "http://thumb1.shutterstock.com/display_pic_with_logo/3548825/379380670/stock-photo-gophers-sitting-in-sunny-sand-379380670.jpg"},
		{Source: "shutterstock", URL: "http://thumb1.shutterstock.com/display_pic_with_logo/3290069/367412768/stock-vector-gopher-silhouette-icon-367412768.jpg"},
		{Source: "shutterstock", URL: "http://thumb1.shutterstock.com/display_pic_with_logo/2861623/310360364/stock-vector-flat-gopher-icon-310360364.jpg"},
		{Source: "shutterstock", URL: "http://thumb101.shutterstock.com/display_pic_with_logo/2448149/373418053/stock-photo-funny-gopher-in-two-feet-in-green-field-373418053.jpg"},
		{Source: "shutterstock", URL: "http://thumb1.shutterstock.com/display_pic_with_logo/51787/229209310/stock-photo-funny-pet-degu-mouse-with-yellow-teeth-standing-lake-a-gopher-isolated-on-white-background-229209310.jpg"},
	}

	actual := ParseShutterstockResult(strings.NewReader(shutterstockResponse))
	if len := len(actual); len != 100 {
		t.Errorf("error: number of images expected\n%#v\ngot:\n%#v\n", 100, len)
	}

	if !reflect.DeepEqual(expected, actual[:6]) {
		t.Errorf("error: expected\n%#v\ngot:\n%#v\n", expected, actual[:6])
	}
}

func TestFlickrParser(t *testing.T) {
	expected := []Image{
		{Source: "flickr", URL: "http://c6.staticflickr.com/1/183/487234645_45a394f7aa.jpg"},
		{Source: "flickr", URL: "http://c2.staticflickr.com/4/3158/2556794937_bb828ea304_n.jpg"},
		{Source: "flickr", URL: "http://c8.staticflickr.com/4/3647/3487148975_1e752f00d0_n.jpg"},
		{Source: "flickr", URL: "http://c3.staticflickr.com/4/3580/3839144394_02ae9f407e_n.jpg"},
		{Source: "flickr", URL: "http://c5.staticflickr.com/3/2612/3861529164_cef701f3b7.jpg"},
		{Source: "flickr", URL: "http://c2.staticflickr.com/4/3517/4079138689_2e0d912f24.jpg"},
	}

	actual := ParseFlickrResult(strings.NewReader(flickrResponse))
	if len := len(actual); len != 96 {
		t.Errorf("error: number of images expected\n%#v\ngot:\n%#v\n", 96, len)
	}

	if !reflect.DeepEqual(expected, actual[:6]) {
		t.Errorf("error: expected\n%#v\ngot:\n%#v\n", expected, actual[:6])
	}
}
