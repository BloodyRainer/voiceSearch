package search

import (
	"context"
	engLog "google.golang.org/appengine/log"
	"sync"
	"github.com/BloodyRainer/voiceSearch/df"
)

const domain = "www.otto.de"
const path = "/p/search/"
const protocol = "https"
const queryPrefix = "articlenumber="


func RequestArticleByArticleNr(ctx context.Context, articleNr string) (*df.Article, error) {
	respBody, url, err := searchArticleByNr(ctx, articleNr)
	if err != nil {
		return nil, err
	}

	a, err := createArticleFromHtml(ctx, respBody)
	if err != nil {
		return nil, err
	}

	a.ArticleNr = articleNr
	a.Link = url

	return a, nil
}

func createArticleFromHtml(ctx context.Context, html string) (*df.Article, error) {

	var err error
	var name string
	var price string
	var imgUrl string

	wg := &sync.WaitGroup{}

	wg.Add(3)

	go func() {
		name, err = getName(html)
		if err != nil {
			engLog.Errorf(ctx, "failed to parse name: "+err.Error())
		}
		wg.Done()
	}()

	go func() {
		price, err = getPrice(html)
		if err != nil {
			engLog.Errorf(ctx, "failed to parse price: "+err.Error())
		}
		wg.Done()
	}()

	go func() {
		imgUrl, err = getImageUrl(html)
		if err != nil {
			engLog.Errorf(ctx, "failed to parse url: "+err.Error())
		}
		wg.Done()
	}()

	wg.Wait()

	if err != nil {
		return nil, err
	}

	a := &df.Article{
		Name:   name,
		Price:  price,
		ImgUrl: imgUrl,
	}

	return a, nil
}
