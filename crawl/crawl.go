package crawl

import (
	"DouBanMovie_Crawl/config"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/gocolly/colly/queue"
	"net/url"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Movie struct {
	Name string
	Director string
	Score string
	Comment string
	Addr string
}


func Get_top() []Movie {
	var movie []Movie
	c := colly.NewCollector()
	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	q, _ := queue.New(
		2,
		&queue.InMemoryQueueStorage{MaxSize: 1000})

	c.OnRequest(func(request *colly.Request) {
	})
	for i := 0; i < 250; {
		n := strconv.Itoa(i)
		s := fmt.Sprintf(config.Defaults["doubian_url"],n)
		q.AddURL(s)

		i = i + 25
	}


	c.OnHTML("ol", func(e *colly.HTMLElement) {

		e.DOM.Find("div.item").Each(func(i int, selection *goquery.Selection) {
			name := Remove(selection.Find("a[href]").Text())
			movie = append(movie, Movie{
				name,
				Remove(selection.Find("div>p").Text()),
				Remove(selection.Find("span.rating_num").Text()),
				Remove(selection.Find("span.inq").Text()),
				"",
			})
		})
	})

	q.Run(c)
	return movie
}

func Get_rrdyw(name string) string {
	var str string
	c := colly.NewCollector()

	ul := "http://www.rrdyw.cc/plus/search.php?q=" + url.QueryEscape(name) + "&pagesize=10&submit="

	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if !(strings.Contains(strings.TrimSpace(e.Text), name)) {
			return
		}
		str = e.Request.AbsoluteURL(link)
	})


	c.Visit(ul)
	c.Wait()
	return str
}

func Remove(str string)  string {
	var out []rune
	for _, r := range str {
		if !unicode.IsSpace(r) {
			out = append(out, r)
		}
	}
	return string(out)
}

func Write(movie []Movie) error {
	path, _ := os.Getwd()
	fmt.Println(path)
	f, err := os.OpenFile(path + "/movie.txt", os.O_WRONLY| os.O_TRUNC, 0777)
	fmt.Println(111)
	if err != nil {
		fmt.Println(222)
		return err
	}
	defer f.Close()
	for _, line := range movie {
		addr := Get_rrdyw(line.Name)
		fmt.Fprintf(f, "%s\n%s\n%s\n%s\n%s\n",
			line.Name,
			line.Director,
			line.Score,
			line.Comment,
			addr)
	}
	return nil
}
