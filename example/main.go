package main

import (
	"fmt"

	scraper "github.com/juby210-PL/google-play-scraper"
)

func main() {
	app, err := scraper.GetApp("com.supercell.brawlstars")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(app.Name)
	fmt.Println(app.Developer)
	fmt.Println(app.Category)
	fmt.Println(app.Video)
	fmt.Println(app.Score)
	fmt.Println("InAppPurchases", app.InAppPurchases)
	fmt.Println("Updated", app.Updated)
	fmt.Println("Size", app.Size)
	fmt.Println("Version", app.Version)
}
