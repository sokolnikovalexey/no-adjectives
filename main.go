package main

import (
	"github.com/grokify/html-strip-tags-go"
	"github.com/ungerik/go-rss"
	"gitlab.com/opennota/morph"
	"html/template"
	"log"
	"net/http"
	"no-adj-news/cmd"
	"strings"
)

func main() {
	if err := morph.Init(); err != nil {
		log.Println(err)
	}

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/fetch-news", fetchNewsRouterHandler)
	http.HandleFunc("/fetch-categories", fetchCategoriesRouterHandler)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Println("can't start server")
	}
}

func fetchNewsRouterHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err == nil {
		categoryId := r.Form.Get("cat")

		_, err := cmd.GetCategoryById(categoryId)
		if err != cmd.UnknownCategoryError {
			tpl, _ := template.ParseFiles("web/tpl/news-item.html")
			url := "https://news.yandex.ru/" + categoryId + ".rss"
			tpl.Execute(w, parseRssUrl(url))
		} else {
			log.Println("Unknown category " + categoryId)
		}
	} else {
		log.Println("can't parse form")
	}

}

func fetchCategoriesRouterHandler(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("web/tpl/category-item.html")
	tpl.Execute(w, cmd.GetAvailableCategories())
}

func parseRssUrl(url string) []NewsFeedItem {
	channel, err := rss.Read(url)
	if err != nil {
		log.Println(err)
	}

	var items []NewsFeedItem
	for _, rssItem := range channel.Item {
		items = append(items, sanitizeNewsFeedItem(rssItem))
	}
	return items
}

func sanitizeNewsFeedItem(item rss.Item) NewsFeedItem {
	descriptionTokens := strings.Split(item.Description, " ")

	description := ""
	for _, token := range descriptionTokens {
		description += " " + replaceAdjs(token)
	}

	return NewsFeedItem{item.Title, strings.TrimSpace(strip.StripTags(description)), item.Link}
}

func replaceAdjs(token string) string {
	isAdj := false
	words, _, tags := morph.Parse(strings.ToLower(token))
	for i := range words {
		if strings.Contains(tags[i], "ADJ") {
			isAdj = true
		}
	}

	if isAdj {
		return "***"
	} else {
		return token
	}
}

type NewsFeedItem struct {
	Title       string
	Description string
	Link        string
}

