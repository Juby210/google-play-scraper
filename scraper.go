package scraper

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// App - Google Play App
type App struct {
	Name      string
	IconURL   string
	Developer string
	Category  string

	Price string
	Free  bool

	Screenshots []string
	Video       string
	VideoImg    string

	Description     string
	DescriptionHTML string

	Score   string
	Reviews string

	ContainsAds    bool
	InAppPurchases bool

	WhatsNew       string
	Updated        string
	Size           string
	Installs       string
	Version        string
	AndroidVersion string
}

// GetApp - Get Google Play App by id
func GetApp(id string) (App, error) {
	app := App{}
	doc, err := goquery.NewDocument("https://play.google.com/store/apps/details?id=" + id + "&hl=en")
	if err != nil {
		return app, err
	}
	app.Name = doc.Find("h1[itemprop=name] span").Text()
	app.IconURL, _ = doc.Find("img[itemprop=image]").First().Attr("src")
	app.Developer = doc.Find("a.hrTbp.R8zArc").First().Text()
	app.Category = doc.Find("a.hrTbp.R8zArc").Last().Text()

	app.Price, _ = doc.Find("meta[itemprop=price]").Attr("content")
	if app.Price == "0" {
		app.Free = true
	}

	doc.Find("button.Q4vdJd img").Each(func(_ int, img *goquery.Selection) {
		src, ok := img.Attr("data-src")
		if ok == false {
			src, _ = img.Attr("src")
		}
		app.Screenshots = append(app.Screenshots, src)
	})

	videl := doc.Find("div.MSLVtf.Q4vdJd")
	app.VideoImg, _ = videl.Find("img").Attr("src")
	app.Video, _ = videl.Find("button").Attr("data-trailer-url")

	deel := doc.Find("div[itemprop=description] div").First()
	app.Description = deel.Text()
	app.DescriptionHTML, _ = deel.Html()

	app.Score = doc.Find("c-wiz div.BHMmbe").Text()
	app.Reviews = doc.Find("c-wiz span.EymY4b span").Last().Text()

	ps := doc.Find("div.bSIuKf").Text()
	if strings.Contains(ps, "Ads") {
		app.ContainsAds = true
	}
	if strings.Contains(ps, "in-app purchases") {
		app.InAppPurchases = true
	}

	app.WhatsNew = doc.Find("div[itemprop=description].DWPxHb").Last().Find("span").Text()
	doc.Find("div.BgcNfc").Each(func(_ int, p *goquery.Selection) {
		val := p.Parent().Find("span").Last().Text()
		switch p.Text() {
		case "Updated":
			app.Updated = val
		case "Size":
			app.Size = val
		case "Installs":
			app.Installs = val
		case "Current Version":
			app.Version = val
		case "Requires Android":
			app.AndroidVersion = val
		}
	})

	return app, nil
}
