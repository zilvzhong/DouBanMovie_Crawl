package main

import (
	"DouBanMovie_Crawl/crawl"
)

func main() {
	movie := crawl.Get_top()
	crawl.Write(movie)
}

