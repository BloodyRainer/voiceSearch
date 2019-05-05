package search

import (
	"google.golang.org/appengine/urlfetch"
	"net/url"
	"io/ioutil"
	"net/http"
	"errors"
	"strings"
	"context"
	"regexp"
	"time"
	"google.golang.org/appengine/log"
	"fmt"
	"github.com/BloodyRainer/voiceSearch/df"
	"golang.org/x/net/html"
	"bytes"
)

var priceReg *regexp.Regexp
var nameReg *regexp.Regexp
var imgUrlReg *regexp.Regexp

func init() {
	priceReg = regexp.MustCompile(`itemprop="price".*content="(\d+\.?\d+)"`)
	nameReg = regexp.MustCompile(`<h1.*itemprop="name".*>(.*)</h1>`)
	//TODO: actual img-url: https://i.otto.de/i/otto/26390776/red-dead-redemption-2-xbox-one.jpg?$formatz$
	//TODO: example of how to customize format https://i.otto.de/i/otto/26390776?h=520&amp;w=384&amp;sm=clamp
	imgUrlReg = regexp.MustCompile(`<meta.*name="twitter:image".*content="(.*?)".*>`)
	//imgUrlReg = regexp.MustCompile(`<img.*id="prd_mainProductImage".*src="(.*?)".*>`)
}

func listArticlesTok(rawHtml []byte) ([]df.Article, error) {

	tokenizer := html.NewTokenizer(bytes.NewReader(rawHtml))
	for {

		t := tokenizer.Next()

		switch {
		case t == html.ErrorToken:
			return nil, errors.New(t.String())
		case t == html.StartTagToken:
			node := tokenizer.Token()

			if node.Data == "article" {
				fmt.Println(node.String())
			}


		}

	}

	return nil, nil
}

func listArticles(rawHtml []byte) ([]df.Article, error) {

	doc, err := html.Parse(bytes.NewReader(rawHtml))
	if err != nil {
		return nil, err
	}

	var f func(node *html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "article" {
			for _, a := range n.Attr {
				if a.Key == "data-productid" {
					fmt.Printf("productId: %v, namespace: %v", a.Val, a.Namespace)
				}
			}
		}

		for e := n.FirstChild; e != nil; e = e.NextSibling {
			f(e)
		}
	}
	f(doc)

	return nil, nil
}



// Appengine needs the original request.
func searchArticleByNr(ctx context.Context, nr string) (string, string, error) {

	client := urlfetch.Client(ctx)

	query := queryPrefix + nr

	url := &url.URL{
		Scheme:   protocol,
		Host:     domain,
		Path:     path,
		RawQuery: query,
	}

	//GET to otto.de
	start := time.Now()
	req := &http.Request{
		Method: "GET",
		URL:    url,
	}
	durStr := fmt.Sprintf("otto GET took %v ms", time.Since(start).Seconds() * 1000)
	log.Infof(ctx, durStr)

	resp, err := client.Do(req)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	return string(body), url.String(), err
}

func getImageUrl(body string) (string, error) {
	imgUrlMatch := imgUrlReg.FindStringSubmatch(string(body))

	if len(imgUrlMatch) < 2 || imgUrlMatch[1] == "" {
		return "", errors.New("no image-url found")
	}

	//imgUrl := customizeImgUrl(imgUrlMatch[1])

	return imgUrlMatch[1], nil
}

func getPrice(body string) (string, error) {

	priceMatch := priceReg.FindStringSubmatch(string(body))
	if len(priceMatch) < 2 || priceMatch[1] == "" {
		return "", errors.New("no price found")
	}

	return priceMatch[1], nil
}

func getName(body string) (string, error) {
	nameMatch := nameReg.FindStringSubmatch(body)
	if len(nameMatch) < 2 || nameMatch[1] == "" {
		return "", errors.New("no name found")
	}

	// TODO: dirty hacks
	name := strings.Replace(nameMatch[1], "&quot;", `'`, -1)
	name = strings.Replace(name, "&amp;", "&", -1)

	return name, nil
}

//may be slow, currently not used
func customizeImgUrl(url string) string {

	const prefix = "https://i.otto.de/i/otto/"
	const postfix = "?h=520&amp;w=384&amp;sm=clamp"

	split := strings.Split(url, "/")

	return prefix + split[5] + postfix
}
