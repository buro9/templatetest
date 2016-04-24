package views

import (
	"net/http"
	"net/url"
)

type ForumPage struct {
	Page

	Forum Forum
}

type Forum struct {
	// Title is the display title for this forum
	Title string

	// URL is the link to this forum
	URL *url.URL

	// Items is the collection of things in the forum, i.e. conversations
	Items []Item
}

type Item struct {
	// Title is the display title for this item within a forum or search results
	Title string

	// URL is the link to this item
	URL *url.URL
}

func GetForum(wr http.ResponseWriter, req *http.Request) {
	conversationURL, _ := url.Parse("/conversation")
	forumURL, _ := url.Parse("/forum")

	data := ForumPage{
		Page: Page{
			Name:        "FORUM",
			Title:       "Forum",
			Description: "This is the forum page",
		},
		Forum: Forum{
			Title: "Forum",
			URL:   forumURL,
			Items: []Item{
				Item{
					Title: "Conversation",
					URL:   conversationURL,
				},
			},
		},
	}

	err := templates[templateForum].Execute(wr, data)
	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
	}
}
