package main

import (
	"context"
	"deepai"
	"deepai/option"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	token := flag.String("t", "", "a valid deep ai token")
	filename := flag.String("f", "", "a valid path to an image file")
	api := flag.String("api", "", "an identifier for a deepai api (nsfw, moderation)")
	flag.Parse()

	ctx := context.Background()

	img, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	defer img.Close()

	client := deepai.NewClient(*token, option.Timeout(time.Second*10))

	switch strings.ToLower(*api) {
	case "nsfw":
		response, err := client.DetectNudity(ctx, img)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v", response)

	case "moderation":
		response, err := client.Moderate(ctx, img)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v", response)

	default:
		log.Fatal("the provided api is not valid or not yet supported")
	}
}
