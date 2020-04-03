package main

import (
	"crawler/fetcher"
	"fmt"
	"regexp"
)

func main() {
	all, err := fetcher.Fetch("https://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	printCityList(all)
}

func printCityList(contents []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		fmt.Printf("City: %s, URL: %s\n", m[2], m[1])
	}
	fmt.Printf("Matches found: %d\n", len(matches))
}
