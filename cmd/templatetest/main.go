package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/buro9/templatetest/views"
)

const interfacePort = ":8080"

func main() {
	views.RegisterFlags()

	flag.Parse()

	views.ParseTemplates()

	http.HandleFunc("/", views.GetHome)
	http.HandleFunc("/forum/", views.GetForum)
	http.HandleFunc("/conversation/", views.GetConversation)
	http.HandleFunc("/comment/", views.GetComment)
	http.HandleFunc("/profiles/", views.GetProfiles)
	http.HandleFunc("/profile/", views.GetProfile)

	fmt.Println("Listening on " + interfacePort)
	http.ListenAndServe(interfacePort, nil)
}
