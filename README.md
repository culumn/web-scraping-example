# web-scraping-sample
This program gets the title and URL of the articles from [here](http://b.hatena.ne.jp/ctop/it) with [goquery](https://github.com/PuerkitoBio/goquery) and writes into JSON file.

## Requirement
``` bash
$ go get github.com/PuerkitoBio/goquery
```

## Usage
``` bash
$ go run main.go
```

Then, create json file `articles.json`.
