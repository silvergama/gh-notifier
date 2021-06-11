package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gen2brain/beeep"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "token"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	orgs, _, err := client.Activity.ListNotifications(ctx, nil)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	for _, o := range orgs {
		fmt.Println(o.GetRepository().GetFullName())
		err := beeep.Notify("Title", "Message body", "assets/information.png")
		if err != nil {
			log.Fatalf("Error %v", err)
		}
	}
}
