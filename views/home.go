package views

import (
	"net/http"
	"net/url"
)

type HomePage struct {
	Page
	Forum Forum
}

func GetHome(wr http.ResponseWriter, req *http.Request) {

	// Junk data just to test the theory
	forumURL, _ := url.Parse("/forum")
	data := HomePage{
		Page: Page{
			Name:        "HOME",
			Title:       "Home",
			Description: "This is the home page",
		},
		Forum: Forum{
			Title: "Forums",
			Items: []Item{
				Item{
					Title: "Forum",
					URL:   forumURL,
				},
			},
		},
	}

	// The bit I'm interested in
	err := templates[templateHome].Execute(wr, data)
	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
	}
}
