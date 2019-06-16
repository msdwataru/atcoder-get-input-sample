package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

var (
	contestName = flag.String("c", "", "name of contest")
	question    = flag.String("q", "a", "which question")
	idx         = flag.Int("i", 1, "index of input sample")
)

func main() {
	flag.Parse()

	baseURL := "https://atcoder.jp/contests/"
	url := baseURL + *contestName + "/tasks/" + *contestName + "_" + *question

	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("pre").Each(func(i int, s *goquery.Selection) {
		if i == (*idx * 2 - 1) {
			fmt.Print(s.Text())
		}
	})

}
