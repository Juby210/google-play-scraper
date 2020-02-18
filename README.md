# google-play-scraper
Golang Google Play Scraper

```
go get github.com/juby210-PL/google-play-scraper
```
or add require in go.mod

## Example

```go
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
}
```
