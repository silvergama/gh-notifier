package main

import (
	"context"
	"log"
	"os"

	"github.com/gen2brain/beeep"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type Notification struct {
	Title string
	Type  string
	URL   string
}

func NewNotification(ghn *github.Notification) *Notification {
	return &Notification{
		Title: ghn.GetRepository().GetName(),
		Type:  ghn.GetSubject().GetType(),
		URL:   ghn.GetSubject().GetURL(),
	}
}

func (n Notification) Notify() {
	err := beeep.Notify(n.Title, n.Type, "")
	if err != nil {
		log.Fatalf("Error %v", err)
	}
}

func main() {
	token := os.Getenv("GH_TOKEN")
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	orgs, _, err := client.Activity.ListNotifications(ctx, nil)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	for _, o := range orgs {
		n := NewNotification(o)
		n.Notify()
	}
}
