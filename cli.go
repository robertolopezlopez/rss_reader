package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mmcdole/gofeed"
	"github.com/urfave/cli/v3"
)

func main() {

	read := &cli.Command{
		Commands: []*cli.Command{
			{
				Name:        "read",
				Description: "reads an RSS feed and prints the items to standard output",
				Usage:       "read [url]",

				Action: func(context.Context, *cli.Command) error {
					fp := gofeed.NewParser()
					url := os.Args[2]
					fmt.Println("Reading RSS feed from URL:", url)
					feed, err := fp.ParseURL(url)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println(feed)
					fmt.Println("reading!")
					return nil
				},
			},
		},
		Name:        "feed reader",
		Usage:       "reads a given RSS feed",
		Description: "A simple command line tool to read RSS feeds.",
	}
	err := read.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
